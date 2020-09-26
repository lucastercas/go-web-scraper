FROM golang
ADD . /go/src/web-scraper
RUN go install web-scraper
ENTRYPOINT [ "/go/bin/web-scraper" ]