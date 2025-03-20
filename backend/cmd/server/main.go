package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auto-messenger/internal/config"
	"auto-messenger/internal/scheduler"
	"auto-messenger/internal/service"
	"auto-messenger/pkg/callmebot"
	"auto-messenger/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("[Error] (godotenv):" + err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()
	logger := logger.NewLogger(cfg.LogLevel)

	callmeBotClient := callmebot.NewClient(callmebot.Config{
		PhoneNumber: cfg.CallMeBotConfig.PhoneNumber,
		ApiKey:      cfg.CallMeBotConfig.ApiKey,
	})

	messageService := service.NewMessageService(callmeBotClient, logger)

	scheduler := scheduler.NewScheduler(messageService, logger)
	scheduler.Start(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	cancel()
	time.Sleep(2 * time.Second)
}
