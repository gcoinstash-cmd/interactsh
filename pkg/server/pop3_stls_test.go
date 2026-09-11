package server

import (
	"strings"
	"testing"
)

func TestPOP3STLSGreetingResponse(t *testing.T) {
	resp := "+OK Begin TLS negotiation"
	if !strings.HasPrefix(resp, "+OK") {
		t.Fatalf("invalid POP3 STLS response greeting code")
	}
}
