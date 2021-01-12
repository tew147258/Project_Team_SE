// Code generated by entc, DO NOT EDIT.

package customer

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Customername applies equality check predicate on the "Customername" field. It's identical to CustomernameEQ.
func Customername(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomername), v))
	})
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Phonenumber applies equality check predicate on the "Phonenumber" field. It's identical to PhonenumberEQ.
func Phonenumber(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhonenumber), v))
	})
}

// CustomernameEQ applies the EQ predicate on the "Customername" field.
func CustomernameEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomername), v))
	})
}

// CustomernameNEQ applies the NEQ predicate on the "Customername" field.
func CustomernameNEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomername), v))
	})
}

// CustomernameIn applies the In predicate on the "Customername" field.
func CustomernameIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCustomername), v...))
	})
}

// CustomernameNotIn applies the NotIn predicate on the "Customername" field.
func CustomernameNotIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCustomername), v...))
	})
}

// CustomernameGT applies the GT predicate on the "Customername" field.
func CustomernameGT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustomername), v))
	})
}

// CustomernameGTE applies the GTE predicate on the "Customername" field.
func CustomernameGTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustomername), v))
	})
}

// CustomernameLT applies the LT predicate on the "Customername" field.
func CustomernameLT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustomername), v))
	})
}

// CustomernameLTE applies the LTE predicate on the "Customername" field.
func CustomernameLTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustomername), v))
	})
}

// CustomernameContains applies the Contains predicate on the "Customername" field.
func CustomernameContains(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustomername), v))
	})
}

// CustomernameHasPrefix applies the HasPrefix predicate on the "Customername" field.
func CustomernameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustomername), v))
	})
}

// CustomernameHasSuffix applies the HasSuffix predicate on the "Customername" field.
func CustomernameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustomername), v))
	})
}

// CustomernameEqualFold applies the EqualFold predicate on the "Customername" field.
func CustomernameEqualFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustomername), v))
	})
}

// CustomernameContainsFold applies the ContainsFold predicate on the "Customername" field.
func CustomernameContainsFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustomername), v))
	})
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// PhonenumberEQ applies the EQ predicate on the "Phonenumber" field.
func PhonenumberEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberNEQ applies the NEQ predicate on the "Phonenumber" field.
func PhonenumberNEQ(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberIn applies the In predicate on the "Phonenumber" field.
func PhonenumberIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhonenumber), v...))
	})
}

// PhonenumberNotIn applies the NotIn predicate on the "Phonenumber" field.
func PhonenumberNotIn(vs ...string) predicate.Customer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Customer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhonenumber), v...))
	})
}

// PhonenumberGT applies the GT predicate on the "Phonenumber" field.
func PhonenumberGT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberGTE applies the GTE predicate on the "Phonenumber" field.
func PhonenumberGTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberLT applies the LT predicate on the "Phonenumber" field.
func PhonenumberLT(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberLTE applies the LTE predicate on the "Phonenumber" field.
func PhonenumberLTE(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberContains applies the Contains predicate on the "Phonenumber" field.
func PhonenumberContains(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberHasPrefix applies the HasPrefix predicate on the "Phonenumber" field.
func PhonenumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberHasSuffix applies the HasSuffix predicate on the "Phonenumber" field.
func PhonenumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberEqualFold applies the EqualFold predicate on the "Phonenumber" field.
func PhonenumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhonenumber), v))
	})
}

// PhonenumberContainsFold applies the ContainsFold predicate on the "Phonenumber" field.
func PhonenumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhonenumber), v))
	})
}

// HasGender applies the HasEdge predicate on the "gender" edge.
func HasGender() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderWith applies the HasEdge predicate on the "gender" edge with a given conditions (other predicates).
func HasGenderWith(preds ...predicate.Gender) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Personal) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTitle applies the HasEdge predicate on the "title" edge.
func HasTitle() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleWith applies the HasEdge predicate on the "title" edge with a given conditions (other predicates).
func HasTitleWith(preds ...predicate.Title) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
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
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		p(s.Not())
	})
}