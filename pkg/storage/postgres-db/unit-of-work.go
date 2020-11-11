package postgres_db

import (
	"context"
	"database/sql"
	"github.com/drprado2/transaction-manager/pkg/storage"
)

type UnitOfWork struct {
	context     context.Context
	transaction *sql.Tx
	db          storage.Database
}

type transactionalDatabase struct {
	unitOfWork *UnitOfWork
}

func (txDb *transactionalDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	return txDb.unitOfWork.transaction.QueryRowContext(txDb.unitOfWork.context, query, args...)
}

func (txDb *transactionalDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return txDb.unitOfWork.transaction.QueryContext(txDb.unitOfWork.context, query, args...)
}

func NewUnitOfWork(db storage.Database) *UnitOfWork {
	ctx := context.Background()
	return &UnitOfWork{
		db:      db,
		context: ctx,
	}
}

func (uow *UnitOfWork) GetTxDb() storage.QueryDatabase {
	return &transactionalDatabase{
		unitOfWork: uow,
	}
}

func (uow *UnitOfWork) BeginTran() {
	if uow.transaction == nil {
		tx, err := uow.db.BeginTx(uow.context, nil)
		if err != nil {
			panic(err)
		}
		uow.transaction = tx
	}
}

func (uow *UnitOfWork) Commit() {
	if uow.transaction != nil {
		uow.transaction.Commit()
	}
}

func (uow *UnitOfWork) Roolback() {
	if uow.transaction != nil {
		uow.transaction.Rollback()
	}
}

func (uow *UnitOfWork) Close() {
	uow.db.Close()
}
