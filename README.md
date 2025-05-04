# CodePulse

**CodePulse** is a Go-powered microservices system that tracks developer activity (commits, PRs, reviews) and provides engineering managers with real-time insights — all via gRPC and PostgreSQL.

> Built with Go · gRPC · PostgreSQL · Docker · JWT

### Build and Start the Application

Use Docker Compose to build and run the application along with the PostgreSQL database:
docker-compose up --build
his will:

Build the Go application

Set up the PostgreSQL database using the postgres:14 image

Run the Go server on http://localhost:8080

The vendor/ directory has been included to ensure that all dependencies are bundled with the project. When you build the Docker image, the Go build process will use the dependencies from the vendor/ directo

### 🔍 Features

- Track commits, PRs, and code reviews
- JWT-authenticated user service
- Developer performance analytics
- Concurrency using goroutines
- Clean gRPC service-to-service communication

---

## 🛠 Services

| Service          | Description                            |
| ---------------- | -------------------------------------- |
| user-service     | Auth + developer profiles (JWT-based)  |
| activity-service | Records developer activity events      |
| insights-service | Serves analytics and leaderboards      |
| event-ingestor   | Sends fake GitHub events (For Testing) |

## 🚀 Stack

- Language: Go 1.21+
- DB: PostgreSQL
- Protocol: gRPC
- Auth: JWT
- Containerized: Docker + docker-compose

# CodePulse

**CodePulse** is a Go-powered microservices system that tracks developer activity (commits, PRs, reviews) and provides engineering managers with real-time insights — all via gRPC and PostgreSQL.

### 🔍 Features

- Track commits, PRs, and code reviews
- JWT-authenticated user service
- Developer performance analytics
- Concurrency using goroutines
- Clean gRPC service-to-service communication

---

## 🛠 Services

| Service          | Description                            |
| ---------------- | -------------------------------------- |
| user-service     | Auth + Developer profiles (JWT-based)  |
| activity-service | Records developer activity events      |
| insights-service | Serves analytics and leaderboards      |
| event-ingestor   | Sends fake GitHub events (for testing) |

---

## ⚙️ Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-repo/CodePulse.git
   cd CodePulse
   ```

2. **Install Docker and Docker Compose** (if not already installed).

3. **Start services using Docker Compose**:

   ```bash
   docker-compose up --build
   ```

   This will set up the database and the main application, which will be accessible at `localhost:8080`.

---

## 🛠 Development

- **User Service**: Handles user registration and authentication via JWT.
- **Activity Service**: Tracks commits, PRs, and reviews, and stores them in PostgreSQL.
- **Insights Service**: Provides analytics and performance metrics for developers.

### Database Schema

The database has tables for:

- **Developers**: Stores user information.
- **Commits**: Records developer commits.
- **PRs**: Tracks pull requests .
- **Reviews**: Stores pull request reviews.

### gRPC Communication

All services communicate over gRPC for efficiency and scalability.

### Concurrency

The system processes data concurrently using Go’s goroutines for improved performance.

---

## 📝 License

MIT License.
