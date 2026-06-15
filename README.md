# Go User API (CRUD + PostgreSQL + SQLC)

A production-ready REST API built using Go (Fiber), PostgreSQL, and SQLC. The project follows clean architecture principles with clear separation of concerns across handler, service, and repository layers. It demonstrates backend engineering practices such as database integration, validation, structured logging, testing, and scalable system design.

## Features

This API provides complete CRUD functionality for user management including creating, retrieving, updating, and deleting users. It supports pagination using limit and offset parameters for efficient data retrieval. PostgreSQL is used as the database layer with SQLC generating type-safe queries at compile time. Input validation is implemented using go-playground/validator to ensure request integrity. Structured logging is handled using Uber Zap for better observability and debugging. The system follows a clean architecture approach where handlers manage HTTP requests, services contain business logic, and repositories handle database interactions. The application also calculates user age dynamically from date of birth. Unit tests are implemented for both handler and service layers to ensure reliability and correctness.

## Tech Stack

Go (Golang), Fiber Web Framework, PostgreSQL, SQLC, Zap Logger, Validator.v10

## Project Structure

cmd/server → Application entry point responsible for starting the server  
db/migrations → Database migration files used to create and update schema  
db/sqlc → SQLC generated type-safe database code  
internal/handler → HTTP layer responsible for request handling and response formatting  
internal/service → Business logic layer containing core application rules  
internal/repository → Database access layer handling all DB interactions  
internal/logger → Centralized logging configuration using Zap  
internal/routes → API route definitions and endpoint mapping  

## Setup Instructions

First, clone the repository using git clone https://github.com/your-username/go-user-api.git and navigate into the project directory using cd go-user-api. Create a PostgreSQL database named userdb using CREATE DATABASE userdb. Run database migrations using migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/userdb?sslmode=disable" up. Start the application using go run cmd/server/main.go. Once started, the server will run on http://localhost:3000.

## API Endpoints

The API exposes the following endpoints:

Create User using POST /users with a JSON body containing name and dob fields.  
Retrieve a single user using GET /users/:id.  
List all users with pagination using GET /users?page=1&limit=10.  
Update user details using PUT /users/:id with updated name and dob values.  
Delete a user using DELETE /users/:id.

## Testing

Unit tests are included for both handler and service layers. You can run all tests using the command go test ./....

## Architecture

The application follows a layered architecture pattern to ensure separation of concerns and maintainability. The handler layer is responsible for handling HTTP requests and formatting responses. The service layer contains business logic including validation and transformations such as age calculation from date of birth. The repository layer is responsible for database operations using SQLC-generated queries. This structure improves scalability, testability, and code clarity.

## Key Highlights

This project uses SQLC for compile-time type-safe SQL query generation which eliminates runtime SQL errors. Context-aware database operations are used for better request handling. Centralized validation ensures data integrity before processing. Structured logging using Zap improves observability and debugging. Pagination is efficiently implemented using SQL limit and offset. The architecture is modular and designed for scalability and future enhancements.

## Author

This project was built as part of a backend engineering assignment using Go, PostgreSQL, and SQLC. The focus was on clean architecture, scalability, maintainability, and production-ready backend development practices.
