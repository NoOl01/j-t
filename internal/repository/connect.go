package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"johny-tuna/internal/config"
	"johny-tuna/internal/models"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4",
		config.Env.DbUser, config.Env.DbPass, config.Env.DbPort, config.Env.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if migrateErr := db.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{}, &models.CartItem{}, &models.Points{}); migrateErr != nil {
		panic(err)
	}

	return db
}
