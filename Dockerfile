FROM golang:1.27 AS builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o path-api .

FROM scratch

WORKDIR /app

COPY --from=builder /app/path-api .
