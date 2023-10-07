# Gin Auth0 Example API

WIP as of Oct 7, 2023.

Goal: Make a functional Gin API with Auth0, Docker, and Postgres.

## Docs and/or Resources

[The doc I followed for project structure](https://github.com/golang-standards/project-layout)

The main executable lives in `/cmd/gin_auth0_example/main.go`.

Code that could be used by 3rd party apps should be in defined packages in `/pkg/<pkg_name>`.

Any code that won't be shared can live in `/internal/<name>/...`.

`/configs` should be self-explaining.

Any OpenAPI Specs, JSON Schemes or protocol definitions should live in `/api`.
