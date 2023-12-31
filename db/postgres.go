package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

const DEFAULT_ENV string = "postgres"

func NewPostgresClient() {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DEFAULT_ENV,
		"5432",
		DEFAULT_ENV,
		DEFAULT_ENV,
		"pass1234",
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article not created")
	}
	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var a Article
	res := db.First(&a, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return &a, nil
}

func ReadArticles() ([]*Article, error) {
	var articles []*Article
	res := db.Find(&articles)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return articles, nil
}

func UpdateArticle(article *Article) (*Article, error) {
	var updateArtile Article
	result := db.Model(&updateArtile).Where(article.ID).Updates(article)
	if result.RowsAffected == 0 {
		return &Article{}, errors.New("article not updated")
	}
	return &updateArtile, nil
}

func DeleteArticle(id string) (*Article, error) {
	var article Article
	res := db.Where(id).Delete(&article)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article data not deleted")
	}
	return &article, nil
}
