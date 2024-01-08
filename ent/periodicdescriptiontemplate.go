// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
)

// PeriodicDescriptionTemplate is the model entity for the PeriodicDescriptionTemplate schema.
type PeriodicDescriptionTemplate struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// StartUseAt holds the value of the "start_use_at" field.
	StartUseAt time.Time `json:"start_use_at,omitempty"`
	// LastUseAt holds the value of the "last_use_at" field.
	LastUseAt time.Time `json:"last_use_at,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash uint64 `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PeriodicDescriptionTemplateQuery when eager-loading is set.
	Edges        PeriodicDescriptionTemplateEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PeriodicDescriptionTemplateEdges holds the relations/edges for other nodes in the graph.
type PeriodicDescriptionTemplateEdges struct {
	// Descriptions holds the value of the descriptions edge.
	Descriptions []*Description `json:"descriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DescriptionsOrErr returns the Descriptions value or an error if the edge
// was not loaded in eager-loading.
func (e PeriodicDescriptionTemplateEdges) DescriptionsOrErr() ([]*Description, error) {
	if e.loadedTypes[0] {
		return e.Descriptions, nil
	}
	return nil, &NotLoadedError{edge: "descriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PeriodicDescriptionTemplate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case periodicdescriptiontemplate.FieldHash:
			values[i] = new(sql.NullInt64)
		case periodicdescriptiontemplate.FieldID, periodicdescriptiontemplate.FieldText:
			values[i] = new(sql.NullString)
		case periodicdescriptiontemplate.FieldStartUseAt, periodicdescriptiontemplate.FieldLastUseAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PeriodicDescriptionTemplate fields.
func (pdt *PeriodicDescriptionTemplate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case periodicdescriptiontemplate.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pdt.ID = value.String
			}
		case periodicdescriptiontemplate.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				pdt.Text = value.String
			}
		case periodicdescriptiontemplate.FieldStartUseAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_use_at", values[i])
			} else if value.Valid {
				pdt.StartUseAt = value.Time
			}
		case periodicdescriptiontemplate.FieldLastUseAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_use_at", values[i])
			} else if value.Valid {
				pdt.LastUseAt = value.Time
			}
		case periodicdescriptiontemplate.FieldHash:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				pdt.Hash = uint64(value.Int64)
			}
		default:
			pdt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PeriodicDescriptionTemplate.
// This includes values selected through modifiers, order, etc.
func (pdt *PeriodicDescriptionTemplate) Value(name string) (ent.Value, error) {
	return pdt.selectValues.Get(name)
}

// QueryDescriptions queries the "descriptions" edge of the PeriodicDescriptionTemplate entity.
func (pdt *PeriodicDescriptionTemplate) QueryDescriptions() *DescriptionQuery {
	return NewPeriodicDescriptionTemplateClient(pdt.config).QueryDescriptions(pdt)
}

// Update returns a builder for updating this PeriodicDescriptionTemplate.
// Note that you need to call PeriodicDescriptionTemplate.Unwrap() before calling this method if this PeriodicDescriptionTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (pdt *PeriodicDescriptionTemplate) Update() *PeriodicDescriptionTemplateUpdateOne {
	return NewPeriodicDescriptionTemplateClient(pdt.config).UpdateOne(pdt)
}

// Unwrap unwraps the PeriodicDescriptionTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pdt *PeriodicDescriptionTemplate) Unwrap() *PeriodicDescriptionTemplate {
	_tx, ok := pdt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PeriodicDescriptionTemplate is not a transactional entity")
	}
	pdt.config.driver = _tx.drv
	return pdt
}

// String implements the fmt.Stringer.
func (pdt *PeriodicDescriptionTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("PeriodicDescriptionTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pdt.ID))
	builder.WriteString("text=")
	builder.WriteString(pdt.Text)
	builder.WriteString(", ")
	builder.WriteString("start_use_at=")
	builder.WriteString(pdt.StartUseAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_use_at=")
	builder.WriteString(pdt.LastUseAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(fmt.Sprintf("%v", pdt.Hash))
	builder.WriteByte(')')
	return builder.String()
}

// PeriodicDescriptionTemplates is a parsable slice of PeriodicDescriptionTemplate.
type PeriodicDescriptionTemplates []*PeriodicDescriptionTemplate
