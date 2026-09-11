package server

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestSessionTokenCorrelationHash(t *testing.T) {
	token := "session-interaction-correlation-token-12345"
	h := sha256.New()
	h.Write([]byte(token))
	hashStr := hex.EncodeToString(h.Sum(nil))

	if len(hashStr) != 64 {
		t.Fatalf("expected 64 char sha256 hash, got %d", len(hashStr))
	}
}
