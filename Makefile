.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot/vk .

start-container:
	docker run --env-file .env -p 8080:8080 telegram-bot/vk