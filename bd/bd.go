package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Db := "personal"

	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Db)

	if err != nil {
		panic(err.Error())
	}

	return connection
}
