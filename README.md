### To start with default settings:

- put file named `file.json` into `_data` directory (file name can be changed by via `IMPORT_FILE` environment variable for `client-api` service in `docker-compose.yml`)
- run `make start`, or `docker-compose up --build`. This will build and start both containers + container with mongodb 
- client-api service will do json import on start and provide a rest service with GET /port/{port-id} endpoint

```shell script
curl -v http://localhost:8080/port/{port-id}

curl -v http://localhost:8080/port/VNPHG 
```

