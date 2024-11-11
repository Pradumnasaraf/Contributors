This document contains the operations that can be performed on the GraphQL API.

## Query

Query operations are used to fetch data from the server.

### Get a contributor

- By directly passing the Argument (userId)

```graphql
query getAContributor {
  getAContributor(userId: "1") {
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

- By passing the Argument (userId) as a variable. `= "1"` is the default value.

```graphql
query getContributor($userId: String! = "1") {
  getContributor(userId: $userId) {
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

```json
{
  "userId": "1"
}
```

### Get all contributors

```graphql
query getAllContributors {
  getAllContributors {
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

### Mutation

Mutation operations are used to modify data on the server.

### Add a contributor - without contributions

```graphql
mutation addAContributor {
  addAContributor(
    input: {
      githubUsername: "Pradumnasaraf"
      name: "Pradumna Saraf"
      email: "pradumnasaraf@gmail.com"
    }
  ) {
    userId
    githubUsername
    name
    email
  }
}
```

### Add a contributor - with contributions

```graphql
mutation addContributor_contributions {
  addAContributor(
    input: {
      githubUsername: "Pradumnasaraf"
      name: "Pradumna Saraf"
      email: "pradumnasaraf@gmail.com"
      contributions: {
        projectName: "Pradumnasaraf/DevOps"
        type: "code"
        date: "2023"
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

### Update a contributor

```graphql
mutation updateAConbributor {
  updateAContributor(
    userId: "UPradumnasaraf"
    input: {
      name: "Pradumna Saraf"
      email: "pradumnasaraf@gmail.com"
      githubUsername: "Pradumnasaraf"
    }
  ) {
    userId
    githubUsername
    name
    email
  }
}
```

### Delete a contributor

```graphql
mutation DeleteAContributor {
  deleteAContributor(userId: "UPradumnasaraf") {
    userId
  }
}
```

### Add a contribution by userId

```graphql
mutation addAcontributions {
  addAContribution(
    userId: "UPradumnasaraf"
    input: { projectName: "UPradumnasaraf/DevOps", type: "code", date: "2023" }
  ) {
    contributionId
  }
}
```

### Delete a contribution by ContributionId

```graphql
mutation deleteContribution {
  deleteAContribution(userId: "UPradumnasaraf", contributionId: "CPradumnasaraf/DevOps") {
    contributionId
  }
}
```
