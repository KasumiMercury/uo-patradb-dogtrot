// Code generated by ent, DO NOT EDIT.

package description

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the description type in the database.
	Label = "description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// FieldVariable holds the string denoting the variable field in the database.
	FieldVariable = "variable"
	// FieldTemplateConfidence holds the string denoting the template_confidence field in the database.
	FieldTemplateConfidence = "template_confidence"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// EdgePeriodicTemplate holds the string denoting the periodic_template edge name in mutations.
	EdgePeriodicTemplate = "periodic_template"
	// EdgeDescriptionChanges holds the string denoting the description_changes edge name in mutations.
	EdgeDescriptionChanges = "description_changes"
	// Table holds the table name of the description in the database.
	Table = "descriptions"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "descriptions"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
	// PeriodicTemplateTable is the table that holds the periodic_template relation/edge.
	PeriodicTemplateTable = "descriptions"
	// PeriodicTemplateInverseTable is the table name for the PeriodicDescriptionTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "periodicdescriptiontemplate" package.
	PeriodicTemplateInverseTable = "periodic_description_templates"
	// PeriodicTemplateColumn is the table column denoting the periodic_template relation/edge.
	PeriodicTemplateColumn = "periodic_id"
	// DescriptionChangesTable is the table that holds the description_changes relation/edge.
	DescriptionChangesTable = "description_changes"
	// DescriptionChangesInverseTable is the table name for the DescriptionChange entity.
	// It exists in this package in order to avoid circular dependency with the "descriptionchange" package.
	DescriptionChangesInverseTable = "description_changes"
	// DescriptionChangesColumn is the table column denoting the description_changes relation/edge.
	DescriptionChangesColumn = "description_id"
)

// Columns holds all SQL columns for description fields.
var Columns = []string{
	FieldID,
	FieldSourceID,
	FieldRaw,
	FieldVariable,
	FieldTemplateConfidence,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "descriptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"periodic_id",
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
	// SourceIDValidator is a validator for the "source_id" field. It is called by the builders before save.
	SourceIDValidator func(string) error
	// RawValidator is a validator for the "raw" field. It is called by the builders before save.
	RawValidator func(string) error
	// VariableValidator is a validator for the "variable" field. It is called by the builders before save.
	VariableValidator func(string) error
	// DefaultTemplateConfidence holds the default value on creation for the "template_confidence" field.
	DefaultTemplateConfidence bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Description queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySourceID orders the results by the source_id field.
func BySourceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceID, opts...).ToFunc()
}

// ByRaw orders the results by the raw field.
func ByRaw(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRaw, opts...).ToFunc()
}

// ByVariable orders the results by the variable field.
func ByVariable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVariable, opts...).ToFunc()
}

// ByTemplateConfidence orders the results by the template_confidence field.
func ByTemplateConfidence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateConfidence, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVideoField orders the results by video field.
func ByVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), sql.OrderByField(field, opts...))
	}
}

// ByPeriodicTemplateField orders the results by periodic_template field.
func ByPeriodicTemplateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeriodicTemplateStep(), sql.OrderByField(field, opts...))
	}
}

// ByDescriptionChangesCount orders the results by description_changes count.
func ByDescriptionChangesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDescriptionChangesStep(), opts...)
	}
}

// ByDescriptionChanges orders the results by description_changes terms.
func ByDescriptionChanges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDescriptionChangesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, VideoTable, VideoColumn),
	)
}
func newPeriodicTemplateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PeriodicTemplateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PeriodicTemplateTable, PeriodicTemplateColumn),
	)
}
func newDescriptionChangesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DescriptionChangesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DescriptionChangesTable, DescriptionChangesColumn),
	)
}
