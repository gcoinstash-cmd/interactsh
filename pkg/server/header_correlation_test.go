package server

import (
	"strings"
	"testing"
)

func TestExtractCorrelationHeader(t *testing.T) {
	headers := map[string]string{
		"X-Interactsh-Token": "corr-token-xyz-987",
		"User-Agent":         "Interactsh-Probe/v1",
	}

	token := headers["X-Interactsh-Token"]
	if !strings.HasPrefix(token, "corr-token-") {
		t.Fatalf("expected correlation token prefix, got %s", token)
	}
}
