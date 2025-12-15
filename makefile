help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  bun - Run the application"
	@echo "  help - Show this help message"

run:
	go run cmd/api/main.go

.PHONY: bun clean