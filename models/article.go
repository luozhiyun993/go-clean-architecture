package models

import (
	"context"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// IArticleService represent the article's usecases
type IArticleService interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) ([]Article, error)
}

// IArticleRepo represent the article's repository contract
type IArticleRepo interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) (res []Article, err error)
}
