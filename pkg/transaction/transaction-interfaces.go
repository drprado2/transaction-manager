package transaction

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
)

type CreateTransactionService interface {
	Execute(model *models.CreateTransactionModel) (int, []models.Error)
}

type ReadRepository interface {
}

type WriteRepository interface {
	Create(transaction *Transaction) entity.ID
}

type Repository interface {
	ReadRepository
	WriteRepository
}
