package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Nivaha Notify\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hey, %s!\n", ps.ByName("name"))
}

func setupRoutes() {
	router = httprouter.New()

	router.GET("/", index)
	router.GET("/hello/:name", hello)
}

func startRouter() {
	log.Printf("Listening on port %v", clOptions.httpPort)

	port := fmt.Sprintf(":%v", clOptions.httpPort)
	log.Fatal(http.ListenAndServe(port, router))
}
