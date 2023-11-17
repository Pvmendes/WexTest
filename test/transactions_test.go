package test

import (
	"my-transaction-app/pkg/config"
	"my-transaction-app/pkg/transactions"
	"testing"
	"time"
)

func TestStoreTransaction(t *testing.T) {
	// load env variables just once in here so can be use in any other place
	config.InitEnvConfigs()
	// Create a mock repository if your service depends on it
	service := transactions.NewService()

	tests := []struct {
		name          string
		transaction   transactions.Transaction
		expectedError bool
	}{
		{
			name: "Valid Transaction",
			transaction: transactions.Transaction{
				ID:              "1",
				Description:     "Test Transaction",
				TransactionDate: time.Now().String(),
				Amount:          100.00,
			},
			expectedError: false,
		},
		{
			name: "Invalid Transaction - Negative Amount",
			transaction: transactions.Transaction{
				ID:              "2",
				Description:     "Test Transaction with Negative Amount",
				TransactionDate: time.Now().String(),
				Amount:          -50.00,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.StoreTransaction(&tt.transaction)
			if (err != nil) != tt.expectedError {
				t.Errorf("%s: expected error: %v, got: %v", tt.name, tt.expectedError, err)
			}
		})
	}
}

func TestRetrieveTransaction(t *testing.T) {
	// load env variables just once in here so can be use in any other place
	config.InitEnvConfigs()
	// Create a mock repository if your service depends on it
	service := transactions.NewService()

	tests := []struct {
		name          string
		id            string
		targetCountry string
		expectedError bool
	}{
		{
			name:          "Retrieve Existing Transaction",
			id:            "existing_id",
			targetCountry: "Brazil",
			expectedError: false,
		},
		{
			name:          "Retrieve Non-Existing Transaction",
			id:            "non_existing_id",
			targetCountry: "XXXX",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.RetrieveTransaction(tt.id, tt.targetCountry)
			if (err != nil) != tt.expectedError {
				t.Errorf("%s: expected error: %v, got: %v", tt.name, tt.expectedError, err)
			}
		})
	}
}
