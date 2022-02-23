package repo

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"my-clean-architecture/models"
	"testing"
	"time"
)

func getSqlMock() (mock sqlmock.Sqlmock, gormDB *gorm.DB) {
	//创建sqlmock
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	//结合gorm、sqlmock
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
	return
}

func Test_mysqlArticleRepository_Fetch(t *testing.T) {
	createAt := time.Now()
	updateAt := time.Now()
	//id,title,content, updated_at, created_at
	var articles = []models.Article{
		{1, "test1", "content", updateAt, createAt},
		{2, "test2", "content2", updateAt, createAt},
	}

	limit := 2
	mock, db := getSqlMock()

	mock.ExpectQuery("SELECT id,title,content, updated_at, created_at FROM `articles` WHERE created_at > ? LIMIT 2").
		WithArgs(createAt).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "created_at"}).
			AddRow(articles[0].ID, articles[0].Title, articles[0].Content, articles[0].UpdatedAt, articles[0].CreatedAt).
			AddRow(articles[1].ID, articles[1].Title, articles[1].Content, articles[1].UpdatedAt, articles[1].CreatedAt))

	repository := NewMysqlArticleRepository(db)
	result, err := repository.Fetch(context.TODO(), createAt, limit)

	assert.Nil(t, err)
	assert.Equal(t, articles, result)
}
