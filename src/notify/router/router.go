package router

import (
	"fmt"
	"log"
	"net/http"

	"notify/event"
	"notify/feed"
	"notify/subscription"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

// Setup creates a router and sets up all the routes
func Setup() {
	router = httprouter.New()

	router.GET("/", index)
	router.GET("/status", status)

	evtHandler := event.NewHandler()
	router.POST("/events", evtHandler.Create)
	router.GET("/events", evtHandler.Index)
	router.GET("/events/:id", evtHandler.Show)

	subHandler := subscription.NewHandler()
	router.POST("/subscriptions", subHandler.Create)
	router.GET("/subscriptions", subHandler.Index)
	router.GET("/subscriptions/:id", subHandler.Show)
	router.DELETE("/subscriptions/:id", subHandler.Destroy)

	router.POST("/feeds", feed.Create)
	router.GET("/feeds", feed.Index)
	router.GET("/feeds/:id", feed.Show)
}

// Start listens on a port and serves data
func Start(port int) {
	log.Printf("Listening on port %v", port)

	portStr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}
