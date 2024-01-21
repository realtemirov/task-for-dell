# Blog and News API
This project is a Golang-based API using the Echo framework and Postgres database for managing and serving Blogs and News content.

## Swagger UI - http://135.181.88.180:8000/swagger/index

#### Table of Contents
* [Features](#features)
* [Prerequisites](#prerequisites)
* [Tools & Technologies](#tool--technologies)
* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)
* [API Endpoints](#api-endpoints)
* [License](#license)
* [Feedback and Support](#feedback-and-support)

## Features
* Create, read, update, and delete blogs and news articles.
* Postgres yordamida doimiy saqlash.
* Swagger UI.
* Docker.
* CI/CD by Github Actions.

## Prerequisites
* Golang - [Installation Guide](https://golang.org/doc/install)
* Docker - [Installation Guide](https://docs.docker.com/engine/install/)

## Tool & Technologies
List of tools and technologies used:
* [echo](https://github.com/labstack/echo) - Web framework
* [swag](https://github.com/swaggo/swag) - Swagger
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [gomock](https://github.com/golang/mock) - Mocking framework
* [testing](https://pkg.go.dev/testing) - Testing
* [require](https://github.com/stretchr/testify/require) - Checking test result
* [Docker](https://www.docker.com/) - Docker
* [Database](https://www.postgresql.org/) - PostgreSQL


## Installation
### 1. Clone the repository:
```bash
git clone https://github.com/realtemirov/task-for-dell.git
cd task-for-dell
```

### 2. Install dependencies:
```bash
go mod download
```
## Configuration
Before running the application, configure the necessary environment variables. Enter the configuration folder and configure the environment you want.
```bash
cd config
nano config-local.yml
```

## Usage
### 1. Run the application with `docker-compose`:
```bash
docker compose up -d            // run containers with docker-compose
```

### 2. Run the application with `container`:
```bash
make start                      // run postgres container and migration-up
make run                        // run app
```

### 3. Run Local app:
```bash
go run cmd/main.go              // equal -> make run
```
### 4. Run Test:
```bash
make test
```

### 5. Swagger UI
http://localhost:8000/swagger/index.html

## API Endpoints
* ### Create a Blog/News Content
  **`POST` /v1/blogs**
  ```json
  {
    "title": "Sample Title",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
  }
  ```

  **`POST` /v1/news**
  ```json
  {
    "title": "Sample Title",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
  }
  ```
* ### Update Content by ID
  **`PUT` /v1/blogs/:id**
  ```json
  {
    "title": "Sample Title",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
  }
  ```

  **`PUT` /v1/news/:id**
  ```json
  {
    "title": "Sample Title",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
  }

* ### Delete Content by ID
  **`DELETE` /v1/blogs/:id**

  **`DELETE` /v1/news/:id**

* ### Get Content by ID
  **`GET` /v1/blogs/:id**

  **`GET` /v1/news/:id**

* ### GetAll Contents
  **`GET` /v1/blogs**

  **`GET` /v1/news**

## License
This project is licensed under the [MIT License](./LICENSE).

## Feedback and Support
For any issues, feedback, or support, please [open an issue](https://github.com/realtemirov/task-for-dell/issues) on GitHub.
