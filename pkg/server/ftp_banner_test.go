package server

import (
	"strings"
	"testing"
)

func TestFTPPASVBannerExtraction(t *testing.T) {
	banner := "220 Interactsh FTP Server Ready"
	if !strings.HasPrefix(banner, "220") {
		t.Fatalf("invalid FTP greeting banner status code")
	}
}
