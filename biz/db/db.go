package db

import "database/sql"

var dbConn *sql.DB

func init0() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gae?charset=utf8")
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
