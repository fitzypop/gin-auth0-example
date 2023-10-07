package api

import (
	"net/http"

	"github.com/fitzypop/gin-auth0-example/db"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"messge": "Hello Gin. Cheers!"})
}

func getArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := db.ReadArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

func getArticles(c *gin.Context) {
	articles, err := db.ReadArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func postArticle(c *gin.Context) {
	var article db.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	res, err := db.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": res})
}

func putArticle(c *gin.Context) {
	var article db.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := db.UpdateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
}

func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "article deleted successfully",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/api/v1/articles/:id", getArticle)
	r.GET("/api/v1/articles", getArticles)
	r.POST("/api/v1/articles", postArticle)
	r.PUT("/api/v1/articles/:id", putArticle)
	r.DELETE("/api/v1/articles/:id", deleteArticle)
	return r
}
