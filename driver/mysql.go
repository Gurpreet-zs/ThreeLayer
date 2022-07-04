package driver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql.
// Connecting sql
func ConnectToMySQL(conf MySQLConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Print(err)
	}

	return db, nil
}
