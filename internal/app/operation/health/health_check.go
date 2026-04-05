package health

import (
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	liveness  = "liveness"
	readiness = "readiness"
)

func Initialize(server *health.Server) {
	Live(server)
	NotReady(server)
}

func Live(server *health.Server) {
	server.SetServingStatus(liveness, grpc_health_v1.HealthCheckResponse_SERVING)
}

func NotLive(server *health.Server) {
	server.SetServingStatus(liveness, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
}

func Ready(server *health.Server) {
	server.SetServingStatus(readiness, grpc_health_v1.HealthCheckResponse_SERVING)
}

func NotReady(server *health.Server) {
	server.SetServingStatus(readiness, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
}
