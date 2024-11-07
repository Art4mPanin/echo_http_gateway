package pkg

import (
	_ "Http-gateway/docs"
	"Http-gateway/internal/config"
	"Http-gateway/internal/services/authservice"
	"Http-gateway/internal/services/infoservice"
	"Http-gateway/pkg/logger"
	"Http-gateway/pkg/singleton"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

// InitServer
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey JWTBearer
// @in header
// @name Authorization
// @description JWT Bearer authentication thru Auth GRPC
//
//go:generate swag init --output ../../
func InitServer() {

	newLogger := logger.SetupLogger()
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	h := singleton.InitializeSingletonHandler()
	h.RegisterSingleton("logger", newLogger)
	h.RegisterSingleton("config", cfg)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/login", authservice.Login)
	e.POST("/register", authservice.Register)
	e.GET("/me", authservice.GetMe)
	e.POST("/validate", authservice.Validate)
	e.POST("/refresh", authservice.Refresh)

	e.POST("/info", infoservice.CreateInfo)

	log.Fatal(e.Start(":5560"))
}
