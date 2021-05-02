package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		viper.GetString("database.connection.username"),
		viper.GetString("database.connection.password"),
		viper.GetString("database.connection.db_name"),
		viper.GetString("database.connection.port"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
