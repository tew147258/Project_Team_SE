// Code generated by entc, DO NOT EDIT.

package gender

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Gendername applies equality check predicate on the "Gendername" field. It's identical to GendernameEQ.
func Gendername(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGendername), v))
	})
}

// GendernameEQ applies the EQ predicate on the "Gendername" field.
func GendernameEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGendername), v))
	})
}

// GendernameNEQ applies the NEQ predicate on the "Gendername" field.
func GendernameNEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGendername), v))
	})
}

// GendernameIn applies the In predicate on the "Gendername" field.
func GendernameIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGendername), v...))
	})
}

// GendernameNotIn applies the NotIn predicate on the "Gendername" field.
func GendernameNotIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGendername), v...))
	})
}

// GendernameGT applies the GT predicate on the "Gendername" field.
func GendernameGT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGendername), v))
	})
}

// GendernameGTE applies the GTE predicate on the "Gendername" field.
func GendernameGTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGendername), v))
	})
}

// GendernameLT applies the LT predicate on the "Gendername" field.
func GendernameLT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGendername), v))
	})
}

// GendernameLTE applies the LTE predicate on the "Gendername" field.
func GendernameLTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGendername), v))
	})
}

// GendernameContains applies the Contains predicate on the "Gendername" field.
func GendernameContains(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGendername), v))
	})
}

// GendernameHasPrefix applies the HasPrefix predicate on the "Gendername" field.
func GendernameHasPrefix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGendername), v))
	})
}

// GendernameHasSuffix applies the HasSuffix predicate on the "Gendername" field.
func GendernameHasSuffix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGendername), v))
	})
}

// GendernameEqualFold applies the EqualFold predicate on the "Gendername" field.
func GendernameEqualFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGendername), v))
	})
}

// GendernameContainsFold applies the ContainsFold predicate on the "Gendername" field.
func GendernameContainsFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGendername), v))
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Personal) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func Not(p predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		p(s.Not())
	})
}