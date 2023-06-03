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

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#github.com/Pradumnasaraf/Contributors)

### Docker Compose

Make sure you have [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed. And you have alredy cloned the repository.

Then, run the following commands to start the server. It will step up a MongoDB container and a Go API container.

```bash
docker compose up
```

## Using the API

You can use the GraphQL Playground to interact with the API. It will be available at `http://localhost:8080/` if you are running the API locally.

### Queries

#### Add a new contributor

The `userId` for a contributor is the `U` + `githubUsername`. For example, if the `githubUsername` is `example`, then the `userId` will be `Uexample`. This is done to make sure that the `userId` is unique. `contributions` is optional in the input.

```graphql
mutation {
  addAContributor(
    input: {
      githubUsername: "example"
      name: "Pradumna Saraf"
      email: "example@example.com"
      contributions: {
        projectName: "example"
        type: "code"
        date: "2021-09-01"
      }
    }
  ) {
    userId
    githubUsername
    name
    email
    contributions {
      contributionId
      projectName
      type
      date
    }
  }
}
```

#### Get all contributors

```graphql
query {
  getAllContributors {
    userId
    githubUsername
    name
    email
  }
}
```

#### Get a contributor by userId

```graphql
query {
  getAContributor(userId: "Uexample") {
    userId
    githubUsername
    name
    email
  }
}
```

#### Update a contributor by userId

In update we don't update the `userId` if `githubUsername` is changed. This is done to make sure that the `userId` is always remain same.

```graphql
mutation {
  updateAContributor(
    userId: "Uexample"
    input: {
      name: "example"
      email: "example@example.com"
      githubUsername: "example"
    }
  ) {
    userId
    githubUsername
    name
    email
  }
}
```

#### Delete a contributor by userId

```graphql
mutation {
  deleteAContributor(userId: "Uexample") {
    userId
    githubUsername
    name
    email
  }
}
```

#### Add a new contribution

The `contributionId` for a contribution is the `C` + `repositoryName`. For example, if the `repositoryName` is `example`, then the `contributionId` will be `Cexample`. This is done to make sure that the `contributionId` is unique.

```graphql
mutation {
  addAContribution(
    userId: "Uexample"
    input: { projectName: "example", type: "code", date: "2021-09-01" }
  ) {
    contributionId
  }
}
```

#### Delete a contribution

```graphql
mutation {
  deleteAContribution(contributionId: "Cexample") {
    contributionId
  }
}
```

####

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) for details.

## Security

If you discover a security vulnerability within this project, please check the [security policy](SECURITY.md) for more information.
