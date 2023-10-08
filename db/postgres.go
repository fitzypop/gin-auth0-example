package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

const DEFAULT_ENV string = "postgres"

func NewPostgresClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var (
		host     = getEnv("DB_HOST", DEFAULT_ENV)
		port     = getEnv("DB_PORT", "5432")
		user     = getEnv("DB_USER", DEFAULT_ENV)
		dbname   = getEnv("DB_NAME", DEFAULT_ENV)
		password = getEnv("DB_PASSWORD", "pass1234")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
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
