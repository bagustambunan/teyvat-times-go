package db

import (
	"database/sql"
	"final-project-backend/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	db       *gorm.DB
	dbConfig = config.Config.DBConfig
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	if config.GetENV("DATABASE_URL", "") != "" {
		sqlDB, dbErr := sql.Open("postgres", config.GetENV("DATABASE_URL", ""))
		if dbErr != nil {
			return dbErr
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			Logger: getLogger(),
		})
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})
	return err

}

func Get() *gorm.DB {
	return db
}
