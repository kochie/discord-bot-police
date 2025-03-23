
.PHONY: dev

TAG_NAME = "discord-bot-police"

build:
	@docker build -t $(TAG_NAME) .

dev:
	@docker run -it $(TAG_NAME)