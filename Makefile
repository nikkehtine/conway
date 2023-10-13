BINARY_NAME=conway
OUTPUT_DIR=bin

build:
	@go build -o $(OUTPUT_DIR)/$(BINARY_NAME) ./

run:
	@./$(OUTPUT_DIR)/$(BINARY_NAME)
