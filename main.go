package main

import "github.com/shariquehaider/ecom-backend/router"

func main() {
	router := router.SetupRouter()
	router.Run(":8000")
}
