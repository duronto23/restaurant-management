# Restaurant Management System

A **Restaurant Management System** written in **Go**, designed to provide a modular and maintainable backend for managing menus, orders, customers, and related restaurant data.

## Features

✔ **Menu Management**

* Create, read, update, delete menu items

✔ **Order Handling**

* Place and track customer orders

✔ **Customer / User Handling**

* Manage customers or user sessions (depending on implementation)

✔ **Modular Structure**

* Clear separation of controllers, models, middleware, and routes

✔ **REST API**

* HTTP API endpoints for all key actions


## Project Structure

````
restaurant-management/
├── controllers/      # API handler logic
├── database/         # Database connection & setup
├── helpers/          # Utility and helper functions
├── middleware/       # Middlewares (auth, logging, etc.)
├── models/           # Database models
├── routes/           # Router and route definitions
├── main.go           # Application entry point
├── go.mod            # Go module file
├── go.sum            # Dependency lock file
└── README.md
````

## Tech Stack

- **Language:** Go  
- **Routing:** `net/http` / Gin / another Go HTTP framework  
- **Database:** (depends on your code/config — could be MySQL, PostgreSQL, SQLite, or MongoDB)

## Requirements

Make sure your system has the following installed:

- Go (1.18 or later recommended)
- Database server (MySQL / PostgreSQL / MongoDB as used in the code)
- Git


## Installation & Running

### 1. **Clone the repository**

```bash
git clone https://github.com/duronto23/restaurant-management.git
cd restaurant-management
````

### 2. **Install dependencies**

```bash
go mod tidy
```

### 3. **Configure environment**

Create a `.env` file (or similar configuration file) with values for:

```
PORT=8080
DB_HOST=…
DB_PORT=…
DB_USER=…
DB_PASS=…
DB_NAME=…
```

### 4. **Run the server**

```bash
go run main.go
```

Your API will start on the configured port (e.g., `http://localhost:8080`).


## API Endpoints

| Method | Endpoint             | Description            |
| ------ | -------------------- | ---------------------- |
| GET    | `/menus`             | List all menu items    |
| POST   | `/menus`             | Create a new menu item |
| GET    | `/orders`            | List all orders        |
| POST   | `/orders`            | Create a new order     |
| GET    | `/customers/:id`     | Get customer by ID     |
| PUT    | `/orders/:id/status` | Update an order status |


## Future Development

* Add **authentication** (JWT or session-based)
* Add **validation** for inputs
* Add **unit tests** for handlers and services
* Document API with **Swagger/OpenAPI**

## Contributing

Contributions, suggestions, and feedback are welcome!
Open an issue or submit a pull request to improve features or documentation.

