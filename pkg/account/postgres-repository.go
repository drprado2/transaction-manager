package account

import (
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
	"github.com/drprado2/transaction-manager/pkg/storage"
)

type PostgresRepository struct {
	database storage.QueryDatabase
}

func NewPostgresRepository(db storage.QueryDatabase) *PostgresRepository {
	return &PostgresRepository{
		database: db,
	}
}

func (repo *PostgresRepository) GetById(id entity.ID) *models.AccountModel {
	query := `
SELECT a.id accountId, a.document_number, a.created_at, t.id transactionId, t.account_id, t.amount, t.event_date, t.operation_type_id, t.created_at, o.description
FROM accounts AS a 
LEFT JOIN transactions AS t ON a.id = t.account_id
LEFT JOIN operationTypes AS o ON t.operation_type_id = o.id
WHERE a.id=$1
`
	accountFound := false
	rows, err := repo.database.Query(query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	account := models.AccountModel{}
	transactions := make([]models.TransactionModel, 0, 10)
	for rows.Next() {
		accountFound = true
		transaction := models.TransactionModel{}
		rows.Scan(
			&account.ID,
			&account.DocumentNumber,
			&account.CreatedAt,
			&transaction.ID,
			&transaction.AccountID,
			&transaction.Amount,
			&transaction.EventDate,
			&transaction.OperationTypeID,
			&transaction.CreatedAt,
			&transaction.OperationTypeDescription,
		)
		if transaction.ID != 0 {
			transactions = append(transactions, transaction)
		}
	}

	if !accountFound {
		return nil
	}

	account.Transactions = transactions
	return &account
}

func (repo *PostgresRepository) Create(account *Account) entity.ID {
	var id entity.ID
	sqlInsert := `
INSERT INTO accounts (document_number) VALUES ($1) RETURNING id`
	if err := repo.database.QueryRow(sqlInsert, account.DocumentNumber).Scan(&id); err != nil {
		panic(err)
	}
	return id
}

func (repo *PostgresRepository) ExistsByDocumentNumber(documentNumber string) bool {
	var exists bool
	query := `
select exists(select 1 from accounts where document_number=$1)`
	if err := repo.database.QueryRow(query, documentNumber).Scan(&exists); err != nil {
		panic(err)
	}
	return exists
}
