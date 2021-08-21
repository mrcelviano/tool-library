package tool_library

import (
	"github.com/gocraft/dbr"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type DataBase struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

func InitConnectionPostgres() *dbr.Connection {
	connectionString := "host=" + configuration.DataBase.Host + " port=" + configuration.DataBase.Port + " user=" + configuration.DataBase.User + " dbname=" + configuration.DataBase.DBName + " sslmode=disable password=" + configuration.DataBase.Password
	conn, err := dbr.Open("postgres", connectionString, nil)
	if err != nil {
		log.Printf("Error dbr: %v \n", err)
	}
	conn.SetMaxOpenConns(6)
	conn.SetMaxIdleConns(2)
	return conn
}
