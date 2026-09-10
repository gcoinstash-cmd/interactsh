package server

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSNITokenParsingW23(t *testing.T) {
	sni := "c123456789abcdef.interactsh.com"
	parts := strings.Split(sni, ".")
	assert.GreaterOrEqual(t, len(parts), 3)
	token := parts[0]
	assert.GreaterOrEqual(t, len(token), 10)
}
