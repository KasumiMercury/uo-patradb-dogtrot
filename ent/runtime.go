// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/categorydescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/channel"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/description"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/descriptionchange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/patchat"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/periodicdescriptiontemplate"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/schema"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/video"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videodisallowrange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videoplayrange"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/videotitlechange"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categorydescriptiontemplateMixin := schema.CategoryDescriptionTemplate{}.Mixin()
	categorydescriptiontemplateMixinFields0 := categorydescriptiontemplateMixin[0].Fields()
	_ = categorydescriptiontemplateMixinFields0
	categorydescriptiontemplateFields := schema.CategoryDescriptionTemplate{}.Fields()
	_ = categorydescriptiontemplateFields
	// categorydescriptiontemplateDescID is the schema descriptor for id field.
	categorydescriptiontemplateDescID := categorydescriptiontemplateMixinFields0[0].Descriptor()
	// categorydescriptiontemplate.DefaultID holds the default value on creation for the id field.
	categorydescriptiontemplate.DefaultID = categorydescriptiontemplateDescID.Default.(func() string)
	// categorydescriptiontemplate.IDValidator is a validator for the "id" field. It is called by the builders before save.
	categorydescriptiontemplate.IDValidator = categorydescriptiontemplateDescID.Validators[0].(func(string) error)
	channelMixin := schema.Channel{}.Mixin()
	channelMixinFields0 := channelMixin[0].Fields()
	_ = channelMixinFields0
	channelFields := schema.Channel{}.Fields()
	_ = channelFields
	// channelDescDisplayName is the schema descriptor for display_name field.
	channelDescDisplayName := channelFields[0].Descriptor()
	// channel.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	channel.DisplayNameValidator = channelDescDisplayName.Validators[0].(func(string) error)
	// channelDescChannelID is the schema descriptor for channel_id field.
	channelDescChannelID := channelFields[1].Descriptor()
	// channel.ChannelIDValidator is a validator for the "channel_id" field. It is called by the builders before save.
	channel.ChannelIDValidator = channelDescChannelID.Validators[0].(func(string) error)
	// channelDescHandle is the schema descriptor for handle field.
	channelDescHandle := channelFields[2].Descriptor()
	// channel.HandleValidator is a validator for the "handle" field. It is called by the builders before save.
	channel.HandleValidator = channelDescHandle.Validators[0].(func(string) error)
	// channelDescThumbnailURL is the schema descriptor for thumbnail_url field.
	channelDescThumbnailURL := channelFields[3].Descriptor()
	// channel.ThumbnailURLValidator is a validator for the "thumbnail_url" field. It is called by the builders before save.
	channel.ThumbnailURLValidator = func() func(string) error {
		validators := channelDescThumbnailURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(thumbnail_url string) error {
			for _, fn := range fns {
				if err := fn(thumbnail_url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// channelDescID is the schema descriptor for id field.
	channelDescID := channelMixinFields0[0].Descriptor()
	// channel.DefaultID holds the default value on creation for the id field.
	channel.DefaultID = channelDescID.Default.(func() string)
	// channel.IDValidator is a validator for the "id" field. It is called by the builders before save.
	channel.IDValidator = channelDescID.Validators[0].(func(string) error)
	descriptionMixin := schema.Description{}.Mixin()
	descriptionMixinFields0 := descriptionMixin[0].Fields()
	_ = descriptionMixinFields0
	descriptionFields := schema.Description{}.Fields()
	_ = descriptionFields
	// descriptionDescRaw is the schema descriptor for raw field.
	descriptionDescRaw := descriptionFields[0].Descriptor()
	// description.RawValidator is a validator for the "raw" field. It is called by the builders before save.
	description.RawValidator = descriptionDescRaw.Validators[0].(func(string) error)
	// descriptionDescCreatedAt is the schema descriptor for created_at field.
	descriptionDescCreatedAt := descriptionFields[3].Descriptor()
	// description.DefaultCreatedAt holds the default value on creation for the created_at field.
	description.DefaultCreatedAt = descriptionDescCreatedAt.Default.(func() time.Time)
	// descriptionDescUpdatedAt is the schema descriptor for updated_at field.
	descriptionDescUpdatedAt := descriptionFields[4].Descriptor()
	// description.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	description.DefaultUpdatedAt = descriptionDescUpdatedAt.Default.(func() time.Time)
	// description.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	description.UpdateDefaultUpdatedAt = descriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// descriptionDescID is the schema descriptor for id field.
	descriptionDescID := descriptionMixinFields0[0].Descriptor()
	// description.DefaultID holds the default value on creation for the id field.
	description.DefaultID = descriptionDescID.Default.(func() string)
	// description.IDValidator is a validator for the "id" field. It is called by the builders before save.
	description.IDValidator = descriptionDescID.Validators[0].(func(string) error)
	descriptionchangeMixin := schema.DescriptionChange{}.Mixin()
	descriptionchangeMixinFields0 := descriptionchangeMixin[0].Fields()
	_ = descriptionchangeMixinFields0
	descriptionchangeFields := schema.DescriptionChange{}.Fields()
	_ = descriptionchangeFields
	// descriptionchangeDescRaw is the schema descriptor for raw field.
	descriptionchangeDescRaw := descriptionchangeFields[0].Descriptor()
	// descriptionchange.RawValidator is a validator for the "raw" field. It is called by the builders before save.
	descriptionchange.RawValidator = descriptionchangeDescRaw.Validators[0].(func(string) error)
	// descriptionchangeDescChangedAt is the schema descriptor for changed_at field.
	descriptionchangeDescChangedAt := descriptionchangeFields[3].Descriptor()
	// descriptionchange.DefaultChangedAt holds the default value on creation for the changed_at field.
	descriptionchange.DefaultChangedAt = descriptionchangeDescChangedAt.Default.(func() time.Time)
	// descriptionchangeDescID is the schema descriptor for id field.
	descriptionchangeDescID := descriptionchangeMixinFields0[0].Descriptor()
	// descriptionchange.DefaultID holds the default value on creation for the id field.
	descriptionchange.DefaultID = descriptionchangeDescID.Default.(func() string)
	// descriptionchange.IDValidator is a validator for the "id" field. It is called by the builders before save.
	descriptionchange.IDValidator = descriptionchangeDescID.Validators[0].(func(string) error)
	patchatMixin := schema.PatChat{}.Mixin()
	patchatMixinFields0 := patchatMixin[0].Fields()
	_ = patchatMixinFields0
	patchatFields := schema.PatChat{}.Fields()
	_ = patchatFields
	// patchatDescMessage is the schema descriptor for message field.
	patchatDescMessage := patchatFields[0].Descriptor()
	// patchat.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	patchat.MessageValidator = patchatDescMessage.Validators[0].(func(string) error)
	// patchatDescMagnitude is the schema descriptor for magnitude field.
	patchatDescMagnitude := patchatFields[1].Descriptor()
	// patchat.MagnitudeValidator is a validator for the "magnitude" field. It is called by the builders before save.
	patchat.MagnitudeValidator = patchatDescMagnitude.Validators[0].(func(float64) error)
	// patchatDescScore is the schema descriptor for score field.
	patchatDescScore := patchatFields[2].Descriptor()
	// patchat.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	patchat.ScoreValidator = patchatDescScore.Validators[0].(func(float64) error)
	// patchatDescIsNegative is the schema descriptor for is_negative field.
	patchatDescIsNegative := patchatFields[3].Descriptor()
	// patchat.DefaultIsNegative holds the default value on creation for the is_negative field.
	patchat.DefaultIsNegative = patchatDescIsNegative.Default.(bool)
	// patchatDescCreatedAt is the schema descriptor for created_at field.
	patchatDescCreatedAt := patchatFields[5].Descriptor()
	// patchat.DefaultCreatedAt holds the default value on creation for the created_at field.
	patchat.DefaultCreatedAt = patchatDescCreatedAt.Default.(func() time.Time)
	// patchatDescID is the schema descriptor for id field.
	patchatDescID := patchatMixinFields0[0].Descriptor()
	// patchat.DefaultID holds the default value on creation for the id field.
	patchat.DefaultID = patchatDescID.Default.(func() string)
	// patchat.IDValidator is a validator for the "id" field. It is called by the builders before save.
	patchat.IDValidator = patchatDescID.Validators[0].(func(string) error)
	periodicdescriptiontemplateMixin := schema.PeriodicDescriptionTemplate{}.Mixin()
	periodicdescriptiontemplateMixinFields0 := periodicdescriptiontemplateMixin[0].Fields()
	_ = periodicdescriptiontemplateMixinFields0
	periodicdescriptiontemplateFields := schema.PeriodicDescriptionTemplate{}.Fields()
	_ = periodicdescriptiontemplateFields
	// periodicdescriptiontemplateDescText is the schema descriptor for text field.
	periodicdescriptiontemplateDescText := periodicdescriptiontemplateFields[0].Descriptor()
	// periodicdescriptiontemplate.TextValidator is a validator for the "text" field. It is called by the builders before save.
	periodicdescriptiontemplate.TextValidator = periodicdescriptiontemplateDescText.Validators[0].(func(string) error)
	// periodicdescriptiontemplateDescID is the schema descriptor for id field.
	periodicdescriptiontemplateDescID := periodicdescriptiontemplateMixinFields0[0].Descriptor()
	// periodicdescriptiontemplate.DefaultID holds the default value on creation for the id field.
	periodicdescriptiontemplate.DefaultID = periodicdescriptiontemplateDescID.Default.(func() string)
	// periodicdescriptiontemplate.IDValidator is a validator for the "id" field. It is called by the builders before save.
	periodicdescriptiontemplate.IDValidator = periodicdescriptiontemplateDescID.Validators[0].(func(string) error)
	videoMixin := schema.Video{}.Mixin()
	videoMixinFields0 := videoMixin[0].Fields()
	_ = videoMixinFields0
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescSourceID is the schema descriptor for source_id field.
	videoDescSourceID := videoFields[0].Descriptor()
	// video.SourceIDValidator is a validator for the "source_id" field. It is called by the builders before save.
	video.SourceIDValidator = videoDescSourceID.Validators[0].(func(string) error)
	// videoDescTitle is the schema descriptor for title field.
	videoDescTitle := videoFields[1].Descriptor()
	// video.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	video.TitleValidator = videoDescTitle.Validators[0].(func(string) error)
	// videoDescDurationSeconds is the schema descriptor for duration_seconds field.
	videoDescDurationSeconds := videoFields[3].Descriptor()
	// video.DurationSecondsValidator is a validator for the "duration_seconds" field. It is called by the builders before save.
	video.DurationSecondsValidator = func() func(int) error {
		validators := videoDescDurationSeconds.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(duration_seconds int) error {
			for _, fn := range fns {
				if err := fn(duration_seconds); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// videoDescIsCollaboration is the schema descriptor for is_collaboration field.
	videoDescIsCollaboration := videoFields[4].Descriptor()
	// video.DefaultIsCollaboration holds the default value on creation for the is_collaboration field.
	video.DefaultIsCollaboration = videoDescIsCollaboration.Default.(bool)
	// videoDescHasTimeRange is the schema descriptor for has_time_range field.
	videoDescHasTimeRange := videoFields[7].Descriptor()
	// video.DefaultHasTimeRange holds the default value on creation for the has_time_range field.
	video.DefaultHasTimeRange = videoDescHasTimeRange.Default.(bool)
	// videoDescCreatedAt is the schema descriptor for created_at field.
	videoDescCreatedAt := videoFields[12].Descriptor()
	// video.DefaultCreatedAt holds the default value on creation for the created_at field.
	video.DefaultCreatedAt = videoDescCreatedAt.Default.(func() time.Time)
	// videoDescUpdatedAt is the schema descriptor for updated_at field.
	videoDescUpdatedAt := videoFields[13].Descriptor()
	// video.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	video.DefaultUpdatedAt = videoDescUpdatedAt.Default.(func() time.Time)
	// video.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	video.UpdateDefaultUpdatedAt = videoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// videoDescID is the schema descriptor for id field.
	videoDescID := videoMixinFields0[0].Descriptor()
	// video.DefaultID holds the default value on creation for the id field.
	video.DefaultID = videoDescID.Default.(func() string)
	// video.IDValidator is a validator for the "id" field. It is called by the builders before save.
	video.IDValidator = videoDescID.Validators[0].(func(string) error)
	videodisallowrangeMixin := schema.VideoDisallowRange{}.Mixin()
	videodisallowrangeMixinFields0 := videodisallowrangeMixin[0].Fields()
	_ = videodisallowrangeMixinFields0
	videodisallowrangeFields := schema.VideoDisallowRange{}.Fields()
	_ = videodisallowrangeFields
	// videodisallowrangeDescStartSeconds is the schema descriptor for start_seconds field.
	videodisallowrangeDescStartSeconds := videodisallowrangeFields[0].Descriptor()
	// videodisallowrange.StartSecondsValidator is a validator for the "start_seconds" field. It is called by the builders before save.
	videodisallowrange.StartSecondsValidator = videodisallowrangeDescStartSeconds.Validators[0].(func(int) error)
	// videodisallowrangeDescEndSeconds is the schema descriptor for end_seconds field.
	videodisallowrangeDescEndSeconds := videodisallowrangeFields[1].Descriptor()
	// videodisallowrange.EndSecondsValidator is a validator for the "end_seconds" field. It is called by the builders before save.
	videodisallowrange.EndSecondsValidator = videodisallowrangeDescEndSeconds.Validators[0].(func(int) error)
	// videodisallowrangeDescID is the schema descriptor for id field.
	videodisallowrangeDescID := videodisallowrangeMixinFields0[0].Descriptor()
	// videodisallowrange.DefaultID holds the default value on creation for the id field.
	videodisallowrange.DefaultID = videodisallowrangeDescID.Default.(func() string)
	// videodisallowrange.IDValidator is a validator for the "id" field. It is called by the builders before save.
	videodisallowrange.IDValidator = videodisallowrangeDescID.Validators[0].(func(string) error)
	videoplayrangeMixin := schema.VideoPlayRange{}.Mixin()
	videoplayrangeMixinFields0 := videoplayrangeMixin[0].Fields()
	_ = videoplayrangeMixinFields0
	videoplayrangeFields := schema.VideoPlayRange{}.Fields()
	_ = videoplayrangeFields
	// videoplayrangeDescStartSeconds is the schema descriptor for start_seconds field.
	videoplayrangeDescStartSeconds := videoplayrangeFields[0].Descriptor()
	// videoplayrange.DefaultStartSeconds holds the default value on creation for the start_seconds field.
	videoplayrange.DefaultStartSeconds = videoplayrangeDescStartSeconds.Default.(int)
	// videoplayrange.StartSecondsValidator is a validator for the "start_seconds" field. It is called by the builders before save.
	videoplayrange.StartSecondsValidator = videoplayrangeDescStartSeconds.Validators[0].(func(int) error)
	// videoplayrangeDescEndSeconds is the schema descriptor for end_seconds field.
	videoplayrangeDescEndSeconds := videoplayrangeFields[1].Descriptor()
	// videoplayrange.EndSecondsValidator is a validator for the "end_seconds" field. It is called by the builders before save.
	videoplayrange.EndSecondsValidator = videoplayrangeDescEndSeconds.Validators[0].(func(int) error)
	// videoplayrangeDescID is the schema descriptor for id field.
	videoplayrangeDescID := videoplayrangeMixinFields0[0].Descriptor()
	// videoplayrange.DefaultID holds the default value on creation for the id field.
	videoplayrange.DefaultID = videoplayrangeDescID.Default.(func() string)
	// videoplayrange.IDValidator is a validator for the "id" field. It is called by the builders before save.
	videoplayrange.IDValidator = videoplayrangeDescID.Validators[0].(func(string) error)
	videotitlechangeMixin := schema.VideoTitleChange{}.Mixin()
	videotitlechangeMixinFields0 := videotitlechangeMixin[0].Fields()
	_ = videotitlechangeMixinFields0
	videotitlechangeFields := schema.VideoTitleChange{}.Fields()
	_ = videotitlechangeFields
	// videotitlechangeDescTitle is the schema descriptor for title field.
	videotitlechangeDescTitle := videotitlechangeFields[0].Descriptor()
	// videotitlechange.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	videotitlechange.TitleValidator = videotitlechangeDescTitle.Validators[0].(func(string) error)
	// videotitlechangeDescChangedAt is the schema descriptor for changed_at field.
	videotitlechangeDescChangedAt := videotitlechangeFields[2].Descriptor()
	// videotitlechange.DefaultChangedAt holds the default value on creation for the changed_at field.
	videotitlechange.DefaultChangedAt = videotitlechangeDescChangedAt.Default.(func() time.Time)
	// videotitlechangeDescID is the schema descriptor for id field.
	videotitlechangeDescID := videotitlechangeMixinFields0[0].Descriptor()
	// videotitlechange.DefaultID holds the default value on creation for the id field.
	videotitlechange.DefaultID = videotitlechangeDescID.Default.(func() string)
	// videotitlechange.IDValidator is a validator for the "id" field. It is called by the builders before save.
	videotitlechange.IDValidator = videotitlechangeDescID.Validators[0].(func(string) error)
}
