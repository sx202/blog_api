package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


//初始化数据库的连接
func LinkDb()(db *sql.DB,err error) {

	db,err = sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		return nil,err
	}
	return db,err
}
