package main

import (
	"shorturl/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
