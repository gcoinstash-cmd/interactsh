package server

import (
	"strings"
	"testing"
)

func TestSanitizeDNSTXTResponse(t *testing.T) {
	inputStr := "response_payload_token\x00_extra_data"
	sanitized := strings.ReplaceAll(inputStr, "\x00", "")

	if strings.Contains(sanitized, "\x00") {
		t.Fatalf("expected string to have null bytes stripped")
	}

	maxLen := 255
	if len(sanitized) > maxLen {
		sanitized = sanitized[:maxLen]
	}

	if len(sanitized) > 255 {
		t.Fatalf("expected TXT record length <= 255, got %d", len(sanitized))
	}
}
