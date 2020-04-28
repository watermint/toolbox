package es_generate

import "testing"

func TestDetectRepositoryRoot(t *testing.T) {
	_, err := DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
	}
}
