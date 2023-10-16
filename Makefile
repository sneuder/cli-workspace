# Variables for output directory names
WINDOWS_64_DIR_NAME = windows-amd64
WINDOWS_32_DIR_NAME = windows-386
LINUX_64_DIR_NAME = linux-amd64
LINUX_32_DIR_NAME = linux-386

# Name of the output binary
OUTPUT_BINARY = workspace

# Variables for output directories
WINDOWS_64_DIR = bin/windows/$(WINDOWS_64_DIR_NAME)
WINDOWS_32_DIR = bin/windows/$(WINDOWS_32_DIR_NAME)
LINUX_64_DIR = bin/linux/$(LINUX_64_DIR_NAME)
LINUX_32_DIR = bin/linux/$(LINUX_32_DIR_NAME)

PATH_MAIN = cmd/workspace/main.go

.PHONY: all windows linux clean

all: windows linux

windows: windows_64 windows_32

linux: linux_64 linux_32

windows_64:
	@GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_64_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_64_DIR_NAME).exe $(PATH_MAIN)
	@make -s message_success MESSAGE=$(WINDOWS_64_DIR_NAME)

windows_32:
	@GOOS=windows GOARCH=386 go build -o $(WINDOWS_32_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_32_DIR_NAME).exe $(PATH_MAIN)
	@make -s message_success MESSAGE=$(WINDOWS_32_DIR_NAME)

linux_64:
	@GOOS=linux GOARCH=amd64 go build -o $(LINUX_64_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_DIR_NAME) $(PATH_MAIN)
	@chmod +x $(LINUX_64_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_DIR_NAME)
	@make -s message_success MESSAGE=$(LINUX_64_DIR_NAME)

linux_32:
	@GOOS=linux GOARCH=386 go build -o $(LINUX_32_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_DIR_NAME) $(PATH_MAIN)
	@chmod +x $(LINUX_32_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_DIR_NAME)
	@make -s message_success MESSAGE=$(LINUX_32_DIR_NAME)

message_success:
	@echo "Build successful for: $(MESSAGE)"

clean:
	rm -f $(WINDOWS_64_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_64_DIR_NAME).exe
	rm -f $(WINDOWS_32_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_32_DIR_NAME).exe
	rm -f $(LINUX_64_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_DIR_NAME)
	rm -f $(LINUX_32_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_DIR_NAME)
	rm -rf bin
	