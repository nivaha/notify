package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"notify/event"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Nivaha Notify\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hey, %s!\n", ps.ByName("name"))
}

func Setup(db *sql.DB) {
	router = httprouter.New()

	router.GET("/", index)
	router.GET("/hello/:name", hello)

	router.GET("/events", event.Index)
	router.POST("/event", event.Create)
}

func Start(port int) {
	log.Printf("Listening on port %v", port)

	portStr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}
