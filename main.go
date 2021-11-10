package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"iam-api-service/config"
	"iam-api-service/handler"
	userHandler "iam-api-service/handler/user"
	userRepo "iam-api-service/repository/user"
	userService "iam-api-service/service/user"
	util "iam-api-service/util/password"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	configData := config.GetConfig()

	dbConnection := config.NewDatabaseConnection(configData)
	util := util.NewUtil()

	userRepo := userRepo.NewGormDBRepository(dbConnection)

	userService := userService.NewService(userRepo, util)

	userHandler := userHandler.NewHandler(userService)

	e := echo.New()

	handler.RegisterPath(e, userHandler)
	go func() {
		address := fmt.Sprintf(":%d", configData.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
