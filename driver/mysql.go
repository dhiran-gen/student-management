package driver

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql.
func ConnectToMySQL(conf MySQLConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	db, err := sql.Open(conf.Driver, connectionString)
	if err != nil {
		return nil, errors.New("cannot connect to sql server")
	}

	return db, nil
}
