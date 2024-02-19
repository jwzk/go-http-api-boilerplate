---
name: go-update-swagger
description: "Describes how to keep the OpenAPI 3.1.1 documentation up to date when creating or modifying endpoints."
---

# go-update-swagger

When modifying or creating new HTTP endpoints, you MUST update the Swagger documentation to reflect those changes.

## Prerequisites
- The OpenAPI specification file is located at the project root: `openapi.yaml`.
- The project uses OpenAPI Specification version 3.1.1.

## Workflow
1. **Identify Changes**: Determine if your work introduces new endpoints, modifies existing routes, changes query parameters, updates request payloads, or modifies response schemas.
2. **Edit `openapi.yaml`**:
   - Open the `openapi.yaml` file located in the root of the project.
   - For new endpoints, add the corresponding paths, methods, operation IDs, request schemas, and responses.
   - For modified endpoints, carefully adjust the existing paths to match the updated implementation.
   - If a new domain model or DTO is introduced, declare it in the `components/schemas` section to ensure consistency.
3. **Validation**: Ensure that the syntax remains valid YAML and fully complies with OpenAPI 3.1.1 standards.
4. **Test with Swagger UI**: The Swagger UI runs via docker-compose. You can start the service with `docker-compose up -d swagger-ui` and navigate to `http://localhost:8080` to verify that your changes are correctly rendered and the specification contains no errors.
