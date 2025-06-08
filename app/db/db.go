package db

import (
	"fmt"
	"tasks_company/app/config"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.User,
		dbPassword,
		config.Cfg.Database.DBName,
		config.Cfg.Database.SSLMode)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.DB{})
	if err != nil {
		logger.LogWithFieldsFatalf("db", "opening", err)
	}

	err = DB.AutoMigrate(&models.Tasks{}, &models.User{})
	if err != nil {
		logger.LogWithFieldsFatalf("db", "migration", err)
	}
}

func CloseDB() {
	postgresqlDB, err := DB.DB()
	if err != nil {
		logger.LogWithFieldsFatalf("db", "instance", err)
	}

	if err := postgresqlDB.Close(); err != nil {
		logger.LogWithFieldsFatalf("db", "closing", err)
	}
}
