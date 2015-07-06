package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
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
	flag.StringVar(&clOptions.dbInfo.databaseName, "dbName", "development", "database name")

	flag.IntVar(&clOptions.httpPort, "httpPort", 8080, "The port to listen on")

	flag.Parse()
}

func main() {
	log.Println(filepath.Base(os.Args[0]))

	dbOpen()

	setupRoutes()

	startRouter()
}
