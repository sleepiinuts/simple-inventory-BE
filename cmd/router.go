package main

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sleepiinuts/simple-inventory-BE/middleware"
)

func getRouter(logger *slog.Logger) {
	router = gin.Default()
	router.Use(middleware.ErrorHandler(logger))
	router.Use(cors.Default())
	router.GET("/product", prodApi.GetAll)

	router.Run("localhost:8080")
}
