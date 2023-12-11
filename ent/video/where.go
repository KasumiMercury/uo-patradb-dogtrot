// Code generated by ent, DO NOT EDIT.

package video

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldID, id))
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldSourceID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldTitle, v))
}

// NormalizedTitle applies equality check predicate on the "normalized_title" field. It's identical to NormalizedTitleEQ.
func NormalizedTitle(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldNormalizedTitle, v))
}

// DurationSeconds applies equality check predicate on the "duration_seconds" field. It's identical to DurationSecondsEQ.
func DurationSeconds(v int) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldDurationSeconds, v))
}

// IsCollaboration applies equality check predicate on the "is_collaboration" field. It's identical to IsCollaborationEQ.
func IsCollaboration(v bool) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldIsCollaboration, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldStatus, v))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldChatID, v))
}

// HasTimeRange applies equality check predicate on the "has_time_range" field. It's identical to HasTimeRangeEQ.
func HasTimeRange(v bool) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldHasTimeRange, v))
}

// ScheduledAt applies equality check predicate on the "scheduled_at" field. It's identical to ScheduledAtEQ.
func ScheduledAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldScheduledAt, v))
}

// ActualStartAt applies equality check predicate on the "actual_start_at" field. It's identical to ActualStartAtEQ.
func ActualStartAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldActualStartAt, v))
}

// ActualEndAt applies equality check predicate on the "actual_end_at" field. It's identical to ActualEndAtEQ.
func ActualEndAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldActualEndAt, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldPublishedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldUpdatedAt, v))
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldSourceID, v))
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldSourceID, v))
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldSourceID, vs...))
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldSourceID, vs...))
}

// SourceIDGT applies the GT predicate on the "source_id" field.
func SourceIDGT(v string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldSourceID, v))
}

// SourceIDGTE applies the GTE predicate on the "source_id" field.
func SourceIDGTE(v string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldSourceID, v))
}

// SourceIDLT applies the LT predicate on the "source_id" field.
func SourceIDLT(v string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldSourceID, v))
}

// SourceIDLTE applies the LTE predicate on the "source_id" field.
func SourceIDLTE(v string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldSourceID, v))
}

// SourceIDContains applies the Contains predicate on the "source_id" field.
func SourceIDContains(v string) predicate.Video {
	return predicate.Video(sql.FieldContains(FieldSourceID, v))
}

// SourceIDHasPrefix applies the HasPrefix predicate on the "source_id" field.
func SourceIDHasPrefix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasPrefix(FieldSourceID, v))
}

// SourceIDHasSuffix applies the HasSuffix predicate on the "source_id" field.
func SourceIDHasSuffix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasSuffix(FieldSourceID, v))
}

// SourceIDEqualFold applies the EqualFold predicate on the "source_id" field.
func SourceIDEqualFold(v string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldSourceID, v))
}

// SourceIDContainsFold applies the ContainsFold predicate on the "source_id" field.
func SourceIDContainsFold(v string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldSourceID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Video {
	return predicate.Video(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldTitle, v))
}

// NormalizedTitleEQ applies the EQ predicate on the "normalized_title" field.
func NormalizedTitleEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldNormalizedTitle, v))
}

// NormalizedTitleNEQ applies the NEQ predicate on the "normalized_title" field.
func NormalizedTitleNEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldNormalizedTitle, v))
}

// NormalizedTitleIn applies the In predicate on the "normalized_title" field.
func NormalizedTitleIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldNormalizedTitle, vs...))
}

// NormalizedTitleNotIn applies the NotIn predicate on the "normalized_title" field.
func NormalizedTitleNotIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldNormalizedTitle, vs...))
}

// NormalizedTitleGT applies the GT predicate on the "normalized_title" field.
func NormalizedTitleGT(v string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldNormalizedTitle, v))
}

