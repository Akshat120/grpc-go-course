# Makefile for gRPC Go Course

# Variables
MODULE_NAME = github.com/Akshat120/grpc-go-course

PROJECT ?= greet
GREET_PROTO_DIR = ./$(PROJECT)/proto
GREET_PROTO_FILES = $(wildcard $(GREET_PROTO_DIR)/*.proto)

# generate proto files
.PHONY: gen-proto
gen-proto:
	@echo "=== generating proto files ==="
	protoc --go_out=. --go_opt=module=$(MODULE_NAME) \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE_NAME) \
		$(GREET_PROTO_FILES)

.PHONY: clean
clean:
	@echo "=== cleaning proto files from $(PROJECT) ==="
	@find ./$(PROJECT)/proto -name "*.pb.go" -type f -delete

.PHONY: clean-all
clean-all:
	@echo "=== cleaning all proto files ==="
	@find . -name "*.pb.go" -type f -delete




