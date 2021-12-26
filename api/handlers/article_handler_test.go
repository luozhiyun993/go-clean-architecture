package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"my-clean-rchitecture/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestArticleHandler_FetchArticle(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	createAt, _ := time.Parse("2006-01-02", "2021-12-26")
	mockService := mock.NewMockIArticleService(ctl)

	gomock.InOrder(
		mockService.EXPECT().Fetch(gomock.Any(), createAt, 10).Return(nil, nil),
	)

	article := NewArticleHandler(mockService)

	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/articles", article.FetchArticle)

	req, err := http.NewRequest(http.MethodGet, "/articles?num=10&create_date=2021-12-26", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
