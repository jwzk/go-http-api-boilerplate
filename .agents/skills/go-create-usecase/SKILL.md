---
name: "go-create-usecase"
description: "Scaffolds a new business logic use case, including the input port interface, and the concrete implementation."
---

# Create UseCase Skill

This skill guides the AI in adding new business logic usecases in accordance with the project's Hexagonal Architecture.

## Workflow

1.  **Define the Interface (Input Port):**
    *   Locate or create the appropriate domain port file in `internal/domain/port/`.
    *   Define an interface that represents the use case (e.g., `type CreateBook interface { Execute(ctx context.Context, b model.Book) (model.Book, error) }`).
    *   Ensure it only depends on standard library types or domain models (`internal/domain/model`). It must never import from `adapter`.
2.  **Create the Implementation:**
    *   Navigate to the appropriate package under `internal/domain/usecase/` (e.g., `book/`).
    *   Create a struct that implements the interface. Name it according to the action (e.g., `type createBook struct { ... }`).
    *   Inject dependencies (like Data Access Objects) via the constructor (e.g., `func NewCreateBook(db port.BookDAO) port.CreateBook`).
    *   Implement the method defined in the interface.
    *   Apply any necessary business rules and validation. Return appropriate domain errors (e.g., `model.ErrNotFound`) if rules are violated.
3.  **Write Tests:**
    *   Create a test file in the usecase package (e.g., `create_test.go`).
    *   Use the AAA (Arrange, Act, Assert) pattern.
    *   Generate and use mocks for injected dependencies using `mockery`. (See `generate-mocks` skill).

## Constraints
- Usecases must not depend on any adapter (HTTP, DB, CLI).
- Rely strictly on dependency injection via interfaces defined in `internal/domain/port/`.
- Handle errors specifically and return standard domain errors defined in `internal/domain/model` if appropriate.
