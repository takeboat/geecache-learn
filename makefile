.Phony: run build clean

BINARY = gee-cache
BUILD_DIR = bin

run:
	@go run main.go
build:
	@go build -o $(BUILD_DIR)/$(BINARY) .

clean:
	@rm -rf $(BUILD_DIR)