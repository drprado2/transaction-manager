package transactions

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
	"github.com/drprado2/transaction-manager/pkg/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	transactionIdTest = 1
)

type mockCreateTransactionService struct{}

func (*mockCreateTransactionService) Execute(model *models.TransactionModel) (int, []models.Error) {
	if model.ID == transactionIdTest {
		return 0, []models.Error{{Message: "Fail"}}
	}
	return transactionIdTest, nil
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

type mockTransactionRepository struct{}

func (*mockTransactionRepository) Create(transaction *transaction.Transaction) entity.ID {
	return transactionIdTest
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

func (*mockServiceProvider) ResolveTransactionRepository(database storage.QueryDatabase) transaction.Repository {
	return &mockTransactionRepository{}
}

func (*mockServiceProvider) ResolveCreateTransactionService(repository transaction.Repository) transaction.CreateTransactionService {
	return &mockCreateTransactionService{}
}

func createServer() *httptest.Server {
	router := gin.Default()
	transactionRouter := NewTransactionRouter(&mockServiceProvider{}, &mockServiceProvider{})
	transactionGroup := router.Group("/transactions")
	{
		transactionRouter.ConfigureRoutes(transactionGroup)
	}
	return httptest.NewServer(router)
}

func TestCreateTransactionWithSuccess(t *testing.T) {
	testServer := createServer()
	defer testServer.Close()

	transactionToCreate := &models.TransactionModel{}
	json, _ := json.Marshal(transactionToCreate)
	requestReader := bytes.NewReader(json)

	resp, err := http.Post(fmt.Sprintf("%s/transactions", testServer.URL), "application/json", requestReader)

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
	accountToCreate.ID = transactionIdTest
	json, _ := json.Marshal(accountToCreate)
	requestReader := bytes.NewReader(json)

	resp, err := http.Post(fmt.Sprintf("%s/transactions", testServer.URL), "application/json", requestReader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Expected status code 400, got %v", resp.StatusCode)
	}
}
