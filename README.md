# Service-To-Service Calls via Phased-Construction Approach POC

This POC shows a lightweight solution for service-to-service calls.

## Problem

Chicken-and-egg problem:
  - constructing service_A relies on an existing instance of service_B, but
  - constructing service_B relies on an existing instance of service_A, but
  - ...

Solution:
  - Defer the dependency injection.
  - The service-to-service dependency is actually at operation-time, not at construction time.
  - Therefore
    - PHASE 1: construct all services from config
    - PHASE 2: wire service-to-service dependencies
    - PHASE 3: operational phase (process requests)

## To run
1. `make run`
2. see tests execute.