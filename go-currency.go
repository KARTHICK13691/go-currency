package main

import (
	"fmt"
	"net/http"

	"github.com/KARTICK13691/go-currency/handlers"
	"github.com/KARTICK13691/go-currency/middlewares"
	"github.com/KARTICK13691/go-currency/models"
	"github.com/gorilla/mux"
)

func main() {
	models.InitDB("postgres://user:password@localhost/gdbname?sslmode=disable")
	router := mux.NewRouter()

	// Favicon handler
	router.HandleFunc("/favicon.ico", handlers.Favicon)

	// API Endpoints
	router.HandleFunc("/latest", middlewares.CORS(handlers.Latest))
	router.HandleFunc("/{date}", middlewares.CORS(handlers.Historical))

	fmt.Println("* Running on http://127.0.0.1:3000/")
	http.ListenAndServe(":3000", middlewares.Log(router))
}
