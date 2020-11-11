package operation_type

import "github.com/drprado2/transaction-manager/pkg/entity"

type OperationType struct {
	entity.BaseEntity
	Description string
	IsDebit     bool
}
