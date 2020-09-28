// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/file"
	"github.com/facebook/ent/entc/integration/ent/filetype"
	"github.com/facebook/ent/schema/field"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	mutation *FileTypeMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ftc *FileTypeCreate) SetName(s string) *FileTypeCreate {
	ftc.mutation.SetName(s)
	return ftc
}

// SetType sets the type field.
func (ftc *FileTypeCreate) SetType(f filetype.Type) *FileTypeCreate {
	ftc.mutation.SetType(f)
	return ftc
}

// SetNillableType sets the type field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableType(f *filetype.Type) *FileTypeCreate {
	if f != nil {
		ftc.SetType(*f)
	}
	return ftc
}

// SetState sets the state field.
func (ftc *FileTypeCreate) SetState(f filetype.State) *FileTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the state field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableState(f *filetype.State) *FileTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// AddFileIDs adds the files edge to File by ids.
func (ftc *FileTypeCreate) AddFileIDs(ids ...int) *FileTypeCreate {
	ftc.mutation.AddFileIDs(ids...)
	return ftc
}

// AddFiles adds the files edges to File.
func (ftc *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftc *FileTypeCreate) Mutation() *FileTypeMutation {
	return ftc.mutation
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	var (
		err  error
		node *FileType
	)
	ftc.defaults()
	if len(ftc.hooks) == 0 {
		if err = ftc.check(); err != nil {
			return nil, err
		}
		node, err = ftc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftc.check(); err != nil {
				return nil, err
			}
			ftc.mutation = mutation
			node, err = ftc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ftc *FileTypeCreate) defaults() {
	if _, ok := ftc.mutation.GetType(); !ok {
		v := filetype.DefaultType
		ftc.mutation.SetType(v)
	}
	if _, ok := ftc.mutation.State(); !ok {
		v := filetype.DefaultState
		ftc.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FileTypeCreate) check() error {
	if _, ok := ftc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ftc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := ftc.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := ftc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required field \"state\"")}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (ftc *FileTypeCreate) sqlSave(ctx context.Context) (*FileType, error) {
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ftc *FileTypeCreate) createSpec() (*FileType, *sqlgraph.CreateSpec) {
	var (
		_node = &FileType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: filetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filetype.FieldID,
			},
		}
	)
	if value, ok := ftc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filetype.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ftc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: filetype.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ftc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: filetype.FieldState,
		})
		_node.State = value
	}
	if nodes := ftc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
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

// FileTypeCreateBulk is the builder for creating a bulk of FileType entities.
type FileTypeCreateBulk struct {
	config
	builders []*FileTypeCreate
}

// Save creates the FileType entities in the database.
func (ftcb *FileTypeCreateBulk) Save(ctx context.Context) ([]*FileType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FileType, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ftcb *FileTypeCreateBulk) SaveX(ctx context.Context) []*FileType {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
