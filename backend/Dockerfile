FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/rest ./cmd/rest/main.go

ENTRYPOINT [ "/app/rest" ]