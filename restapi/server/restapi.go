package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-algorithm-api/restapi/service"
)

func StartRESTServer() {
	g := gin.Default()

	beefService, err := service.InitBeefService("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to initialize BeefService: %v", err)
	}

	// health check
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g.GET("/beef/summary", beefService.GetBeefMap)

	serverAddress := fmt.Sprintf(":8082")
	log.Printf("Starting REST server on %s...", serverAddress)
	if err := g.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
