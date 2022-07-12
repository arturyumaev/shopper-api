.PHONY: default

tag = shopper-api
image_name = arturyumaev/public:$(tag)

default: build run

build:
	docker build -t $(image_name) .

run:
	docker run --publish 8888:8080 $(image_name)
