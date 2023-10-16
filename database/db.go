package database

import (
	"log"
	"os"
	"pizza-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func createDb() {
	err := os.Remove("dataBase14.sql")
	if err != nil {
		log.Fatal(err)
	}

	newdb, err := gorm.Open(sqlite.Open("dataBase14.sql"), &gorm.Config{
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
		panic("errore creazione utente")
	}
}

func InitDb() (*gorm.DB, error) {
	var err error

	createDb()

	db, err := gorm.Open(sqlite.Open("dataBase14.sql"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
