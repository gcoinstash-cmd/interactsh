package server

import (
	"testing"
)

func TestSMBNTLMChallengePayload(t *testing.T) {
	challenge := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	if len(challenge) != 8 {
		t.Fatalf("NTLM server challenge nonce must be exactly 8 bytes")
	}
}
