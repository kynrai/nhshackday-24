$(shell cp -n .env.example .env)
include .env
export

all:
	@go run main.go
.PHONY: all