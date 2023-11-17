package transactions

import (
	"my-transaction-app/pkg/currency"

	guuid "github.com/google/uuid"
)

type Service struct {
	repo      *Repository
	converter *currency.Converter
}

func NewService() *Service {
	return &Service{
		repo:      NewInMemoryRepository(),
		converter: currency.NewConverter(),
	}
}

func (s *Service) StoreTransaction(t *Transaction) (string, error) {

	err := t.Validate()
	if err != nil {
		return "", err
	}

	// Generate a unique ID
	id := guuid.New()

	// Store the transaction using the repo
	t.ID = id.String()
	s.repo.Store(t)

	// Return error if any
	return id.String(), nil
}

func (s *Service) RetrieveTransaction(id string, targetCountry string) (*TransactionResponse, error) {

	// Retrieve transaction
	trans, errRetrive := s.repo.Retrieve(id)
	if errRetrive != nil {
		return new(TransactionResponse), errRetrive
	}

	if len(targetCountry) != 0 {
		// Convert the amount using the currency.Converter
		converted, errConvert := s.converter.Convert(trans.Amount, targetCountry, trans.TransactionDate)
		
		if errConvert != nil {
			return new(TransactionResponse), errConvert
		}

		return NewTransactionResponse(trans, converted.ConvertedAmount, targetCountry, converted.ExchangeRate), nil
	}

	// Return the modified transaction or an error
	return NewTransactionResponse(trans, 0.0, targetCountry, 0.0), nil
}
