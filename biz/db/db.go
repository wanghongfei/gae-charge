package db

import (
	"database/sql"
	"gaecharge/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbConn *sql.DB

func InitDb() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.AppConfig.Database.Username,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.Db)

	log.Println("db connection string: " + connString)
	db, err := sql.Open("mysql", connString)
	if nil != err {
		panic(err)
	}

	// test connection
	err = db.Ping()
	if nil != err {
		panic(err)
	}

	dbConn = db

	log.Printf("db initialized")
}

func ExecuteUpdate(sql string, args ...interface{}) (int64, error) {
	stmt, err := dbConn.Prepare(sql)
	if nil != err {
		return -1, err
	}

	result, err := stmt.Exec(args)
	affected, err := result.RowsAffected()
	if nil != err {
		return -1, err
	}

	return affected, nil
}
