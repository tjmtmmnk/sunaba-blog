package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tjmtmmnk/blog/graph/generated"
	"github.com/tjmtmmnk/blog/graph/model"
)

func (r *queryResolver) ArticlesByUserID(ctx context.Context, userID string) ([]*model.Article, error) {
	return r.getArticlesByUserID(ctx, userID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
