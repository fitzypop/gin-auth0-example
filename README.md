# Gin Auth0 Example API

WIP as of Oct 7, 2023.

Goal: Make a functional Gin API with Auth0, Docker, and Postgres.

## Docs and/or Resources

[The doc I followed for project structure](https://github.com/golang-standards/project-layout)

I'm ignoring the `/cmd` structure, because this is not a true "executable". The goal is to run this api in docker, not natively.

Code that could be used by 3rd party apps should be in `/pkg/<pkg_name>`.

Any code that won't be shared can live in `/internal/<name>/...`.

`/configs` should be self-explaining.

Any OpenAPI Specs, JSON Schemes or protocol definitions should live in `/api`.

[Some Example I Found](https://www.golinuxcloud.com/golang-gin/)
