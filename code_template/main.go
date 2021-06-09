package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-tutorial/code_template/internal"
	"go-tutorial/code_template/internal/handler"
	"os"
	"os/signal"
	"time"
)

func main() {
	//service := os.Getenv("SERVICE")
	port := os.Getenv("PORT")
	state := os.Getenv("STATE")

	cv := internal.Configs{ConfigPath: os.Getenv("CONFIG_PATH")}
	if err := cv.InitAllConfigurations(state); err != nil {
		log.Errorf("init all configurations %s error: %s", state, err)
		return
	}

	e := echo.New()
	e.Use(handler.Recover)

	closes, err := handler.NewRoutes(e, &cv)
	if err != nil {
		log.Panic("new routes error:", err)
		return
	}

	if port == "" {
		port = "8888"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	srvCtx, srvCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer srvCancel()

	// background services should be shutdown after http server so that
	// no more request should be come while shutting down the services
	for _, close := range closes {
		if err := close(); err != nil {
			log.Errorf("failed to cleanup resources: %v", err)
		}
	}

	log.Info("shutting down http server...")
	if err := e.Server.Shutdown(srvCtx); err != nil {
		log.Panic("http server shutdown with error:", err)
	}

}
