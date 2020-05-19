##########################################
### Step 1: building executable binary ###
##########################################

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/marcomicera/sayhi

# Enabling Go modules
ENV GO111MODULE=on

# Project information is injected as build-time vars
ARG GIT_COMMIT
ARG PROJECT_NAME

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X github.com/marcomicera/sayhi/go.GitCommit=${GIT_COMMIT} -X github.com/marcomicera/sayhi/go.ProjectName=${PROJECT_NAME}" -o /go/bin/sayhi .

######################################
### Step 2: building a small image ###
######################################

FROM scratch
COPY --from=builder /go/bin/sayhi /go/bin/sayhi
EXPOSE 8080
ENTRYPOINT ["/go/bin/sayhi"]
