package connection

import (
	"errors"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewDatabase(connectionName string) (db *gorm.DB, err error) {
	if connectionName == "" {
		connectionName = viper.GetString("database.use")
	}

	switch connectionName {
	case "mssql":
		return NewMsSqlDB()
	case "postgres":
		return NewPostgresDB()
	case "sqlite":
		return NewSqliteDB()
	default:
		err = errors.New("connection name: " + connectionName + ", database connection args does not exist in config file")
	}

	return
}
