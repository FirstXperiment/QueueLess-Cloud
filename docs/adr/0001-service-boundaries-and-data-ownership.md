# ADR-0001: Service Boundaries and Data Ownership

- Status: Accepted
- Date: 2026-10-20
- Decision owners: QueueLess Cloud maintainers

## Context

QueueLess Cloud needs independent evolution of identity, tenant administration, scheduling, queue operations, resources, and notifications. Splitting code by technical concern alone would create distributed coupling without clear ownership.

## Decision

The platform uses business-capability services with strict data ownership:

- Identity owns authentication and identity records.
- Tenant owns organizations, branches, memberships, and tenant configuration.
- Scheduling owns availability and appointment consistency.
- Queue owns queue entries and queue state transitions.
- Resource owns staff, rooms, counters, and capacity.
- Notification owns delivery attempts and channel adapters.

Services communicate through synchronous APIs for immediate commands and versioned asynchronous events for facts and downstream work. No service directly accesses another service's database.

## Consequences

### Positive

- Business rules have a clear authoritative owner.
- Services can scale and deploy independently.
- Failures in notification or analytics do not invalidate bookings.
- Data access and authorization boundaries are reviewable.

### Negative

- Cross-service queries require APIs, events, or read models.
- Distributed workflows require retries, idempotency, tracing, and operational tooling.
- Local development is more complex than a single process.

## Rejected Alternatives

### Shared database across services

Rejected because it creates hidden coupling and allows business rules to bypass service boundaries.

### One service per database table

Rejected because technical decomposition does not create meaningful ownership and would increase network coordination.

### Kafka from the first commit

Rejected for the initial slice. The event contract and outbox pattern are defined first; the broker can be introduced when operationally justified.

