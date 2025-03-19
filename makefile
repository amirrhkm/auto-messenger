PHONY: migrate-up
migrate-up:
	migrate -path ./backend/migrations -database "mysql://root:root@tcp(127.0.0.1:3307)/auto-messenger?parseTime=true" up

PHONY: migrate-down
migrate-down:
	migrate -path ./backend/migrations -database "mysql://root:root@tcp(127.0.0.1:3307)/auto-messenger?parseTime=true" down
