// Code generated by ent, DO NOT EDIT.

package pat_chat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the pat_chat type in the database.
	Label = "pat_chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldMagnitude holds the string denoting the magnitude field in the database.
	FieldMagnitude = "magnitude"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldIsNegative holds the string denoting the is_negative field in the database.
	FieldIsNegative = "is_negative"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the pat_chat in the database.
	Table = "pat_chats"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "pat_chats"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
)

// Columns holds all SQL columns for pat_chat fields.
var Columns = []string{
	FieldID,
	FieldMessage,
	FieldMagnitude,
	FieldScore,
	FieldIsNegative,
	FieldPublishedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pat_chats"
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
	// DefaultIsNegative holds the default value on creation for the "is_negative" field.
	DefaultIsNegative bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the Pat_chat queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByMagnitude orders the results by the magnitude field.
func ByMagnitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMagnitude, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByIsNegative orders the results by the is_negative field.
func ByIsNegative(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsNegative, opts...).ToFunc()
}

// ByPublishedAt orders the results by the published_at field.
func ByPublishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
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
