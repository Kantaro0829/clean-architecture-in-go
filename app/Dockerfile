FROM golang:1.17.7-alpine
RUN apk update && apk add git
COPY go.mod ./app/go.mod
COPY go.sum ./app/go.sum
COPY main.go ./app/main.go
WORKDIR /app
RUN ls

CMD go run main.go