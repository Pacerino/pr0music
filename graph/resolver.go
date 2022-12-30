package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Pacerino/pr0music/ent"
	"github.com/Pacerino/pr0music/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
