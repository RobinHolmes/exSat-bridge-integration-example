package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/exsat-network/exSat-bridge-integration-example/config"
	"github.com/exsat-network/exSat-bridge-integration-example/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize EOS service
	eosService, err := services.NewEosService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize EOS service: %v", err)
	}

	// Initialize API service
	apiService := services.NewApiService(cfg, eosService)

	// Start HTTP server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting: http://localhost%s", serverAddr)
	log.Printf("EOS account: %s", cfg.EosAccount)
	log.Fatal(http.ListenAndServe(serverAddr, apiService.Router()))
}
