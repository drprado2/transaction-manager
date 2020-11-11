package entity

import "time"

type ID int

type BaseEntity struct {
	ID        ID
	CreatedAt time.Time
}
