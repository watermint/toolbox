package es_cert

import "testing"

func TestCreateSelfSigned(t *testing.T) {
	cert, key, err := CreateSelfSigned(365)
	if err != nil {
		t.Error(err)
	}
	t.Log("cert", string(cert))
	t.Log("key", string(key))
}
