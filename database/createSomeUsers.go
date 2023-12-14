package database

import (
	"pizza-backend/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateAdmin(newdb *gorm.DB) {

	roles := make([]models.UserRole, 1, 10)
	roles[0] = "ADMIN"

	preferences := &models.UserPreferences{
		UseIdle:          false,
		IdleTimeout:      100,
		UseIdlePassword:  false,
		IdlePin:          "9999",
		UseDirectLogin:   false,
		UseQuadcodeLogin: false,
		SendNoticesMail:  false,
		Language:         "it",
	}

	details := &models.UserDetails{
		Title:     "Sig.",
		FirstName: "Carlito",
		LastName:  "Prada",
		Address:   "Strada Regina 12d",
		City:      "Manno",
		ZipCode:   "6928",
		Country:   "Switzerland",
		Phone:     "078 824 10 89",
	}

	user := &models.User{
		Email:    "carlito@prova.com",
		Name:     "Carlito",
		Roles:    roles,
		Types:    "SITE",
		Password: models.HashPassword("password"),
	}

	err := newdb.Create(&user).Error
	if err != nil {
		log.Fatalln("-----> errore creazione utente")
	}
	preferences.UserId = user.ID
	user.Preferences = preferences
	details.UserID = user.ID
	user.Details = details
	err = newdb.Save(&user).Error
	if err != nil {
		log.Fatalln("-----> errore creazione utente")
	}
}

func CreateUser(newdb *gorm.DB) {

	roles := make([]models.UserRole, 1, 10)
	roles[0] = "USER"

	preferences := &models.UserPreferences{
		UseIdle:          false,
		IdleTimeout:      100,
		UseIdlePassword:  false,
		IdlePin:          "9999",
		UseDirectLogin:   false,
		UseQuadcodeLogin: false,
		SendNoticesMail:  false,
		Language:         "it",
	}

	details := &models.UserDetails{
		Title:     "Sig.",
		FirstName: "Pippo",
		LastName:  "De Pippis",
		Address:   "Via Pipponia",
		City:      "Pippola",
		ZipCode:   "5555",
		Country:   "Pippolandia",
		Phone:     "0000 567 567",
	}

	user := &models.User{
		Email:    "pippo@mail.com",
		Name:     "Pippo",
		Roles:    roles,
		Types:    "SITE",
		Password: models.HashPassword("password"),
	}

	err := newdb.Create(&user).Error
	if err != nil {
		log.Fatalln("-----> errore creazione utente")
	}
	preferences.UserId = user.ID
	user.Preferences = preferences
	details.UserID = user.ID
	user.Details = details
	err = newdb.Save(&user).Error
	if err != nil {
		log.Fatalln("-----> errore creazione utente")
	}
}
