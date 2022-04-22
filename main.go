package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/santosfilipe/eeveentory/api"
)

func main() {
	log.Println("eeveentory is up & running!")
	router := gin.Default()
	router.GET("/assets", api.GetAssets)
	router.Run("localhost:8888")
}
