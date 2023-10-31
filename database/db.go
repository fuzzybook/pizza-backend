package database

import (
	"os"
	"pizza-backend/config"
	"pizza-backend/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func createDb() {
	conf := config.GetYamlValues()

	err := os.Remove(conf.SqliteConfig.Database)
	if err != nil {
		log.Fatal(err)
	}

	newdb, err := gorm.Open(sqlite.Open(conf.SqliteConfig.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error)
	}

	newdb.AutoMigrate(
		&models.User{},
	)

	roles := make([]models.UserRole, 1, 10)
	roles[0] = "ADMIN"

	user := &models.User{
		Email:    "pippo@mail.ccom",
		Name:     "Pippo",
		Roles:    roles,
		Types:    "SITE",
		Password: models.HashPassword("password"),
	}

	err = newdb.Create(&user).Error
	if err != nil {
		log.Fatalln("-----> errore creazione utente")
	}
}

func InitDb() (*gorm.DB, error) {
	var err error
	conf := config.GetYamlValues()

	createDb()

	db, err := gorm.Open(sqlite.Open(conf.SqliteConfig.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
