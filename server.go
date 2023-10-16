package main

import (
	"context"
	"fmt"
	"path/filepath"
	"pizza-backend/common"
	"pizza-backend/database"
	"pizza-backend/models"
	"pizza-backend/resolvers"
	"strings"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"
const storagePath = "./STORAGE"

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	customCtx := &common.CustomContext{
		Database: db,
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//AllowCredentials: true,
		Debug: true,
	}).Handler)

	router.Use(models.Middleware(db))
	c := resolvers.Config{Resolvers: &resolvers.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.UserRole) (interface{}, error) {
		var ur models.UserRoles = roles

		if user := models.ForContext(ctx); user == nil || !ur.CheckRoles(user.Roles) {
			return nil, fmt.Errorf("Access denied")
		}
		// or let it pass through
		return next(ctx)
	}

	FileServer(router, "/IMAGES", http.Dir(filepath.Join(storagePath, "IMAGES")))

	srv := handler.NewDefaultServer(resolvers.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
