// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pacerino/pr0music/ent/comments"
	"github.com/Pacerino/pr0music/ent/items"
)

// ItemsCreate is the builder for creating a Items entity.
type ItemsCreate struct {
	config
	mutation *ItemsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ic *ItemsCreate) SetCreatedAt(t time.Time) *ItemsCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableCreatedAt(t *time.Time) *ItemsCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ItemsCreate) SetUpdatedAt(t time.Time) *ItemsCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableUpdatedAt(t *time.Time) *ItemsCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetItemID sets the "item_id" field.
func (ic *ItemsCreate) SetItemID(i int) *ItemsCreate {
	ic.mutation.SetItemID(i)
	return ic
}

// SetTitle sets the "title" field.
func (ic *ItemsCreate) SetTitle(s string) *ItemsCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableTitle(s *string) *ItemsCreate {
	if s != nil {
		ic.SetTitle(*s)
	}
	return ic
}

// SetAlbum sets the "album" field.
func (ic *ItemsCreate) SetAlbum(s string) *ItemsCreate {
	ic.mutation.SetAlbum(s)
	return ic
}

// SetNillableAlbum sets the "album" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableAlbum(s *string) *ItemsCreate {
	if s != nil {
		ic.SetAlbum(*s)
	}
	return ic
}

// SetArtist sets the "artist" field.
func (ic *ItemsCreate) SetArtist(s string) *ItemsCreate {
	ic.mutation.SetArtist(s)
	return ic
}

// SetNillableArtist sets the "artist" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableArtist(s *string) *ItemsCreate {
	if s != nil {
		ic.SetArtist(*s)
	}
	return ic
}

// SetURL sets the "url" field.
func (ic *ItemsCreate) SetURL(s string) *ItemsCreate {
	ic.mutation.SetURL(s)
	return ic
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableURL(s *string) *ItemsCreate {
	if s != nil {
		ic.SetURL(*s)
	}
	return ic
}

// SetAcrID sets the "acr_id" field.
func (ic *ItemsCreate) SetAcrID(s string) *ItemsCreate {
	ic.mutation.SetAcrID(s)
	return ic
}

// SetNillableAcrID sets the "acr_id" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableAcrID(s *string) *ItemsCreate {
	if s != nil {
		ic.SetAcrID(*s)
	}
	return ic
}

// SetSpotifyURL sets the "spotify_url" field.
func (ic *ItemsCreate) SetSpotifyURL(s string) *ItemsCreate {
	ic.mutation.SetSpotifyURL(s)
	return ic
}

// SetNillableSpotifyURL sets the "spotify_url" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableSpotifyURL(s *string) *ItemsCreate {
	if s != nil {
		ic.SetSpotifyURL(*s)
	}
	return ic
}

// SetSpotifyID sets the "spotify_id" field.
func (ic *ItemsCreate) SetSpotifyID(s string) *ItemsCreate {
	ic.mutation.SetSpotifyID(s)
	return ic
}

// SetNillableSpotifyID sets the "spotify_id" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableSpotifyID(s *string) *ItemsCreate {
	if s != nil {
		ic.SetSpotifyID(*s)
	}
	return ic
}

// SetYoutubeURL sets the "youtube_url" field.
func (ic *ItemsCreate) SetYoutubeURL(s string) *ItemsCreate {
	ic.mutation.SetYoutubeURL(s)
	return ic
}

// SetNillableYoutubeURL sets the "youtube_url" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableYoutubeURL(s *string) *ItemsCreate {
	if s != nil {
		ic.SetYoutubeURL(*s)
	}
	return ic
}

// SetYoutubeID sets the "youtube_id" field.
func (ic *ItemsCreate) SetYoutubeID(s string) *ItemsCreate {
	ic.mutation.SetYoutubeID(s)
	return ic
}

// SetNillableYoutubeID sets the "youtube_id" field if the given value is not nil.
func (ic *ItemsCreate) SetNillableYoutubeID(s *string) *ItemsCreate {
	if s != nil {
		ic.SetYoutubeID(*s)
	}
	return ic
}

// AddCommentIDs adds the "comments" edge to the Comments entity by IDs.
func (ic *ItemsCreate) AddCommentIDs(ids ...int) *ItemsCreate {
	ic.mutation.AddCommentIDs(ids...)
	return ic
}

