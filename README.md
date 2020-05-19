# `sayhi`: a simple Golang HTTP service

First, define the web service port (defaults to `8080`):
```
PORT=12345
```

## Run locally
```
make run [PORT=$PORT]
```

## Run in container
```
make run-image [PORT=$PORT] 
```
