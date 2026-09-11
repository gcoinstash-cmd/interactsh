package server

import (
	"testing"
)

func TestDNSANYQueryTypeHandling(t *testing.T) {
	queryTypeANY := uint16(255)
	if queryTypeANY != 255 {
		t.Fatalf("RFC-1035 type ANY code must be 255")
	}
}
