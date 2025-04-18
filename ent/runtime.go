// Code generated by ent, DO NOT EDIT.

package ent

import (
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/ent/schema"
	"newdeal/ent/student"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	hymnFields := schema.Hymn{}.Fields()
	_ = hymnFields
	// hymnDescID is the schema descriptor for id field.
	hymnDescID := hymnFields[0].Descriptor()
	// hymn.IDValidator is a validator for the "id" field. It is called by the builders before save.
	hymn.IDValidator = hymnDescID.Validators[0].(func(int64) error)
	hymnsworkFields := schema.HymnsWork{}.Fields()
	_ = hymnsworkFields
	// hymnsworkDescWorkID is the schema descriptor for work_id field.
	hymnsworkDescWorkID := hymnsworkFields[0].Descriptor()
	// hymnswork.WorkIDValidator is a validator for the "work_id" field. It is called by the builders before save.
	hymnswork.WorkIDValidator = hymnsworkDescWorkID.Validators[0].(func(int64) error)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentFields[0].Descriptor()
	// student.IDValidator is a validator for the "id" field. It is called by the builders before save.
	student.IDValidator = studentDescID.Validators[0].(func(int64) error)
}
