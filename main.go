// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/ecommerce-app/config"
	"github.com/yourusername/ecommerce-app/router"
)

func main() {
	// Load configuration from YAML file
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database connection
	db, err := config.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize router
	r := router.NewRouter(db)

	// Start HTTP server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
