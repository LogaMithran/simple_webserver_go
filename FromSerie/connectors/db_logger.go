package connector

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetCustomDbLogger() logger.Interface {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	return logger
}
