package database

import (
	"database/sql"
	"errors"
	"log"

	"api.default.marincor/clients/postgres"
	"api.default.marincor/pkg/app"
	"gorm.io/gorm"
)

var errConnectDB = errors.New("error to connect to database")

// Using generics.
func Query[T interface{}](query string, outputType T, args ...interface{}) (T, error) { //nolint:ireturn
	gormConn, conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	defer conn.Close()

	err := gormConn.Raw(query, args...).Scan(&outputType).Error

	return outputType, err
}

func Exec(query string, args ...interface{}) error {
	_, conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	defer conn.Close()

	err := conn.QueryRow(query, args...).Err()

	return err
}

func QueryCount(query string, args ...interface{}) (int, error) {
	var count int

	_, conn := Connect(app.Inst.Config.DBString, app.Inst.Config.DBLogMode, app.Inst.Config.Debug)
	defer conn.Close()

	err := conn.QueryRow(query, args...).Scan(&count)

	return count, err
}

func Connect(dbString string, logLevel int, debug bool) (*gorm.DB, *sql.DB) {
	gormDB, databaseConnection, err := postgres.Connect(dbString, logLevel, debug)
	if err != nil {
		log.Panicln(errConnectDB, err)
	}

	return gormDB, databaseConnection
}
