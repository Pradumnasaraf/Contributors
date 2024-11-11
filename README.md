## Contributors

Contributors is a GraphQL API written in Go. It stores data in MongoDB and uses Redis to cache and rate-limit requests. For monitoring, it uses Prometheus (Custom Metrics) and Grafana. The purpose of this API is store open source contributors and their contributions to different projects.





## Using and developing

### Local Development

#### Prerequisites

- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)
- [Prometheus and Grafana](https://prometheus.io/docs/visualization/grafana/) - Only if you need monitoring.

> [Note]
> Redis, Prometheus and Grafana can we run using Docker. It's not feasible to run half of services in Docker and half of them locally. So, it's better use Docker Compose to run all the services.


First, copy the `.env.example` file to `.env` and change the values to your own. You can use below bash command to do that.

```bash
cp .env.example .env
```

Then, run the following commands to start the server.

```bash
go mod download
go run main.go
```

Now, you can access the API at `http://localhost:8080/`.

### Docker Compose

#### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

If you are using Docker for Desktop, Compose comes with it. 

Then, run the following commands to start the server. It will step up a MongoDB container and a Go API container. Monitoring is also configured with Prometheus and Grafana.

```bash
docker compose up
```

If you using lases version of Docker Compose, you can use Compose Watch feature, it will automatically rebuild the application when you make changes to the code.

```bash
docker compose up --watch
docker compose up --build --watch
```

Now, you can access the API at `http://localhost:8080/`. Prometheus is available at `http://localhost:9090/` and Grafana is available at `http://localhost:3000/`.

## Using the API

You can use the GraphQL Playground to interact with the API. It will be available at `http://localhost:8080/` if you are running the API locally. Otherwise you can use `/query` endpoint to interact with the API. All the **Query** and **Mutation** operations are defined in the [operation.md](operation.md) file.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) for details.

## Security

If you discover a security vulnerability within this project, please check the [security policy](SECURITY.md) for more information.
