// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/art-Hasan/gqlgen-todos/ent/todo"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Done holds the value of the "done" field.
	Done      bool `json:"done,omitempty"`
	user_todo *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullString{}, // text
		&sql.NullBool{},   // done
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Todo) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_todo
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(values ...interface{}) error {
	if m, n := len(values), len(todo.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		t.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field text", values[2])
	} else if value.Valid {
		t.Text = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field done", values[3])
	} else if value.Valid {
		t.Done = value.Bool
	}
	values = values[4:]
	if len(values) == len(todo.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_todo", value)
		} else if value.Valid {
			t.user_todo = new(int)
			*t.user_todo = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Todo.
// Note that, you need to call Todo.Unwrap() before calling this method, if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return (&TodoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", text=")
	builder.WriteString(t.Text)
	builder.WriteString(", done=")
	builder.WriteString(fmt.Sprintf("%v", t.Done))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo

func (t Todos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}