# QueueLess Cloud Development Roadmap

This document defines the planned implementation sequence for QueueLess Cloud. It is a living project plan and may be updated as requirements, architecture, and priorities evolve.

The application will be developed in functional slices. Each phase should result in a usable and testable capability before the next phase begins.

## Phase 1: Project Foundation

- Finalize the initial technology choices.
- Establish the monorepo structure.
- Add local Docker Compose development support.
- Configure PostgreSQL and Redis.
- Define API conventions, error responses, configuration, and logging.
- Add formatting, linting, testing, and CI checks.
- Document local development and contribution workflows.

## Phase 2: Authentication and Identity

- User registration and login.
- Secure password storage and password reset.
- Access and refresh tokens.
- Email verification.
- User profile management.
- Role-based permissions for platform administrators, organization owners, branch managers, staff, and customers.
- Authentication middleware for backend services.
- Audit logging for security-sensitive actions.

## Phase 3: Organizations and Tenant Management

- Organization creation and settings.
- Branch creation and management.
- User invitations.
- Role assignment.
- Tenant-isolated data access.
- Branch timezone and business hours.
- Organization and branch APIs.

## Phase 4: Services and Resources

- Create and manage bookable service types.
- Configure service duration and buffer time.
- Add staff members.
- Add rooms, counters, and other resources.
- Assign staff and resources to services.
- Configure resource availability.
- Validate resource assignments and conflicts.

## Phase 5: Scheduling and Appointments

- Define availability rules.
- Generate available time slots.
- Create appointments.
- Prevent overlapping bookings.
- Support idempotent booking requests.
- Cancel appointments.
- Reschedule appointments.
- Support timezone-aware scheduling.
- Maintain appointment history.
- Provide customer and staff appointment views.

## Phase 6: Customer Application

- Customer registration and login.
- Browse organizations and branches.
- Select a service, date, and time.
- Book an appointment.
- View upcoming and historical appointments.
- Cancel or reschedule appointments.
- Display booking confirmations.
- Support public booking links for organizations.

## Phase 7: Staff Application

- Staff dashboard.
- Daily appointment view.
- Customer check-in.
- Mark appointments as completed.
- Mark customers as no-show.
- Start and finish service.
- View the current queue.
- Reassign appointments when required.

## Phase 8: Queue Management

- Generate queue tokens for walk-in customers.
- Allow customers to join a queue.
- Display queue position.
- Estimate waiting time.
- Call the next customer.
- Skip, recall, and complete queue entries.
- Support appointment and walk-in queues.
- Add configurable queue-priority rules.
- Add real-time queue updates.

## Phase 9: Notifications

- Appointment confirmation notifications.
- Cancellation and rescheduling notifications.
- Appointment reminders.
- Queue position notifications.
- Staff notifications.
- Start with email delivery, followed by SMS and push notifications.
- Use an outbox and background worker for reliable delivery.
- Add retries and dead-letter handling.

## Phase 10: Administration and Reporting

- Organization administration dashboard.
- Appointment volume reports.
- No-show rate reports.
- Average waiting-time reports.
- Service utilization reports.
- Staff utilization reports.
- Branch comparison.
- Report export.

## Phase 11: Security and Production Readiness

- Rate limiting.
- Input validation.
- CORS and security headers.
- Tenant-isolation tests.
- Audit-log review tools.
- Database backup and recovery procedures.
- Service health checks.
- Metrics and distributed tracing.
- Error monitoring.
- Load and performance testing.
- Deployment configuration.

## Phase 12: Release and Operations

- Deploy a staging environment.
- Run end-to-end and contract tests.
- Perform a security review.
- Publish onboarding documentation.
- Add demo organization and sample data.
- Deploy the first production release.
- Monitor service health and collect user feedback.

## Minimum Viable Product

The first usable release should include:

1. Authentication and identity.
2. Organization and branch management.
3. Service and staff configuration.
4. Appointment availability.
5. Appointment booking.
6. Cancellation and rescheduling.
7. A basic staff dashboard.
8. Email notifications.

Advanced queue management, analytics, SMS delivery, Kafka, Kubernetes, and complex resource orchestration should follow the MVP unless a concrete product requirement moves them earlier.

## Current Priority

The next implementation priority is **Authentication and Identity**. After that, implement **Organizations and Tenant Management**, then connect the Scheduling Service to authenticated users, branches, services, and permissions.

## Roadmap Maintenance

This roadmap should be updated when a phase is completed, its scope changes, or new product requirements are approved. Completed work should be reflected in the main README and the relevant service documentation.
