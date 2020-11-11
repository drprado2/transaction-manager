package storage

type DbConnection interface {
	CreateConnection() Database
}
