package api

import (
	"github.com/gin-gonic/gin"
	"my-clean-architecture/api/handlers"
)

type Router struct {
	article handlers.ArticleHandler
}

func (r *Router) With(engine *gin.Engine) {
	engine.GET("/articles", r.article.FetchArticle)
}

func NewRouter(article handlers.ArticleHandler) *Router {
	router := &Router{
		article: article,
	}
	return router
}
