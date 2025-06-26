package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "gorm.io/driver/mysql"
	"fmt"
	routespckg "go-bookstore-mysql-crud/pkg/routes"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	routespckg.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
