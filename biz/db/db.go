package db

import (
	"database/sql"
	"gaecharge/config"
	"fmt"
)

var dbConn *sql.DB

func init0() {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.AppConfig.Database.Username,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.Db)

	db, err := sql.Open("mysql", connString)
	if nil != err {
		panic(err)
	}

	dbConn = db
}

func ExecuteUpdate(sql string, args []interface{}) (int64, error) {
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
