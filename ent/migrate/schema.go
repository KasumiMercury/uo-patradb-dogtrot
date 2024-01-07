// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "display_name", Type: field.TypeString, Size: 100},
		{Name: "channel_id", Type: field.TypeString, Unique: true, Size: 26},
		{Name: "handle", Type: field.TypeString, Size: 30},
		{Name: "thumbnail_url", Type: field.TypeString, Size: 140},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
	}
	// DescriptionsColumns holds the columns for the "descriptions" table.
	DescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "source_id", Type: field.TypeString, Unique: true, Size: 12},
		{Name: "raw", Type: field.TypeString, Size: 5000},
		{Name: "variable", Type: field.TypeString, Nullable: true},
		{Name: "template_confidence", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "periodic_id", Type: field.TypeString, Nullable: true, Size: 26},
		{Name: "video_id", Type: field.TypeString, Unique: true, Size: 26},
	}
	// DescriptionsTable holds the schema information for the "descriptions" table.
	DescriptionsTable = &schema.Table{
		Name:       "descriptions",
		Columns:    DescriptionsColumns,
		PrimaryKey: []*schema.Column{DescriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "descriptions_periodic_description_templates_periodic_template",
				Columns:    []*schema.Column{DescriptionsColumns[7]},
				RefColumns: []*schema.Column{PeriodicDescriptionTemplatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "descriptions_videos_descriptions",
				Columns:    []*schema.Column{DescriptionsColumns[8]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DescriptionChangesColumns holds the columns for the "description_changes" table.
	DescriptionChangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "raw", Type: field.TypeString, Size: 5000},
		{Name: "variable", Type: field.TypeString, Nullable: true},
		{Name: "changed_at", Type: field.TypeTime},
		{Name: "description_id", Type: field.TypeString, Size: 26},
	}
	// DescriptionChangesTable holds the schema information for the "description_changes" table.
	DescriptionChangesTable = &schema.Table{
		Name:       "description_changes",
		Columns:    DescriptionChangesColumns,
		PrimaryKey: []*schema.Column{DescriptionChangesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "description_changes_descriptions_description_changes",
				Columns:    []*schema.Column{DescriptionChangesColumns[4]},
				RefColumns: []*schema.Column{DescriptionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PatChatsColumns holds the columns for the "pat_chats" table.
	PatChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "message", Type: field.TypeString, Size: 200},
		{Name: "magnitude", Type: field.TypeFloat64},
		{Name: "score", Type: field.TypeFloat64},
		{Name: "is_negative", Type: field.TypeBool, Default: false},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "from_freechat", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "video_id", Type: field.TypeString, Nullable: true, Size: 26},
	}
	// PatChatsTable holds the schema information for the "pat_chats" table.
	PatChatsTable = &schema.Table{
		Name:       "pat_chats",
		Columns:    PatChatsColumns,
		PrimaryKey: []*schema.Column{PatChatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pat_chats_videos_pat_chats",
				Columns:    []*schema.Column{PatChatsColumns[8]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PeriodicDescriptionTemplatesColumns holds the columns for the "periodic_description_templates" table.
	PeriodicDescriptionTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "text", Type: field.TypeString, Size: 5000},
		{Name: "start_use_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_use_at", Type: field.TypeTime, Nullable: true},
	}
	// PeriodicDescriptionTemplatesTable holds the schema information for the "periodic_description_templates" table.
	PeriodicDescriptionTemplatesTable = &schema.Table{
		Name:       "periodic_description_templates",
		Columns:    PeriodicDescriptionTemplatesColumns,
		PrimaryKey: []*schema.Column{PeriodicDescriptionTemplatesColumns[0]},
	}
	// StreamSchedulesColumns holds the columns for the "stream_schedules" table.
	StreamSchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "scheduled_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 100},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "video_id", Type: field.TypeString, Nullable: true, Size: 26},
	}
	// StreamSchedulesTable holds the schema information for the "stream_schedules" table.
	StreamSchedulesTable = &schema.Table{
		Name:       "stream_schedules",
		Columns:    StreamSchedulesColumns,
		PrimaryKey: []*schema.Column{StreamSchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stream_schedules_videos_videos",
				Columns:    []*schema.Column{StreamSchedulesColumns[5]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "source_id", Type: field.TypeString, Unique: true, Size: 12},
		{Name: "title", Type: field.TypeString, Size: 100},
		{Name: "duration_seconds", Type: field.TypeInt, Nullable: true},
		{Name: "is_collaboration", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeString},
		{Name: "chat_id", Type: field.TypeString, Nullable: true},
		{Name: "has_time_range", Type: field.TypeBool, Default: false},
		{Name: "capture_permission", Type: field.TypeBool, Default: true},
		{Name: "scheduled_at", Type: field.TypeTime, Nullable: true},
		{Name: "actual_start_at", Type: field.TypeTime, Nullable: true},
		{Name: "actual_end_at", Type: field.TypeTime, Nullable: true},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "numbering", Type: field.TypeInt, Default: 1, SchemaType: map[string]string{"mysql": "TINYINT"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
	}
	// VideoDisallowRangesColumns holds the columns for the "video_disallow_ranges" table.
	VideoDisallowRangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "start_seconds", Type: field.TypeInt},
		{Name: "end_seconds", Type: field.TypeInt},
		{Name: "video_id", Type: field.TypeString, Size: 26},
	}
	// VideoDisallowRangesTable holds the schema information for the "video_disallow_ranges" table.
	VideoDisallowRangesTable = &schema.Table{
		Name:       "video_disallow_ranges",
		Columns:    VideoDisallowRangesColumns,
		PrimaryKey: []*schema.Column{VideoDisallowRangesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "video_disallow_ranges_videos_video_disallow_ranges",
				Columns:    []*schema.Column{VideoDisallowRangesColumns[3]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// VideoPlayRangesColumns holds the columns for the "video_play_ranges" table.
	VideoPlayRangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "start_seconds", Type: field.TypeInt, Default: 0},
		{Name: "end_seconds", Type: field.TypeInt, Nullable: true},
		{Name: "video_id", Type: field.TypeString, Size: 26},
	}
	// VideoPlayRangesTable holds the schema information for the "video_play_ranges" table.
	VideoPlayRangesTable = &schema.Table{
		Name:       "video_play_ranges",
		Columns:    VideoPlayRangesColumns,
		PrimaryKey: []*schema.Column{VideoPlayRangesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "video_play_ranges_videos_video_play_ranges",
				Columns:    []*schema.Column{VideoPlayRangesColumns[3]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// VideoTagsColumns holds the columns for the "video_tags" table.
	VideoTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "title", Type: field.TypeString, Size: 50},
		{Name: "normalized_title", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "series_numbering", Type: field.TypeInt, Default: 1, SchemaType: map[string]string{"mysql": "TINYINT"}},
	}
	// VideoTagsTable holds the schema information for the "video_tags" table.
	VideoTagsTable = &schema.Table{
		Name:       "video_tags",
		Columns:    VideoTagsColumns,
		PrimaryKey: []*schema.Column{VideoTagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "videotag_title",
				Unique:  false,
				Columns: []*schema.Column{VideoTagsColumns[1]},
			},
		},
	}
	// VideoTitleChangesColumns holds the columns for the "video_title_changes" table.
	VideoTitleChangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 26},
		{Name: "title", Type: field.TypeString, Size: 100},
		{Name: "changed_at", Type: field.TypeTime},
		{Name: "video_id", Type: field.TypeString, Size: 26},
	}
	// VideoTitleChangesTable holds the schema information for the "video_title_changes" table.
	VideoTitleChangesTable = &schema.Table{
		Name:       "video_title_changes",
		Columns:    VideoTitleChangesColumns,
		PrimaryKey: []*schema.Column{VideoTitleChangesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "video_title_changes_videos_video_title_changes",
				Columns:    []*schema.Column{VideoTitleChangesColumns[3]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// VideoChannelColumns holds the columns for the "video_channel" table.
	VideoChannelColumns = []*schema.Column{
		{Name: "video_id", Type: field.TypeString, Size: 26},
		{Name: "channel_id", Type: field.TypeString, Size: 26},
	}
	// VideoChannelTable holds the schema information for the "video_channel" table.
	VideoChannelTable = &schema.Table{
		Name:       "video_channel",
		Columns:    VideoChannelColumns,
		PrimaryKey: []*schema.Column{VideoChannelColumns[0], VideoChannelColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "video_channel_video_id",
				Columns:    []*schema.Column{VideoChannelColumns[0]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "video_channel_channel_id",
				Columns:    []*schema.Column{VideoChannelColumns[1]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VideoTagVideosColumns holds the columns for the "video_tag_videos" table.
	VideoTagVideosColumns = []*schema.Column{
		{Name: "video_tag_id", Type: field.TypeString, Size: 26},
		{Name: "video_id", Type: field.TypeString, Size: 26},
	}
	// VideoTagVideosTable holds the schema information for the "video_tag_videos" table.
	VideoTagVideosTable = &schema.Table{
		Name:       "video_tag_videos",
		Columns:    VideoTagVideosColumns,
		PrimaryKey: []*schema.Column{VideoTagVideosColumns[0], VideoTagVideosColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "video_tag_videos_video_tag_id",
				Columns:    []*schema.Column{VideoTagVideosColumns[0]},
				RefColumns: []*schema.Column{VideoTagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "video_tag_videos_video_id",
				Columns:    []*schema.Column{VideoTagVideosColumns[1]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		DescriptionsTable,
		DescriptionChangesTable,
		PatChatsTable,
		PeriodicDescriptionTemplatesTable,
		StreamSchedulesTable,
		VideosTable,
		VideoDisallowRangesTable,
		VideoPlayRangesTable,
		VideoTagsTable,
		VideoTitleChangesTable,
		VideoChannelTable,
		VideoTagVideosTable,
	}
)

func init() {
	DescriptionsTable.ForeignKeys[0].RefTable = PeriodicDescriptionTemplatesTable
	DescriptionsTable.ForeignKeys[1].RefTable = VideosTable
	DescriptionChangesTable.ForeignKeys[0].RefTable = DescriptionsTable
	PatChatsTable.ForeignKeys[0].RefTable = VideosTable
	StreamSchedulesTable.ForeignKeys[0].RefTable = VideosTable
	VideoDisallowRangesTable.ForeignKeys[0].RefTable = VideosTable
	VideoPlayRangesTable.ForeignKeys[0].RefTable = VideosTable
	VideoTitleChangesTable.ForeignKeys[0].RefTable = VideosTable
	VideoChannelTable.ForeignKeys[0].RefTable = VideosTable
	VideoChannelTable.ForeignKeys[1].RefTable = ChannelsTable
	VideoTagVideosTable.ForeignKeys[0].RefTable = VideoTagsTable
	VideoTagVideosTable.ForeignKeys[1].RefTable = VideosTable
}
