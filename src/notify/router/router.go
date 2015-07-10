package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"notify/event"
	"notify/subscription"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func Setup(db *sql.DB) {
	router = httprouter.New()

	router.GET("/", index)
	router.GET("/status", status)

	router.POST("/events", event.Create)
	router.GET("/events", event.Index)
	router.GET("/events/:id", event.Show)

	router.POST("/subscriptions", subscription.Create)
	router.GET("/subscriptions", subscription.Index)
	router.GET("/subscriptions/:id", subscription.Show)
	router.DELETE("/subscriptions/:id", subscription.Destroy)
}

func Start(port int) {
	log.Printf("Listening on port %v", port)

	portStr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}