// NormalizedTitleGTE applies the GTE predicate on the "normalized_title" field.
func NormalizedTitleGTE(v string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldNormalizedTitle, v))
}

// NormalizedTitleLT applies the LT predicate on the "normalized_title" field.
func NormalizedTitleLT(v string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldNormalizedTitle, v))
}

// NormalizedTitleLTE applies the LTE predicate on the "normalized_title" field.
func NormalizedTitleLTE(v string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldNormalizedTitle, v))
}

// NormalizedTitleContains applies the Contains predicate on the "normalized_title" field.
func NormalizedTitleContains(v string) predicate.Video {
	return predicate.Video(sql.FieldContains(FieldNormalizedTitle, v))
}

// NormalizedTitleHasPrefix applies the HasPrefix predicate on the "normalized_title" field.
func NormalizedTitleHasPrefix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasPrefix(FieldNormalizedTitle, v))
}

// NormalizedTitleHasSuffix applies the HasSuffix predicate on the "normalized_title" field.
func NormalizedTitleHasSuffix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasSuffix(FieldNormalizedTitle, v))
}

// NormalizedTitleIsNil applies the IsNil predicate on the "normalized_title" field.
func NormalizedTitleIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldNormalizedTitle))
}

// NormalizedTitleNotNil applies the NotNil predicate on the "normalized_title" field.
func NormalizedTitleNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldNormalizedTitle))
}

// NormalizedTitleEqualFold applies the EqualFold predicate on the "normalized_title" field.
func NormalizedTitleEqualFold(v string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldNormalizedTitle, v))
}

// NormalizedTitleContainsFold applies the ContainsFold predicate on the "normalized_title" field.
func NormalizedTitleContainsFold(v string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldNormalizedTitle, v))
}

// DurationSecondsEQ applies the EQ predicate on the "duration_seconds" field.
func DurationSecondsEQ(v int) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldDurationSeconds, v))
}

// DurationSecondsNEQ applies the NEQ predicate on the "duration_seconds" field.
func DurationSecondsNEQ(v int) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldDurationSeconds, v))
}

// DurationSecondsIn applies the In predicate on the "duration_seconds" field.
func DurationSecondsIn(vs ...int) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldDurationSeconds, vs...))
}

// DurationSecondsNotIn applies the NotIn predicate on the "duration_seconds" field.
func DurationSecondsNotIn(vs ...int) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldDurationSeconds, vs...))
}

// DurationSecondsGT applies the GT predicate on the "duration_seconds" field.
func DurationSecondsGT(v int) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldDurationSeconds, v))
}

// DurationSecondsGTE applies the GTE predicate on the "duration_seconds" field.
func DurationSecondsGTE(v int) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldDurationSeconds, v))
}

// DurationSecondsLT applies the LT predicate on the "duration_seconds" field.
func DurationSecondsLT(v int) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldDurationSeconds, v))
}

// DurationSecondsLTE applies the LTE predicate on the "duration_seconds" field.
func DurationSecondsLTE(v int) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldDurationSeconds, v))
}

// DurationSecondsIsNil applies the IsNil predicate on the "duration_seconds" field.
func DurationSecondsIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldDurationSeconds))
}

// DurationSecondsNotNil applies the NotNil predicate on the "duration_seconds" field.
func DurationSecondsNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldDurationSeconds))
}

// IsCollaborationEQ applies the EQ predicate on the "is_collaboration" field.
func IsCollaborationEQ(v bool) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldIsCollaboration, v))
}

// IsCollaborationNEQ applies the NEQ predicate on the "is_collaboration" field.
func IsCollaborationNEQ(v bool) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldIsCollaboration, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Video {
	return predicate.Video(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldStatus, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v string) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...string) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v string) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v string) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v string) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v string) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldChatID, v))
}

// ChatIDContains applies the Contains predicate on the "chat_id" field.
func ChatIDContains(v string) predicate.Video {
	return predicate.Video(sql.FieldContains(FieldChatID, v))
}

