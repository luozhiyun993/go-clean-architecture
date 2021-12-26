package repo

import (
	"context"
	"gorm.io/gorm"
	"my-clean-rchitecture/models"
	"time"
)

type mysqlArticleRepository struct {
	DB *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(DB *gorm.DB) IArticleRepo {
	return &mysqlArticleRepository{DB}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, createdDate time.Time,
	num int) (res []models.Article, err error) {

	err = m.DB.WithContext(ctx).Model(&models.Article{}).
		Select("id,title,content, updated_at, created_at").
		Where("created_at > ?", createdDate).Limit(num).Find(&res).Error
	return
}
