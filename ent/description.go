// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/categorydescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// Description is the model entity for the Description schema.
type Description struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Raw holds the value of the "raw" field.
	Raw string `json:"raw,omitempty"`
	// Variable holds the value of the "variable" field.
	Variable string `json:"variable,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TemplateConfidence holds the value of the "template_confidence" field.
	TemplateConfidence bool `json:"template_confidence,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DescriptionQuery when eager-loading is set.
	Edges        DescriptionEdges `json:"edges"`
	periodic_id  *string
	category_id  *string
	video_id     *string
	selectValues sql.SelectValues
}

// DescriptionEdges holds the relations/edges for other nodes in the graph.
type DescriptionEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// PeriodicTemplate holds the value of the periodic_template edge.
	PeriodicTemplate *PeriodicDescriptionTemplate `json:"periodic_template,omitempty"`
	// CategoryTemplate holds the value of the category_template edge.
	CategoryTemplate *CategoryDescriptionTemplate `json:"category_template,omitempty"`
	// DescriptionChanges holds the value of the description_changes edge.
	DescriptionChanges []*DescriptionChange `json:"description_changes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DescriptionEdges) VideoOrErr() (*Video, error) {
	if e.loadedTypes[0] {
		if e.Video == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: video.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// PeriodicTemplateOrErr returns the PeriodicTemplate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DescriptionEdges) PeriodicTemplateOrErr() (*PeriodicDescriptionTemplate, error) {
	if e.loadedTypes[1] {
		if e.PeriodicTemplate == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: periodicdescriptiontemplate.Label}
		}
		return e.PeriodicTemplate, nil
	}
	return nil, &NotLoadedError{edge: "periodic_template"}
}

// CategoryTemplateOrErr returns the CategoryTemplate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DescriptionEdges) CategoryTemplateOrErr() (*CategoryDescriptionTemplate, error) {
	if e.loadedTypes[2] {
		if e.CategoryTemplate == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: categorydescriptiontemplate.Label}
		}
		return e.CategoryTemplate, nil
	}
	return nil, &NotLoadedError{edge: "category_template"}
}

// DescriptionChangesOrErr returns the DescriptionChanges value or an error if the edge
// was not loaded in eager-loading.
func (e DescriptionEdges) DescriptionChangesOrErr() ([]*DescriptionChange, error) {
	if e.loadedTypes[3] {
		return e.DescriptionChanges, nil
	}
	return nil, &NotLoadedError{edge: "description_changes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Description) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case description.FieldTemplateConfidence:
			values[i] = new(sql.NullBool)
		case description.FieldID, description.FieldRaw, description.FieldVariable:
			values[i] = new(sql.NullString)
		case description.FieldCreatedAt, description.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case description.ForeignKeys[0]: // periodic_id
			values[i] = new(sql.NullString)
		case description.ForeignKeys[1]: // category_id
			values[i] = new(sql.NullString)
		case description.ForeignKeys[2]: // video_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Description fields.
func (d *Description) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case description.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case description.FieldRaw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw", values[i])
			} else if value.Valid {
				d.Raw = value.String
			}
		case description.FieldVariable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field variable", values[i])
			} else if value.Valid {
				d.Variable = value.String
			}
		case description.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case description.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case description.FieldTemplateConfidence:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field template_confidence", values[i])
			} else if value.Valid {
				d.TemplateConfidence = value.Bool
			}
		case description.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field periodic_id", values[i])
			} else if value.Valid {
				d.periodic_id = new(string)
				*d.periodic_id = value.String
			}
		case description.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				d.category_id = new(string)
				*d.category_id = value.String
			}
		case description.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				d.video_id = new(string)
				*d.video_id = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Description.
// This includes values selected through modifiers, order, etc.
func (d *Description) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the Description entity.
func (d *Description) QueryVideo() *VideoQuery {
	return NewDescriptionClient(d.config).QueryVideo(d)
}

// QueryPeriodicTemplate queries the "periodic_template" edge of the Description entity.
func (d *Description) QueryPeriodicTemplate() *PeriodicDescriptionTemplateQuery {
	return NewDescriptionClient(d.config).QueryPeriodicTemplate(d)
}

// QueryCategoryTemplate queries the "category_template" edge of the Description entity.
func (d *Description) QueryCategoryTemplate() *CategoryDescriptionTemplateQuery {
	return NewDescriptionClient(d.config).QueryCategoryTemplate(d)
}

// QueryDescriptionChanges queries the "description_changes" edge of the Description entity.
func (d *Description) QueryDescriptionChanges() *DescriptionChangeQuery {
	return NewDescriptionClient(d.config).QueryDescriptionChanges(d)
}

// Update returns a builder for updating this Description.
// Note that you need to call Description.Unwrap() before calling this method if this Description
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Description) Update() *DescriptionUpdateOne {
	return NewDescriptionClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Description entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Description) Unwrap() *Description {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Description is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Description) String() string {
	var builder strings.Builder
	builder.WriteString("Description(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("raw=")
	builder.WriteString(d.Raw)
	builder.WriteString(", ")
	builder.WriteString("variable=")
	builder.WriteString(d.Variable)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("template_confidence=")
	builder.WriteString(fmt.Sprintf("%v", d.TemplateConfidence))
	builder.WriteByte(')')
	return builder.String()
}

// Descriptions is a parsable slice of Description.
type Descriptions []*Description
