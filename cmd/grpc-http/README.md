## Overview

The `grpc-http` command serves as a minimal example showing:

- Starting application
- Starting gRPC server
- Starting http server
- Report readiness
- Waiting for signal
- Report not readiness
- Shutdown http server
- Shutdown gRPC server
- Shutdown application

## Usage

```bash
go run ./cmd/grpc-http
```

## Readiness probe
```bash
grpcurl -plaintext -d '{"service":"readiness"}' localhost:9090 grpc.health.v1.Health/Check
```