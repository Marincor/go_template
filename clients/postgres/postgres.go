package postgres

import (
	"database/sql"
	"time"

	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgres struct {
	DBString              string
	DBLogMode             logger.LogLevel
	maxIdleConnection     int
	maxOpenConnection     int
	connectionMaxLifetime int
	Debug                 bool
}
type Options func(*postgres)

func Connect(dbString string, dbLogMode int, debug bool, sqloptions ...Options) (*gorm.DB, *sql.DB, error) {
	const maxIdleConnection = 5
	const maxOpenConnection = 30
	const connectionMaxLifetime = 60

	databaseConnection := &postgres{
		DBString:              dbString,
		DBLogMode:             logger.LogLevel(dbLogMode),
		maxIdleConnection:     maxIdleConnection,
		maxOpenConnection:     maxOpenConnection,
		connectionMaxLifetime: connectionMaxLifetime,
		Debug:                 debug,
	}

	for _, o := range sqloptions {
		o(databaseConnection)
	}

	return connect(databaseConnection)
}

func SetMaxIdleConns(connections int) Options {
	return func(c *postgres) {
		if connections > 0 {
			c.maxIdleConnection = connections
		}
	}
}

func SetMaxOpenConn(connections int) Options {
	return func(c *postgres) {
		if connections > 0 {
			c.maxOpenConnection = connections
		}
	}
}

func SetMaxLifetime(lifetime int) Options {
	return func(c *postgres) {
		if lifetime > 0 {
			c.connectionMaxLifetime = lifetime
		}
	}
}

func connect(data *postgres) (*gorm.DB, *sql.DB, error) {
	var gormConfig gorm.Config

	if data.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(data.DBLogMode)
	}

	gormConfig.SkipDefaultTransaction = true
	gormConfig.PrepareStmt = true

	databaseConnection, err := gorm.Open(postgresDriver.Open(data.DBString), &gormConfig)
	if err != nil {
		return nil, nil, err
	}

	pgDB, _ := databaseConnection.DB()
	pgDB.SetConnMaxLifetime(time.Duration(data.connectionMaxLifetime) * time.Minute)
	pgDB.SetMaxOpenConns(data.maxOpenConnection)
	pgDB.SetMaxIdleConns(data.maxIdleConnection)

	return databaseConnection, pgDB, nil
}
