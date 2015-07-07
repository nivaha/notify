package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type dbAccess struct {
	username     string
	password     string
	host         string
	port         int
	databaseName string
}

var db *sql.DB

func dbConnectURL(info dbAccess) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		info.username, info.password,
		info.host, info.port,
		info.databaseName)
}

func dbOpen() *sql.DB {
	url := dbConnectURL(clOptions.dbInfo)
	log.Println("Connecting to:", url)

	var err error

	db, err = sql.Open("postgres", url)
	fatalIfError(err)

	err = db.Ping()
	fatalIfError(err)

	return db
}

func dbClose() {
	db.Close()
}
