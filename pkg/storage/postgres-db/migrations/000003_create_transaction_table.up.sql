CREATE TABLE IF NOT EXISTS transactions(
    id SERIAL PRIMARY KEY,
    operation_type_id int NOT NULL,
    account_id int NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    event_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_operationType
      FOREIGN KEY(operation_type_id)
      REFERENCES operationTypes(id)
      ON DELETE RESTRICT,
    CONSTRAINT fk_account
      FOREIGN KEY(account_id)
      REFERENCES accounts(id)
      ON DELETE RESTRICT
);
