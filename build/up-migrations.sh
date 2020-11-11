export POSTGRESQL_URL="postgres://postgres:admin123@localhost:5432/transaction-manager-db?sslmode=disable"
migrate -database ${POSTGRESQL_URL} -path ../pkg/storage/postgres-db/migrations up