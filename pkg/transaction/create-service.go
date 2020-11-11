package transaction

import (
	"github.com/drprado2/transaction-manager/pkg/models"
	"time"
)

type CreateTransaction struct {
	repository Repository
}

func NewCreateTransactionService(repository Repository) *CreateTransaction {
	return &CreateTransaction{
		repository: repository,
	}
}

func (svc *CreateTransaction) Execute(model *models.CreateTransactionModel) (int, []models.Error) {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fullModel := &models.TransactionModel{
		CreateTransactionModel: *model,
	}
	transaction := TransactionOf(fullModel)
	transaction.EventDate = now
	id := svc.repository.Create(transaction)
	return int(id), nil
}
