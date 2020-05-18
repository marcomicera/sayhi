# `sayhi`: a simple Golang HTTP service

## Run locally
```
PORT=12345
make run [PORT=$PORT]
```

## Run in container
```
make image
PORT=12345
docker run --rm -p $PORT:8080 marcomicera/cyclapp
```