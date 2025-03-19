package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auto-messenger/internal/config"
	"auto-messenger/internal/repository"
	"auto-messenger/internal/scheduler"
	"auto-messenger/internal/service"
	"auto-messenger/pkg/database"
	"auto-messenger/pkg/logger"
	"auto-messenger/pkg/whatsapp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger := logger.NewLogger(cfg.LogLevel)

	// Initialize WhatsApp client
	whatsappClient := whatsapp.NewClient(whatsapp.Config{
		AccessToken:   cfg.WhatsappConfig.AccessToken,
		PhoneNumberID: cfg.WhatsappConfig.PhoneNumberID,
	})

	// Initialize database connection
	dbConfig := database.DatabaseConfig{
		Driver:   cfg.Database.Driver,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := database.Connect(dbConfig)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return
	}
	defer db.Close()

	// Initialize repository
	messageRepo := repository.NewMessageRepository(db)

	// Initialize service
	messageService := service.NewMessageService(messageRepo, whatsappClient, logger)

	// Initialize scheduler
	scheduler := scheduler.NewScheduler(messageService, logger)
	scheduler.Start(ctx)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	cancel()
	time.Sleep(2 * time.Second)
}
