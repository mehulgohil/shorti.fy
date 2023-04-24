# shorti.fy - Writer

This is the writer microservice, that will help to register and shorten the long url

## Get Started

### Set Env Variables
```shell
export APP_PORT=80
export DYNAMO_DB_URL=http://localhost:8000
```

### Backend Setup
```shell
go mod download
go run .
```

### Database Setup
We've used DynamoDB as the database to store the URLs.
Initial Configuration will require you to set up the DynamoDB in your local.

1. Install Docker Desktop
2. ```shell
   docker run -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb
   ```
   