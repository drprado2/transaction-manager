package account

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
)

type Account struct {
	entity.BaseEntity
	DocumentNumber string
}

func AccountOf(model *models.AccountModel) *Account {
	return &Account{
		BaseEntity: entity.BaseEntity{
			ID:        entity.ID(model.ID),
			CreatedAt: model.CreatedAt,
		},
		DocumentNumber: model.DocumentNumber,
	}
}
