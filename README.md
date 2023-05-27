## Contributors

Contributors is a GraphQL API written in Go. It uses MongoDB as a database. It is a simple API to manage Open Source Contributors and their contributions.

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

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#github.com/pradumnasaraf/Contributors)

### Docker Compose

Make sure you have [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed. And you have alredy cloned the repository.

Then, run the following commands to start the server. It will step up a MongoDB container and a Go API container.

```bash
docker compose up
```

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) for details.

## Security

If you discover a security vulnerability within this project, please check the [security policy](SECURITY.md) for more information.
