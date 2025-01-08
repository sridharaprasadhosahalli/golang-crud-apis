# Building and Deploying a Simple CRUD API with Golang,Docker, and Kubernetes

This is a simple CRUD (Create, Read, Update, Delete) REST API in Go using the Gin web framework.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

- Go (at least version 1.23.4)

### Cloning the Project

```bash
git clone https://github.com/sridharaprasadhosahalli/golang-crud-apis.git
cd golang-crud-apis
```

### Initializing the Project

```bash
go mod init github.com/sridharaprasadhosahalli/golang-crud-apis
```

### Downloading Dependencies

```bash
go mod tidy
```

## Running the Application

```bash
go run main.go
```

The application will start running at `http://localhost:8080`.


## API Endpoints

- **GET /items**: Fetch all items.
- **GET /items/:id**: Fetch a item by ID.
- **POST /items**: Create a new item. 
- **PUT /items/:id**: Update a item by ID.
- **DELETE /items/:id**: Delete a item by ID.

## Testing the API

You can use a tool like curl or Postman to test the API. Here are some examples:

- Fetch all items:

curl -X GET http://localhost:8080/items

- Fetch a item by ID:

curl -X GET http://localhost:8080/items/1


- Create a new item:

 curl http://localhost:8080/items \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","name": "potato","type": "vegetable","price": 49.99}'


- Update a item by ID:

curl -X PUT -H "Content-Type: application/json" -d '{"id": "4","name": "cabbage","type": "vegetable","price": 49.99}' http:/localhost:8080/items/3


- Delete a item by ID:

curl -X DELETE http://localhost:8080/items/3
