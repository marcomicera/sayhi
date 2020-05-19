# Step 1: building executable binary
FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/marcomicera/sayhi
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/sayhi .

# Step 2: building a small image
FROM scratch
COPY --from=builder /go/bin/sayhi /go/bin/sayhi
EXPOSE 8080
ENTRYPOINT ["/go/bin/sayhi"]
