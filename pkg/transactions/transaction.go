package transactions

import (
	"errors"
	"math"
	"time"
)

type Transaction struct {
	ID              string  `json:"id"`
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	Amount          float64 `json:"amount"`
}

// Validate fields here
// Description max 50 length, TransactionDate format, Amount positivity
// Return an error if validation fails
func (t *Transaction) Validate() error {
	if len(t.Description) > 50 {
		return errors.New("Description: must not exceed 50 characters")
	}

	_, err := time.Parse("2006-01-02", t.TransactionDate)
	if err != nil {
		return err
	}

	if math.Signbit(t.Amount) {
		return errors.New("Purchase amount: must be a valid positive")
	}

	return nil
}

// TransactionResponse represents the data structure for a transaction response.
type TransactionResponse struct {
	ID               string  `json:"id"`
	Description      string  `json:"description"`
	TransactionDate  string  `json:"transaction_date"`
	OriginalAmount   float64 `json:"original_amount"`
	ConvertedAmount  float64 `json:"converted_amount,omitempty"` // omitempty if no conversion is performed
	OriginalCurrency string  `json:"original_currency"`
	TargetCurrency   string  `json:"target_currency,omitempty"` // omitempty if no conversion is performed
	ExchangeRate     float64 `json:"exchange_rate,omitempty"`   // omitempty if no conversion is performed
}

// NewTransactionResponse creates and returns a new TransactionResponse.
func NewTransactionResponse(transaction *Transaction, convertedAmount float64, targetCurrency string, exchangeRate float64) *TransactionResponse {
	return &TransactionResponse{
		ID:               transaction.ID,
		Description:      transaction.Description,
		TransactionDate:  transaction.TransactionDate,
		OriginalAmount:   transaction.Amount,
		ConvertedAmount:  convertedAmount,
		OriginalCurrency: "USD", // Assuming original currency is always USD
		TargetCurrency:   targetCurrency,
		ExchangeRate:     exchangeRate,
	}
}
