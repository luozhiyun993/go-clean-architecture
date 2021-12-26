package service

import (
	"context"
	"my-clean-rchitecture/models"
	"time"
)

// IArticleService represent the article's usecases
type IArticleService interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) ([]models.Article, error)
}
