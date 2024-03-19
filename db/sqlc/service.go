package db

import (
	// "context"
	"database/sql"
	// "fmt"
)

// Service provides all functions to execute db queries and transaction
type Service interface {
	Querier
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// Service provides functions to execute SQL queries and transactions
type SQLService struct {
	*Queries
	db *sql.DB
}

// newService creates a new service
func NewService(db *sql.DB) Service {
	return &SQLService{
		db:      db,
		Queries: New(db),
	}
}
