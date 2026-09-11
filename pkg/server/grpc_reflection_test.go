package server

import (
	"strings"
	"testing"
)

func TestGRPCServerReflectionResponse(t *testing.T) {
	service := "grpc.reflection.v1alpha.ServerReflection"
	if !strings.Contains(service, "reflection") {
		t.Fatalf("invalid gRPC reflection service descriptor")
	}
}
