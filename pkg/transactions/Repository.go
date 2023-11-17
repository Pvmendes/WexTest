package transactions

import (
    "errors"
    "sync"
)

/*
type Repository interface {
	Store(transaction *Transaction) error
	Retrieve(id string) (*Transaction, error)
}
*/

// InMemoryRepository is an in-memory implementation of the Repository interface.
type Repository struct {
    mu           sync.Mutex // To handle concurrent accesses
    transactions map[string]*Transaction
}

// NewInMemoryRepository creates a new instance of InMemoryRepository.
func NewInMemoryRepository() *Repository {
    return &Repository{
        transactions: make(map[string]*Transaction),
    }
}

// Store saves a transaction to the repository.
func (r *Repository) Store(transaction *Transaction) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.transactions[transaction.ID]; exists {
        return errors.New("transaction already exists with the given ID")
    }

    r.transactions[transaction.ID] = transaction
    return nil
}

// Retrieve fetches a transaction by ID.
func (r *Repository) Retrieve(id string) (*Transaction, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    transaction, exists := r.transactions[id]
    if !exists {
        return nil, errors.New("transaction not found")
    }

    return transaction, nil
}