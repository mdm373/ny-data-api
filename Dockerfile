FROM golang:1.11-alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache curl
ADD ./app /go/src/github.com/mdm373/ny-data-api/app
ADD ./scripts /go/src/github.com/mdm373/ny-data-api/scripts
ADD ./static ./static

RUN /go/src/github.com/mdm373/ny-data-api/scripts/go-get-deps.sh
RUN CGO_ENABLED=0 GOOS=linux /go/src/github.com/mdm373/ny-data-api/scripts/go-build-app.sh --prod

EXPOSE 8000

ENTRYPOINT ["/go/src/github.com/mdm373/ny-data-api/scripts/bin-run.sh"]