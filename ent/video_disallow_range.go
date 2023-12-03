// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video_disallow_range"
)

// Video_disallow_range is the model entity for the Video_disallow_range schema.
type Video_disallow_range struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// StartSeconds holds the value of the "start_seconds" field.
	StartSeconds int `json:"start_seconds,omitempty"`
	// EndSeconds holds the value of the "end_seconds" field.
	EndSeconds int `json:"end_seconds,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Video_disallow_rangeQuery when eager-loading is set.
	Edges                       Video_disallow_rangeEdges `json:"edges"`
	video_video_disallow_ranges *pulid.ID
	selectValues                sql.SelectValues
}

// Video_disallow_rangeEdges holds the relations/edges for other nodes in the graph.
type Video_disallow_rangeEdges struct {
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Video_disallow_rangeEdges) VideoOrErr() (*Video, error) {
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
func (*Video_disallow_range) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case video_disallow_range.FieldStartSeconds, video_disallow_range.FieldEndSeconds:
			values[i] = new(sql.NullInt64)
		case video_disallow_range.FieldID:
			values[i] = new(sql.NullString)
		case video_disallow_range.ForeignKeys[0]: // video_video_disallow_ranges
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video_disallow_range fields.
func (vdr *Video_disallow_range) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video_disallow_range.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vdr.ID = pulid.ID(value.String)
			}
		case video_disallow_range.FieldStartSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_seconds", values[i])
			} else if value.Valid {
				vdr.StartSeconds = int(value.Int64)
			}
		case video_disallow_range.FieldEndSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_seconds", values[i])
			} else if value.Valid {
				vdr.EndSeconds = int(value.Int64)
			}
		case video_disallow_range.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_video_disallow_ranges", values[i])
			} else if value.Valid {
				vdr.video_video_disallow_ranges = new(pulid.ID)
				*vdr.video_video_disallow_ranges = pulid.ID(value.String)
			}
		default:
			vdr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Video_disallow_range.
// This includes values selected through modifiers, order, etc.
func (vdr *Video_disallow_range) Value(name string) (ent.Value, error) {
	return vdr.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the Video_disallow_range entity.
func (vdr *Video_disallow_range) QueryVideo() *VideoQuery {
	return NewVideoDisallowRangeClient(vdr.config).QueryVideo(vdr)
}

// Update returns a builder for updating this Video_disallow_range.
// Note that you need to call Video_disallow_range.Unwrap() before calling this method if this Video_disallow_range
// was returned from a transaction, and the transaction was committed or rolled back.
func (vdr *Video_disallow_range) Update() *VideoDisallowRangeUpdateOne {
	return NewVideoDisallowRangeClient(vdr.config).UpdateOne(vdr)
}

// Unwrap unwraps the Video_disallow_range entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vdr *Video_disallow_range) Unwrap() *Video_disallow_range {
	_tx, ok := vdr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video_disallow_range is not a transactional entity")
	}
	vdr.config.driver = _tx.drv
	return vdr
}

// String implements the fmt.Stringer.
func (vdr *Video_disallow_range) String() string {
	var builder strings.Builder
	builder.WriteString("Video_disallow_range(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vdr.ID))
	builder.WriteString("start_seconds=")
	builder.WriteString(fmt.Sprintf("%v", vdr.StartSeconds))
	builder.WriteString(", ")
	builder.WriteString("end_seconds=")
	builder.WriteString(fmt.Sprintf("%v", vdr.EndSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// Video_disallow_ranges is a parsable slice of Video_disallow_range.
type Video_disallow_ranges []*Video_disallow_range