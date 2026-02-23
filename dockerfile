FROM golang:1.24 AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /kannacraftbot ./cmd/main.go

FROM alpine:latest

COPY --from=build /kannacraftbot /kannacraftbot

CMD ["/kannacraftbot"]
