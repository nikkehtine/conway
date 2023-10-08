BINARY_NAME=conway
OUTPUT_DIR=bin

build:
	@go build -o $(OUTPUT_DIR)/$(BINARY_NAME) ./src

run:
	@./$(OUTPUT_DIR)/$(BINARY_NAME)
