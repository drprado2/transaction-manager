package account

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
	"testing"
)

const (
	accountId                     = 1
	alreadyExistentDocumentNumber = "25336987"
)

type mockRepository struct{}

func (*mockRepository) GetById(id entity.ID) *models.AccountModel {
	return nil
}

func (*mockRepository) ExistsByDocumentNumber(documentNumber string) bool {
	if documentNumber == alreadyExistentDocumentNumber {
		return true
	}
	return false
}

func (*mockRepository) Create(account *Account) entity.ID {
	return accountId
}

func createService() *CreateAccount {
	return NewCreateAccountService(&mockRepository{})
}

func TestExecuteWithInvalidModel(t *testing.T) {
	model := &models.AccountModel{
		DocumentNumber: "",
	}
	service := createService()

	id, errors := service.Execute(model)

	if id != 0 {
		t.Fatalf("Expected id to be 0, got %v", id)
	}

	if len(errors) != 1 {
		t.Fatalf("Expected to contain one error, got %v", len(errors))
	}
}

func TestExecuteWithDocumentAlreadyInUse(t *testing.T) {
	model := &models.AccountModel{
		DocumentNumber: alreadyExistentDocumentNumber,
	}
	service := createService()

	id, errors := service.Execute(model)

	if id != 0 {
		t.Fatalf("Expected id to be 0, got %v", id)
	}

	if len(errors) != 1 {
		t.Fatalf("Expected to contain one error, got %v", len(errors))
	}
}

func TestExecuteWithSuccess(t *testing.T) {
	model := &models.AccountModel{
		DocumentNumber: "7788",
	}
	service := createService()

	id, errors := service.Execute(model)

	if id != accountId {
		t.Fatalf("Expected id to be %v, got %v", accountId, id)
	}

	if len(errors) > 0 {
		t.Fatalf("Expected to contain 0 errors, got %v", len(errors))
	}

}
