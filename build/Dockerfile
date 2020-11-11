FROM golang:1.14.6-alpine3.12 as builder
COPY ../go.mod go.sum /go/src/github.com/drprado2/transaction-manager/
WORKDIR github.com/drprado2/transaction-manager
RUN go mod download
COPY .. github.com/drprado2/transaction-manager
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/transaction-manager github.com/drprado2/transaction-manager

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/drprado2/transaction-manager/build/transaction-manager /usr/bin/transaction-manager
EXPOSE 9000 9000
ENTRYPOINT ["/usr/bin/transaction-manager"]