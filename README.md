# `sayhi`: a simple Golang HTTP service

### Run the web service
- Locally:
    ```bash
    make run [PORT=12345]
    ```
- In a container:
    ```bash
    make run-image [PORT=12345] 
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