// ChatIDHasPrefix applies the HasPrefix predicate on the "chat_id" field.
func ChatIDHasPrefix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasPrefix(FieldChatID, v))
}

// ChatIDHasSuffix applies the HasSuffix predicate on the "chat_id" field.
func ChatIDHasSuffix(v string) predicate.Video {
	return predicate.Video(sql.FieldHasSuffix(FieldChatID, v))
}

// ChatIDIsNil applies the IsNil predicate on the "chat_id" field.
func ChatIDIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldChatID))
}

// ChatIDNotNil applies the NotNil predicate on the "chat_id" field.
func ChatIDNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldChatID))
}

// ChatIDEqualFold applies the EqualFold predicate on the "chat_id" field.
func ChatIDEqualFold(v string) predicate.Video {
	return predicate.Video(sql.FieldEqualFold(FieldChatID, v))
}

// ChatIDContainsFold applies the ContainsFold predicate on the "chat_id" field.
func ChatIDContainsFold(v string) predicate.Video {
	return predicate.Video(sql.FieldContainsFold(FieldChatID, v))
}

// HasTimeRangeEQ applies the EQ predicate on the "has_time_range" field.
func HasTimeRangeEQ(v bool) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldHasTimeRange, v))
}

// HasTimeRangeNEQ applies the NEQ predicate on the "has_time_range" field.
func HasTimeRangeNEQ(v bool) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldHasTimeRange, v))
}

// ScheduledAtEQ applies the EQ predicate on the "scheduled_at" field.
func ScheduledAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldScheduledAt, v))
}

// ScheduledAtNEQ applies the NEQ predicate on the "scheduled_at" field.
func ScheduledAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldScheduledAt, v))
}

// ScheduledAtIn applies the In predicate on the "scheduled_at" field.
func ScheduledAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldScheduledAt, vs...))
}

// ScheduledAtNotIn applies the NotIn predicate on the "scheduled_at" field.
func ScheduledAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldScheduledAt, vs...))
}

// ScheduledAtGT applies the GT predicate on the "scheduled_at" field.
func ScheduledAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldScheduledAt, v))
}

// ScheduledAtGTE applies the GTE predicate on the "scheduled_at" field.
func ScheduledAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldScheduledAt, v))
}

// ScheduledAtLT applies the LT predicate on the "scheduled_at" field.
func ScheduledAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldScheduledAt, v))
}

// ScheduledAtLTE applies the LTE predicate on the "scheduled_at" field.
func ScheduledAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldScheduledAt, v))
}

// ScheduledAtIsNil applies the IsNil predicate on the "scheduled_at" field.
func ScheduledAtIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldScheduledAt))
}

// ScheduledAtNotNil applies the NotNil predicate on the "scheduled_at" field.
func ScheduledAtNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldScheduledAt))
}

// ActualStartAtEQ applies the EQ predicate on the "actual_start_at" field.
func ActualStartAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldActualStartAt, v))
}

// ActualStartAtNEQ applies the NEQ predicate on the "actual_start_at" field.
func ActualStartAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldActualStartAt, v))
}

// ActualStartAtIn applies the In predicate on the "actual_start_at" field.
func ActualStartAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldActualStartAt, vs...))
}

// ActualStartAtNotIn applies the NotIn predicate on the "actual_start_at" field.
func ActualStartAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldActualStartAt, vs...))
}

// ActualStartAtGT applies the GT predicate on the "actual_start_at" field.
func ActualStartAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldActualStartAt, v))
}

// ActualStartAtGTE applies the GTE predicate on the "actual_start_at" field.
func ActualStartAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldActualStartAt, v))
}

// ActualStartAtLT applies the LT predicate on the "actual_start_at" field.
func ActualStartAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldActualStartAt, v))
}

// ActualStartAtLTE applies the LTE predicate on the "actual_start_at" field.
func ActualStartAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldActualStartAt, v))
}

