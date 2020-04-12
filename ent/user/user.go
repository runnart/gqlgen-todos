// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID   = "id" // FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"

	// Table holds the table name of the user in the database.
	Table = "users"
	// TodoTable is the table the holds the todo relation/edge.
	TodoTable = "todos"
	// TodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodoInverseTable = "todos"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "user_todo"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)