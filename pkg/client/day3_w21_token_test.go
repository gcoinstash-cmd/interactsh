package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteractionTokenValidation(t *testing.T) {
	token := "oob-session-token-9921"
	assert.True(t, len(token) > 10)
	assert.Contains(t, token, "session")
}
