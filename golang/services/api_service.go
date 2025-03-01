package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/eoscanada/eos-go"
	"github.com/gorilla/mux"

	"github.com/exsat-network/exSat-bridge-integration-example/config"
	"github.com/exsat-network/exSat-bridge-integration-example/utils"
)

// ApiService handles HTTP API requests
type ApiService struct {
	router     *mux.Router
	eosService *EosService
	cfg        *config.Config
}

// NewApiService creates a new API service instance
func NewApiService(cfg *config.Config, eosService *EosService) *ApiService {
	service := &ApiService{
		router:     mux.NewRouter(),
		eosService: eosService,
		cfg:        cfg,
	}
	service.setupRoutes()
	return service
}

// setupRoutes configures API routes
func (s *ApiService) setupRoutes() {
	api := s.router.PathPrefix("/api").Subrouter()

	// Health check endpoint
	api.HandleFunc("/health", s.healthHandler).Methods("GET")

	// Apply for BTC deposit address
	api.HandleFunc("/brdgmng/appaddrmap", s.applyAddrMapHandler).Methods("POST")

	// Get BTC deposit address
	api.HandleFunc("/brdgmng/deposit-address/{recipientAddress}", s.getDepositAddressHandler).Methods("GET")

	// Root path handler
	s.router.HandleFunc("/", s.indexHandler).Methods("GET")

	// CORS middleware
	s.router.Use(s.corsMiddleware)
}

// Router returns the configured router
func (s *ApiService) Router() *mux.Router {
	return s.router
}

// corsMiddleware handles CORS
func (s *ApiService) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// healthHandler health check handler function
func (s *ApiService) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":    "ok",
		"timestamp": time.Now(),
	})
}

// indexHandler root path handler function
func (s *ApiService) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"name":    "ExSat Bridge Integration API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"healthCheck":            "/api/health",
			"applyBtcDepositAddress": "/api/brdgmng/appaddrmap",
			"getBtcDepositAddress":   "/api/brdgmng/deposit-address",
		},
	})
}

// ApplyAddrMapRequest request structure
type ApplyAddrMapRequest struct {
	RecipientAddress string `json:"recipient_address"`
	Remark           string `json:"remark"`
}

// applyAddrMapHandler handles requests to apply for a BTC deposit address
func (s *ApiService) applyAddrMapHandler(w http.ResponseWriter, r *http.Request) {
	var req ApplyAddrMapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required parameters
	if req.RecipientAddress == "" {
		http.Error(w, "Missing required parameter: recipient_address", http.StatusBadRequest)
		return
	}
	actionData := eos.NewActionData(struct {
		Actor                eos.AccountName `json:"actor"`
		PermissionID         uint64          `json:"permission_id"`
		RecipientAddress     string          `json:"recipient_address"`
		Remark               string          `json:"remark"`
		AssignDepositAddress string          `json:"assign_deposit_address"`
	}{
		Actor:                eos.AN(s.cfg.EosAccount),
		PermissionID:         s.cfg.BrdgmngPermissionId,
		RecipientAddress:     req.RecipientAddress,
		Remark:               req.Remark,
		AssignDepositAddress: "",
	})
	// Call contract
	result, err := s.eosService.ExecuteAction(s.cfg.BtcBridgeContract, "appaddrmap", actionData)
	if err != nil {
		log.Printf("Contract execution error: %v", err)
		http.Error(w, fmt.Sprintf("Contract execution failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"success":     true,
		"transaction": result,
	})
}

// getDepositAddressHandler handles requests to get a BTC deposit address
func (s *ApiService) getDepositAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipientAddress := vars["recipientAddress"]

	// Calculate key value
	key := utils.ComputeId(recipientAddress)

	// Prepare request parameters
	params := eos.GetTableRowsRequest{
		JSON:       true,
		Code:       s.cfg.BtcBridgeContract,
		Scope:      strconv.FormatUint(s.cfg.BrdgmngPermissionId, 10),
		Table:      "addrmappings",
		Index:      utils.IndexTertiary,
		KeyType:    utils.KeyTypeSha256,
		LowerBound: key,
		UpperBound: key,
		Limit:      1,
	}

	// Query table data
	rows, err := s.eosService.GetTableRows(params)
	if err != nil {
		log.Printf("Error getting table data: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get table data: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if result was found
	if len(rows) > 0 {
		// Extract BTC address
		row := rows[0].(map[string]any)
		btcAddress, ok := row["btc_address"].(string)
		if !ok {
			http.Error(w, "Unable to parse BTC address", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"recipientAddress": recipientAddress,
			"depositAddress":   btcAddress,
		})
	} else {
		http.Error(w, "Deposit address not found", http.StatusNotFound)
	}
}
