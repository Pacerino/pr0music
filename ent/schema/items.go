package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Items holds the schema definition for the Items entity.
type Items struct {
	ent.Schema
}

// Fields of the Items.
func (Items) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("item_id").
			Immutable().
			Unique(),
		field.Text("title").
			Optional(),
		field.Text("album").
			Optional(),
		field.Text("artist").
			Optional(),
		field.Text("url").
			Optional(),
		field.Text("acr_id").
			Optional(),
		field.Text("spotify_url").
			Optional(),
		field.Text("spotify_id").
			Optional(),
		field.Text("youtube_url").
			Optional(),
		field.Text("youtube_id").
			Optional(),
	}
}

// Edges of the Items.
func (Items) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comments.Type),
	}
}
