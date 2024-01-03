package main

import (
	"context"
	"fmt"
	"path/filepath"
	"pizza-backend/common"
	"pizza-backend/config"
	"pizza-backend/database"
	"pizza-backend/models"
	"pizza-backend/resolvers"
	"pizza-backend/schema"

	"strings"

	"net/http"

	log "github.com/sirupsen/logrus"

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

type myFormatter struct {
	log.TextFormatter
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
	// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 31 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("\x1b[%dm [%s] - %s - %s\x1b[0m\n", levelColor, entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func main() {

	log.SetFormatter(&myFormatter{
		log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
	},
	)

	config.Initialize("./config/config.yaml")
	conf := config.GetYamlValues()

	port := conf.ServerConfig.Port
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
		AllowedOrigins: []string{"https://pizza.dyndns.winner4ever.com", "https://localhost:9000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//AllowCredentials: true,
		Debug: true,
	}).Handler)

	router.Use(models.Middleware(db))
	c := resolvers.Config{Resolvers: &resolvers.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.UserRole) (interface{}, error) {
		var userRoles models.UserRoles = roles
		if user := models.ForContext(ctx); user == nil || !userRoles.CheckRoles(user.Roles) {
			return nil, fmt.Errorf("access denied")
		}
		// or let it pass through
		return next(ctx)
	}
	c.Directives.NeedAuth = func(ctx context.Context, obj interface{}, next graphql.Resolver, need *bool) (interface{}, error) {
		if !*need {
			return next(ctx)
		}
		if user := models.ForContext(ctx); user == nil {
			return nil, fmt.Errorf("access denied")
		}
		// or let it pass through
		return next(ctx)
	}

	FileServer(router, "/IMAGES", http.Dir(filepath.Join(storagePath, "IMAGES")))
	FileServer(router, "/schema", http.Dir("."))

	srv := handler.NewDefaultServer(resolvers.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", common.CreateContext(customCtx, srv))

	schema.String()

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln("-----> not able to connect: ", port)
	}

}
