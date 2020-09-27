.PHONY: run

run:
	go run ./cmd/web-parser/main.go https://www.google.com/

docker-build:
	docker build -t lucastercas/web-scraper .