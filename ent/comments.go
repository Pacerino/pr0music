// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Pacerino/pr0music/ent/comments"
	"github.com/Pacerino/pr0music/ent/items"
)

// Comments is the model entity for the Comments schema.
type Comments struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ItemID holds the value of the "item_id" field.
	ItemID int `json:"item_id,omitempty"`
	// CommentID holds the value of the "comment_id" field.
	CommentID int `json:"comment_id,omitempty"`
	// Up holds the value of the "up" field.
	Up int `json:"up,omitempty"`
	// Down holds the value of the "down" field.
	Down int `json:"down,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Thumb holds the value of the "thumb" field.
	Thumb string `json:"thumb,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentsQuery when eager-loading is set.
	Edges          CommentsEdges `json:"edges"`
	items_comments *int
}

// CommentsEdges holds the relations/edges for other nodes in the graph.
type CommentsEdges struct {
	// Items holds the value of the items edge.
	Items *Items `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentsEdges) ItemsOrErr() (*Items, error) {
	if e.loadedTypes[0] {
		if e.Items == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: items.Label}
		}
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comments) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comments.FieldID, comments.FieldItemID, comments.FieldCommentID, comments.FieldUp, comments.FieldDown:
			values[i] = new(sql.NullInt64)
		case comments.FieldContent, comments.FieldThumb:
			values[i] = new(sql.NullString)
		case comments.FieldCreatedAt, comments.FieldUpdatedAt, comments.FieldCreated:
			values[i] = new(sql.NullTime)
		case comments.ForeignKeys[0]: // items_comments
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comments", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comments fields.
func (c *Comments) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comments.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comments.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comments.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case comments.FieldItemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field item_id", values[i])
			} else if value.Valid {
				c.ItemID = int(value.Int64)
			}
		case comments.FieldCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_id", values[i])
			} else if value.Valid {
				c.CommentID = int(value.Int64)
			}
		case comments.FieldUp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field up", values[i])
			} else if value.Valid {
				c.Up = int(value.Int64)
			}
		case comments.FieldDown:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field down", values[i])
			} else if value.Valid {
				c.Down = int(value.Int64)
			}
		case comments.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comments.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				c.Created = value.Time
			}
		case comments.FieldThumb:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumb", values[i])
			} else if value.Valid {
				c.Thumb = value.String
			}
		case comments.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field items_comments", value)
			} else if value.Valid {
				c.items_comments = new(int)
				*c.items_comments = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryItems queries the "items" edge of the Comments entity.
func (c *Comments) QueryItems() *ItemsQuery {
	return (&CommentsClient{config: c.config}).QueryItems(c)
}

// Update returns a builder for updating this Comments.
// Note that you need to call Comments.Unwrap() before calling this method if this Comments
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comments) Update() *CommentsUpdateOne {
	return (&CommentsClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comments entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comments) Unwrap() *Comments {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comments is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comments) String() string {
	var builder strings.Builder
	builder.WriteString("Comments(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("item_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ItemID))
	builder.WriteString(", ")
	builder.WriteString("comment_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CommentID))
	builder.WriteString(", ")
	builder.WriteString("up=")
	builder.WriteString(fmt.Sprintf("%v", c.Up))
	builder.WriteString(", ")
	builder.WriteString("down=")
	builder.WriteString(fmt.Sprintf("%v", c.Down))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(c.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("thumb=")
	builder.WriteString(c.Thumb)
	builder.WriteByte(')')
	return builder.String()
}

// CommentsSlice is a parsable slice of Comments.
type CommentsSlice []*Comments

func (c CommentsSlice) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}