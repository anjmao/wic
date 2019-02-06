FROM golang:1.11 AS builder
WORKDIR /src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /src/app/do ./cmd/server/*.go

FROM alpine
COPY --from=builder /src/app/do /src/app/do
COPY --from=builder /src/app/certs /src/app/certs
WORKDIR /src/app
USER nobody:nobody
ENTRYPOINT ["./do"]