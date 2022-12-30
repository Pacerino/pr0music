// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "item_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "up", Type: field.TypeInt, Nullable: true},
		{Name: "down", Type: field.TypeInt, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created", Type: field.TypeTime, Nullable: true},
		{Name: "thumb", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "items_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_items_comments",
				Columns:    []*schema.Column{CommentsColumns[10]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "comments_item_id_items_comments",
				Unique:  true,
				Columns: []*schema.Column{CommentsColumns[3], CommentsColumns[10]},
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "item_id", Type: field.TypeInt, Unique: true},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "album", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "artist", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "acr_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "spotify_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "spotify_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "youtube_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "youtube_id", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		ItemsTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = ItemsTable
}
