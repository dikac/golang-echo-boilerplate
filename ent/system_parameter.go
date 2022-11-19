// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"myapp/ent/system_parameter"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// System_parameter is the model entity for the System_parameter schema.
type System_parameter struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*System_parameter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case system_parameter.FieldID:
			values[i] = new(sql.NullInt64)
		case system_parameter.FieldKey, system_parameter.FieldValue:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type System_parameter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the System_parameter fields.
func (sp *System_parameter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case system_parameter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case system_parameter.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				sp.Key = value.String
			}
		case system_parameter.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sp.Value = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this System_parameter.
// Note that you need to call System_parameter.Unwrap() before calling this method if this System_parameter
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *System_parameter) Update() *SystemParameterUpdateOne {
	return (&System_parameterClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the System_parameter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *System_parameter) Unwrap() *System_parameter {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: System_parameter is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *System_parameter) String() string {
	var builder strings.Builder
	builder.WriteString("System_parameter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("key=")
	builder.WriteString(sp.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(sp.Value)
	builder.WriteByte(')')
	return builder.String()
}

// System_parameters is a parsable slice of System_parameter.
type System_parameters []*System_parameter

func (sp System_parameters) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}
