package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-clean-rchitecture/models"
	"net/http"
	"strconv"
	"time"
)

// ArticleHandler
type ArticleHandler struct {
	ArticleService models.IArticleService
}

// NewArticleHandler
func NewArticleHandler(as models.IArticleService) ArticleHandler {
	handler := ArticleHandler{
		ArticleService: as,
	}
	return handler
}

// FetchArticle will fetch the article based on given params
func (a *ArticleHandler) FetchArticle(c *gin.Context) {
	numS := c.Query("num")
	num, _ := strconv.Atoi(numS)
	createDate := c.Query("create_date")
	parseDate, _ := time.Parse("2006-01-02", createDate)
	listAr, err := a.ArticleService.Fetch(c, parseDate, num)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": listAr})
}
