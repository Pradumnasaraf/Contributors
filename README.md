## GO API

A simple API to do CRUD operations on a database.

## Tech Stack

- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Gorilla Mux](github.com/gorilla/mux)

## Using the API

Make sure you have [Go](https://golang.org/) and [MongoDB](https://www.mongodb.com/) installed. And you have alredy cloned the repository.

First, copy the `.env.example` file to `.env` and change the values to your own.

```bash
cp .env.example .env
```

Then, run the following commands to start the server.

```bash
go mod download
go run main.go
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