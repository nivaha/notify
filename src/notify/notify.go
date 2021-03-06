package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"notify/event"
	"notify/feed"
	"notify/router"
	"notify/subscription"
)

var clOptions struct {
	dbInfo   dbAccess
	httpPort int
}

func init() {
	flag.StringVar(&clOptions.dbInfo.username, "dbUsername", "postgres", "database username")
	flag.StringVar(&clOptions.dbInfo.password, "dbPassword", "", "database password")
	flag.StringVar(&clOptions.dbInfo.host, "dbHost", "localhost", "database host")
	flag.IntVar(&clOptions.dbInfo.port, "dbPort", 5004, "database port")
	flag.StringVar(&clOptions.dbInfo.databaseName, "dbName", "notify_dev", "database name")

	flag.IntVar(&clOptions.httpPort, "httpPort", 8080, "The port to listen on")

	flag.Parse()
}

func main() {
	log.Println(filepath.Base(os.Args[0]))

	db := dbOpen()
	defer dbClose()

	fatalIfError(event.CreateDB(db))
	fatalIfError(feed.CreateDB(db))
	fatalIfError(subscription.CreateDB(db))

	router.Setup()
	router.Start(clOptions.httpPort)
}
