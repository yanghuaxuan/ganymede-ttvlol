// Code generated by ent, DO NOT EDIT.

package playback

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/internal/utils"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldID, id))
}

// VodID applies equality check predicate on the "vod_id" field. It's identical to VodIDEQ.
func VodID(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldVodID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldUserID, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldTime, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldCreatedAt, v))
}

// VodIDEQ applies the EQ predicate on the "vod_id" field.
func VodIDEQ(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldVodID, v))
}

// VodIDNEQ applies the NEQ predicate on the "vod_id" field.
func VodIDNEQ(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldVodID, v))
}

// VodIDIn applies the In predicate on the "vod_id" field.
func VodIDIn(vs ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldVodID, vs...))
}

// VodIDNotIn applies the NotIn predicate on the "vod_id" field.
func VodIDNotIn(vs ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldVodID, vs...))
}

// VodIDGT applies the GT predicate on the "vod_id" field.
func VodIDGT(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldVodID, v))
}

// VodIDGTE applies the GTE predicate on the "vod_id" field.
func VodIDGTE(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldVodID, v))
}

// VodIDLT applies the LT predicate on the "vod_id" field.
func VodIDLT(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldVodID, v))
}

// VodIDLTE applies the LTE predicate on the "vod_id" field.
func VodIDLTE(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldVodID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldUserID, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v utils.PlaybackStatus) predicate.Playback {
	vc := v
	return predicate.Playback(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v utils.PlaybackStatus) predicate.Playback {
	vc := v
	return predicate.Playback(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...utils.PlaybackStatus) predicate.Playback {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Playback(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...utils.PlaybackStatus) predicate.Playback {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Playback(sql.FieldNotIn(FieldStatus, v...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Playback {
	return predicate.Playback(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Playback {
	return predicate.Playback(sql.FieldNotNull(FieldStatus))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Playback {
	return predicate.Playback(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Playback) predicate.Playback {
	return predicate.Playback(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Playback) predicate.Playback {
	return predicate.Playback(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Playback) predicate.Playback {
	return predicate.Playback(func(s *sql.Selector) {
		p(s.Not())
	})
}
