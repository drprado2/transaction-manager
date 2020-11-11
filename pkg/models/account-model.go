package models

import (
	"time"
)

type CreateAccountModel struct {
	DocumentNumber string `json:"document_number,omitempty"`
}

type AccountModel struct {
	CreateAccountModel
	ID           int                `json:"id,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty"`
	Transactions []TransactionModel `json:"transactions,omitempty"`
}
