.PHONY: run

run:
	go run ./cmd/web-parser/main.go https://www.google.com/

build_docker:
	docker build -t lucastercas/web-scraper .