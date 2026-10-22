# ApiRedisCache

A small Go REST API that uses Redis as a cache in front of product data, to reduce repeated reads and speed up responses.

Built with [Gin](https://github.com/gin-gonic/gin) and Redis.

## How it works

- `GET` product listing checks Redis for a cached copy under the `getAllProducts` key.
- A separate `getAllProducts:validation` key tracks whether the cache is still valid.
- If the cache is missing or stale, the handler reloads the products from the database, refreshes the Redis entry, and returns the fresh data.
- `POST` creates a new product and stores it in the database.

## Project structure

```
.
├── controllers/
│   └── products.go      # HTTP handlers: GetAllProducts, CreateProduct
├── db/
│   ├── client.go        # Database connection
│   └── redis/
│       └── client.go    # Redis client setup
├── models/
│   └── product.go       # Product model and CreateProductInput
├── main.go              # Entry point, router setup
├── docker-compose.yml   # Local services (API + Redis)
├── test.db              # Local SQLite database file
├── go.mod
└── go.sum
```

## Requirements

- Go 1.21+ (adjust to match `go.mod`)
- Redis
- Docker and Docker Compose (optional, for running everything together)

## Getting started

### With Docker Compose

```bash
docker-compose up --build
```

### Without Docker

1. Start a Redis instance locally, for example:
```bash
   docker run -p 6379:6379 redis:7
```
2. Install dependencies and run the API:
```bash
   go mod download
   go run main.go
```

## API

| Method | Path        | Description                          |
|--------|-------------|--------------------------------------|
| GET    | `/products` | List all products (served from cache when valid) |
| POST   | `/products` | Create a product                     |

Example request body for `POST /products`:

```json
{
  "name": "Example product",
  "value": 10.5
}
```

## Configuration

<!-- Fill in from docker-compose.yml / db/redis/client.go -->

| Variable     | Default          | Description           |
|--------------|------------------|-----------------------|
| `REDIS_ADDR` | `localhost:6379` | Redis address         |
| `PORT`       | `8080`           | HTTP port for the API |

## Notes

- The cache is refreshed when it is stale, so data can be briefly out of date between refreshes.
- `test.db` is a local development database and should not be committed for production use.

## License

<!-- Add a license -->
