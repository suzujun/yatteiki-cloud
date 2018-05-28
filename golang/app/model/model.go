package model

import (
	"time"

	gorp "gopkg.in/gorp.v1"
)

// Model ...
type Model struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// PreInsert is previous insert func
func (m *Model) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().Round(time.Second)
	m.UpdatedAt = now
	m.CreatedAt = now
	return nil
}

// PreUpdate is previous update func
func (m *Model) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().Round(time.Second)
	m.UpdatedAt = now
	return nil
}
