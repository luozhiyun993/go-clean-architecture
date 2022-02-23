package repo

import (
	"context"
	"my-clean-architecture/models"
	"time"
)

// IArticleRepo represent the article's repository contract
type IArticleRepo interface {
	Fetch(ctx context.Context, createdDate time.Time, num int) (res []models.Article, err error)
}
