Purpose of this file is to provide the information about the operations that can be performed with the GraphQL API.

## Query

### Get a contributor

- By directly passing the Argument (userId)

```graphql
query getContributor {
  getContributor(userId: "1") {
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

### Add a contributor - without contributions

```graphql
mutation addContributor {
  addAContributor(
    input: {
      githubUsername: "user1"
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
      githubUsername: "user1"
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
mutation updateConbributor {
  updateAContributor(
    userId: "UPradummnasaraf"
    input: {
      name: "Pradumna Saraf"
      email: "pradumnasaraf@gmail.com"
      githubUsername: "Pradummnasaraf"
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
mutation deleteContributor {
  deleteAContributor(userId: "UPradummnasaraf") {
    userId
  }
}
```

### Add a contribution by ContributorId

```graphql
mutation addcontributions {
  addAContribution(
    userId: "Pradumnasaraf"
    input: { projectName: "UPradumnasaraf/DevOps", type: "code", date: "2023" }
  ) {
    contributionId
  }
}
```
