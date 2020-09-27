FROM golang
ADD ./cmd/ /go/src/
RUN go install web-scraper
ENTRYPOINT [ "/go/bin/web-scraper" ]