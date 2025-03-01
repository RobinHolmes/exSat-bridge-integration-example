package config

import (
	"log"
	"os"
	"strconv"

	"github.com/exsat-network/exSat-bridge-integration-example/utils"
	"github.com/joho/godotenv"
)

// Config stores application configuration
type Config struct {
	Port                     string
	EosNodeURL               string
	EosAccount               string
	EosPrivateKey            string
	ResourcePayment          bool
	BtcBridgeContract        string
	BrdgmngPermissionId      uint64
	MultichainBridgeContract string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	_ = godotenv.Load()

	// Initialize default configuration
	config := &Config{
		Port:                     "3000",
		EosNodeURL:               "https://rpc-sg.exsat.network",
		BtcBridgeContract:        utils.ContractBrdgmng,
		MultichainBridgeContract: utils.ContractCbridge,
	}

	// Override configuration from environment variables
	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}

	if url := os.Getenv("EOS_NODE_URL"); url != "" {
		config.EosNodeURL = url
	}

	// Required configurations
	config.EosAccount = os.Getenv("EOS_ACCOUNT")
	config.EosPrivateKey = os.Getenv("EOS_PRIVATE_KEY")

	// Resource payment configuration
	if resourcePayment := os.Getenv("RESOURCE_PAYMENT"); resourcePayment == "true" {
		config.ResourcePayment = true
	}

	// BrdgmngPermissionId configuration
	if permissionId := os.Getenv("BRDGMNG_PERMISSION_ID"); permissionId != "" {
		id, err := strconv.ParseUint(permissionId, 10, 64)
		if err == nil {
			config.BrdgmngPermissionId = id
		}
	}

	// Validate required configurations
	if config.EosAccount == "" || config.EosPrivateKey == "" {
		log.Fatal("Error: EOS_ACCOUNT and EOS_PRIVATE_KEY environment variables are required")
	}

	return config
}
