# Payment Gateway Exercise - Authentication Service

Please check out documentation for the whole system [here](https://github.com/gustavooferreira/pgw-docs).

This repository structure follows [this convention](https://github.com/golang-standards/project-layout).

---

# Build

To build a binary, run:

```bash
make build
```

The `api-server` binary will be placed inside the `bin/` folder.

To build the docker image, run:

```bash
make build-docker
```

The docker image is called `pgw/auth-api-server`.

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

## Tip

> If you run `make` without any targets, it will display all options available on the makefile followed by a short description.

# Design

This service serves as a light dependency for the payment gateway service. It does not perform authorization, only supports authentication.

This service reads credentials from a yaml file and has only one endpoint that can be used by the payment gateway service to validate whether a merchant has provided valid credentials.

It's not meant to be production ready by any means. I'm not using a database to simplify the service, as this service was only created so the payment gateway can simulate talking to an external system for credentials validation.

The OpenAPI spec is located in the `openapi` folder. The API is made up of one endpoint only.

To view the spec in Swagger UI, click [this link](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/gustavooferreira/pgw-auth-service/master/openapi/spec.yaml).

This service should require authentication/authorization. I have not implemented this to keep the service simple.
