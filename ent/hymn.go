// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/ent/student"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Hymn is the model entity for the Hymn schema.
type Hymn struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// NameJp holds the value of the "name_jp" field.
	NameJp string `json:"name_jp,omitempty"`
	// NameKr holds the value of the "name_kr" field.
	NameKr string `json:"name_kr,omitempty"`
	// Link holds the value of the "link" field.
	Link *string `json:"link,omitempty"`
	// UpdatedTime holds the value of the "updated_time" field.
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	// UpdatedUser holds the value of the "updated_user" field.
	UpdatedUser string `json:"updated_user,omitempty"`
	// Serif holds the value of the "serif" field.
	Serif *string `json:"serif,omitempty"`
	// VisibleFlg holds the value of the "visible_flg" field.
	VisibleFlg bool `json:"visible_flg,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HymnQuery when eager-loading is set.
	Edges         HymnEdges `json:"edges"`
	student_hymns *int64
	selectValues  sql.SelectValues
}

// HymnEdges holds the relations/edges for other nodes in the graph.
type HymnEdges struct {
	// Students holds the value of the students edge.
	Students *Student `json:"students,omitempty"`
	// HymnsWork holds the value of the hymns_work edge.
	HymnsWork *HymnsWork `json:"hymns_work,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HymnEdges) StudentsOrErr() (*Student, error) {
	if e.Students != nil {
		return e.Students, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: student.Label}
	}
	return nil, &NotLoadedError{edge: "students"}
}

// HymnsWorkOrErr returns the HymnsWork value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HymnEdges) HymnsWorkOrErr() (*HymnsWork, error) {
	if e.HymnsWork != nil {
		return e.HymnsWork, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: hymnswork.Label}
	}
	return nil, &NotLoadedError{edge: "hymns_work"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hymn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hymn.FieldVisibleFlg:
			values[i] = new(sql.NullBool)
		case hymn.FieldID:
			values[i] = new(sql.NullInt64)
		case hymn.FieldNameJp, hymn.FieldNameKr, hymn.FieldLink, hymn.FieldUpdatedUser, hymn.FieldSerif:
			values[i] = new(sql.NullString)
		case hymn.FieldUpdatedTime:
			values[i] = new(sql.NullTime)
		case hymn.ForeignKeys[0]: // student_hymns
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hymn fields.
func (h *Hymn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hymn.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int64(value.Int64)
		case hymn.FieldNameJp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_jp", values[i])
			} else if value.Valid {
				h.NameJp = value.String
			}
		case hymn.FieldNameKr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_kr", values[i])
			} else if value.Valid {
				h.NameKr = value.String
			}
		case hymn.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				h.Link = new(string)
				*h.Link = value.String
			}
		case hymn.FieldUpdatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_time", values[i])
			} else if value.Valid {
				h.UpdatedTime = value.Time
			}
		case hymn.FieldUpdatedUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_user", values[i])
			} else if value.Valid {
				h.UpdatedUser = value.String
			}
		case hymn.FieldSerif:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serif", values[i])
			} else if value.Valid {
				h.Serif = new(string)
				*h.Serif = value.String
			}
		case hymn.FieldVisibleFlg:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field visible_flg", values[i])
			} else if value.Valid {
				h.VisibleFlg = value.Bool
			}
		case hymn.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field student_hymns", value)
			} else if value.Valid {
				h.student_hymns = new(int64)
				*h.student_hymns = int64(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hymn.
// This includes values selected through modifiers, order, etc.
func (h *Hymn) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryStudents queries the "students" edge of the Hymn entity.
func (h *Hymn) QueryStudents() *StudentQuery {
	return NewHymnClient(h.config).QueryStudents(h)
}

// QueryHymnsWork queries the "hymns_work" edge of the Hymn entity.
func (h *Hymn) QueryHymnsWork() *HymnsWorkQuery {
	return NewHymnClient(h.config).QueryHymnsWork(h)
}

// Update returns a builder for updating this Hymn.
// Note that you need to call Hymn.Unwrap() before calling this method if this Hymn
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hymn) Update() *HymnUpdateOne {
	return NewHymnClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hymn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hymn) Unwrap() *Hymn {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hymn is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hymn) String() string {
	var builder strings.Builder
	builder.WriteString("Hymn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name_jp=")
	builder.WriteString(h.NameJp)
	builder.WriteString(", ")
	builder.WriteString("name_kr=")
	builder.WriteString(h.NameKr)
	builder.WriteString(", ")
	if v := h.Link; v != nil {
		builder.WriteString("link=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("updated_time=")
	builder.WriteString(h.UpdatedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_user=")
	builder.WriteString(h.UpdatedUser)
	builder.WriteString(", ")
	if v := h.Serif; v != nil {
		builder.WriteString("serif=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("visible_flg=")
	builder.WriteString(fmt.Sprintf("%v", h.VisibleFlg))
	builder.WriteByte(')')
	return builder.String()
}

// Hymns is a parsable slice of Hymn.
type Hymns []*Hymn
