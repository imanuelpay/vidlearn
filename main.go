package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"vidlearn-final-projcect/api"
	"vidlearn-final-projcect/api/middleware"
	"vidlearn-final-projcect/app/modules"
	"vidlearn-final-projcect/config"
	"vidlearn-final-projcect/util"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	config := config.GetConfig()

	dbCon := util.CreateDatabaseConnection(config)
	dbLog := util.CreateLogConnection(config)
	defer dbCon.CloseConnection()
	defer dbLog.CloseConnection()

	controllers := modules.RegisterModules(dbCon, config)

	e := echo.New()
	middleware.Logger(e, dbLog)
	e.HTTPErrorHandler = middleware.ErrorHandler
	e.HideBanner = true

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	api.RegistrationPath(e, controllers, config)

	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
