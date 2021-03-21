# Payment Gateway Exercise - Authentication Service

Please check out documentation for the whole system [here](https://github.com/gustavooferreira/pgw-docs).

This repository structure follows [this convention](https://github.com/golang-standards/project-layout).

---

## Tip

> If you run `make` without any targets, it will display all options available on the makefile followed by a short description.

# Build

To build a binary, run:

```bash
make build
```

The `api-server` binary will be placed inside the `bin/` folder.

---

# Tests

To run tests:

```bash
make test
```

To get coverage:

```bash
make coverage
```

# Docker

To build the docker image, run:

```bash
make build-docker
```

The docker image is named `pgw/auth-api-server`.

Create a file with some credentials, like this:

```yaml
credentials:
  bill: pass1
  john: pass2
  adam: pass3
```

And start a docker container like this:

```bash
docker run --rm --name pgw-auth-service -p 127.0.0.1:9000:8080/tcp -v "$(pwd)"/db_creds.yaml:/db_creds.yaml:ro -e PGW_AUTH_APP_DATABASE_FILENAME=/db_creds.yaml pgw/auth-api-server
```

This assumes the yaml file created is called `db_creds.yaml` and is placed in the current directory.

Once the container is running, you can make a request like this:

```bash
curl -i -X POST http://localhost:9000/api/v1/auth -d '{"username":"bill", "password": "pass1"}'
```

# Design

This service serves as a light dependency for the payment gateway service. It does not perform authorization, only supports authentication.

This service reads credentials from a yaml file and has only one endpoint that can be used by the payment gateway service to validate whether a merchant has provided valid credentials.

It's not meant to be production ready by any means. I'm not using a database to simplify the service, as this service was only created so the payment gateway can simulate talking to an external system for credentials validation.

The OpenAPI spec is located in the `openapi` folder.

To view the spec in the Swagger UI [click this link](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/gustavooferreira/pgw-auth-service/master/openapi/spec.yaml).

The requests to this service should go through an authentication/authorization process as well. I have not implemented this to keep the service simple.