// ActualStartAtIsNil applies the IsNil predicate on the "actual_start_at" field.
func ActualStartAtIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldActualStartAt))
}

// ActualStartAtNotNil applies the NotNil predicate on the "actual_start_at" field.
func ActualStartAtNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldActualStartAt))
}

// ActualEndAtEQ applies the EQ predicate on the "actual_end_at" field.
func ActualEndAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldActualEndAt, v))
}

// ActualEndAtNEQ applies the NEQ predicate on the "actual_end_at" field.
func ActualEndAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldActualEndAt, v))
}

// ActualEndAtIn applies the In predicate on the "actual_end_at" field.
func ActualEndAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldActualEndAt, vs...))
}

// ActualEndAtNotIn applies the NotIn predicate on the "actual_end_at" field.
func ActualEndAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldActualEndAt, vs...))
}

// ActualEndAtGT applies the GT predicate on the "actual_end_at" field.
func ActualEndAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldActualEndAt, v))
}

// ActualEndAtGTE applies the GTE predicate on the "actual_end_at" field.
func ActualEndAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldActualEndAt, v))
}

// ActualEndAtLT applies the LT predicate on the "actual_end_at" field.
func ActualEndAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldActualEndAt, v))
}

// ActualEndAtLTE applies the LTE predicate on the "actual_end_at" field.
func ActualEndAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldActualEndAt, v))
}

// ActualEndAtIsNil applies the IsNil predicate on the "actual_end_at" field.
func ActualEndAtIsNil() predicate.Video {
	return predicate.Video(sql.FieldIsNull(FieldActualEndAt))
}

// ActualEndAtNotNil applies the NotNil predicate on the "actual_end_at" field.
func ActualEndAtNotNil() predicate.Video {
	return predicate.Video(sql.FieldNotNull(FieldActualEndAt))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldPublishedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Video {
	return predicate.Video(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Video {
	return predicate.Video(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasDescriptions applies the HasEdge predicate on the "descriptions" edge.
func HasDescriptions() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DescriptionsTable, DescriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDescriptionsWith applies the HasEdge predicate on the "descriptions" edge with a given conditions (other predicates).
func HasDescriptionsWith(preds ...predicate.Description) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newDescriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChannelTable, ChannelPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newChannelStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVideoPlayRanges applies the HasEdge predicate on the "video_play_ranges" edge.
func HasVideoPlayRanges() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideoPlayRangesTable, VideoPlayRangesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoPlayRangesWith applies the HasEdge predicate on the "video_play_ranges" edge with a given conditions (other predicates).
func HasVideoPlayRangesWith(preds ...predicate.VideoPlayRange) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newVideoPlayRangesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVideoDisallowRanges applies the HasEdge predicate on the "video_disallow_ranges" edge.
func HasVideoDisallowRanges() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideoDisallowRangesTable, VideoDisallowRangesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoDisallowRangesWith applies the HasEdge predicate on the "video_disallow_ranges" edge with a given conditions (other predicates).
func HasVideoDisallowRangesWith(preds ...predicate.VideoDisallowRange) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newVideoDisallowRangesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVideoTitleChanges applies the HasEdge predicate on the "video_title_changes" edge.
func HasVideoTitleChanges() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideoTitleChangesTable, VideoTitleChangesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoTitleChangesWith applies the HasEdge predicate on the "video_title_changes" edge with a given conditions (other predicates).
func HasVideoTitleChangesWith(preds ...predicate.VideoTitleChange) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newVideoTitleChangesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatChats applies the HasEdge predicate on the "Pat_chats" edge.
func HasPatChats() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatChatsTable, PatChatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatChatsWith applies the HasEdge predicate on the "Pat_chats" edge with a given conditions (other predicates).
func HasPatChatsWith(preds ...predicate.PatChat) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := newPatChatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Video) predicate.Video {
	return predicate.Video(sql.NotPredicates(p))
}
