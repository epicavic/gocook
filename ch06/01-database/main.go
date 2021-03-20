package main

import (
	"context"
	"main/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	db, err := database.Setup(ctx)
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
