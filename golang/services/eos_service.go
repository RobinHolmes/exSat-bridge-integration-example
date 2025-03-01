package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/eoscanada/eos-go"
	"github.com/exsat-network/exSat-bridge-integration-example/config"
	"github.com/exsat-network/exSat-bridge-integration-example/utils"
)

// EosService provides functionality for interacting with the EOS blockchain
type EosService struct {
	api *eos.API
	cfg *config.Config
}

// NewEosService creates a new EOS service instance
func NewEosService(cfg *config.Config) (*EosService, error) {
	// Create EOS API client
	api := eos.New(cfg.EosNodeURL)

	// Configure private key
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(context.Background(), cfg.EosPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to import private key: %w", err)
	}
	api.SetSigner(keyBag)

	return &EosService{
		api: api,
		cfg: cfg,
	}, nil
}

// ExecuteAction executes a contract action
func (s *EosService) ExecuteAction(account string, name string, actionData eos.ActionData) (*eos.PushTransactionFullResp, error) {
	// Prepare authorization
	authorization := []eos.PermissionLevel{
		{Actor: eos.AccountName(s.cfg.EosAccount), Permission: eos.PermissionName("active")},
	}

	// If resource payment is configured, add extra authorization
	if s.cfg.ResourcePayment {
		authorization = append([]eos.PermissionLevel{
			{Actor: eos.AccountName(utils.ContractRes), Permission: eos.PermissionName("bridge")},
		}, authorization...)
	}

	// Prepare action
	action := &eos.Action{
		Account:       eos.AccountName(account),
		Name:          eos.ActionName(name),
		Authorization: authorization,
		ActionData:    actionData,
	}
	response, err := s.api.SignPushActions(context.Background(), action)
	if err != nil {
		return nil, fmt.Errorf("transaction execution failed: %w", err)
	}
	log.Printf("Transaction successfully broadcast, Transaction ID: %s", response.TransactionID)
	return response, nil
}

// GetTableRows retrieves table data
func (s *EosService) GetTableRows(params eos.GetTableRowsRequest) ([]any, error) {
	resp, err := s.api.GetTableRows(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("failed to get table data: %w", err)
	}

	var rows []any
	err = json.Unmarshal(resp.Rows, &rows)
	if err != nil {
		return nil, fmt.Errorf("failed to parse table data: %w", err)
	}

	return rows, nil
}
