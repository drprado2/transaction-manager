package dependencyInjection

import (
	"github.com/drprado2/transaction-manager/pkg/account"
	"github.com/drprado2/transaction-manager/pkg/storage"
	postgres_db "github.com/drprado2/transaction-manager/pkg/storage/postgres-db"
	"github.com/drprado2/transaction-manager/pkg/transaction"
)

type StorageServiceProvider interface {
	ResolveDbConnection() storage.DbConnection
	ResolveDb() storage.Database
	ResolveUnitOfWork() storage.UnitOfWork
}

type AccountServiceProvider interface {
	ResolveAccountRepository(database storage.QueryDatabase) account.Repository
	ResolveCreateAccountService(repository account.Repository) account.CreateAccountService
}

type TransactionServiceProvider interface {
	ResolveTransactionRepository(database storage.QueryDatabase) transaction.Repository
	ResolveCreateTransactionService(repository transaction.Repository) transaction.CreateTransactionService
}

type ServiceProvider interface {
	StorageServiceProvider
	AccountServiceProvider
	TransactionServiceProvider
}

type AppServiceProvider struct{}

func (*AppServiceProvider) ResolveDbConnection() storage.DbConnection {
	return &postgres_db.PostgresConnection{}
}

func (provider *AppServiceProvider) ResolveDb() storage.Database {
	return provider.ResolveDbConnection().CreateConnection()
}

func (provider *AppServiceProvider) ResolveUnitOfWork() storage.UnitOfWork {
	uow := postgres_db.NewUnitOfWork(provider.ResolveDb())
	return uow
}

func (*AppServiceProvider) ResolveAccountRepository(database storage.QueryDatabase) account.Repository {
	return account.NewPostgresRepository(database)
}

func (*AppServiceProvider) ResolveCreateAccountService(repository account.Repository) account.CreateAccountService {
	return account.NewCreateAccountService(repository)
}

func (*AppServiceProvider) ResolveTransactionRepository(database storage.QueryDatabase) transaction.Repository {
	return transaction.NewPostgresRepository(database)
}

func (*AppServiceProvider) ResolveCreateTransactionService(repository transaction.Repository) transaction.CreateTransactionService {
	return transaction.NewCreateTransactionService(repository)
}
