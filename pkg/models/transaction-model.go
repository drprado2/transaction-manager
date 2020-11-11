package models

import (
	"time"
)

type CreateTransactionModel struct {
	OperationTypeID int     `json:"operation_type_id,omitempty"`
	AccountID       int     `json:"account_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}

type TransactionModel struct {
	CreateTransactionModel
	ID                       int           `json:"id,omitempty"`
	OperationTypeDescription string        `json:"operation_type_description,omitempty"`
	Account                  *AccountModel `json:"account,omitempty"`
	EventDate                time.Time     `json:"event_date,omitempty"`
	CreatedAt                time.Time     `json:"created_at,omitempty"`
}
