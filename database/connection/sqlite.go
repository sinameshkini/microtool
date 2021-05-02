package connection

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqliteDB() (db *gorm.DB, err error) {
	dbFile := viper.GetString("database.sqlite.file")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	return gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: newLogger,
	})
}
