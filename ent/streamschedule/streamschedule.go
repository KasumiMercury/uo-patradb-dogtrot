// Code generated by ent, DO NOT EDIT.

package streamschedule

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the streamschedule type in the database.
	Label = "stream_schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScheduledAt holds the string denoting the scheduled_at field in the database.
	FieldScheduledAt = "scheduled_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// Table holds the table name of the streamschedule in the database.
	Table = "stream_schedules"
	// VideosTable is the table that holds the videos relation/edge.
	VideosTable = "stream_schedules"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "video_id"
)

// Columns holds all SQL columns for streamschedule fields.
var Columns = []string{
	FieldID,
	FieldScheduledAt,
	FieldTitle,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "stream_schedules"
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the StreamSchedule queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScheduledAt orders the results by the scheduled_at field.
func ByScheduledAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduledAt, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVideosField orders the results by videos field.
func ByVideosField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideosStep(), sql.OrderByField(field, opts...))
	}
}
func newVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, VideosTable, VideosColumn),
	)
}
