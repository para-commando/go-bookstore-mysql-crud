// Package main Go Bookstore API
//
// This is a sample REST API for a bookstore application.
//
//	Schemes: http, https
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_key:
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	// Import the generated Swagger docs
	_ "go-bookstore-mysql-crud/docs"

	// "gorm.io/driver/mysql"
	"fmt"
	routespckg "go-bookstore-mysql-crud/pkg/routes"
)

// @title Go Bookstore API
// @version 1.0
// @description A REST API for managing books in a bookstore
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey api_key
// @in header
// @name Authorization
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize the router
	router := mux.NewRouter()

	// Register API routes
	routespckg.RegisterBookstoreRoutes(router)

	// Swagger documentation route
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	http.Handle("/", router)
	fmt.Println("Starting server on :8080...")
	fmt.Println("Swagger documentation available at: http://localhost:8080/swagger/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
