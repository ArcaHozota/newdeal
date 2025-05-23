// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"newdeal/ent/book"
	"newdeal/ent/chapter"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Chapter is the model entity for the Chapter schema.
type Chapter struct {
	config `json:"-"`
	// ID of the ent.
	// 章節ID
	ID int32 `json:"id,omitempty"`
	// 章節名
	Name string `json:"name,omitempty"`
	// 章節日本語名
	NameJp string `json:"name_jp,omitempty"`
	// 書別ID
	BookID int16 `json:"book_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChapterQuery when eager-loading is set.
	Edges        ChapterEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChapterEdges holds the relations/edges for other nodes in the graph.
type ChapterEdges struct {
	// ToPhrase holds the value of the to_phrase edge.
	ToPhrase []*Phrase `json:"to_phrase,omitempty"`
	// ChapterBook holds the value of the chapter_book edge.
	ChapterBook *Book `json:"chapter_book,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ToPhraseOrErr returns the ToPhrase value or an error if the edge
// was not loaded in eager-loading.
func (e ChapterEdges) ToPhraseOrErr() ([]*Phrase, error) {
	if e.loadedTypes[0] {
		return e.ToPhrase, nil
	}
	return nil, &NotLoadedError{edge: "to_phrase"}
}

// ChapterBookOrErr returns the ChapterBook value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChapterEdges) ChapterBookOrErr() (*Book, error) {
	if e.ChapterBook != nil {
		return e.ChapterBook, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: book.Label}
	}
	return nil, &NotLoadedError{edge: "chapter_book"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chapter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chapter.FieldID, chapter.FieldBookID:
			values[i] = new(sql.NullInt64)
		case chapter.FieldName, chapter.FieldNameJp:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chapter fields.
func (c *Chapter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chapter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int32(value.Int64)
		case chapter.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case chapter.FieldNameJp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_jp", values[i])
			} else if value.Valid {
				c.NameJp = value.String
			}
		case chapter.FieldBookID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field book_id", values[i])
			} else if value.Valid {
				c.BookID = int16(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Chapter.
// This includes values selected through modifiers, order, etc.
func (c *Chapter) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryToPhrase queries the "to_phrase" edge of the Chapter entity.
func (c *Chapter) QueryToPhrase() *PhraseQuery {
	return NewChapterClient(c.config).QueryToPhrase(c)
}

// QueryChapterBook queries the "chapter_book" edge of the Chapter entity.
func (c *Chapter) QueryChapterBook() *BookQuery {
	return NewChapterClient(c.config).QueryChapterBook(c)
}

// Update returns a builder for updating this Chapter.
// Note that you need to call Chapter.Unwrap() before calling this method if this Chapter
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chapter) Update() *ChapterUpdateOne {
	return NewChapterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Chapter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chapter) Unwrap() *Chapter {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chapter is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chapter) String() string {
	var builder strings.Builder
	builder.WriteString("Chapter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("name_jp=")
	builder.WriteString(c.NameJp)
	builder.WriteString(", ")
	builder.WriteString("book_id=")
	builder.WriteString(fmt.Sprintf("%v", c.BookID))
	builder.WriteByte(')')
	return builder.String()
}

// Chapters is a parsable slice of Chapter.
type Chapters []*Chapter
