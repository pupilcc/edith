FROM golang:1.21.1 as builder
WORKDIR /go/src/edith
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/edith ./cmd/edith

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/edith .
EXPOSE 1455
CMD ["./edith"]