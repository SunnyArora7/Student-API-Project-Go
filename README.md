# 🎓 Student Project (Go REST API)

A clean, modular **Go (Golang) REST API** project built using **Clean Architecture principles**.
This project demonstrates how to structure a real-world backend service with clear separation of concerns, scalability, and maintainability.

---

## 📌 Key Features

* Clean Architecture folder structure
* SQLite database integration
* RESTful APIs for Student management
* Request validation
* Standardized JSON responses
* Easy to extend and test

---

## 📁 Project Folder Structure

```
STUDENT_PROJECT/
├── cmd/
│   └── student-project/
│       └── main.go
│
├── config/
│   └── local.yaml
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go
│   │
│   ├── storage/
│   │   ├── Sqllite/
│   │   │   └── sqlLite.go
│   │   └── storage.go
│   │
│   ├── type/
│   │   └── types.go
│   │
│   └── utils/
│       └── response/
│           └── response.go
│
├── storage/
│   └── storage.db
│
├── go.mod
├── go.sum
└── .gitignore
```

---

## 📂 Folder Structure Explanation

### 🔹 `cmd/student-project/`

**Application entry point**

* Contains `main.go`
* Responsibilities:

  * Load configuration
  * Initialize database
  * Setup HTTP routes
  * Start the server

> Each executable in a Go project lives inside the `cmd/` directory.

---

### 🔹 `config/`

**Environment configuration files**

* `local.yaml` → Local development configuration
* Can include:

  * Server port
  * Database path
  * Environment variables

---

### 🔹 `internal/`

Private application code (cannot be imported by other projects).

---

### 🔹 `internal/config/`

**Configuration loader**

* Reads YAML / ENV values
* Maps them to Go structs
* Centralized config management

---

### 🔹 `internal/http/handlers/student/`

**HTTP Layer (Controllers)**

* Handles incoming HTTP requests
* Parses request body & path params
* Validates request data
* Calls storage/business logic
* Sends standardized JSON responses

> ❌ No database logic here

---

### 🔹 `internal/storage/`

**Data access layer**

#### `storage.go`

* Defines interfaces for storage operations
* Enables dependency injection and testing

```go
type Storage interface {
	CreateStudent(name, email string, age int) (int64, error)
	GetStudent(id int) (Student, error)
	GetStudents() ([]Student, error)
}
```

#### `Sqllite/sqlLite.go`

* SQLite implementation of the storage interface
* Contains SQL queries
* Uses `database/sql`

---

### 🔹 `internal/type/`

**Domain models / DTOs**

* Contains shared data structures like `Student`
* Used across handlers and storage layers

---

### 🔹 `internal/utils/response/`

**Response utilities**

* Standard JSON response format
* Error and validation helpers
* Keeps handlers clean and consistent

---

### 🔹 `storage/storage.db`

**SQLite database file**

* Local development database
* Can be replaced with Postgres/MySQL later

---

## 🔄 Request Flow (Architecture Chart)

```
HTTP Request
   ↓
Handler (internal/http/handlers)
   ↓
Storage Interface (internal/storage)
   ↓
SQLite Implementation
   ↓
Database
```

---

## 🚀 How to Run the Project

```bash
go run cmd/student-project/main.go
```



