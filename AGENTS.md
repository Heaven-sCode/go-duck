## GO-DUCK: Evolutionary Go Code Generator

### Specialist Roles
*   **Architect & Code Generator**: Handles core MVC, Multi-tenancy, and stateful GDL parsing.
*   **Security & Audit Officer**: Implements JWT, Audit logic, and Metering.
*   **API Integrator**: Manages GraphQL, PostgREST, and Swagger UI.
*   **WebSocket & Encryption Specialist**: Generates "REST-over-WS" with HMAC integrity and OIDC validation.
*   **Resilience & Performance Specialist**: Implements Distributed Redis Caching and Circuit Breakers.
*   **OpenTelemetry Integration Specialist**: Injects full-stack tracing (HTTP -> Controller -> DB) and OTel Collector K8s configs.

### Implementation Strategy & CLI Workflow
- **GO-DUCK-CLI**: The main generation engine, installable via `npm link` and executable as `go-duck`.
- **Command Architecture**:
  - `create`: Sets up a fresh microservice with full infrastructure.
  - `import-gdl`: Performs incremental updates, detecting deltas in entities and relationships.
- **Template System**: Uses Handlebars for Go code and Liquibase XML generation.
- **Stateful Persistence**: Stores entity snapshots in `.go-duck/` for safe incremental schema evolution.

### Implemented Features (The 200% Milestone)

#### 1. Core & Generic Layers
- **Full CRUD REST APIs**: Automatically generated with pagination and filtering.
- **Generic Search Layer**: PostgREST-like RPC endpoint (`/api/rpc/:table`).
- **Audit & Metering**: Automatic capturing of entity changes and per-tenant usage tracking.

#### 2. Real-Time & Performance
- **REST-over-WS**: WebSocket dispatcher with **Traced-Envelopes** (OTel supported).
- **Distributed Caching (Redis)**: Multi-tenant aware (Tenant-Prefixed) with Cache-Aside strategy.
- **Event Streaming (MQTT)**: Real-time CRUD notifications for webhooks/audit.

#### 3. Resilience & Security
- **Circuit Breakers**: Sony/Gobreaker integration for Redis/MQTT/DB calls.
- **OIDC/JWT Security**: Keycloak validation with dynamic context-aware multi-tenancy.
- **Digital Signatures**: HMAC-SHA256 verification for high-integrity WebSocket message payloads.
- **CORS & Rate Limiting**: Property-driven policies for security and burst protection.

#### 4. Observability (Full-Stack)
- **OpenTelemetry (OTel)**: Distributed tracing from Router (otelgin) to Database (otelpgx plugin).
- **Datadog Logging**: Environment-driven log streaming and monitoring.
- **Statsd Metrics**: Infrastructure performance tracing and custom metric pushing.

#### 5. Deployment & Cloud-Native
- **K8s Configs**: Auto-generated ConfigMaps for OpenTelemetry Collectors.
- **Environment Profiles**: Switchable via `GO_PROFILE` (`dev`, `prod`).

### Technology Stack
- **Language**: Go
- **Web**: Gin Gonic + Gorilla WS
- **ORM**: GORM (PostgreSQL)
- **Identity**: Keycloak (OIDC)
- **Caching**: Redis
- **Observability**: OpenTelemetry + Datadog
- **Messaging**: MQTT
- **Resilience**: Sony/Gobreaker
