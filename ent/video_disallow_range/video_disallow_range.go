// Code generated by ent, DO NOT EDIT.

package video_disallow_range

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the video_disallow_range type in the database.
	Label = "video_disallow_range"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartSeconds holds the string denoting the start_seconds field in the database.
	FieldStartSeconds = "start_seconds"
	// FieldEndSeconds holds the string denoting the end_seconds field in the database.
	FieldEndSeconds = "end_seconds"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the video_disallow_range in the database.
	Table = "video_disallow_ranges"
	// VideoTable is the table that holds the video relation/edge. The primary key declared below.
	VideoTable = "video_video_disallow_ranges"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
)

// Columns holds all SQL columns for video_disallow_range fields.
var Columns = []string{
	FieldID,
	FieldStartSeconds,
	FieldEndSeconds,
}

var (
	// VideoPrimaryKey and VideoColumn2 are the table columns denoting the
	// primary key for the video relation (M2M).
	VideoPrimaryKey = []string{"video_id", "video_disallow_range_id"}
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

// OrderOption defines the ordering options for the Video_disallow_range queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStartSeconds orders the results by the start_seconds field.
func ByStartSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartSeconds, opts...).ToFunc()
}

// ByEndSeconds orders the results by the end_seconds field.
func ByEndSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndSeconds, opts...).ToFunc()
}

// ByVideoCount orders the results by video count.
func ByVideoCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoStep(), opts...)
	}
}

// ByVideo orders the results by video terms.
func ByVideo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, VideoTable, VideoPrimaryKey...),
	)
}
