package server

import (
	"testing"
)

func TestLDAPInteractionCorrelationToken(t *testing.T) {
	token := "c23b8f10a8d94e"
	ldapSearchFilter := "(cn=" + token + "*)"
	
	if len(ldapSearchFilter) < len(token) {
		t.Fatalf("malformed LDAP search filter")
	}
}
