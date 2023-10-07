package main

import (
	"github.com/fitzypop/gin-auth0-example/api"
	"github.com/fitzypop/gin-auth0-example/db"
)

func init() {
	db.NewPostgresClient()
}

func main() {
	r := api.SetupRouter()
	r.Run()
}
