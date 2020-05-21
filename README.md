# `sayhi`: a simple Golang HTTP service

### Getting started

First, you can optionally se the web service port (defaults to `8080`):
```bash
PORT=12345
```

You can launch the web service
- locally:
    ```bash
    make run [PORT=$PORT]
    ```
- in a container:
    ```bash
    make run-image [PORT=$PORT] 
    ```

Look at the web service saying hi!
```
$ curl localhost:${PORT}/helloworld?name=AlfredENeumann
Hello Alfred E Neumann
```

### Run tests
```
make test 
```

##### Override environment variables
Simply edit [`config.env`](config.env) like so:
```dotenv
TEST_ENV_VAR=a non-significant value
```
