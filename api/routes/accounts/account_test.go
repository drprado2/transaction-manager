package accounts

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/drprado2/transaction-manager/pkg/account"
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
	"github.com/drprado2/transaction-manager/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	accountIdTest      = 1
	documentNumberTest = "test"
)

type mockCreateAccountService struct{}

func (*mockCreateAccountService) Execute(model *models.AccountModel) (int, []models.Error) {
	if model.ID == accountIdTest {
		return 0, []models.Error{{Message: "Fail"}}
	}
	return accountIdTest, nil
}

type mockDb struct{}

func (*mockDb) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (*mockDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (*mockDb) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return nil, nil
}

func (*mockDb) Close() error {
	return nil
}

type mockUnitOfWork struct{}

func (*mockUnitOfWork) GetTxDb() storage.QueryDatabase {
	return nil
}

func (*mockUnitOfWork) BeginTran() {
}

func (*mockUnitOfWork) Commit() {
}

func (*mockUnitOfWork) Roolback() {
}

func (*mockUnitOfWork) Close() {
}

type mockAccountRepository struct{}

func (*mockAccountRepository) GetById(id entity.ID) *models.AccountModel {
	if id == accountIdTest {
		return &models.AccountModel{
			ID:             accountIdTest,
			DocumentNumber: "test",
			CreatedAt:      time.Now(),
			Transactions:   nil,
		}
	}
	return nil
}

func (*mockAccountRepository) ExistsByDocumentNumber(documentNumber string) bool {
	if documentNumber == documentNumberTest {
		return true
	}
	return false
}

func (*mockAccountRepository) Create(account *account.Account) entity.ID {
	return accountIdTest
}

type mockServiceProvider struct{}

func (*mockServiceProvider) ResolveDbConnection() storage.DbConnection {
	return nil
}

func (provider *mockServiceProvider) ResolveDb() storage.Database {
	return &mockDb{}
}

func (provider *mockServiceProvider) ResolveUnitOfWork() storage.UnitOfWork {
	return &mockUnitOfWork{}
}

func (*mockServiceProvider) ResolveAccountRepository(database storage.QueryDatabase) account.Repository {
	return &mockAccountRepository{}
}

func (*mockServiceProvider) ResolveCreateAccountService(repository account.Repository) account.CreateAccountService {
	return &mockCreateAccountService{}
}

func createServer() *httptest.Server {
	router := gin.Default()
	accountRouter := NewAccountRouter(&mockServiceProvider{}, &mockServiceProvider{})
	accountsGroup := router.Group("/accounts")
	{
		accountRouter.ConfigureRoutes(accountsGroup)
	}
	return httptest.NewServer(router)
}

func TestGetByIdWithExistentId(t *testing.T) {
	testServer := createServer()
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/accounts/%v", testServer.URL, accountIdTest))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestGetByIdWithInexistentId(t *testing.T) {
	testServer := createServer()
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/accounts/%v", testServer.URL, 5))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 404 {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}
}

func TestCreateAccountWithSuccess(t *testing.T) {
	testServer := createServer()
	defer testServer.Close()

	accountToCreate := &account.Account{}
	json, _ := json.Marshal(accountToCreate)
	requestReader := bytes.NewReader(json)

	resp, err := http.Post(fmt.Sprintf("%s/accounts", testServer.URL), "application/json", requestReader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}
}

func TestCreateAccountWithErrors(t *testing.T) {
	testServer := createServer()
	defer testServer.Close()

	accountToCreate := &account.Account{}
	accountToCreate.ID = accountIdTest
	json, _ := json.Marshal(accountToCreate)
	requestReader := bytes.NewReader(json)

	resp, err := http.Post(fmt.Sprintf("%s/accounts", testServer.URL), "application/json", requestReader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Expected status code 400, got %v", resp.StatusCode)
	}
}
