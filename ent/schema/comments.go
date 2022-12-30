package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Comments holds the schema definition for the Comments entity.
type Comments struct {
	ent.Schema
}

// Fields of the Comments.
func (Comments) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("item_id").
			Immutable(),
		field.Int("comment_id").
			Optional().
			Unique(),
		field.Int("up").
			Optional(),
		field.Int("down").
			Optional(),
		field.Text("content").
			Optional(),
		field.Time("created").
			Optional(),
		field.Text("thumb").
			Optional(),
	}
}

// Edges of the Comments.
func (Comments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("items", Items.Type).
			Ref("comments").
			Unique(),
	}
}

// Indexes of the Comments.
func (Comments) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("item_id").
			Edges("items").
			Unique(),
	}
}
