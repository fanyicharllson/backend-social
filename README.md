# Social Backend (Go)

A professional-grade RESTful API built with **Go (Golang)**. This project demonstrates how to build a scalable backend using **Clean Layered Architecture** and the **Repository Pattern**, moving away from "magic" frameworks and focusing on standard library performance.

## 🚀 Key Features
- **Clean Architecture:** Separation of concerns between Handlers (Controllers), Services (Business Logic), and Repositories (Data Access).
- **Concurrency:** Utilizing Goroutines and Channels for non-blocking background tasks (e.g., notifications, analytics).
- **Database Persistence:** Robust integration with **PostgreSQL** using the `database/sql` package.
- **Environment Management:** Secure configuration using `.env` files.
- **RESTful Design:** Standardized JSON encoding/decoding for seamless frontend integration.

## 🏗️ Architecture
This project follows the **Clean Layered Architecture** pattern to ensure testability and maintainability:
- **cmd/api**: The entry point of the application.
- **internal/store**: Data access layer (Repository Pattern) for PostgreSQL.
- **internal/service**: Business logic and orchestrations.
- **internal/handler**: HTTP routing and JSON request/response handling.

## 🛠️ Tech Stack
- **Language:** Go (1.22+)
- **Database:** PostgreSQL
- **Routing:** [Chi Router](https://github.com) (or net/http)
- **Dependency Management:** Go Modules

## 🚦 Getting Started

### Prerequisites
- Go installed (1.22 or higher)
- PostgreSQL running (Local or Docker)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/fanyicharllson/backend-social
   cd social-backend
