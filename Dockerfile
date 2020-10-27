# First Stage
FROM golang:alpine as base

RUN apk update && apk add --no-cache git

RUN mkdir /src/app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux

COPY . /app

# Air for hot reload
RUN go get -u github.com/cosmtrek/air

# Last Stage
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=base /app .

EXPOSE 8080

CMD ["./server.go"]