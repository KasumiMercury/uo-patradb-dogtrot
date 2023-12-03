// Code generated by ent, DO NOT EDIT.

package category_description_template

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the category_description_template type in the database.
	Label = "category_description_template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldStartUseAt holds the string denoting the start_use_at field in the database.
	FieldStartUseAt = "start_use_at"
	// FieldLastUseAt holds the string denoting the last_use_at field in the database.
	FieldLastUseAt = "last_use_at"
	// EdgeDescriptions holds the string denoting the descriptions edge name in mutations.
	EdgeDescriptions = "descriptions"
	// Table holds the table name of the category_description_template in the database.
	Table = "category_description_templates"
	// DescriptionsTable is the table that holds the descriptions relation/edge.
	DescriptionsTable = "descriptions"
	// DescriptionsInverseTable is the table name for the Description entity.
	// It exists in this package in order to avoid circular dependency with the "description" package.
	DescriptionsInverseTable = "descriptions"
	// DescriptionsColumn is the table column denoting the descriptions relation/edge.
	DescriptionsColumn = "category_template_id"
)

// Columns holds all SQL columns for category_description_template fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldStartUseAt,
	FieldLastUseAt,
}

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

// OrderOption defines the ordering options for the Category_description_template queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByStartUseAt orders the results by the start_use_at field.
func ByStartUseAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartUseAt, opts...).ToFunc()
}

// ByLastUseAt orders the results by the last_use_at field.
func ByLastUseAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUseAt, opts...).ToFunc()
}

// ByDescriptionsCount orders the results by descriptions count.
func ByDescriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDescriptionsStep(), opts...)
	}
}

// ByDescriptions orders the results by descriptions terms.
func ByDescriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDescriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDescriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DescriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DescriptionsTable, DescriptionsColumn),
	)
}