# shorti.fy - Redirect

This is the redirect microservice, that will redirect the short url to long url

## Get Started
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
   