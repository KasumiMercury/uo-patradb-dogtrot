// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/patchat"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
)

// PatChat is the model entity for the PatChat schema.
type PatChat struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Magnitude holds the value of the "magnitude" field.
	Magnitude float64 `json:"magnitude,omitempty"`
	// Score holds the value of the "score" field.
	Score float64 `json:"score,omitempty"`
	// IsNegative holds the value of the "is_negative" field.
	IsNegative bool `json:"is_negative,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// FromFreechat holds the value of the "from_freechat" field.
	FromFreechat bool `json:"from_freechat,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatChatQuery when eager-loading is set.
	Edges        PatChatEdges `json:"edges"`
	video_id     *string
	selectValues sql.SelectValues
}

// PatChatEdges holds the relations/edges for other nodes in the graph.
type PatChatEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatChatEdges) VideoOrErr() (*Video, error) {
	if e.loadedTypes[0] {
		if e.Video == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: video.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PatChat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case patchat.FieldIsNegative, patchat.FieldFromFreechat:
			values[i] = new(sql.NullBool)
		case patchat.FieldMagnitude, patchat.FieldScore:
			values[i] = new(sql.NullFloat64)
		case patchat.FieldID, patchat.FieldMessage:
			values[i] = new(sql.NullString)
		case patchat.FieldPublishedAt, patchat.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case patchat.ForeignKeys[0]: // video_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PatChat fields.
func (pc *PatChat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case patchat.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pc.ID = value.String
			}
		case patchat.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				pc.Message = value.String
			}
		case patchat.FieldMagnitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field magnitude", values[i])
			} else if value.Valid {
				pc.Magnitude = value.Float64
			}
		case patchat.FieldScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				pc.Score = value.Float64
			}
		case patchat.FieldIsNegative:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_negative", values[i])
			} else if value.Valid {
				pc.IsNegative = value.Bool
			}
		case patchat.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				pc.PublishedAt = value.Time
			}
		case patchat.FieldFromFreechat:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field from_freechat", values[i])
			} else if value.Valid {
				pc.FromFreechat = value.Bool
			}
		case patchat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case patchat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				pc.video_id = new(string)
				*pc.video_id = value.String
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PatChat.
// This includes values selected through modifiers, order, etc.
func (pc *PatChat) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the PatChat entity.
func (pc *PatChat) QueryVideo() *VideoQuery {
	return NewPatChatClient(pc.config).QueryVideo(pc)
}

// Update returns a builder for updating this PatChat.
// Note that you need to call PatChat.Unwrap() before calling this method if this PatChat
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PatChat) Update() *PatChatUpdateOne {
	return NewPatChatClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PatChat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PatChat) Unwrap() *PatChat {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PatChat is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PatChat) String() string {
	var builder strings.Builder
	builder.WriteString("PatChat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("message=")
	builder.WriteString(pc.Message)
	builder.WriteString(", ")
	builder.WriteString("magnitude=")
	builder.WriteString(fmt.Sprintf("%v", pc.Magnitude))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", pc.Score))
	builder.WriteString(", ")
	builder.WriteString("is_negative=")
	builder.WriteString(fmt.Sprintf("%v", pc.IsNegative))
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(pc.PublishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("from_freechat=")
	builder.WriteString(fmt.Sprintf("%v", pc.FromFreechat))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PatChats is a parsable slice of PatChat.
type PatChats []*PatChat
