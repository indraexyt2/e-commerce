package helpers

import (
	"e-commerce-ums/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("Error connecting to database: ", err)
		return
	}

	DB.Exec("CREATE TYPE user_role AS ENUM ('user', 'admin')")
	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		Logger.Error("Error migrating database: ", err)
		return
	}
	Logger.Info("Connected to database")
}
