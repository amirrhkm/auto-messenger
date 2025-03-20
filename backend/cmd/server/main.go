package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"auto-messenger/internal/config"
	"auto-messenger/internal/scheduler"
	"auto-messenger/internal/service"
	"auto-messenger/pkg/callmebot"
	"auto-messenger/pkg/logger"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients      = make(map[*websocket.Conn]bool)
	clientsMutex = sync.Mutex{}
)

func BroadcastLog(message string) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, logger logger.Logger) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("(handleWebSocket):" + err.Error())
		return
	}

	// Register new client
	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	logger.Info("[Success] (handleWebSocket): New WebSocket client connected")

	conn.WriteMessage(websocket.TextMessage, []byte("[Success] (handleWebSocket): Connected to Auto Messenger Logger"))

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			clientsMutex.Lock()
			delete(clients, conn)
			clientsMutex.Unlock()
			logger.Info("[Success] (handleWebSocket): WebSocket client disconnected")
			break
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("(godotenv):" + err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()
	logger := logger.NewLogger(cfg.LogLevel)

	logger.SetBroadcastFunc(BroadcastLog)

	callmeBotClient := callmebot.NewClient(callmebot.Config{
		PhoneNumber: cfg.CallMeBotConfig.PhoneNumber,
		ApiKey:      cfg.CallMeBotConfig.ApiKey,
	})

	messageService := service.NewMessageService(callmeBotClient, logger)

	scheduler := scheduler.NewScheduler(messageService, logger)
	scheduler.Start(ctx)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(w, r, logger)
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Error("(websocket server):" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	cancel()
	time.Sleep(2 * time.Second)
}
