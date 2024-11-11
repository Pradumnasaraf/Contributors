## Contributors

**Contributors** is a GraphQL API written in Go. It stores data in MongoDB and uses Redis to cache and rate-limit requests. For monitoring, it integrates with Prometheus (Custom Metrics) and Grafana. The purpose of this API is to store open-source contributors and their contributions to various projects.

https://github.com/user-attachments/assets/49bb79f6-5d63-4ab5-b7b3-92b9bcb783db

## Using and Developing

### Local Development

#### Prerequisites

To run the project locally, ensure you have the following installed:

- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)
- [Prometheus and Grafana](https://prometheus.io/docs/visualization/grafana/) (Only if you need monitoring)

> **Note:**  
> Redis, Prometheus, and Grafana can be run using Docker. It's not feasible to run some services in Docker and others locally. For consistency, it is recommended to use Docker Compose to run all services.

#### Steps for Local Setup

1. Copy the `.env.example` file to `.env` and update the values with your own configuration. Use the following bash command:
   ```bash
   cp .env.example .env
   ```

2. Install dependencies and start the server:
   ```bash
   go mod download
   go run main.go
   ```

3. Access the API at `http://localhost:8080/`.

### Docker Compose

#### Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) (Docker Compose is included with Docker Desktop)

#### Steps for Docker Compose Setup

1. Run the following command to start the services, including MongoDB, the Go API, Prometheus, and Grafana:
   ```bash
   docker compose up
   ```

2. If you're using the latest version of Docker Compose, you can leverage the Compose Watch feature. This will automatically rebuild the application when code changes are made:
   ```bash
   docker compose up --watch
   docker compose up --build --watch
   ```

3. You can now access the following:
   - API: `http://localhost:8080/`
   - Prometheus: `http://localhost:9090/`
   - Grafana: `http://localhost:3000/`

## Using the API

You can interact with the API using the **GraphQL Playground**, which will be available at `http://localhost:8080/` when running locally. Alternatively, you can interact with the API via the `/query` endpoint.

All **Query** and **Mutation** operations are defined in the [operations.md](operations.md) file.

If you are using Grafana, you can import the pre-built dashboard from the `grafana` directory.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

## Security

If you discover a security vulnerability within this project, please refer to the [security policy](SECURITY.md) for more information.
