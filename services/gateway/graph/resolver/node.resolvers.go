package resolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/yuorei/hackathon/graph/generated"
	"github.com/yuorei/hackathon/graph/model"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	splitID := strings.Split(id, "_")

	kind := splitID[0]
	id = splitID[1]
	switch kind {
	case "user":
		return r.User(ctx, id)
	default:
		return nil, fmt.Errorf("そのようなIDは定義されていません")
	}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
