FROM golang:1.11-alpine AS builder
RUN apk add --no-cache git
# RUN go get github.com/mdm373/ny-data-api // get project dependencies
WORKDIR /project

COPY ./scripts/build.sh ./scripts/
COPY ./main.go .

RUN CGO_ENABLED=0 GOOS=linux ./scripts/build.sh --prod

EXPOSE 443

ENTRYPOINT ["/project/.temp/ny-data-api"]