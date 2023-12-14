package database

import (
	"pizza-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDb() {
	//conf := config.GetYamlValues()

	/* err := os.Remove(conf.SqliteConfig.Database)
	if err != nil {
		log.Fatal(err)
	} */

	dsn := "host=localhost user=pizzaiolo password=pizzaiolo dbname=pizzeria port=5432 sslmode=disable TimeZone=Europe/Rome"
	newdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	/* newdb, err := gorm.Open(sqlite.Open(conf.SqliteConfig.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err.Error)
	} */

	newdb.AutoMigrate(
		&models.User{},
		&models.UserDetails{},
		&models.UserPreferences{},
		&models.Session{},

		&models.MenuIngredient{},
		&models.MenuDough{},
		&models.MenuCondiment{},
		&models.MenuCategory{},
		&models.MenuItem{},
		&models.MenuOrder{},
		&models.MenuTimes{},
	)

	//CreateAdmin(newdb)
	//CreateUser(newdb)

}

func InitDb() (*gorm.DB, error) {
	var err error
	//conf := config.GetYamlValues()

	createDb()

	dsn := "host=localhost user=pizzaiolo password=pizzaiolo dbname=pizzeria port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	/* db, err := gorm.Open(sqlite.Open(conf.SqliteConfig.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	} */

	return db, nil
}
