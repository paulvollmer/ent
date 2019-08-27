// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleNode() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the node's edges.
	n1 := client.Node.
		Create().
		SetValue(1).
		SaveX(ctx)
	log.Println("node created:", n1)

	// create node vertex with its edges.
	n := client.Node.
		Create().
		SetValue(1).
		AddChildren(n1).
		SaveX(ctx)
	log.Println("node created:", n)

	// query edges.

	n1, err = n.QueryChildren().First(ctx)
	if err != nil {
		log.Fatalf("failed querying children: %v", err)
	}
	log.Println("children found:", n1)

	// Output:
}
