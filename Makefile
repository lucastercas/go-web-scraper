.PHONY: run

run:
	go run ./cmd/web-parser/main.go

docker-build:
	docker build -t lucastercas/web-scraper .