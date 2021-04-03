// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// Joined applies equality check predicate on the "joined" field. It's identical to JoinedEQ.
func Joined(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoined), v))
	})
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// Exp applies equality check predicate on the "exp" field. It's identical to ExpEQ.
func Exp(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExp), v))
	})
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// Banned applies equality check predicate on the "banned" field. It's identical to BannedEQ.
func Banned(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBanned), v))
	})
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// JoinedEQ applies the EQ predicate on the "joined" field.
func JoinedEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJoined), v))
	})
}

// JoinedNEQ applies the NEQ predicate on the "joined" field.
func JoinedNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJoined), v))
	})
}

// JoinedIn applies the In predicate on the "joined" field.
func JoinedIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJoined), v...))
	})
}

// JoinedNotIn applies the NotIn predicate on the "joined" field.
func JoinedNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJoined), v...))
	})
}

// JoinedGT applies the GT predicate on the "joined" field.
func JoinedGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJoined), v))
	})
}

// JoinedGTE applies the GTE predicate on the "joined" field.
func JoinedGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJoined), v))
	})
}

// JoinedLT applies the LT predicate on the "joined" field.
func JoinedLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJoined), v))
	})
}

// JoinedLTE applies the LTE predicate on the "joined" field.
func JoinedLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJoined), v))
	})
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoints), v))
	})
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...uint) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPoints), v...))
	})
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...uint) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPoints), v...))
	})
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPoints), v))
	})
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPoints), v))
	})
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPoints), v))
	})
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v uint) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPoints), v))
	})
}

// ExpEQ applies the EQ predicate on the "exp" field.
func ExpEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExp), v))
	})
}

// ExpNEQ applies the NEQ predicate on the "exp" field.
func ExpNEQ(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExp), v))
	})
}

// ExpIn applies the In predicate on the "exp" field.
func ExpIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExp), v...))
	})
}

// ExpNotIn applies the NotIn predicate on the "exp" field.
func ExpNotIn(vs ...uint64) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExp), v...))
	})
}

// ExpGT applies the GT predicate on the "exp" field.
func ExpGT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExp), v))
	})
}

// ExpGTE applies the GTE predicate on the "exp" field.
func ExpGTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExp), v))
	})
}

// ExpLT applies the LT predicate on the "exp" field.
func ExpLT(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExp), v))
	})
}

// ExpLTE applies the LTE predicate on the "exp" field.
func ExpLTE(v uint64) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExp), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExternalID), v...))
	})
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExternalID), v...))
	})
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExternalID), v))
	})
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExternalID), v))
	})
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExternalID), v))
	})
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExternalID), v))
	})
}

// BannedEQ applies the EQ predicate on the "banned" field.
func BannedEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBanned), v))
	})
}

// BannedNEQ applies the NEQ predicate on the "banned" field.
func BannedNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBanned), v))
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
