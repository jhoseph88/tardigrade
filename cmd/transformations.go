package cmd

import (
	"context"
	"fmt"
	"log"
	"path"
	"reflect"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Transformer struct{}

func (Transformer) BackfillLogoFileType(ctx context.Context, client *spanner.Client) error {
	// file_uri starting with /tmp indicates the logo is either actively being uploaded or the upload errored out
	stmt := spanner.Statement{SQL: `SELECT id, file_uri FROM logos where file_type IS NULL and NOT STARTS_WITH(file_uri, "/tmp")`}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		var id, fileURI string
		if err := row.Columns(&id, &fileURI); err != nil {
			fmt.Println("Couldn't get columns for some reason")
			fmt.Println(err)
			return err
		}
		// try deriving the file type from the path (everything after the dot)
		fileType := path.Ext(fileURI)[1:]
		_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			return txn.BufferWrite([]*spanner.Mutation{
				spanner.Update("logos", []string{"id", "file_type"}, []interface{}{id, fileType}),
			})
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("Updated logo, set file_type = %s where id = %s\n", fileType, id)
	}
}

func RunTransformation(db, fName string) {
	ctx := context.Background()
	// TODO - make db an argument
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	transformer := reflect.ValueOf(Transformer{})
	transformFn := transformer.MethodByName(fName)
	transformFn.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(client),
	})
}
