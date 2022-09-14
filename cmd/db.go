package cmd

import (
	"context"

	"cloud.google.com/go/spanner"
)

func createClient(db string) (*spanner.Client, error) {
	return spanner.NewClient(context.Background(), db)
}
