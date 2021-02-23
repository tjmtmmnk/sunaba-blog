package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = os.Getenv("MYSQL_USER")
	config.Passwd = os.Getenv("MYSQL_USER_PASSWORD")
	config.Addr = os.Getenv("MYSQL_ADDRESS")
	config.DBName = os.Getenv("MYSQL_DBNAME")
	connector, err := mysql.NewConnector(config)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
