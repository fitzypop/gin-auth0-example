package api

import (
	"net/http"

	"github.com/fitzypop/gin-auth0-example/db"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"messge": "Hello Gin. Cheers!"})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := db.ReadArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

func GetArticles(c *gin.Context) {
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

func PostArticle(c *gin.Context) {
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

func PutArticle(c *gin.Context) {
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

func DeleteArticle(c *gin.Context) {
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
