# Wex test

# Transaction Management Application

This application is designed to handle financial transactions, including storing transaction details and retrieving them with the option of currency conversion based on the Treasury Reporting Rates of Exchange API.

## Project Structure
The project is organized into several packages, each with a specific responsibility:

### cmd/main.go
The entry point of the application. It sets up and starts the HTTP server, routes, and integrates all application components.

### pkg/transactions
Contains the core business logic for handling transactions.

- `transaction.go`: Defines the Transaction struct and any methods related to it.
- `service.go`: Contains the Service struct with methods for storing and retrieving.
- `transactions.go` It interacts with the repository and the currency converter.
- `repository.go`: Defines the Repository interface and includes an in-memory implementation for storing transactions.

### pkg/currency
Handles currency conversion logic.

- `converter.go`: Contains the Converter struct with methods to perform currency conversions using exchange rates from an external API.

### pkg/api
Responsible for external API interactions.

- `client.go`: Defines the Client struct with methods to interact with external APIs, such as fetching currency exchange rates.

### pkg/utils
Includes utility functions and custom error definitions.

- `errors.go`: Contains custom error types and utility functions related to error handling.


## Installation and Running

```Bash
go build ./cmd/main.go
./main
```

## API Endpoints
Describe the available HTTP endpoints, their methods, request parameters, and expected responses. For example:

- `POST /store`: Stores a new transaction.
- `GET /retrieve/{id}`: Retrieves a transaction by ID with no conversion.
- `GET /retrieve/{id}/{targetCountry}`: Retrieves a transaction by ID and converts it to the specified currency.