---
name: "go-create-endpoint"
description: "Scaffolds a new HTTP endpoint, including request/response DTOs, handler function, and routing registration."
---

# Create HTTP Endpoint Skill

This skill guides the AI in safely adding new HTTP endpoints according to the project's Hexagonal Architecture and conventions.

## Workflow

1.  **Locate the Adapter:** Identify the correct API package under `internal/adapter/http/` (e.g., `bookapi/book`). If the resource does not exist, create a new directory for it.
2.  **Define DTOs (if applicable):** Create or update `dto.go` in the resource package.
    - DTOs should be local to the HTTP package (e.g., `type createBookDTO struct {...}`).
    - Use JSON struct tags.
    - Add validation tags if needed by `pkg/validator` (e.g., `validate:"required"`).
    - Include a method to map the DTO to a domain `model` (e.g., `func (d *createBookDTO) Model() model.Book`).
3.  **Create Handler:** In `handler.go`, create a handler function.
    - The function should return an `http.HandlerFunc`.
    - It should be a method on the resource's Router struct (e.g., `func (b *BookRouter) createBook() http.HandlerFunc`).
    - Extract path variables using Go 1.26.4's `r.PathValue("id")`.
    - Validate the body using `validator.Validate(r.Context(), r.Body, &dto)`.
    - Call the appropriate usecase method.
    - Respond using `writer.JSON(r.Context(), w, response, err)`.
4.  **Register Route:** In `router.go` of the resource package, register the new handler in the `New...Router` function using Go 1.26.4 routing patterns (e.g., `router.HandleFunc("POST /", b.createBook())`).
5.  **Write Tests:** Create or update `handler_test.go` using the AAA pattern. Mock the usecase.

## Constraints

- Never put business logic in the HTTP handler.
- Always use the `pkg/http/writer` or `internal/adapter/http/internal/writer` for responses.
- Ensure the handler handles errors appropriately without panicking.
