package connection

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMsSqlDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		viper.GetString("database.mssql.username"),
		viper.GetString("database.mssql.password"),
		viper.GetString("database.mssql.host"),
		viper.GetString("database.mssql.port"),
		viper.GetString("database.mssql.db_name"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		return
	}

	return
}
