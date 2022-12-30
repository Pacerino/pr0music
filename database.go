package main

import (
	"context"

	ent "github.com/Pacerino/pr0music/ent"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB(dsn string) (*ent.Client, error) {

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return client, nil
}
