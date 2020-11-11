package account

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
)

type CreateAccountService interface {
	Execute(model *models.CreateAccountModel) (int, []models.Error)
}

type ReadRepository interface {
	GetById(id entity.ID) *models.AccountModel
	ExistsByDocumentNumber(documentNumber string) bool
}

type WriteRepository interface {
	Create(account *Account) entity.ID
}

type Repository interface {
	ReadRepository
	WriteRepository
}
