// Code generated by ent, DO NOT EDIT.

package student

import (
	"newdeal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// LoginAccount applies equality check predicate on the "login_account" field. It's identical to LoginAccountEQ.
func LoginAccount(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldLoginAccount, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPassword, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUsername, v))
}

// DateOfBirth applies equality check predicate on the "date_of_birth" field. It's identical to DateOfBirthEQ.
func DateOfBirth(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldDateOfBirth, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEmail, v))
}

// UpdatedTime applies equality check predicate on the "updated_time" field. It's identical to UpdatedTimeEQ.
func UpdatedTime(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUpdatedTime, v))
}

// VisibleFlg applies equality check predicate on the "visible_flg" field. It's identical to VisibleFlgEQ.
func VisibleFlg(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldVisibleFlg, v))
}

// LoginAccountEQ applies the EQ predicate on the "login_account" field.
func LoginAccountEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldLoginAccount, v))
}

// LoginAccountNEQ applies the NEQ predicate on the "login_account" field.
func LoginAccountNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldLoginAccount, v))
}

// LoginAccountIn applies the In predicate on the "login_account" field.
func LoginAccountIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldLoginAccount, vs...))
}

// LoginAccountNotIn applies the NotIn predicate on the "login_account" field.
func LoginAccountNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldLoginAccount, vs...))
}

// LoginAccountGT applies the GT predicate on the "login_account" field.
func LoginAccountGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldLoginAccount, v))
}

// LoginAccountGTE applies the GTE predicate on the "login_account" field.
func LoginAccountGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldLoginAccount, v))
}

// LoginAccountLT applies the LT predicate on the "login_account" field.
func LoginAccountLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldLoginAccount, v))
}

// LoginAccountLTE applies the LTE predicate on the "login_account" field.
func LoginAccountLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldLoginAccount, v))
}

// LoginAccountContains applies the Contains predicate on the "login_account" field.
func LoginAccountContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldLoginAccount, v))
}

// LoginAccountHasPrefix applies the HasPrefix predicate on the "login_account" field.
func LoginAccountHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldLoginAccount, v))
}

// LoginAccountHasSuffix applies the HasSuffix predicate on the "login_account" field.
func LoginAccountHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldLoginAccount, v))
}

// LoginAccountEqualFold applies the EqualFold predicate on the "login_account" field.
func LoginAccountEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldLoginAccount, v))
}

// LoginAccountContainsFold applies the ContainsFold predicate on the "login_account" field.
func LoginAccountContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldLoginAccount, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldPassword, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldUsername, v))
}

// DateOfBirthEQ applies the EQ predicate on the "date_of_birth" field.
func DateOfBirthEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldDateOfBirth, v))
}

// DateOfBirthNEQ applies the NEQ predicate on the "date_of_birth" field.
func DateOfBirthNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldDateOfBirth, v))
}

// DateOfBirthIn applies the In predicate on the "date_of_birth" field.
func DateOfBirthIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldDateOfBirth, vs...))
}

// DateOfBirthNotIn applies the NotIn predicate on the "date_of_birth" field.
func DateOfBirthNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldDateOfBirth, vs...))
}

// DateOfBirthGT applies the GT predicate on the "date_of_birth" field.
func DateOfBirthGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldDateOfBirth, v))
}

// DateOfBirthGTE applies the GTE predicate on the "date_of_birth" field.
func DateOfBirthGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldDateOfBirth, v))
}

// DateOfBirthLT applies the LT predicate on the "date_of_birth" field.
func DateOfBirthLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldDateOfBirth, v))
}

// DateOfBirthLTE applies the LTE predicate on the "date_of_birth" field.
func DateOfBirthLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldDateOfBirth, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldEmail, v))
}

// UpdatedTimeEQ applies the EQ predicate on the "updated_time" field.
func UpdatedTimeEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUpdatedTime, v))
}

// UpdatedTimeNEQ applies the NEQ predicate on the "updated_time" field.
func UpdatedTimeNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldUpdatedTime, v))
}

// UpdatedTimeIn applies the In predicate on the "updated_time" field.
func UpdatedTimeIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldUpdatedTime, vs...))
}

// UpdatedTimeNotIn applies the NotIn predicate on the "updated_time" field.
func UpdatedTimeNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldUpdatedTime, vs...))
}

// UpdatedTimeGT applies the GT predicate on the "updated_time" field.
func UpdatedTimeGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldUpdatedTime, v))
}

// UpdatedTimeGTE applies the GTE predicate on the "updated_time" field.
func UpdatedTimeGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldUpdatedTime, v))
}

// UpdatedTimeLT applies the LT predicate on the "updated_time" field.
func UpdatedTimeLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldUpdatedTime, v))
}

// UpdatedTimeLTE applies the LTE predicate on the "updated_time" field.
func UpdatedTimeLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldUpdatedTime, v))
}

// VisibleFlgEQ applies the EQ predicate on the "visible_flg" field.
func VisibleFlgEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldVisibleFlg, v))
}

// VisibleFlgNEQ applies the NEQ predicate on the "visible_flg" field.
func VisibleFlgNEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldVisibleFlg, v))
}

// HasHymns applies the HasEdge predicate on the "hymns" edge.
func HasHymns() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HymnsTable, HymnsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHymnsWith applies the HasEdge predicate on the "hymns" edge with a given conditions (other predicates).
func HasHymnsWith(preds ...predicate.Hymn) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := newHymnsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
