package database

import (
	"database/sql"
	"go_base_project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var isConnected = false
var dbConnection *dbc

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

// connectToMySql : Simplify method and reusable method to connect with mysql driver
func connectToMySql(username string, password string, host string, port string, name string, useTimeStamp string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=" + strings.ToLower(useTimeStamp)
	sqlDB, err1 := sql.Open(MYSQL, dsn)

	if err1 != nil {
		panic(err1.Error())
	}

	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn:                     sqlDB,
		DisableDatetimePrecision: false,
	}), &gorm.Config{})

	if err2 != nil {
		panic(err2.Error())
	}

	return gormDB
}

// mysqlOpenConnection : To open connection with mysql driver
func mysqlOpenConnection() *dbc {
	if isConnected {
		return dbConnection
	}

	var db = config.Database()

	gormDB := connectToMySql(db.Username, db.Password, db.Host, db.Port, db.Name, db.UseTimestamp)

	var replicas []*gorm.DB

	// Do connection on replicas mysql database
	if len(db.Replicas) > 0 {
		for _, replica := range db.Replicas {
			if replica.Connection == MYSQL {
				replicaConnect := connectToMySql(replica.Username, replica.Password, replica.Host, replica.Port, replica.Name, replica.UseTimestamp)
				replicas = append(replicas, replicaConnect)
			}
		}
	}

	// Set initiator variable flag to connected for singleton reason
	isConnected = true
	dbConnection = &dbc{
		main:     gormDB,
		replicas: replicas,
	}

	return dbConnection
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