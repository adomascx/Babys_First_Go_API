.PHONY: default build clean

OUTPUT_DIR = build
MAIN_DIR = ./cmd/api

default: build

build:
	mkdir -p $(OUTPUT_DIR)/

	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/skelbiu-api_1.1.0_linux_amd64 $(MAIN_DIR)/
	GOOS=linux GOARCH=arm64 go build -o $(OUTPUT_DIR)/skelbiu-api_1.1.0_linux_arm64 $(MAIN_DIR)/
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/skelbiu-api_1.1.0_windows_amd64.exe $(MAIN_DIR)/
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/skelbiu-api_1.1.0_darwin_amd64 $(MAIN_DIR)/
	GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/skelbiu-api_1.1.0_darwin_arm64 $(MAIN_DIR)/

clean:
	rm -rf $(OUTPUT_DIR)/