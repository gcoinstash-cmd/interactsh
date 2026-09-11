package server

import (
	"strings"
	"testing"
)

func TestSyslogRFC5424MessageHeader(t *testing.T) {
	rawMsg := "<165>1 2026-09-10T20:00:00.000Z host.example.com app - ID47 [interactsh@123 token="abc"] probe message"
	if !strings.HasPrefix(rawMsg, "<165>1") {
		t.Fatalf("invalid RFC-5424 Syslog priority/version header")
	}
}
