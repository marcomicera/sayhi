# `sayhi`: a simple Golang HTTP service

### Getting started

First, you can optionally set the web service port (defaults to `8080`):
```bash
$ PORT=12345
```

You can launch the web service
- locally:
    ```
    $ make run [PORT=$PORT]
    ```
- in a container:
    ```
    $ make run-image [PORT=$PORT] 
    ```

Look at the web service saying hi!
```
$ curl localhost:${PORT}/helloworld?name=AlfredENeumann
Hello Alfred E Neumann
```

### Run tests
```
$ make test 
```

### Override environment variables
Simply edit [`config.env`](config.env) like so:
```dotenv
TEST_ENV_VAR=a non-significant value
```

### Contribute
Take a look at the [`CONTRIBUTING.md`](CONTRIBUTING.md) file.
