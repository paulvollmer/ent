// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/m2mbidi/ent"
	"github.com/facebookincubator/ent/examples/m2mbidi/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:o2o2types?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	client := ent.NewClient(ent.Driver(db))
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SaveX(ctx)
	nati := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		AddFriends(a8m).
		SaveX(ctx)

	// Query friends. Unlike `All`, `AllX` panics if an error occurs.
	friends := nati.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m)]

	friends = a8m.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=2, age=28, name=nati)]

	// Query the graph:
	friends = client.User.
		Query().
		Where(user.HasFriends()).
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
