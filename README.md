# CRUD API
## CRUD API with GO and POSTGRESQL

This is a simple product CRUD API built with Go and PostgreSQL. The API allows users to perform basic operations such as Create, Read, Update, and Delete on a PostgreSQL database.

## Setup
1. Clone the repository:
```bash
    git clone https://github.com/danialcodes/go_learnings.git
    cd go_learnings
    git checkout crud_api_server
```
2. Install dependencies:
```bash
    go mod tidy
```
3. Environment Variables:
- Create a `.env` file in the root directory and copy the contents of `.env.example` into it. Update the values as needed.
```bash
    cp .env.example .env
```
4. Seeding the database:
- Run the following command to seed the database with sample data:
```bash
    go run scripts/seed.go
```
5. Run the server:
```bash
    go run .
```


## API Endpoints

### Products

| Method | URL                    | Description                |
|--------|------------------------|----------------------------|
| GET    | `/api/v1/products`        | List all products          |
| GET    | `/api/v1/products/:id`    | Get a specific product     |
| POST   | `/api/v1/products`        | Create a new product       |
| PUT    | `/api/v1/products/:id`    | Update an existing product |
| DELETE | `/api/v1/products/:id`    | Delete a product           |

## Data Model

### Product

```json
{
  "id": 1,
  "name": "Example Product",
  "price": 99.99,
  "quantity": 100,
  "created_at": "2025-04-20T10:30:00Z",
  "updated_at": "2025-04-20T10:30:00Z"
}
```

## API Usage Examples

### List Products

```bash
curl -X GET "http://localhost:8080/api/products?page=1&limit=10"
```

Response:
```json
{
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "price": 1299.99,
      "quantity": 10,
      "created_at": "2025-04-20T10:30:00Z",
      "updated_at": "2025-04-20T10:30:00Z"
    },
    {
      "id": 2,
      "name": "Smartphone",
      "price": 699.99,
      "quantity": 20,
      "created_at": "2025-04-20T10:30:00Z",
      "updated_at": "2025-04-20T10:30:00Z"
    }
  ],
  "total": 5,
  "page": 1,
  "limit": 10,
  "totalPages": 1
}
```

### Get a Specific Product

```bash
curl -X GET "http://localhost:8080/api/products/1"
```

Response:
```json
{
    "data": {
      "id": 1,
      "name": "Laptop",
      "price": 1299.99,
      "quantity": 10,
      "created_at": "2025-04-20T10:30:00Z",
      "updated_at": "2025-04-20T10:30:00Z"
    },
    "message": "Product retrieved successfully"
}
```

### Create a Product

```bash
curl -X POST "http://localhost:8080/api/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Product",
    "price": 49.99,
    "quantity": 25
  }'
```

Response:
```json

{       
    "data": {
            "id": 6,
            "name": "New Product",
            "price": 49.99,
            "quantity": 25,
            "created_at": "2025-04-21T15:30:00Z",
            "updated_at": "2025-04-21T15:30:00Z"
            },
	"message": "Product created successfully"}
```

### Update a Product

```bash
curl -X PUT "http://localhost:8080/api/products/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Laptop",
    "price": 1399.99,
    "quantity": 15
  }'
```

Response:
```json
{
    "data": {
        "id": 1,
        "name": "Updated Laptop",
        "price": 1399.99,
        "quantity": 15,
        "created_at": "2025-04-20T10:30:00Z",
        "updated_at": "2025-04-21T15:30:00Z"
        },
    "message": "Product updated successfully"
}
```

### Delete a Product

```bash
curl -X DELETE "http://localhost:8080/api/products/1"
```

Response:
```json
{
  "message": "Product deleted successfully"
}
```


### Using Docker

```bash
# Build docker image
docker-compose up -d
```

## Development

### Project Structure

```
├── main.go                 # Application entry point
├── .env.example            # Environment variables template
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── cmd/                    # Application commands
│   └── api/                # API command
│       └── api.go         # API entrypoint
├── internal/               # Private application code
│   ├── config/             # Configuration
│   ├── api/                # API-related code
│   │   ├── routes.go       # Route definitions
│   │   ├── handlers/       # Request handlers
│   │   └── middleware/     # Middleware
│   ├── db/                 # Database setup
│   ├── models/             # Data models
│   └── repository/         # Data access
└── scripts/                # Utility scripts
    └── seed.go             # Database seeder
```
