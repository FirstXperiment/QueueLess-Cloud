# QueueLess Cloud — System Context

## Scope

QueueLess Cloud is a multi-tenant scheduling and queue orchestration platform. The first production capability is appointment booking for an organization operating one or more branches.

The architecture is organized around business capabilities, not technical layers. A service owns its business rules and data; no service reads another service's database directly.

## Bounded Contexts

| Context | Responsibility | Owns | Initial boundary |
|---|---|---|---|
| Identity | Authentication and user identity | Users, credentials, sessions, roles | `identity-service` |
| Tenant | Organizations, branches, tenant configuration | Organizations, branches, memberships | `tenant-service` |
| Scheduling | Availability and appointment lifecycle | Staff availability, schedules, appointments | `scheduling-service` |
| Queue | Walk-ins, tokens, queue progression | Queue entries, priorities, state transitions | `queue-service` |
| Resource | Staff, rooms, counters, capacity | Resource definitions and assignments | `resource-service` |
| Notification | Delivery of user-facing messages | Notification intents, delivery attempts | `notification-service` |

## Ownership Rules

1. A service is the sole writer of its owned data.
2. Cross-service references use stable IDs, never foreign keys across databases.
3. A service may cache another service's data only with an explicit freshness policy.
4. Commands target one owning service.
5. Events describe facts that already happened; consumers must be idempotent.
6. API contracts and event schemas are versioned independently from implementations.

## Critical Booking Invariants

- An appointment belongs to exactly one tenant and branch.
- A resource cannot have overlapping confirmed appointments.
- A booking request must be safe to retry with the same idempotency key.
- A cancelled appointment cannot transition back to confirmed.
- Authorization must verify both actor permissions and tenant scope.
- The scheduling service is authoritative for booking consistency.

## Initial Request Flow

```text
Client
  -> API Gateway
  -> Scheduling Service
      -> validates tenant context through a trusted identity claim
      -> checks availability and booking constraints
      -> commits appointment and outbox record atomically
  -> returns booking result

Outbox Relay
  -> AppointmentCreated event
      -> Notification Service
      -> Analytics/read-model consumers (later)
```

The booking response must not depend on notification delivery. Notification failure is retried asynchronously and does not roll back a successful booking.

## Consistency Model

Strong consistency is required inside the scheduling service for slot reservation, appointment state transitions, and idempotency records. Cross-service workflows are eventually consistent and use durable events, retries, and dead-letter handling.

## Non-Goals for the First Slice

- Kubernetes and multi-region deployment.
- Kafka before event volume or operational needs justify it.
- Distributed transactions across services.
- A shared database used by all services.
- Generic service extraction without a business ownership boundary.

