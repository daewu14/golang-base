package database

import (
	"github.com/daewu14/golang-base/config"
	"gorm.io/gorm"
)

var isConnected = false
var dbConnection *dbc

type IDbConnection interface {
	Main() *gorm.DB
	Replicas() []*gorm.DB
	Open() *dbc
}

type dbc struct {
	main     *gorm.DB
	replicas []*gorm.DB
}

// Main database connection
func (d dbc) Main() *gorm.DB {
	return d.main
}

// Replicas database connection
func (d dbc) Replicas() []*gorm.DB {
	return d.replicas
}

// Open database connection
func Open() *dbc {
	return openConnection()
}

// Main database uses
func Main() *gorm.DB {
	return Open().Main()
}

// Replicas database uses
func Replicas() []*gorm.DB {
	return Open().Replicas()
}

// Close database connection
func Close() {
	isConnected = false
	db, err := openConnection().Main().DB()
	if err != nil {
		return
	}
	db.Close()
}

// openConnection : To open connection with GORM
func openConnection() *dbc {
	err := checkDatabaseConnectionSupport()
	if err != nil {
		panic(err.Error())
	}

	var db = config.Database()

	var connection *dbc
	switch db.Connection {
	case MYSQL:
		connection = mysqlOpenConnection()
	default:
		connection = mysqlOpenConnection()
	}
	return connection
}
