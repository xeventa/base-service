# Base Service (Go)

## Architecture (Domain-based)
- app: dependency setup/wiring (DI container)
- core: domain core files (config, repository interfaces/impl, DB connection, use cases)
- src: endpoints and server routing
- cmd/base-service: application entrypoint

## Quick start
1. Ensure Go 1.22+
2. Optional environment variables:
   - HTTP_PORT (default: 8080)
   - DATABASE_URL (optional; if set, will attempt to open a postgres connection)
3. Run:

```
cd /d D:\project-xeventa\base-service
go run .\cmd\base-service
```

## Endpoints
- GET /health -> 200 ok
- GET /greet/{name} -> returns greeting message

## Notes
- DI is managed in `app/container.go`.
- Core domain elements in `core/` are easy to extend (swap repo impl, add use cases).
- `src/` will host endpoints and route handlers for future features.
