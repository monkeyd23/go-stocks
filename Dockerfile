FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .
RUN apk add git
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/stocks
ENTRYPOINT ["/go/bin/stocks"]

# Runner
FROM alpine
COPY --from=builder /go/bin/stocks /go/bin/stocks
ENTRYPOINT ["/go/bin/stocks"]