package followschema

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestGetterHaserPattern(t *testing.T) {
	resolvers := &Stub{}
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolvers})))

	t.Run("returns value when haser returns true", func(t *testing.T) {
		name := "Alice"
		age := 30
		resolvers.QueryResolver.PersonWithGetterHaser = func(ctx context.Context) (*PersonWithGetterHaser, error) {
			return &PersonWithGetterHaser{
				name: &name,
				age:  &age,
			}, nil
		}

		var resp struct {
			PersonWithGetterHaser struct {
				Name *string
				Age  *int
			}
		}

		c.MustPost(`{ personWithGetterHaser { name age } }`, &resp)
		require.NotNil(t, resp.PersonWithGetterHaser.Name)
		require.Equal(t, "Alice", *resp.PersonWithGetterHaser.Name)
		require.NotNil(t, resp.PersonWithGetterHaser.Age)
		require.Equal(t, 30, *resp.PersonWithGetterHaser.Age)
	})

	t.Run("returns nil when haser returns false", func(t *testing.T) {
		resolvers.QueryResolver.PersonWithGetterHaser = func(ctx context.Context) (*PersonWithGetterHaser, error) {
			return &PersonWithGetterHaser{
				name: nil,
				age:  nil,
			}, nil
		}

		var resp struct {
			PersonWithGetterHaser struct {
				Name *string
				Age  *int
			}
		}

		c.MustPost(`{ personWithGetterHaser { name age } }`, &resp)
		require.Nil(t, resp.PersonWithGetterHaser.Name)
		require.Nil(t, resp.PersonWithGetterHaser.Age)
	})
}
