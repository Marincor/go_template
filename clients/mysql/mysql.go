package mysql

import (
	"database/sql"
	"fmt"
	"log"

	// Import mysql.
	_ "github.com/go-sql-driver/mysql"
)

func Connect(dbString string) *sql.DB {
	const maxIdleConnection = 5
	const maxOpenConnection = 30
	const connectionMaxLifetime = 60

	pool, err := sql.Open("mysql", fmt.Sprintf("%s?%s", dbString, "charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	pool.SetMaxIdleConns(maxIdleConnection)
	pool.SetMaxOpenConns(maxOpenConnection)
	pool.SetConnMaxLifetime(connectionMaxLifetime)

	if err := pool.Ping(); err != nil {
		panic(err)
	}

	log.Println("mysql database connected successfully")

	return pool
}
