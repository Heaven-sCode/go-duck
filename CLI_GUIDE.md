# GO-DUCK-CLI Usage Guide

The `GO-DUCK-CLI` is a powerful Go code generator that transforms GDL files into a production-ready microservice with multi-tenancy, auditing, and Liquibase migrations.

## Installation

To install the CLI locally for development:

1. Navigate to the CLI directory:
   ```bash
   cd GO-DUCK-CLI
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Link the package globally (optional):
   ```bash
   npm link
   ```

## Usage

Once installed, you can generate a base Go application using a single command:

```bash
go-duck create --config <config-path> --output <output-path>
```

### Options:

| Option | Shorthand | Default | Description |
| :--- | :--- | :--- | :--- |
| `--config` | `-c` | `../CONFIG/config.yaml` | Path to your `config.yaml` file. |
| `--output` | `-o` | `../SAMPLE-GO-FUNCTION` | Directory where the app will be generated. |
| `--gdl` | `-g` | `../GDL` | Directory containing your `.gdl` files. |

### Example:

```bash
go-duck create -c ./CONFIG/config.yaml -o ./MyGeneratedApp
```

## Features Generated:

*   **Models**: GORM models with `@AutoWired` audit fields.
*   **Multi-tenancy**: Middleware for dynamic DB routing based on Keycloak roles.
*   **Auditing**: Global row-wise audit log tracking Keycloak IDs and client IPs.
*   **Migrations**: Complete Liquibase changelogs for schema management.
*   **Management API**: Endpoint to create and migrate databases at runtime.
