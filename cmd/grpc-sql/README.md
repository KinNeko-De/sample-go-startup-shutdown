## Overview

The `grpc-sql` command serves as a minimal example showing:

- Starting application
- Starting gRPC server
- Connecting to sql server
- Report readiness
- Waiting for signal
- Report not readiness
- Disconnect from sql server
- Shutdown gRPC server
- Shutdown application

## Usage

```bash
(cd scripts; ./postgres-run.sh)
go run ./cmd/grpc-sql
```

## Readiness probe
```bash
grpcurl -plaintext -d '{"service":"readiness"}' localhost:9090 grpc.health.v1.Health/Check
```