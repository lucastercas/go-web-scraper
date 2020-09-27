FROM golang:1.15.2-buster
ADD ./cmd/ /go/src/
RUN go install web-scraper
ENTRYPOINT [ "/go/bin/web-scraper" ]