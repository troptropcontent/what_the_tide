FROM golang:1.22

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

COPY . .

ENV GO_ENV=production

RUN go build -o /app/web cmd/web/main.go

EXPOSE 3001
CMD ["/app/web"]