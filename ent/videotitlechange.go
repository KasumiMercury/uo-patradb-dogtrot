// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotitlechange"
)

// VideoTitleChange is the model entity for the VideoTitleChange schema.
type VideoTitleChange struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// ChangedAt holds the value of the "changed_at" field.
	ChangedAt time.Time `json:"changed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoTitleChangeQuery when eager-loading is set.
	Edges        VideoTitleChangeEdges `json:"edges"`
	video_id     *string
	selectValues sql.SelectValues
}

// VideoTitleChangeEdges holds the relations/edges for other nodes in the graph.
type VideoTitleChangeEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoTitleChangeEdges) VideoOrErr() (*Video, error) {
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
func (*VideoTitleChange) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videotitlechange.FieldID, videotitlechange.FieldTitle:
			values[i] = new(sql.NullString)
		case videotitlechange.FieldChangedAt:
			values[i] = new(sql.NullTime)
		case videotitlechange.ForeignKeys[0]: // video_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoTitleChange fields.
func (vtc *VideoTitleChange) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videotitlechange.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vtc.ID = value.String
			}
		case videotitlechange.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				vtc.Title = value.String
			}
		case videotitlechange.FieldChangedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field changed_at", values[i])
			} else if value.Valid {
				vtc.ChangedAt = value.Time
			}
		case videotitlechange.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				vtc.video_id = new(string)
				*vtc.video_id = value.String
			}
		default:
			vtc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoTitleChange.
// This includes values selected through modifiers, order, etc.
func (vtc *VideoTitleChange) Value(name string) (ent.Value, error) {
	return vtc.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the VideoTitleChange entity.
func (vtc *VideoTitleChange) QueryVideo() *VideoQuery {
	return NewVideoTitleChangeClient(vtc.config).QueryVideo(vtc)
}

// Update returns a builder for updating this VideoTitleChange.
// Note that you need to call VideoTitleChange.Unwrap() before calling this method if this VideoTitleChange
// was returned from a transaction, and the transaction was committed or rolled back.
func (vtc *VideoTitleChange) Update() *VideoTitleChangeUpdateOne {
	return NewVideoTitleChangeClient(vtc.config).UpdateOne(vtc)
}

// Unwrap unwraps the VideoTitleChange entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vtc *VideoTitleChange) Unwrap() *VideoTitleChange {
	_tx, ok := vtc.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoTitleChange is not a transactional entity")
	}
	vtc.config.driver = _tx.drv
	return vtc
}

// String implements the fmt.Stringer.
func (vtc *VideoTitleChange) String() string {
	var builder strings.Builder
	builder.WriteString("VideoTitleChange(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vtc.ID))
	builder.WriteString("title=")
	builder.WriteString(vtc.Title)
	builder.WriteString(", ")
	builder.WriteString("changed_at=")
	builder.WriteString(vtc.ChangedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// VideoTitleChanges is a parsable slice of VideoTitleChange.
type VideoTitleChanges []*VideoTitleChange
