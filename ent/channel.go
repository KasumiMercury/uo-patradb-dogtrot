// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/channel"
)

// Channel is the model entity for the Channel schema.
type Channel struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID string `json:"channel_id,omitempty"`
	// Handle holds the value of the "handle" field.
	Handle string `json:"handle,omitempty"`
	// ThumbnailURL holds the value of the "thumbnail_url" field.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChannelQuery when eager-loading is set.
	Edges        ChannelEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChannelEdges holds the relations/edges for other nodes in the graph.
type ChannelEdges struct {
	// Videos holds the value of the videos edge.
	Videos []*Video `json:"videos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VideosOrErr returns the Videos value or an error if the edge
// was not loaded in eager-loading.
func (e ChannelEdges) VideosOrErr() ([]*Video, error) {
	if e.loadedTypes[0] {
		return e.Videos, nil
	}
	return nil, &NotLoadedError{edge: "videos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Channel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case channel.FieldID, channel.FieldDisplayName, channel.FieldChannelID, channel.FieldHandle, channel.FieldThumbnailURL:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Channel fields.
func (c *Channel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channel.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case channel.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case channel.FieldChannelID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				c.ChannelID = value.String
			}
		case channel.FieldHandle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field handle", values[i])
			} else if value.Valid {
				c.Handle = value.String
			}
		case channel.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				c.ThumbnailURL = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Channel.
// This includes values selected through modifiers, order, etc.
func (c *Channel) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryVideos queries the "videos" edge of the Channel entity.
func (c *Channel) QueryVideos() *VideoQuery {
	return NewChannelClient(c.config).QueryVideos(c)
}

// Update returns a builder for updating this Channel.
// Note that you need to call Channel.Unwrap() before calling this method if this Channel
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Channel) Update() *ChannelUpdateOne {
	return NewChannelClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Channel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Channel) Unwrap() *Channel {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Channel is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Channel) String() string {
	var builder strings.Builder
	builder.WriteString("Channel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("display_name=")
	builder.WriteString(c.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(c.ChannelID)
	builder.WriteString(", ")
	builder.WriteString("handle=")
	builder.WriteString(c.Handle)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(c.ThumbnailURL)
	builder.WriteByte(')')
	return builder.String()
}

// Channels is a parsable slice of Channel.
type Channels []*Channel
