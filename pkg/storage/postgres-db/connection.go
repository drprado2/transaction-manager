package postgres_db

import (
	"database/sql"
	"fmt"
	"github.com/drprado2/transaction-manager/configs"
	"github.com/drprado2/transaction-manager/pkg/storage"
	_ "github.com/lib/pq"
	"log"
)

var ErrNoMatch = fmt.Errorf("no matching record")

type PostgresConnection struct{}

func (*PostgresConnection) CreateConnection() storage.Database {
	configuration := configs.GetConfig()
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configuration.DatabaseHost, configuration.DatabasePort, configuration.DatabaseUser, configuration.DatabasePassword, configuration.DatabaseName)
	conn, err := sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established successfully")
	return conn
}
