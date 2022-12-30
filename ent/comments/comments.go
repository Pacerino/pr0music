// Code generated by ent, DO NOT EDIT.

package comments

import (
	"time"
)

const (
	// Label holds the string label denoting the comments type in the database.
	Label = "comments"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldItemID holds the string denoting the item_id field in the database.
	FieldItemID = "item_id"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// FieldUp holds the string denoting the up field in the database.
	FieldUp = "up"
	// FieldDown holds the string denoting the down field in the database.
	FieldDown = "down"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldThumb holds the string denoting the thumb field in the database.
	FieldThumb = "thumb"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// Table holds the table name of the comments in the database.
	Table = "comments"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "comments"
	// ItemsInverseTable is the table name for the Items entity.
	// It exists in this package in order to avoid circular dependency with the "items" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "items_comments"
)

// Columns holds all SQL columns for comments fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldItemID,
	FieldCommentID,
	FieldUp,
	FieldDown,
	FieldContent,
	FieldCreated,
	FieldThumb,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"items_comments",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
