package main

import (
	"github.com/fitzypop/gin-auth0-example/api"
	"github.com/fitzypop/gin-auth0-example/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.NewPostgresClient()
}

func main() {
	r := gin.Default()
	r.GET("/", api.Home)
	r.GET("/api/v1/articles/:id", api.GetArticle)
	r.GET("/api/v1/articles", api.GetArticles)
	r.POST("/api/v1/articles", api.PostArticle)
	r.PUT("/api/v1/articles/:id", api.PutArticle)
	r.DELETE("/api/v1/articles/:id", api.DeleteArticle)
	r.Run()
}
