# Sample Go Startup/Shutdown

This repository shows how to handle both local shutdown signals (like pressing **Ctrl+C** in your terminal) and signals sent by container orchestrators (such as **Kubernetes**), ensuring your application can shut down gracefully.

The samples listen for:
- **SIGINT** (triggered by Ctrl+C locally)
- **SIGTERM** (sent by the OS or Kubernetes when stopping a pod)

## Execution Order

The samples should be reviewed in this order:

1. [**nothing**](cmd/nothing/README.md) - Simple startup and graceful shutdown using a channel
2. [**context**](cmd/context/README.md) - Simple startup and graceful shutdown using a context
2. [**grpc**](cmd/grpc/README.md) - Startup and graceful shutdown of a gRPC server with health probe
2. [**grpc-http**](cmd/grpc-http/README.md) - Startup and graceful shutdown of a gRPC and http server
