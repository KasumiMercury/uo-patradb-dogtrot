// Code generated by ent, DO NOT EDIT.

package channel

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the channel type in the database.
	Label = "channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldHandle holds the string denoting the handle field in the database.
	FieldHandle = "handle"
	// FieldThumbnailURL holds the string denoting the thumbnail_url field in the database.
	FieldThumbnailURL = "thumbnail_url"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// Table holds the table name of the channel in the database.
	Table = "channels"
	// VideosTable is the table that holds the videos relation/edge. The primary key declared below.
	VideosTable = "video_channel"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
)

// Columns holds all SQL columns for channel fields.
var Columns = []string{
	FieldID,
	FieldDisplayName,
	FieldChannelID,
	FieldHandle,
	FieldThumbnailURL,
}

var (
	// VideosPrimaryKey and VideosColumn2 are the table columns denoting the
	// primary key for the videos relation (M2M).
	VideosPrimaryKey = []string{"video_id", "channel_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the Channel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByChannelID orders the results by the channel_id field.
func ByChannelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelID, opts...).ToFunc()
}

// ByHandle orders the results by the handle field.
func ByHandle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandle, opts...).ToFunc()
}

// ByThumbnailURL orders the results by the thumbnail_url field.
func ByThumbnailURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailURL, opts...).ToFunc()
}

// ByVideosCount orders the results by videos count.
func ByVideosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideosStep(), opts...)
	}
}

// ByVideos orders the results by videos terms.
func ByVideos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, VideosTable, VideosPrimaryKey...),
	)
}
