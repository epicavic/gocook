package database

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSetup(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Setup(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
