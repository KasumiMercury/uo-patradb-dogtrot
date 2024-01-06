// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotag"
)

// VideoTag is the model entity for the VideoTag schema.
type VideoTag struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// NormalizedTitle holds the value of the "normalized_title" field.
	NormalizedTitle string `json:"normalized_title,omitempty"`
	// SeriesNumbering holds the value of the "series_numbering" field.
	SeriesNumbering int `json:"series_numbering,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoTagQuery when eager-loading is set.
	Edges        VideoTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VideoTagEdges holds the relations/edges for other nodes in the graph.
type VideoTagEdges struct {
	// Videos holds the value of the videos edge.
	Videos []*Video `json:"videos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideosOrErr returns the Videos value or an error if the edge
// was not loaded in eager-loading.
func (e VideoTagEdges) VideosOrErr() ([]*Video, error) {
	if e.loadedTypes[0] {
		return e.Videos, nil
	}
	return nil, &NotLoadedError{edge: "videos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VideoTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videotag.FieldSeriesNumbering:
			values[i] = new(sql.NullInt64)
		case videotag.FieldID, videotag.FieldTitle, videotag.FieldNormalizedTitle:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoTag fields.
func (vt *VideoTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videotag.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vt.ID = value.String
			}
		case videotag.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				vt.Title = value.String
			}
		case videotag.FieldNormalizedTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field normalized_title", values[i])
			} else if value.Valid {
				vt.NormalizedTitle = value.String
			}
		case videotag.FieldSeriesNumbering:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field series_numbering", values[i])
			} else if value.Valid {
				vt.SeriesNumbering = int(value.Int64)
			}
		default:
			vt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoTag.
// This includes values selected through modifiers, order, etc.
func (vt *VideoTag) Value(name string) (ent.Value, error) {
	return vt.selectValues.Get(name)
}

// QueryVideos queries the "videos" edge of the VideoTag entity.
func (vt *VideoTag) QueryVideos() *VideoQuery {
	return NewVideoTagClient(vt.config).QueryVideos(vt)
}

// Update returns a builder for updating this VideoTag.
// Note that you need to call VideoTag.Unwrap() before calling this method if this VideoTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (vt *VideoTag) Update() *VideoTagUpdateOne {
	return NewVideoTagClient(vt.config).UpdateOne(vt)
}

// Unwrap unwraps the VideoTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vt *VideoTag) Unwrap() *VideoTag {
	_tx, ok := vt.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoTag is not a transactional entity")
	}
	vt.config.driver = _tx.drv
	return vt
}

// String implements the fmt.Stringer.
func (vt *VideoTag) String() string {
	var builder strings.Builder
	builder.WriteString("VideoTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vt.ID))
	builder.WriteString("title=")
	builder.WriteString(vt.Title)
	builder.WriteString(", ")
	builder.WriteString("normalized_title=")
	builder.WriteString(vt.NormalizedTitle)
	builder.WriteString(", ")
	builder.WriteString("series_numbering=")
	builder.WriteString(fmt.Sprintf("%v", vt.SeriesNumbering))
	builder.WriteByte(')')
	return builder.String()
}

// VideoTags is a parsable slice of VideoTag.
type VideoTags []*VideoTag
