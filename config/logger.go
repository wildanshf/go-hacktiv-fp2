package config

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

//InitLog make log for database query
func InitLog() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)

	return newLogger
}
