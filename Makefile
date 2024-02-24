$(shell cp -n .env.example .env)
include .env
export

all: air
.PHONY: all

air:
	@air
.PHONY: air