// AddComments adds the "comments" edges to the Comments entity.
func (ic *ItemsCreate) AddComments(c ...*Comments) *ItemsCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ic.AddCommentIDs(ids...)
}

// Mutation returns the ItemsMutation object of the builder.
func (ic *ItemsCreate) Mutation() *ItemsMutation {
	return ic.mutation
}

// Save creates the Items in the database.
func (ic *ItemsCreate) Save(ctx context.Context) (*Items, error) {
	var (
		err  error
		node *Items
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Items)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ItemsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemsCreate) SaveX(ctx context.Context) *Items {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemsCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemsCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemsCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := items.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := items.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemsCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Items.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Items.updated_at"`)}
	}
	if _, ok := ic.mutation.ItemID(); !ok {
		return &ValidationError{Name: "item_id", err: errors.New(`ent: missing required field "Items.item_id"`)}
	}
	return nil
}

func (ic *ItemsCreate) sqlSave(ctx context.Context) (*Items, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *ItemsCreate) createSpec() (*Items, *sqlgraph.CreateSpec) {
	var (
		_node = &Items{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: items.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: items.FieldID,
			},
		}
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(items.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(items.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.ItemID(); ok {
		_spec.SetField(items.FieldItemID, field.TypeInt, value)
		_node.ItemID = value
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.SetField(items.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ic.mutation.Album(); ok {
		_spec.SetField(items.FieldAlbum, field.TypeString, value)
		_node.Album = value
	}
	if value, ok := ic.mutation.Artist(); ok {
		_spec.SetField(items.FieldArtist, field.TypeString, value)
		_node.Artist = value
	}
	if value, ok := ic.mutation.URL(); ok {
		_spec.SetField(items.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ic.mutation.AcrID(); ok {
		_spec.SetField(items.FieldAcrID, field.TypeString, value)
		_node.AcrID = value
	}
	if value, ok := ic.mutation.SpotifyURL(); ok {
		_spec.SetField(items.FieldSpotifyURL, field.TypeString, value)
		_node.SpotifyURL = value
	}
	if value, ok := ic.mutation.SpotifyID(); ok {
		_spec.SetField(items.FieldSpotifyID, field.TypeString, value)
		_node.SpotifyID = value
	}
	if value, ok := ic.mutation.YoutubeURL(); ok {
		_spec.SetField(items.FieldYoutubeURL, field.TypeString, value)
		_node.YoutubeURL = value
	}
	if value, ok := ic.mutation.YoutubeID(); ok {
		_spec.SetField(items.FieldYoutubeID, field.TypeString, value)
		_node.YoutubeID = value
	}
	if nodes := ic.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   items.CommentsTable,
			Columns: []string{items.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comments.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Items.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemsUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ic *ItemsCreate) OnConflict(opts ...sql.ConflictOption) *ItemsUpsertOne {
	ic.conflict = opts
	return &ItemsUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Items.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *ItemsCreate) OnConflictColumns(columns ...string) *ItemsUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &ItemsUpsertOne{
		create: ic,
	}
}

type (
	// ItemsUpsertOne is the builder for "upsert"-ing
	//  one Items node.
	ItemsUpsertOne struct {
		create *ItemsCreate
	}

	// ItemsUpsert is the "OnConflict" setter.
	ItemsUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemsUpsert) SetUpdatedAt(v time.Time) *ItemsUpsert {
	u.Set(items.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateUpdatedAt() *ItemsUpsert {
	u.SetExcluded(items.FieldUpdatedAt)
	return u
}

// SetTitle sets the "title" field.
func (u *ItemsUpsert) SetTitle(v string) *ItemsUpsert {
	u.Set(items.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateTitle() *ItemsUpsert {
	u.SetExcluded(items.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *ItemsUpsert) ClearTitle() *ItemsUpsert {
	u.SetNull(items.FieldTitle)
	return u
}

// SetAlbum sets the "album" field.
func (u *ItemsUpsert) SetAlbum(v string) *ItemsUpsert {
	u.Set(items.FieldAlbum, v)
	return u
}

// UpdateAlbum sets the "album" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateAlbum() *ItemsUpsert {
	u.SetExcluded(items.FieldAlbum)
	return u
}

// ClearAlbum clears the value of the "album" field.
func (u *ItemsUpsert) ClearAlbum() *ItemsUpsert {
	u.SetNull(items.FieldAlbum)
	return u
}

// SetArtist sets the "artist" field.
func (u *ItemsUpsert) SetArtist(v string) *ItemsUpsert {
	u.Set(items.FieldArtist, v)
	return u
}

// UpdateArtist sets the "artist" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateArtist() *ItemsUpsert {
	u.SetExcluded(items.FieldArtist)
	return u
}

// ClearArtist clears the value of the "artist" field.
func (u *ItemsUpsert) ClearArtist() *ItemsUpsert {
	u.SetNull(items.FieldArtist)
	return u
}

// SetURL sets the "url" field.
func (u *ItemsUpsert) SetURL(v string) *ItemsUpsert {
	u.Set(items.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateURL() *ItemsUpsert {
	u.SetExcluded(items.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *ItemsUpsert) ClearURL() *ItemsUpsert {
	u.SetNull(items.FieldURL)
	return u
}

// SetAcrID sets the "acr_id" field.
func (u *ItemsUpsert) SetAcrID(v string) *ItemsUpsert {
	u.Set(items.FieldAcrID, v)
	return u
}

// UpdateAcrID sets the "acr_id" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateAcrID() *ItemsUpsert {
	u.SetExcluded(items.FieldAcrID)
	return u
}

// ClearAcrID clears the value of the "acr_id" field.
func (u *ItemsUpsert) ClearAcrID() *ItemsUpsert {
	u.SetNull(items.FieldAcrID)
	return u
}

// SetSpotifyURL sets the "spotify_url" field.
func (u *ItemsUpsert) SetSpotifyURL(v string) *ItemsUpsert {
	u.Set(items.FieldSpotifyURL, v)
	return u
}

// UpdateSpotifyURL sets the "spotify_url" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateSpotifyURL() *ItemsUpsert {
	u.SetExcluded(items.FieldSpotifyURL)
	return u
}

// ClearSpotifyURL clears the value of the "spotify_url" field.
func (u *ItemsUpsert) ClearSpotifyURL() *ItemsUpsert {
	u.SetNull(items.FieldSpotifyURL)
	return u
}

// SetSpotifyID sets the "spotify_id" field.
func (u *ItemsUpsert) SetSpotifyID(v string) *ItemsUpsert {
	u.Set(items.FieldSpotifyID, v)
	return u
}

// UpdateSpotifyID sets the "spotify_id" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateSpotifyID() *ItemsUpsert {
	u.SetExcluded(items.FieldSpotifyID)
	return u
}

// ClearSpotifyID clears the value of the "spotify_id" field.
func (u *ItemsUpsert) ClearSpotifyID() *ItemsUpsert {
	u.SetNull(items.FieldSpotifyID)
	return u
}

// SetYoutubeURL sets the "youtube_url" field.
func (u *ItemsUpsert) SetYoutubeURL(v string) *ItemsUpsert {
	u.Set(items.FieldYoutubeURL, v)
	return u
}

// UpdateYoutubeURL sets the "youtube_url" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateYoutubeURL() *ItemsUpsert {
	u.SetExcluded(items.FieldYoutubeURL)
	return u
}

// ClearYoutubeURL clears the value of the "youtube_url" field.
func (u *ItemsUpsert) ClearYoutubeURL() *ItemsUpsert {
	u.SetNull(items.FieldYoutubeURL)
	return u
}

// SetYoutubeID sets the "youtube_id" field.
func (u *ItemsUpsert) SetYoutubeID(v string) *ItemsUpsert {
	u.Set(items.FieldYoutubeID, v)
	return u
}

// UpdateYoutubeID sets the "youtube_id" field to the value that was provided on create.
func (u *ItemsUpsert) UpdateYoutubeID() *ItemsUpsert {
	u.SetExcluded(items.FieldYoutubeID)
	return u
}

// ClearYoutubeID clears the value of the "youtube_id" field.
func (u *ItemsUpsert) ClearYoutubeID() *ItemsUpsert {
	u.SetNull(items.FieldYoutubeID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Items.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ItemsUpsertOne) UpdateNewValues() *ItemsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(items.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.ItemID(); exists {
			s.SetIgnore(items.FieldItemID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Items.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemsUpsertOne) Ignore() *ItemsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemsUpsertOne) DoNothing() *ItemsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemsCreate.OnConflict
// documentation for more info.
func (u *ItemsUpsertOne) Update(set func(*ItemsUpsert)) *ItemsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemsUpsertOne) SetUpdatedAt(v time.Time) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateUpdatedAt() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *ItemsUpsertOne) SetTitle(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateTitle() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ItemsUpsertOne) ClearTitle() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearTitle()
	})
}

// SetAlbum sets the "album" field.
func (u *ItemsUpsertOne) SetAlbum(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetAlbum(v)
	})
}

// UpdateAlbum sets the "album" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateAlbum() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateAlbum()
	})
}

