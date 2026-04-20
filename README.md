# Enrollment System Backend (Golang)

This is the skeleton for the backend of the school enrollment system.

## Tech Stack
- **Go** 1.21+
- **GoFiber** (Web Framework)
- **GORM** (ORM)
- **MySQL** (Database)
- **golang-migrate** (Database Migrations)
- **JWT** (Authentication)
- **bcrypt** (Password Hashing)
- **zap** (Logging)
- **Cloudflare R2** (Storage - Planned)

## Project Structure

```text
.
├── cmd/server           # Main applications for this project
├── internal             # Private application and library code
│   ├── common           # Common structs like Response, Pagination
│   ├── config           # Configuration loader
│   ├── database         # Database connection logic
│   ├── middleware       # HTTP middlewares (CORS, Logger, Recovery)
│   └── modules          # Feature modules (e.g., health, auth, users)
├── migrations           # SQL migration files
├── pkg                  # Public library code (e.g., Logger wrapper)
└── README.md
```

## How to run locally

## Database Setup & Migrations

This project explicitly uses explicitly defined SQL schema, NOT GORM AutoMigrate. 
Your target MySQL database is structured via `golang-migrate`.

### 1. Requirements
Ensure you have the following CLI installed:
- Go (1.21+)
- MySQL server
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) CLI tool

### 2. Setup Empty Database
Create a database in your local MySQL instance:
```sql
CREATE DATABASE enrollment_db;
```

### 3. Run Migrations
Be sure your `.env` contains the correct DB credentials, then run the migration command passing the credentials identical to your environment (Replace `root`, `secret`, `127.0.0.1:3306`, and `enrollment_db` if you use different ones):

**Run All Up Migrations (Apply to Database):**
```sh
migrate -path migrations -database "mysql://root:secret@tcp(127.0.0.1:3306)/enrollment_db" up
```

**Run Down Migrations (Revert):**
```sh
migrate -path migrations -database "mysql://root:secret@tcp(127.0.0.1:3306)/enrollment_db" down 1
```

### 4. Create further migrations
```sh
migrate create -ext sql -dir migrations -seq add_some_column
```

### 5. Running the Application
```sh
air
```
