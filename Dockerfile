FROM golang:1.16.5-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o tv-market
WORKDIR /app
RUN cp /src/tv-market .
RUN mkdir -p /app/config; cp /src/config/db.env /app/config/db.env
EXPOSE 3000

CMD ["/app/tv-market"]