// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/entc/integration/gremlin/ent/user"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges CardEdges `json:"edges"`

	// StaticField defined by templates.
	StaticField string `json:"boring,omitempty"`
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Spec holds the value of the spec edge.
	Spec []*Spec
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// SpecOrErr returns the Spec value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) SpecOrErr() ([]*Spec, error) {
	if e.loadedTypes[1] {
		return e.Spec, nil
	}
	return nil, &NotLoadedError{edge: "spec"}
}

// FromResponse scans the gremlin response data into Card.
func (c *Card) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanc struct {
		ID         string `json:"id,omitempty"`
		CreateTime int64  `json:"create_time,omitempty"`
		UpdateTime int64  `json:"update_time,omitempty"`
		Number     string `json:"number,omitempty"`
		Name       string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanc); err != nil {
		return err
	}
	c.ID = scanc.ID
	c.CreateTime = time.Unix(0, scanc.CreateTime)
	c.UpdateTime = time.Unix(0, scanc.UpdateTime)
	c.Number = scanc.Number
	c.Name = scanc.Name
	return nil
}

// QueryOwner queries the owner edge of the Card.
func (c *Card) QueryOwner() *UserQuery {
	return (&CardClient{config: c.config}).QueryOwner(c)
}

// QuerySpec queries the spec edge of the Card.
func (c *Card) QuerySpec() *SpecQuery {
	return (&CardClient{config: c.config}).QuerySpec(c)
}

// Update returns a builder for updating this Card.
// Note that, you need to call Card.Unwrap() before calling this method, if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", number=")
	builder.WriteString(c.Number)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card

// FromResponse scans the gremlin response data into Cards.
func (c *Cards) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanc []struct {
		ID         string `json:"id,omitempty"`
		CreateTime int64  `json:"create_time,omitempty"`
		UpdateTime int64  `json:"update_time,omitempty"`
		Number     string `json:"number,omitempty"`
		Name       string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanc); err != nil {
		return err
	}
	for _, v := range scanc {
		*c = append(*c, &Card{
			ID:         v.ID,
			CreateTime: time.Unix(0, v.CreateTime),
			UpdateTime: time.Unix(0, v.UpdateTime),
			Number:     v.Number,
			Name:       v.Name,
		})
	}
	return nil
}

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
