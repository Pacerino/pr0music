// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/Pacerino/pr0music/ent/comments"
	"github.com/Pacerino/pr0music/ent/items"
	"github.com/Pacerino/pr0music/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   comments.Table,
			Columns: comments.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comments.FieldID,
			},
		},
		Type: "Comments",
		Fields: map[string]*sqlgraph.FieldSpec{
			comments.FieldCreatedAt: {Type: field.TypeTime, Column: comments.FieldCreatedAt},
			comments.FieldUpdatedAt: {Type: field.TypeTime, Column: comments.FieldUpdatedAt},
			comments.FieldItemID:    {Type: field.TypeInt, Column: comments.FieldItemID},
			comments.FieldCommentID: {Type: field.TypeInt, Column: comments.FieldCommentID},
			comments.FieldUp:        {Type: field.TypeInt, Column: comments.FieldUp},
			comments.FieldDown:      {Type: field.TypeInt, Column: comments.FieldDown},
			comments.FieldContent:   {Type: field.TypeString, Column: comments.FieldContent},
			comments.FieldCreated:   {Type: field.TypeTime, Column: comments.FieldCreated},
			comments.FieldThumb:     {Type: field.TypeString, Column: comments.FieldThumb},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   items.Table,
			Columns: items.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: items.FieldID,
			},
		},
		Type: "Items",
		Fields: map[string]*sqlgraph.FieldSpec{
			items.FieldCreatedAt:  {Type: field.TypeTime, Column: items.FieldCreatedAt},
			items.FieldUpdatedAt:  {Type: field.TypeTime, Column: items.FieldUpdatedAt},
			items.FieldItemID:     {Type: field.TypeInt, Column: items.FieldItemID},
			items.FieldTitle:      {Type: field.TypeString, Column: items.FieldTitle},
			items.FieldAlbum:      {Type: field.TypeString, Column: items.FieldAlbum},
			items.FieldArtist:     {Type: field.TypeString, Column: items.FieldArtist},
			items.FieldURL:        {Type: field.TypeString, Column: items.FieldURL},
			items.FieldAcrID:      {Type: field.TypeString, Column: items.FieldAcrID},
			items.FieldSpotifyURL: {Type: field.TypeString, Column: items.FieldSpotifyURL},
			items.FieldSpotifyID:  {Type: field.TypeString, Column: items.FieldSpotifyID},
			items.FieldYoutubeURL: {Type: field.TypeString, Column: items.FieldYoutubeURL},
			items.FieldYoutubeID:  {Type: field.TypeString, Column: items.FieldYoutubeID},
		},
	}
	graph.MustAddE(
		"items",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comments.ItemsTable,
			Columns: []string{comments.ItemsColumn},
			Bidi:    false,
		},
		"Comments",
		"Items",
	)
	graph.MustAddE(
		"comments",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   items.CommentsTable,
			Columns: []string{items.CommentsColumn},
			Bidi:    false,
		},
		"Items",
		"Comments",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CommentsQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CommentsQuery builder.
func (cq *CommentsQuery) Filter() *CommentsFilter {
	return &CommentsFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CommentsMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CommentsMutation builder.
func (m *CommentsMutation) Filter() *CommentsFilter {
	return &CommentsFilter{config: m.config, predicateAdder: m}
}

// CommentsFilter provides a generic filtering capability at runtime for CommentsQuery.
type CommentsFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CommentsFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CommentsFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(comments.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *CommentsFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(comments.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *CommentsFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(comments.FieldUpdatedAt))
}

// WhereItemID applies the entql int predicate on the item_id field.
func (f *CommentsFilter) WhereItemID(p entql.IntP) {
	f.Where(p.Field(comments.FieldItemID))
}

// WhereCommentID applies the entql int predicate on the comment_id field.
func (f *CommentsFilter) WhereCommentID(p entql.IntP) {
	f.Where(p.Field(comments.FieldCommentID))
}

// WhereUp applies the entql int predicate on the up field.
func (f *CommentsFilter) WhereUp(p entql.IntP) {
	f.Where(p.Field(comments.FieldUp))
}

// WhereDown applies the entql int predicate on the down field.
func (f *CommentsFilter) WhereDown(p entql.IntP) {
	f.Where(p.Field(comments.FieldDown))
}

// WhereContent applies the entql string predicate on the content field.
func (f *CommentsFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(comments.FieldContent))
}

