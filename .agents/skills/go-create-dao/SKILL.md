---
name: "go-create-dao"
description: "Scaffolds a new database adapter (DAO), implementing a domain port, with basic CRUD operations."
---

# Create DAO (Data Access Object) Skill

This skill guides the AI in safely implementing a new Data Access Object layer that adheres to the project's Hexagonal Architecture, ensuring the database logic is isolated.

## Workflow

1.  **Define the Interface (Output Port):**
    *   Locate or create the appropriate domain port file in `internal/domain/port/` (e.g., `book.go`).
    *   Define the interface for data access (e.g., `type BookDAO interface { Save(ctx context.Context, b model.Book) (model.Book, error) }`).
    *   This interface must use types from `internal/domain/model`.
2.  **Create the Implementation:**
    *   Navigate to the correct package under `internal/adapter/dao/` (e.g., `book/postgresql/`).
    *   Create a struct that implements the interface. Name it according to the underlying technology (e.g., `type bookDAO struct { db *sql.DB }`).
    *   Provide a constructor function (e.g., `func NewBookDAO(db *sql.DB) port.BookDAO`).
    *   Implement the interface methods.
    *   Perform queries securely using parameterized queries to avoid SQL injection.
    *   Map database specific errors to standard domain errors defined in `internal/domain/model` (e.g., map a `sql.ErrNoRows` to `model.ErrNotFound`).
3.  **Generate Mocks:**
    *   After the interface in `internal/domain/port/` is finalized, invoke the `generate-mocks` skill to update the mocks so that usecases can be tested.
4.  **Write Tests:**
    *   Create integration tests for the DAO in the same package to verify actual queries against a test database instance (if applicable).

## Constraints
- A DAO must implement an interface from `internal/domain/port`.
- Never leak database specific types (like `sql.NullString`) out of the DAO layer. Always map to domain models before returning.
- Ensure all queries are resilient to SQL injection.
