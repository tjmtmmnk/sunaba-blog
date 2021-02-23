package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tjmtmmnk/blog/graph/model"
	"github.com/tjmtmmnk/blog/pkg/db"
	"github.com/tjmtmmnk/blog/pkg/db/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *sql.DB
}

func NewResolver() (*Resolver, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	return &Resolver{
		db: db,
	}, nil
}

func (r *Resolver) getArticlesByUserID(ctx context.Context, userID string) ([]*model.Article, error) {
	title := "title"
	user, err := models.FindUser(context.Background(), r.db, 0)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)

	return []*model.Article{
		{
			ID:    "11",
			Title: &title,
			User: &model.User{
				ID:   "22",
				Name: "bob",
			},
			Body:      "Hello",
			CreatedAt: "2021",
		},
	}, nil
}
