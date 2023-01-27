FROM golang:1.19-alpine as builder

WORKDIR /go/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go

FROM gcr.io/distroless/base
COPY --from=builder /go/app/ .

# Run the service binary.
CMD ["/app"]