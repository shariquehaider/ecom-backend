package main

import (
	"github.com/shariquehaider/ecom-backend/models"
	"github.com/shariquehaider/ecom-backend/router"
)

func main() {
	models.InitDB()
	router := router.SetupRouter()
	router.Run(":8000")
}
