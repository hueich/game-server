package db

import (
	"context"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

// NewDatastoreClient creates a new Google Datastore client.
// For both projectID and credsFile, they can alternatively be provided through environment variables
// DATASTORE_PROJECT_ID and GOOGLE_APPLICATION_CREDENTIALS, respectively, in which case the corresponding argument can be left empty.
func NewDatastoreClient(ctx context.Context, projectID, credsFile string) (*datastore.Client, error) {
	opts := []option.ClientOption{}
	if credsFile != "" {
		opts = append(opts, option.WithServiceAccountFile(credsFile))
	}
	c, err := datastore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return c, nil
}
