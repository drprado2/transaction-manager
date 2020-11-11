package storage

type UnitOfWork interface {
	GetTxDb() QueryDatabase
	BeginTran()
	Commit()
	Roolback()
	Close()
}
