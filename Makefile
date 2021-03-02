development:
	@go build -o bin/sandpiper

docker:
	@docker build -t endigma/sandpiper:latest .

compose:
	@docker-compose build
	@docker-compose up