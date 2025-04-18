// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"newdeal/ent/student"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// LoginAccount holds the value of the "login_account" field.
	LoginAccount string `json:"login_account,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// UpdatedTime holds the value of the "updated_time" field.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
	// VisibleFlg holds the value of the "visible_flg" field.
	VisibleFlg bool `json:"visible_flg,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges        StudentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Hymns holds the value of the hymns edge.
	Hymns []*Hymn `json:"hymns,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HymnsOrErr returns the Hymns value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) HymnsOrErr() ([]*Hymn, error) {
	if e.loadedTypes[0] {
		return e.Hymns, nil
	}
	return nil, &NotLoadedError{edge: "hymns"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldVisibleFlg:
			values[i] = new(sql.NullBool)
		case student.FieldID:
			values[i] = new(sql.NullInt64)
		case student.FieldLoginAccount, student.FieldPassword, student.FieldUsername, student.FieldEmail:
			values[i] = new(sql.NullString)
		case student.FieldDateOfBirth, student.FieldUpdatedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case student.FieldLoginAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login_account", values[i])
			} else if value.Valid {
				s.LoginAccount = value.String
			}
		case student.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				s.Password = value.String
			}
		case student.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				s.Username = value.String
			}
		case student.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value.Valid {
				s.DateOfBirth = value.Time
			}
		case student.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = new(string)
				*s.Email = value.String
			}
		case student.FieldUpdatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_time", values[i])
			} else if value.Valid {
				s.UpdatedTime = new(time.Time)
				*s.UpdatedTime = value.Time
			}
		case student.FieldVisibleFlg:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field visible_flg", values[i])
			} else if value.Valid {
				s.VisibleFlg = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Student.
// This includes values selected through modifiers, order, etc.
func (s *Student) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryHymns queries the "hymns" edge of the Student entity.
func (s *Student) QueryHymns() *HymnQuery {
	return NewStudentClient(s.config).QueryHymns(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("login_account=")
	builder.WriteString(s.LoginAccount)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(s.Password)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(s.Username)
	builder.WriteString(", ")
	builder.WriteString("date_of_birth=")
	builder.WriteString(s.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.UpdatedTime; v != nil {
		builder.WriteString("updated_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("visible_flg=")
	builder.WriteString(fmt.Sprintf("%v", s.VisibleFlg))
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
