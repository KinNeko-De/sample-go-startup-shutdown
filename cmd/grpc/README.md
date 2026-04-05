## Overview

The `grpc` command serves as a minimal example showing:

- Starting application
- Starting gRPC server
- Report readiness
- Waiting for signal
- Report not readiness
- Shutdown gRPC server
- Shutdown application

## Usage

```bash
go run ./cmd/grpc
```

## Readiness probe
```bash
grpcurl -plaintext -d '{"service":"readiness"}' localhost:9090 grpc.health.v1.Health/Check
```