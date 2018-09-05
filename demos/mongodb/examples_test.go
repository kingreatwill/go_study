package mongodb

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/examples/documentation_examples"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/stretchr/testify/require"
)

func TestDocumentationExamples(t *testing.T) {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	require.NoError(t, err)

	db := client.Database("documentation_examples")

	documentation_examples.InsertExamples(t, db)
	documentation_examples.QueryToplevelFieldsExamples(t, db)
	documentation_examples.QueryEmbeddedDocumentsExamples(t, db)
	documentation_examples.QueryArraysExamples(t, db)
	documentation_examples.QueryArrayEmbeddedDocumentsExamples(t, db)
	documentation_examples.QueryNullMissingFieldsExamples(t, db)
	documentation_examples.ProjectionExamples(t, db)
	documentation_examples.UpdateExamples(t, db)
	documentation_examples.DeleteExamples(t, db)
}
