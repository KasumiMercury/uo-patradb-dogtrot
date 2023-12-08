// Code generated by ent, DO NOT EDIT.

package videodisallowrange

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the videodisallowrange type in the database.
	Label = "video_disallow_range"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartSeconds holds the string denoting the start_seconds field in the database.
	FieldStartSeconds = "start_seconds"
	// FieldEndSeconds holds the string denoting the end_seconds field in the database.
	FieldEndSeconds = "end_seconds"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the videodisallowrange in the database.
	Table = "video_disallow_ranges"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "video_disallow_ranges"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
)

// Columns holds all SQL columns for videodisallowrange fields.
var Columns = []string{
	FieldID,
	FieldStartSeconds,
	FieldEndSeconds,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "video_disallow_ranges"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"video_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// StartSecondsValidator is a validator for the "start_seconds" field. It is called by the builders before save.
	StartSecondsValidator func(int) error
	// EndSecondsValidator is a validator for the "end_seconds" field. It is called by the builders before save.
	EndSecondsValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the VideoDisallowRange queries.
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

// ByVideoField orders the results by video field.
func ByVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), sql.OrderByField(field, opts...))
	}
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
	)
}
