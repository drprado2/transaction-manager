package transaction

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/storage"
)

type PostgresRepository struct {
	connection storage.QueryDatabase
}

func NewPostgresRepository(connection storage.QueryDatabase) *PostgresRepository {
	return &PostgresRepository{
		connection: connection,
	}
}

func (repo *PostgresRepository) Create(transaction *Transaction) entity.ID {
	var id entity.ID
	sqlInsert := `
INSERT INTO transactions (operation_type_id, account_id, amount, event_date) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := repo.connection.QueryRow(sqlInsert, transaction.OperationTypeID, transaction.AccountID, transaction.Amount, transaction.EventDate).Scan(&id); err != nil {
		panic(err)
	}
	return id
}
