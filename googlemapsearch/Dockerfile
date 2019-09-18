FROM golang:1.13.0-alpine3.10

COPY ./go.mod /googlemapsearch/go.mod
COPY ./app /googlemapsearch/app

RUN apk add --update --no-cache gcc git bind-dev && rm -rf /var/cache/apk/*
RUN go get -u github.com/oxequa/realize

WORKDIR /googlemapsearch/app

RUN CGO_ENABLED=0 go mod tidy
