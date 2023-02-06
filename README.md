# Cookiecutter Golang API

A base template for creating a API with golang. Uses [GIN](https://gin-gonic.com/) and [GORM](https://gorm.io/) packages.

## Installation

```
cookiecutter https://github.com/Mr-Destructive/cookiecutter-golang-api
```

```
$ cookiecutter https://github.com/Mr-Destructive/cookiecutter-golang-api

project_name [go_test_api]:
project_slug [core]:
description [GO GIN API Template]:
author_name [Someone]:
domain_name [example.com]:
email [meet-gor@example.com]:
version [0.0.1]:
```

## Features

- Base API Route (JSON)
- Template Route (HTML)
- SQLite support with gorm
- Structured project layout

## TODO
- Integrate other API template
  - Vanilla GO API
  - Gorilla MUX API
  - Fiber API
- Other Database integrations
  - PostgreSQL
  - MySQL
  - MongoDB
- Base Authentication Views
- CRUD APIs
