package account

import (
	"fmt"
	"github.com/drprado2/transaction-manager/pkg/models"
)

type CreateAccount struct {
	repository Repository
}

func NewCreateAccountService(repository Repository) *CreateAccount {
	return &CreateAccount{
		repository: repository,
	}
}

func (svc *CreateAccount) Execute(model *models.CreateAccountModel) (int, []models.Error) {
	fullModel := &models.AccountModel{
		CreateAccountModel: *model,
	}
	if errors := svc.validate(fullModel); len(errors) > 0 {
		return 0, errors
	}

	id := svc.repository.Create(AccountOf(fullModel))

	return int(id), nil
}

func (svc *CreateAccount) validate(model *models.AccountModel) []models.Error {
	modelResult := svc.validateAccountModel(model)
	uniqueDocResult := svc.checkDocumentNumberIsUnique(model.DocumentNumber)
	return append(modelResult, uniqueDocResult...)
}

func (svc *CreateAccount) validateAccountModel(model *models.AccountModel) []models.Error {
	errors := make([]models.Error, 0)
	if len(model.DocumentNumber) <= 0 {
		errors = append(errors, models.Error{Message: "Document number is required"})
	}

	return errors
}

func (svc *CreateAccount) checkDocumentNumberIsUnique(documentNumber string) []models.Error {
	if svc.repository.ExistsByDocumentNumber(documentNumber) {
		return []models.Error{
			{Message: fmt.Sprintf("Already exists an account with the document number %v", documentNumber)},
		}
	}
	return nil
}
