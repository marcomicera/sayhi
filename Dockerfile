# Step 1: building executable binary
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o sayhi .

# Step 2: building a small image
FROM scratch
COPY --from=builder /app/sayhi /go/bin/sayhi
EXPOSE 8080
ENTRYPOINT ["/go/bin/sayhi"]
