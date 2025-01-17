// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videoplayrange"
)

// VideoPlayRange is the model entity for the VideoPlayRange schema.
type VideoPlayRange struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// StartSeconds holds the value of the "start_seconds" field.
	StartSeconds int `json:"start_seconds,omitempty"`
	// EndSeconds holds the value of the "end_seconds" field.
	EndSeconds int `json:"end_seconds,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoPlayRangeQuery when eager-loading is set.
	Edges        VideoPlayRangeEdges `json:"edges"`
	video_id     *string
	selectValues sql.SelectValues
}

// VideoPlayRangeEdges holds the relations/edges for other nodes in the graph.
type VideoPlayRangeEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoPlayRangeEdges) VideoOrErr() (*Video, error) {
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
func (*VideoPlayRange) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videoplayrange.FieldStartSeconds, videoplayrange.FieldEndSeconds:
			values[i] = new(sql.NullInt64)
		case videoplayrange.FieldID:
			values[i] = new(sql.NullString)
		case videoplayrange.ForeignKeys[0]: // video_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoPlayRange fields.
func (vpr *VideoPlayRange) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videoplayrange.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vpr.ID = value.String
			}
		case videoplayrange.FieldStartSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_seconds", values[i])
			} else if value.Valid {
				vpr.StartSeconds = int(value.Int64)
			}
		case videoplayrange.FieldEndSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_seconds", values[i])
			} else if value.Valid {
				vpr.EndSeconds = int(value.Int64)
			}
		case videoplayrange.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				vpr.video_id = new(string)
				*vpr.video_id = value.String
			}
		default:
			vpr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoPlayRange.
// This includes values selected through modifiers, order, etc.
func (vpr *VideoPlayRange) Value(name string) (ent.Value, error) {
	return vpr.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the VideoPlayRange entity.
func (vpr *VideoPlayRange) QueryVideo() *VideoQuery {
	return NewVideoPlayRangeClient(vpr.config).QueryVideo(vpr)
}

// Update returns a builder for updating this VideoPlayRange.
// Note that you need to call VideoPlayRange.Unwrap() before calling this method if this VideoPlayRange
// was returned from a transaction, and the transaction was committed or rolled back.
func (vpr *VideoPlayRange) Update() *VideoPlayRangeUpdateOne {
	return NewVideoPlayRangeClient(vpr.config).UpdateOne(vpr)
}

// Unwrap unwraps the VideoPlayRange entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vpr *VideoPlayRange) Unwrap() *VideoPlayRange {
	_tx, ok := vpr.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoPlayRange is not a transactional entity")
	}
	vpr.config.driver = _tx.drv
	return vpr
}

// String implements the fmt.Stringer.
func (vpr *VideoPlayRange) String() string {
	var builder strings.Builder
	builder.WriteString("VideoPlayRange(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vpr.ID))
	builder.WriteString("start_seconds=")
	builder.WriteString(fmt.Sprintf("%v", vpr.StartSeconds))
	builder.WriteString(", ")
	builder.WriteString("end_seconds=")
	builder.WriteString(fmt.Sprintf("%v", vpr.EndSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// VideoPlayRanges is a parsable slice of VideoPlayRange.
type VideoPlayRanges []*VideoPlayRange
