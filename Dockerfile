FROM golang:1.15
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build ./cmd/main.go
CMD ["/app/main"]