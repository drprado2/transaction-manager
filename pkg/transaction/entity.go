package transaction

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
	"time"
)

type Transaction struct {
	entity.BaseEntity
	OperationTypeID entity.ID
	AccountID       entity.ID
	Amount          float64
	EventDate       time.Time
}

func TransactionOf(model *models.TransactionModel) *Transaction {
	return &Transaction{
		BaseEntity: entity.BaseEntity{
			ID:        entity.ID(model.ID),
			CreatedAt: model.CreatedAt,
		},
		OperationTypeID: entity.ID(model.OperationTypeID),
		AccountID:       entity.ID(model.AccountID),
		Amount:          model.Amount,
		EventDate:       model.EventDate,
	}
}
