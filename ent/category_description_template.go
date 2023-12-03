// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/category_description_template"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

// Category_description_template is the model entity for the Category_description_template schema.
type Category_description_template struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// StartUseAt holds the value of the "start_use_at" field.
	StartUseAt time.Time `json:"start_use_at,omitempty"`
	// LastUseAt holds the value of the "last_use_at" field.
	LastUseAt time.Time `json:"last_use_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Category_description_templateQuery when eager-loading is set.
	Edges        Category_description_templateEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Category_description_templateEdges holds the relations/edges for other nodes in the graph.
type Category_description_templateEdges struct {
	// Descriptions holds the value of the descriptions edge.
	Descriptions []*Description `json:"descriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DescriptionsOrErr returns the Descriptions value or an error if the edge
// was not loaded in eager-loading.
func (e Category_description_templateEdges) DescriptionsOrErr() ([]*Description, error) {
	if e.loadedTypes[0] {
		return e.Descriptions, nil
	}
	return nil, &NotLoadedError{edge: "descriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category_description_template) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category_description_template.FieldID, category_description_template.FieldText:
			values[i] = new(sql.NullString)
		case category_description_template.FieldStartUseAt, category_description_template.FieldLastUseAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category_description_template fields.
func (cdt *Category_description_template) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category_description_template.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				cdt.ID = pulid.ID(value.String)
			}
		case category_description_template.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				cdt.Text = value.String
			}
		case category_description_template.FieldStartUseAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_use_at", values[i])
			} else if value.Valid {
				cdt.StartUseAt = value.Time
			}
		case category_description_template.FieldLastUseAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_use_at", values[i])
			} else if value.Valid {
				cdt.LastUseAt = value.Time
			}
		default:
			cdt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Category_description_template.
// This includes values selected through modifiers, order, etc.
func (cdt *Category_description_template) Value(name string) (ent.Value, error) {
	return cdt.selectValues.Get(name)
}

// QueryDescriptions queries the "descriptions" edge of the Category_description_template entity.
func (cdt *Category_description_template) QueryDescriptions() *DescriptionQuery {
	return NewCategoryDescriptionTemplateClient(cdt.config).QueryDescriptions(cdt)
}

// Update returns a builder for updating this Category_description_template.
// Note that you need to call Category_description_template.Unwrap() before calling this method if this Category_description_template
// was returned from a transaction, and the transaction was committed or rolled back.
func (cdt *Category_description_template) Update() *CategoryDescriptionTemplateUpdateOne {
	return NewCategoryDescriptionTemplateClient(cdt.config).UpdateOne(cdt)
}

// Unwrap unwraps the Category_description_template entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cdt *Category_description_template) Unwrap() *Category_description_template {
	_tx, ok := cdt.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category_description_template is not a transactional entity")
	}
	cdt.config.driver = _tx.drv
	return cdt
}

// String implements the fmt.Stringer.
func (cdt *Category_description_template) String() string {
	var builder strings.Builder
	builder.WriteString("Category_description_template(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cdt.ID))
	builder.WriteString("text=")
	builder.WriteString(cdt.Text)
	builder.WriteString(", ")
	builder.WriteString("start_use_at=")
	builder.WriteString(cdt.StartUseAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_use_at=")
	builder.WriteString(cdt.LastUseAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Category_description_templates is a parsable slice of Category_description_template.
type Category_description_templates []*Category_description_template