BUILD_FILE=notifyd
CMD_DIR=cmd
BUILD_DIR=bin
BUILD_WIN=$(BUILD_DIR)/win

CONFIG_PATH=data/config_dev.ini

.PHONY: build run alert dev
dev:
	go run $(CMD_DIR)/$(BUILD_FILE).go --config=$(CONFIG_PATH)

run: build
	$(BUILD_DIR)/$(BUILD_FILE) --config=$(CONFIG_PATH)

build:
	[ -d $(BUILD_DIR) ] || mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BUILD_FILE) $(CMD_DIR)/$(BUILD_FILE).go

build_win:
	[ -d $(BUILD_WIN) ] || mkdir -p $(BUILD_WIN)
	GOOS=windows GOARCH=386 go1.10 build -o $(BUILD_WIN)/$(BUILD_FILE)_xp.exe $(CMD_DIR)/$(BUILD_FILE).go
	GOOS=windows GOARCH=386 go build -o $(BUILD_WIN)/$(BUILD_FILE)_w10.exe $(CMD_DIR)/$(BUILD_FILE).go

build_all: build build_win