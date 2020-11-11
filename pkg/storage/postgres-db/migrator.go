package postgres_db

import (
	"fmt"
	"github.com/drprado2/transaction-manager/configs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func MigrateDb(config *configs.Configuration) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	log.Println("Aplicar migrations na conn string", connString)
	m, err := migrate.New("file://"+config.MigrationsPath, connString)
	if err != nil {
		panic(err)
	}
	m.Steps(4)
}
