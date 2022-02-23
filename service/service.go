package service

import (
	"context"
	"my-clean-architecture/models"
	"time"
)

// IArticleService represent the article's usecases
type IArticleService interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) ([]models.Article, error)
}
