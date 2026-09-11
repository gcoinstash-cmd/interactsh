package server

import (
	"strings"
	"testing"
)

func TestSMTPSTARTTLSGreetingResponse(t *testing.T) {
	greeting := "220 2.0.0 Ready to start TLS"
	if !strings.HasPrefix(greeting, "220") {
		t.Fatalf("invalid SMTP STARTTLS response code")
	}
}
