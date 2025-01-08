# Dockerfile for Go
FROM golang:1.20

WORKDIR /app

COPY . .
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]