// WhereCreated applies the entql time.Time predicate on the created field.
func (f *CommentsFilter) WhereCreated(p entql.TimeP) {
	f.Where(p.Field(comments.FieldCreated))
}

// WhereThumb applies the entql string predicate on the thumb field.
func (f *CommentsFilter) WhereThumb(p entql.StringP) {
	f.Where(p.Field(comments.FieldThumb))
}

// WhereHasItems applies a predicate to check if query has an edge items.
func (f *CommentsFilter) WhereHasItems() {
	f.Where(entql.HasEdge("items"))
}

// WhereHasItemsWith applies a predicate to check if query has an edge items with a given conditions (other predicates).
func (f *CommentsFilter) WhereHasItemsWith(preds ...predicate.Items) {
	f.Where(entql.HasEdgeWith("items", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (iq *ItemsQuery) addPredicate(pred func(s *sql.Selector)) {
	iq.predicates = append(iq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ItemsQuery builder.
func (iq *ItemsQuery) Filter() *ItemsFilter {
	return &ItemsFilter{config: iq.config, predicateAdder: iq}
}

// addPredicate implements the predicateAdder interface.
func (m *ItemsMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ItemsMutation builder.
func (m *ItemsMutation) Filter() *ItemsFilter {
	return &ItemsFilter{config: m.config, predicateAdder: m}
}

// ItemsFilter provides a generic filtering capability at runtime for ItemsQuery.
type ItemsFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ItemsFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ItemsFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(items.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *ItemsFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(items.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *ItemsFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(items.FieldUpdatedAt))
}

// WhereItemID applies the entql int predicate on the item_id field.
func (f *ItemsFilter) WhereItemID(p entql.IntP) {
	f.Where(p.Field(items.FieldItemID))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *ItemsFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(items.FieldTitle))
}

// WhereAlbum applies the entql string predicate on the album field.
func (f *ItemsFilter) WhereAlbum(p entql.StringP) {
	f.Where(p.Field(items.FieldAlbum))
}

// WhereArtist applies the entql string predicate on the artist field.
func (f *ItemsFilter) WhereArtist(p entql.StringP) {
	f.Where(p.Field(items.FieldArtist))
}

// WhereURL applies the entql string predicate on the url field.
func (f *ItemsFilter) WhereURL(p entql.StringP) {
	f.Where(p.Field(items.FieldURL))
}

// WhereAcrID applies the entql string predicate on the acr_id field.
func (f *ItemsFilter) WhereAcrID(p entql.StringP) {
	f.Where(p.Field(items.FieldAcrID))
}

// WhereSpotifyURL applies the entql string predicate on the spotify_url field.
func (f *ItemsFilter) WhereSpotifyURL(p entql.StringP) {
	f.Where(p.Field(items.FieldSpotifyURL))
}

// WhereSpotifyID applies the entql string predicate on the spotify_id field.
func (f *ItemsFilter) WhereSpotifyID(p entql.StringP) {
	f.Where(p.Field(items.FieldSpotifyID))
}

// WhereYoutubeURL applies the entql string predicate on the youtube_url field.
func (f *ItemsFilter) WhereYoutubeURL(p entql.StringP) {
	f.Where(p.Field(items.FieldYoutubeURL))
}

// WhereYoutubeID applies the entql string predicate on the youtube_id field.
func (f *ItemsFilter) WhereYoutubeID(p entql.StringP) {
	f.Where(p.Field(items.FieldYoutubeID))
}

// WhereHasComments applies a predicate to check if query has an edge comments.
func (f *ItemsFilter) WhereHasComments() {
	f.Where(entql.HasEdge("comments"))
}

// WhereHasCommentsWith applies a predicate to check if query has an edge comments with a given conditions (other predicates).
func (f *ItemsFilter) WhereHasCommentsWith(preds ...predicate.Comments) {
	f.Where(entql.HasEdgeWith("comments", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
