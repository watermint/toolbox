package es_project

import "testing"

func TestDetectRepositoryRoot(t *testing.T) {
	_, err := DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
	}
}