// ClearAlbum clears the value of the "album" field.
func (u *ItemsUpsertOne) ClearAlbum() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearAlbum()
	})
}

// SetArtist sets the "artist" field.
func (u *ItemsUpsertOne) SetArtist(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetArtist(v)
	})
}

// UpdateArtist sets the "artist" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateArtist() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateArtist()
	})
}

// ClearArtist clears the value of the "artist" field.
func (u *ItemsUpsertOne) ClearArtist() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearArtist()
	})
}

// SetURL sets the "url" field.
func (u *ItemsUpsertOne) SetURL(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *ItemsUpsertOne) ClearURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearURL()
	})
}

// SetAcrID sets the "acr_id" field.
func (u *ItemsUpsertOne) SetAcrID(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetAcrID(v)
	})
}

// UpdateAcrID sets the "acr_id" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateAcrID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateAcrID()
	})
}

// ClearAcrID clears the value of the "acr_id" field.
func (u *ItemsUpsertOne) ClearAcrID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearAcrID()
	})
}

// SetSpotifyURL sets the "spotify_url" field.
func (u *ItemsUpsertOne) SetSpotifyURL(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetSpotifyURL(v)
	})
}

// UpdateSpotifyURL sets the "spotify_url" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateSpotifyURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateSpotifyURL()
	})
}

