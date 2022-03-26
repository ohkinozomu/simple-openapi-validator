# simple-openapi-validator

## Usage

```sh
$ simple-openapi-validator validate -f [file]
```

## Motivation

There are already great tools like https://github.com/IBM/openapi-validator and https://github.com/Redocly/openapi-cli, but I wanted a simple tool that only inspects for compliance with the spec.

## Supported OpenAPI version

v3.0, v3.1

## License

Apache-2.0

- `pkg/validator/schemas/v3.0.json` was from https://github.com/OAI/OpenAPI-Specification/blob/a1facce1b3621df3630cb692e9fbe18a7612ea6d/schemas/v3.0/schema.json

- `pkg/validator/schemas/v3.1.json` was from https://github.com/OAI/OpenAPI-Specification/blob/a1facce1b3621df3630cb692e9fbe18a7612ea6d/schemas/v3.1/schema.json

- `test/non-oauth-scopes.json` was from https://github.com/OAI/OpenAPI-Specification/blob/a1facce1b3621df3630cb692e9fbe18a7612ea6d/examples/v3.1/non-oauth-scopes.json

- `test/petstore.json` was from https://github.com/OAI/OpenAPI-Specification/blob/a1facce1b3621df3630cb692e9fbe18a7612ea6d/examples/v3.0/petstore.json

- `test/petstore.yaml` was from https://github.com/OAI/OpenAPI-Specification/blob/a1facce1b3621df3630cb692e9fbe18a7612ea6d/examples/v3.0/petstore.yaml