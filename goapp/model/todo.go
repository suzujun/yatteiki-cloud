package model

type (
	// Todo is model
	Todo struct {
		ID        int    `db:"id" json:"id"`
		Title     string `db:"title" json:"title"`
		Completed bool   `db:"completed" json:"completed"`
		Model
	}
)

// TableName is get name
func (m Todo) Name() string {
	return "todos"
}

// PrimaryKeys is get primary keys for table
func (m Todo) PrimaryKeys() []string {
	return []string{"id"}
}

// ColumnNames is get columns for table
func (m Todo) ColumnNames() []string {
	return []string{"id", "title", "completed", "created_at", "updated_at"}
}
