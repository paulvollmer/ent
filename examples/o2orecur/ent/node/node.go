// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package node

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value vertex property in the database.
	FieldValue = "value"

	// Table holds the table name of the node in the database.
	Table = "nodes"
	// PrevTable is the table the holds the prev relation/edge.
	PrevTable = "nodes"
	// PrevColumn is the table column denoting the prev relation/edge.
	PrevColumn = "prev_id"
	// NextTable is the table the holds the next relation/edge.
	NextTable = "nodes"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "prev_id"
)

// Columns holds all SQL columns are node fields.
var Columns = []string{
	FieldID,
	FieldValue,
}
