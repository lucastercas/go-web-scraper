.PHONY: run

project_name="web-scraper"

run:
	go run ./cmd/$(project_name)/main.go

docker-build:
	docker build -t lucastercas/$(project_name) .

docker-push:
	docker push lucastercas/$(project_name)

docker: docker-build docker-push