// ClearSpotifyURL clears the value of the "spotify_url" field.
func (u *ItemsUpsertOne) ClearSpotifyURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearSpotifyURL()
	})
}

// SetSpotifyID sets the "spotify_id" field.
func (u *ItemsUpsertOne) SetSpotifyID(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetSpotifyID(v)
	})
}

// UpdateSpotifyID sets the "spotify_id" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateSpotifyID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateSpotifyID()
	})
}

// ClearSpotifyID clears the value of the "spotify_id" field.
func (u *ItemsUpsertOne) ClearSpotifyID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearSpotifyID()
	})
}

// SetYoutubeURL sets the "youtube_url" field.
func (u *ItemsUpsertOne) SetYoutubeURL(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetYoutubeURL(v)
	})
}

// UpdateYoutubeURL sets the "youtube_url" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateYoutubeURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateYoutubeURL()
	})
}

// ClearYoutubeURL clears the value of the "youtube_url" field.
func (u *ItemsUpsertOne) ClearYoutubeURL() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearYoutubeURL()
	})
}

// SetYoutubeID sets the "youtube_id" field.
func (u *ItemsUpsertOne) SetYoutubeID(v string) *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.SetYoutubeID(v)
	})
}

// UpdateYoutubeID sets the "youtube_id" field to the value that was provided on create.
func (u *ItemsUpsertOne) UpdateYoutubeID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateYoutubeID()
	})
}

// ClearYoutubeID clears the value of the "youtube_id" field.
func (u *ItemsUpsertOne) ClearYoutubeID() *ItemsUpsertOne {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearYoutubeID()
	})
}

