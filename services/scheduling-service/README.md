# Scheduling Service

The Scheduling Service owns appointment availability and booking consistency.

## Responsibilities

- Validate tenant and branch scope supplied by the authenticated request.
- Reserve an available slot without allowing overlapping confirmed appointments.
- Provide retry-safe booking through an idempotency key.
- Persist an outbox event in the same transaction as the appointment.

## Consistency contract

The database is authoritative for booking concurrency. Application-level checks are not sufficient: the PostgreSQL exclusion constraint is the final protection against overlapping appointments.

The booking transaction must write the appointment, idempotency record, and `appointment.created` outbox event atomically. Notification delivery is downstream and must not be part of the booking transaction.

## Database

Apply migrations in order from `migrations/`. The migration expects PostgreSQL with the `btree_gist` extension available.

## Local commands

Go uses `go.mod` instead of a `requirements.txt` file. Runtime dependencies are declared in `go.mod`; Go downloads them with `go mod download`.

Create local configuration from the template:

```bash
cp .env.example .env
```

The real `.env` file is ignored by Git and must not contain committed secrets. Update `DATABASE_URL` if your local PostgreSQL credentials or database name differ.

Install the migration CLI into this service's local `bin/` directory and download Go dependencies:

```bash
make deps
```

Set `DATABASE_URL` if PostgreSQL is not running with the local default, then run:

```bash
make migrate-up       # apply pending migrations
make migrate-status   # show migration state
make migrate-down     # roll back the latest migration
make test              # run Go tests
```

The Makefile uses a project-local installation equivalent to a Python virtual environment:

```bash
GOBIN=$(pwd)/bin go install github.com/pressly/goose/v3/cmd/goose@v3.24.1
```

The generated `bin/` directory is ignored by Git. The migration commands call this local binary directly, so no global `PATH` change is required.

## Current implementation boundary

This first slice contains the domain contract and schema. HTTP transport, authentication integration, and database adapter wiring are intentionally added only after the invariants are covered by tests.
