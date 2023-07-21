## Tes Data Required (In MongoDB)

### getAllContributor

To retrieve all contributors, you need to have the following documents in your MongoDB collection:

```json
{
  "_id": "U1",
  "githubusername": "1",
  "name": "user1",
  "email": "test@test.com",
  "contributions": []
}
```

```json
{
  "_id": "U2",
  "githubusername": "2",
  "name": "user2",
  "email": "test@test.com",
  "contributions": []
}
```

```json
{
  "_id": "U3",
  "githubusername": "3",
  "name": "user3",
  "email": "test@test.com",
  "contributions": []
}
```

### getAContributor

To retrieve a specific contributor, such as U1, you can use the document created in the above step (getAllContributor).

### addAContributor

To add a new contributor, create the following document in your MongoDB collection:

```json
{
  "_id": "U4",
  "githubusername": "4",
  "name": "user4",
  "email": "test@test.com",
  "contributions": []
}
```

```json
{
  "_id": "U5",
  "githubusername": "5",
  "name": "user5",
  "email": "test@test.com",
  "contributions": []
}
```

After adding these documents, make sure to delete them as mentioned in the next sections.

### addAContribution

To add a contribution for a contributor (e.g., U6), use the following document in your MongoDB collection:

```json
{
  "_id": "U6",
  "githubusername": "6",
  "name": "user6",
  "email": "test@test.com",
  "contributions": []
}
```

### deleteAContributor

To delete a contributor (e.g., U7), ensure that you create the following document in your MongoDB collection:

```json
{
  "_id": "U7",
  "githubusername": "7",
  "name": "user7",
  "email": "test@test.com",
  "contributions": []
}
```

### deleteAContribution

To delete a contribution for a contributor (e.g., U8), create the following document in your MongoDB collection:

```json
{
  "_id": "U8",
  "githubusername": "8",
  "name": "user8",
  "email": "test@test.com",
  "contributions": [
    {
      "contributionid": "C1",
      "projectname": "1",
      "type": "code",
      "date": "2021"
    }
  ]
}
```

### updateAContributor

To update a contributor (e.g., U9), create the following document in your MongoDB collection:

```json
{
  "_id": "U9",
  "githubusername": "9",
  "name": "user9",
  "email": "test@test.com",
  "contributions": []
}
```

Or you can import the test data directly in MongoDB from the file `tests/testdata.json`.