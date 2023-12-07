// Code generated by ent, DO NOT EDIT.

package video

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVideoID holds the string denoting the video_id field in the database.
	FieldVideoID = "video_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldNormalizedTitle holds the string denoting the normalized_title field in the database.
	FieldNormalizedTitle = "normalized_title"
	// FieldDurationSeconds holds the string denoting the duration_seconds field in the database.
	FieldDurationSeconds = "duration_seconds"
	// FieldIsCollaboration holds the string denoting the is_collaboration field in the database.
	FieldIsCollaboration = "is_collaboration"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// FieldHasTimeRange holds the string denoting the has_time_range field in the database.
	FieldHasTimeRange = "has_time_range"
	// FieldScheduledAt holds the string denoting the scheduled_at field in the database.
	FieldScheduledAt = "scheduled_at"
	// FieldActualStartAt holds the string denoting the actual_start_at field in the database.
	FieldActualStartAt = "actual_start_at"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeDescriptions holds the string denoting the descriptions edge name in mutations.
	EdgeDescriptions = "descriptions"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeVideoPlayRanges holds the string denoting the video_play_ranges edge name in mutations.
	EdgeVideoPlayRanges = "video_play_ranges"
	// EdgeVideoDisallowRanges holds the string denoting the video_disallow_ranges edge name in mutations.
	EdgeVideoDisallowRanges = "video_disallow_ranges"
	// EdgeVideoTitleChanges holds the string denoting the video_title_changes edge name in mutations.
	EdgeVideoTitleChanges = "video_title_changes"
	// EdgePatChats holds the string denoting the pat_chats edge name in mutations.
	EdgePatChats = "Pat_chats"
	// Table holds the table name of the video in the database.
	Table = "videos"
	// DescriptionsTable is the table that holds the descriptions relation/edge.
	DescriptionsTable = "descriptions"
	// DescriptionsInverseTable is the table name for the Description entity.
	// It exists in this package in order to avoid circular dependency with the "description" package.
	DescriptionsInverseTable = "descriptions"
	// DescriptionsColumn is the table column denoting the descriptions relation/edge.
	DescriptionsColumn = "video_id"
	// ChannelTable is the table that holds the channel relation/edge. The primary key declared below.
	ChannelTable = "video_channel"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// VideoPlayRangesTable is the table that holds the video_play_ranges relation/edge.
	VideoPlayRangesTable = "video_play_ranges"
	// VideoPlayRangesInverseTable is the table name for the VideoPlayRange entity.
	// It exists in this package in order to avoid circular dependency with the "videoplayrange" package.
	VideoPlayRangesInverseTable = "video_play_ranges"
	// VideoPlayRangesColumn is the table column denoting the video_play_ranges relation/edge.
	VideoPlayRangesColumn = "video_id"
	// VideoDisallowRangesTable is the table that holds the video_disallow_ranges relation/edge.
	VideoDisallowRangesTable = "video_disallow_ranges"
	// VideoDisallowRangesInverseTable is the table name for the VideoDisallowRange entity.
	// It exists in this package in order to avoid circular dependency with the "videodisallowrange" package.
	VideoDisallowRangesInverseTable = "video_disallow_ranges"
	// VideoDisallowRangesColumn is the table column denoting the video_disallow_ranges relation/edge.
	VideoDisallowRangesColumn = "video_id"
	// VideoTitleChangesTable is the table that holds the video_title_changes relation/edge.
	VideoTitleChangesTable = "video_title_changes"
	// VideoTitleChangesInverseTable is the table name for the VideoTitleChange entity.
	// It exists in this package in order to avoid circular dependency with the "videotitlechange" package.
	VideoTitleChangesInverseTable = "video_title_changes"
	// VideoTitleChangesColumn is the table column denoting the video_title_changes relation/edge.
	VideoTitleChangesColumn = "video_id"
	// PatChatsTable is the table that holds the Pat_chats relation/edge.
	PatChatsTable = "pat_chats"
	// PatChatsInverseTable is the table name for the PatChat entity.
	// It exists in this package in order to avoid circular dependency with the "patchat" package.
	PatChatsInverseTable = "pat_chats"
	// PatChatsColumn is the table column denoting the Pat_chats relation/edge.
	PatChatsColumn = "video_id"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldVideoID,
	FieldTitle,
	FieldNormalizedTitle,
	FieldDurationSeconds,
	FieldIsCollaboration,
	FieldStatus,
	FieldChatID,
	FieldHasTimeRange,
	FieldScheduledAt,
	FieldActualStartAt,
	FieldPublishedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// ChannelPrimaryKey and ChannelColumn2 are the table columns denoting the
	// primary key for the channel relation (M2M).
	ChannelPrimaryKey = []string{"video_id", "channel_id"}
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
	// DefaultHasTimeRange holds the default value on creation for the "has_time_range" field.
	DefaultHasTimeRange bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Video queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVideoID orders the results by the video_id field.
func ByVideoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByNormalizedTitle orders the results by the normalized_title field.
func ByNormalizedTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalizedTitle, opts...).ToFunc()
}

// ByDurationSeconds orders the results by the duration_seconds field.
func ByDurationSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDurationSeconds, opts...).ToFunc()
}

// ByIsCollaboration orders the results by the is_collaboration field.
func ByIsCollaboration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCollaboration, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByChatID orders the results by the chat_id field.
func ByChatID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatID, opts...).ToFunc()
}

// ByHasTimeRange orders the results by the has_time_range field.
func ByHasTimeRange(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasTimeRange, opts...).ToFunc()
}

// ByScheduledAt orders the results by the scheduled_at field.
func ByScheduledAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduledAt, opts...).ToFunc()
}

// ByActualStartAt orders the results by the actual_start_at field.
func ByActualStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActualStartAt, opts...).ToFunc()
}

// ByPublishedAt orders the results by the published_at field.
func ByPublishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDescriptionsField orders the results by descriptions field.
func ByDescriptionsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDescriptionsStep(), sql.OrderByField(field, opts...))
	}
}

// ByChannelCount orders the results by channel count.
func ByChannelCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChannelStep(), opts...)
	}
}

// ByChannel orders the results by channel terms.
func ByChannel(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChannelStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVideoPlayRangesCount orders the results by video_play_ranges count.
func ByVideoPlayRangesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoPlayRangesStep(), opts...)
	}
}

// ByVideoPlayRanges orders the results by video_play_ranges terms.
func ByVideoPlayRanges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoPlayRangesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVideoDisallowRangesCount orders the results by video_disallow_ranges count.
func ByVideoDisallowRangesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoDisallowRangesStep(), opts...)
	}
}

// ByVideoDisallowRanges orders the results by video_disallow_ranges terms.
func ByVideoDisallowRanges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoDisallowRangesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVideoTitleChangesCount orders the results by video_title_changes count.
func ByVideoTitleChangesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoTitleChangesStep(), opts...)
	}
}

// ByVideoTitleChanges orders the results by video_title_changes terms.
func ByVideoTitleChanges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoTitleChangesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPatChatsCount orders the results by Pat_chats count.
func ByPatChatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPatChatsStep(), opts...)
	}
}

// ByPatChats orders the results by Pat_chats terms.
func ByPatChats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPatChatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDescriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DescriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, DescriptionsTable, DescriptionsColumn),
	)
}
func newChannelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChannelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ChannelTable, ChannelPrimaryKey...),
	)
}
func newVideoPlayRangesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoPlayRangesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VideoPlayRangesTable, VideoPlayRangesColumn),
	)
}
func newVideoDisallowRangesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoDisallowRangesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VideoDisallowRangesTable, VideoDisallowRangesColumn),
	)
}
func newVideoTitleChangesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoTitleChangesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VideoTitleChangesTable, VideoTitleChangesColumn),
	)
}
func newPatChatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PatChatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PatChatsTable, PatChatsColumn),
	)
}
