package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/santosfilipe/eeveentory/api"
)

func ConfigureLogger() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	log.Println("IMPORTANT: eeveentory is up & running!")
	log.Println("IMPORTANT: Deployment done through Github Actions!")
	ConfigureLogger()

	router := gin.Default()
	router.GET("/assets", api.GetAssets)
	router.GET("/assets/:ip", api.GetAssetsByIp)
	router.Run(":8888")
}
