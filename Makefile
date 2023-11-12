# Variables for output directory names
WINDOWS_64_NAME = windows-amd64
WINDOWS_32_NAME = windows-386
LINUX_64_NAME = linux-amd64
LINUX_32_NAME = linux-386

# Name of the output binary
OUTPUT_BINARY = workspace

# Version
VERSION = 1.0.0

# Variables for output directories
OUTPUT_DIR = bin/$(VERSION)

PATH_MAIN = cmd/workspace/main.go

.PHONY: all windows linux clean

# Processes

all: windows linux

windows: windows_64 windows_32

linux: linux_64 linux_32

windows_64:
	@GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_64_NAME).exe $(PATH_MAIN)
	@make -s message_success MESSAGE=$(WINDOWS_64_NAME)

windows_32:
	@GOOS=windows GOARCH=386 go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_32_NAME).exe $(PATH_MAIN)
	@make -s message_success MESSAGE=$(WINDOWS_32_NAME)

linux_64:
	@GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_NAME) $(PATH_MAIN)
	@chmod +x $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_NAME)
	@make -s message_success MESSAGE=$(LINUX_64_NAME)

linux_32:
	@GOOS=linux GOARCH=386 go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_NAME) $(PATH_MAIN)
	@chmod +x $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_NAME)
	@make -s message_success MESSAGE=$(LINUX_32_NAME)

message_success:
	@echo "Build successful for: $(MESSAGE)"

# zipping files

zip: clean-zip
	@for file in $(wildcard $(OUTPUT_DIR)/*); do \
		filename=$$(basename $$file); \
		zip -r $(OUTPUT_DIR)/$$filename.zip $$file; \
		echo "Zip archive created: $$filename.zip"; \
	done

clean-zip:
	@rm -f $(OUTPUT_DIR)/*.zip
	@echo "All zip files files removed"

clean-compile:
	rm -f $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_64_NAME).exe
	rm -f $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(WINDOWS_32_NAME).exe
	rm -f $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_64_NAME)
	rm -f $(OUTPUT_DIR)/$(OUTPUT_BINARY)-$(LINUX_32_NAME)
	rm -rf bin
