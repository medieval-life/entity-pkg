PHONY:=help

help:
	@echo "Pass an argument"

generate_proto:
	@protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative grpc/health/health.proto
	@protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative grpc/entity/mob.proto

format:
	@go fmt ./...
