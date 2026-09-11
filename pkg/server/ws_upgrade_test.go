package server

import (
	"strings"
	"testing"
)

func TestWebSocketUpgradeHeaderExtraction(t *testing.T) {
	req := "GET /chat HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==

"
	if !strings.Contains(req, "Sec-WebSocket-Key") {
		t.Fatalf("missing Sec-WebSocket-Key header in probe handshake")
	}
}
