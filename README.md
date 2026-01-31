# Service-To-Service Calls via Phased-Construction Approach POC

This POC shows a straightforward solution for service-to-service calls.

## Problems
- service_A calls service_B over HTTP, which is slow. Services should call each other directly.

- chicken-and-egg problem:
  - the problem: 
    - a new service_A relies on an existing instance of service_B, but
    - a new service_B relies on an existing instance of service_A.
  - the solution:
    - Defer the dependency injection
    - The service-to-service dependency is actually at operation-time, not at construction time.
    - Therefore, PHASE 1: construct all services from config; PHASE 2: wire service-to-service dependencies.

## To run
1. `make run`
2. see tests execute.