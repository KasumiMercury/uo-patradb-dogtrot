// Code generated by ent, DO NOT EDIT.

package descriptionchange

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the descriptionchange type in the database.
	Label = "description_change"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// FieldVariable holds the string denoting the variable field in the database.
	FieldVariable = "variable"
	// FieldNormalizedVariable holds the string denoting the normalized_variable field in the database.
	FieldNormalizedVariable = "normalized_variable"
	// FieldChangedAt holds the string denoting the changed_at field in the database.
	FieldChangedAt = "changed_at"
	// EdgeDescription holds the string denoting the description edge name in mutations.
	EdgeDescription = "description"
	// Table holds the table name of the descriptionchange in the database.
	Table = "description_changes"
	// DescriptionTable is the table that holds the description relation/edge.
	DescriptionTable = "description_changes"
	// DescriptionInverseTable is the table name for the Description entity.
	// It exists in this package in order to avoid circular dependency with the "description" package.
	DescriptionInverseTable = "descriptions"
	// DescriptionColumn is the table column denoting the description relation/edge.
	DescriptionColumn = "description_id"
)

// Columns holds all SQL columns for descriptionchange fields.
var Columns = []string{
	FieldID,
	FieldRaw,
	FieldVariable,
	FieldNormalizedVariable,
	FieldChangedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "description_changes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"description_id",
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

// OrderOption defines the ordering options for the DescriptionChange queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRaw orders the results by the raw field.
func ByRaw(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRaw, opts...).ToFunc()
}

// ByVariable orders the results by the variable field.
func ByVariable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVariable, opts...).ToFunc()
}

// ByNormalizedVariable orders the results by the normalized_variable field.
func ByNormalizedVariable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalizedVariable, opts...).ToFunc()
}

// ByChangedAt orders the results by the changed_at field.
func ByChangedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChangedAt, opts...).ToFunc()
}

// ByDescriptionField orders the results by description field.
func ByDescriptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDescriptionStep(), sql.OrderByField(field, opts...))
	}
}
func newDescriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DescriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DescriptionTable, DescriptionColumn),
	)
}
