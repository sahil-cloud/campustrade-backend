package main

// /usr/local/go/bin/go run main.go

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sahil-cloud/backend/blockchain"
	"github.com/sahil-cloud/backend/constants"
	"github.com/sahil-cloud/backend/middleware"
	"github.com/sahil-cloud/backend/routes"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// initialize connection to blockchain and get contract
	blockchain.Initialize()

	// router
	router := gin.New()
	router.Use(gin.Logger())
	routes.PublicRoutes(router)
	apiRoutes := router.Group("/api", middleware.Authenticate())
	routes.ServiceRoutes(apiRoutes)
	// routes.OrderRoutes(apiRoutes)

	ctx := context.Background()
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain(constants.NGROK_DOMAIN_NAME),
		),
		ngrok.WithAuthtokenFromEnv())
	if err != nil {
		log.Fatalln(err)
	}

	// router.Run(":" + os.Getenv("APP_PORT"))

	if err := router.RunListener(listener); err != nil {
		log.Fatalln(err)
	}
}
