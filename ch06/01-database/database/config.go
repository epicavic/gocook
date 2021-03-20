package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sethvargo/go-retry"
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database connection pool
func Setup(ctx context.Context) (*sql.DB, error) {
	dbuser := os.Getenv("MYSQLUSERNAME")
	dbpass := os.Getenv("MYSQLPASSWORD")
	dbbase := os.Getenv("MYSQLDATABASE")
	dburl := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbuser, dbpass, dbbase)
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		return nil, err
	}

	// retry database connection while mysql is booting up
	fmt.Println("Trying to connect to database")
	err = retry.Constant(ctx, 5*time.Second, func(ctx context.Context) error {
		if err := db.PingContext(ctx); err != nil {
			// This marks the error as retryable
			return retry.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")
	return db, nil
}
