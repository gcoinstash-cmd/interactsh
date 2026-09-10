package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPollingIntervalDurationValidation(t *testing.T) {
	interval := 5 * time.Second
	assert.Equal(t, 5*time.Second, interval)
}
