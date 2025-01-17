// Code generated by ent, DO NOT EDIT.

package videotag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the videotag type in the database.
	Label = "video_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldNormalizedTitle holds the string denoting the normalized_title field in the database.
	FieldNormalizedTitle = "normalized_title"
	// FieldSeriesNumbering holds the string denoting the series_numbering field in the database.
	FieldSeriesNumbering = "series_numbering"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// Table holds the table name of the videotag in the database.
	Table = "video_tags"
	// VideosTable is the table that holds the videos relation/edge. The primary key declared below.
	VideosTable = "video_tag_videos"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
)

// Columns holds all SQL columns for videotag fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldNormalizedTitle,
	FieldSeriesNumbering,
}

var (
	// VideosPrimaryKey and VideosColumn2 are the table columns denoting the
	// primary key for the videos relation (M2M).
	VideosPrimaryKey = []string{"video_tag_id", "video_id"}
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// NormalizedTitleValidator is a validator for the "normalized_title" field. It is called by the builders before save.
	NormalizedTitleValidator func(string) error
	// DefaultSeriesNumbering holds the default value on creation for the "series_numbering" field.
	DefaultSeriesNumbering int
	// SeriesNumberingValidator is a validator for the "series_numbering" field. It is called by the builders before save.
	SeriesNumberingValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the VideoTag queries.
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

// BySeriesNumbering orders the results by the series_numbering field.
func BySeriesNumbering(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeriesNumbering, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.M2M, false, VideosTable, VideosPrimaryKey...),
	)
}
