## API for transactions

## Description

This API is used to save and read transactions.

## Functionality
+ Saving transactions of three types: receiving ballot, voting and counting results
+ Getting all transactions
+ Getting transactions by time period
+ Getting result by vote transactions

## Getting Started

1. Make sure you have Go installed. If not, you can download it from [official website](https://golang.org/dl/)

2. Install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

2. Clone the repository to your computer:
```
https://github.com/fur1ouswolf/transactions-api
```

3. Run the application. Execute the command `make run`. By default, the project will be available at `http://localhost:8080`

## Endpoints
- `POST /transaction/ballot` - save ballot transaction
- `POST /transaction/vote` - save vote transaction
- `POST /transaction/result` - save result transaction
- `GET /transactions` - get all transactions
- `GET /transactions/time` - get transactions by time period
- `GET /results` - get result by vote transactions

For each request, the API expects a JSON object with data in the request body and returns a JSON response with the result of the operation.

#### Examples
Use Postman or any other tool to send requests to the API.
##### Saving ballot transaction
```
POST http://localhost:8080/api/v1/transaction/ballot

{
    "ballot_id": 1,
    "signature": "12345678",
}
```

##### Saving vote transaction
```
POST http://localhost:8080/api/v1/transaction/vote

{
    "ballot_id": 1,
    "candidate_id": 1,
    "signature": "12345678",
}
```

##### Saving result transaction
```
POST http://localhost:8080/api/v1/transaction/result

{
    "candidat_id": 1,
    "votes_count": 1,
}
```

##### Getting all transactions
```
GET http://localhost:8080/api/v1/transactions
```

##### Getting transactions by time period
```
GET http://localhost:8080/api/v1/transactions/time

{
    "start_time": "2023-01-01T00:00:00.000UTC",
    "end_time": "2023-12-31T23:59:59.000UTC"
}
```

##### Getting result by vote transactions
```
GET http://localhost:8080/api/v1/results
```

## License

This project is distributed under the MIT license. See the ``LICENSE`` file for details.