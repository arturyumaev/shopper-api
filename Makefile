.PHONY: default

image_name = shopper-api

default: build run

build:
	docker build -t $(image_name) .

run:
	docker run --publish 8888:8080 $(image_name)
