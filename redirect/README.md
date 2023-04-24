# shorti.fy - Redirect

This is the redirect microservice, that will redirect the short url to long url

## Get Started

### Set Env Variables
```shell
export APP_PORT=80
export DYNAMO_DB_URL=http://localhost:8000 #if dynamodb is running on local
export REDIS_HOST=localhost:6379
export REDIS_PASSWORD=
export AWS_ACCESS_KEY_ID=<AWS ACCESS ID>
export AWS_SECRET_ACCESS_TOKEN=<AWS SECRET ACCESS TOKEN>
export AWS_REGION=ap-south-1
```

### Backend Setup
```shell
go mod download
go run .
```

### Database Setup
We've used DynamoDB as the database to store the URLs.
If you want to run dynamodb on local, follow the below steps

1. Install Docker Desktop
2. ```shell
   docker run -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb
   ```

We've used local redis to cache the data

1. Install Docker Desktop
2. ```shell
    docker run --name redis-local -p 6379:6379 -d redis
    ```