// Exec executes the query.
func (u *ItemsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemsCreateBulk is the builder for creating many Items entities in bulk.
type ItemsCreateBulk struct {
	config
	builders []*ItemsCreate
	conflict []sql.ConflictOption
}

// Save creates the Items entities in the database.
func (icb *ItemsCreateBulk) Save(ctx context.Context) ([]*Items, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Items, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemsCreateBulk) SaveX(ctx context.Context) []*Items {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemsCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemsCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Items.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemsUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (icb *ItemsCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemsUpsertBulk {
	icb.conflict = opts
	return &ItemsUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Items.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *ItemsCreateBulk) OnConflictColumns(columns ...string) *ItemsUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &ItemsUpsertBulk{
		create: icb,
	}
}

// ItemsUpsertBulk is the builder for "upsert"-ing
// a bulk of Items nodes.
type ItemsUpsertBulk struct {
	create *ItemsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Items.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ItemsUpsertBulk) UpdateNewValues() *ItemsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(items.FieldCreatedAt)
			}
			if _, exists := b.mutation.ItemID(); exists {
				s.SetIgnore(items.FieldItemID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Items.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemsUpsertBulk) Ignore() *ItemsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemsUpsertBulk) DoNothing() *ItemsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemsCreateBulk.OnConflict
// documentation for more info.
func (u *ItemsUpsertBulk) Update(set func(*ItemsUpsert)) *ItemsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemsUpsertBulk) SetUpdatedAt(v time.Time) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateUpdatedAt() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *ItemsUpsertBulk) SetTitle(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateTitle() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ItemsUpsertBulk) ClearTitle() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearTitle()
	})
}

// SetAlbum sets the "album" field.
func (u *ItemsUpsertBulk) SetAlbum(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetAlbum(v)
	})
}

// UpdateAlbum sets the "album" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateAlbum() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateAlbum()
	})
}

// ClearAlbum clears the value of the "album" field.
func (u *ItemsUpsertBulk) ClearAlbum() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearAlbum()
	})
}

// SetArtist sets the "artist" field.
func (u *ItemsUpsertBulk) SetArtist(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetArtist(v)
	})
}

// UpdateArtist sets the "artist" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateArtist() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateArtist()
	})
}

// ClearArtist clears the value of the "artist" field.
func (u *ItemsUpsertBulk) ClearArtist() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearArtist()
	})
}

// SetURL sets the "url" field.
func (u *ItemsUpsertBulk) SetURL(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *ItemsUpsertBulk) ClearURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearURL()
	})
}

// SetAcrID sets the "acr_id" field.
func (u *ItemsUpsertBulk) SetAcrID(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetAcrID(v)
	})
}

// UpdateAcrID sets the "acr_id" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateAcrID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateAcrID()
	})
}

// ClearAcrID clears the value of the "acr_id" field.
func (u *ItemsUpsertBulk) ClearAcrID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearAcrID()
	})
}

// SetSpotifyURL sets the "spotify_url" field.
func (u *ItemsUpsertBulk) SetSpotifyURL(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetSpotifyURL(v)
	})
}

// UpdateSpotifyURL sets the "spotify_url" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateSpotifyURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateSpotifyURL()
	})
}

// ClearSpotifyURL clears the value of the "spotify_url" field.
func (u *ItemsUpsertBulk) ClearSpotifyURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearSpotifyURL()
	})
}

// SetSpotifyID sets the "spotify_id" field.
func (u *ItemsUpsertBulk) SetSpotifyID(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetSpotifyID(v)
	})
}

// UpdateSpotifyID sets the "spotify_id" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateSpotifyID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateSpotifyID()
	})
}

// ClearSpotifyID clears the value of the "spotify_id" field.
func (u *ItemsUpsertBulk) ClearSpotifyID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearSpotifyID()
	})
}

// SetYoutubeURL sets the "youtube_url" field.
func (u *ItemsUpsertBulk) SetYoutubeURL(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetYoutubeURL(v)
	})
}

// UpdateYoutubeURL sets the "youtube_url" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateYoutubeURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateYoutubeURL()
	})
}

// ClearYoutubeURL clears the value of the "youtube_url" field.
func (u *ItemsUpsertBulk) ClearYoutubeURL() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearYoutubeURL()
	})
}

// SetYoutubeID sets the "youtube_id" field.
func (u *ItemsUpsertBulk) SetYoutubeID(v string) *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.SetYoutubeID(v)
	})
}

// UpdateYoutubeID sets the "youtube_id" field to the value that was provided on create.
func (u *ItemsUpsertBulk) UpdateYoutubeID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.UpdateYoutubeID()
	})
}

// ClearYoutubeID clears the value of the "youtube_id" field.
func (u *ItemsUpsertBulk) ClearYoutubeID() *ItemsUpsertBulk {
	return u.Update(func(s *ItemsUpsert) {
		s.ClearYoutubeID()
	})
}

// Exec executes the query.
func (u *ItemsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
