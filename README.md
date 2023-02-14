## GO API

A simple API to do CRUD operations on a database.

## Tech Stack

- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Gorilla Mux](github.com/gorilla/mux)

## Using and developing

### Local Development

Make sure you have [Go](https://golang.org/) and [MongoDB](https://www.mongodb.com/) installed. And you have alredy cloned the repository.

First, copy the `.env.example` file to `.env` and change the values to your own. You can use below bash command to do that.

```bash
cp .env.example .env
```

Then, run the following commands to start the server.

```bash
go mod download
go run main.go
```

### Gitpod

The easiest way to run this project in cloud with use of [Gitpod](https://www.gitpod.io/). Just click on the button below to start the project in Gitpod.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#github.com/Pradumnasaraf/go-api)

### Docker Compose

Make sure you have [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed. And you have alredy cloned the repository.   

Then, run the following commands to start the server. It will step up a MongoDB container and a Go API container.

```bash
docker compose up
``` 

## API Endpoints 

- `GET /` - Homepage
- `GET /api/movie/{id}` - Get a movie
- `GET /api/movies` - Get all movies
- `POST /api/movie` - Create a movie
- `PUT /api/movie/{id}` - Mark a movie as watched
- `DELETE /api/movie/{id}` - Delete a movie
- `DELETE /api/movies` - Delete all movies


## License 

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Security 

If you discover a security vulnerability within this project, please check the [security policy](SECURITY.md) for more information.