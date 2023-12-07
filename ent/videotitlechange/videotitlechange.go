// Code generated by ent, DO NOT EDIT.

package videotitlechange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the videotitlechange type in the database.
	Label = "video_title_change"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldNormalizedTitle holds the string denoting the normalized_title field in the database.
	FieldNormalizedTitle = "normalized_title"
	// FieldChangedAt holds the string denoting the changed_at field in the database.
	FieldChangedAt = "changed_at"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the videotitlechange in the database.
	Table = "video_title_changes"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "video_title_changes"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
)

// Columns holds all SQL columns for videotitlechange fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldNormalizedTitle,
	FieldChangedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "video_title_changes"
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
	// DefaultChangedAt holds the default value on creation for the "changed_at" field.
	DefaultChangedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the VideoTitleChange queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByNormalizedTitle orders the results by the normalized_title field.
func ByNormalizedTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalizedTitle, opts...).ToFunc()
}

// ByChangedAt orders the results by the changed_at field.
func ByChangedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChangedAt, opts...).ToFunc()
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
