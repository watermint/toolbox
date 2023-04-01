package sv_project

import (
	"strings"
	"testing"
)

func TestVerifyTeamId(t *testing.T) {
	if r, _ := VerifyTeamId("1234"); r != VerifyTeamIdLooksOkay {
		t.Error(r)
	}
	if r, _ := VerifyTeamId("@#$%"); r != VerifyTeamInvalidChar {
		t.Error(r)
	}
	if r, _ := VerifyTeamId(strings.Repeat("x", TeamIdMaxLength+1)); r != VerifyTeamIdTooLong {
		t.Error(r)
	}
}

func TestVerifyProjectId(t *testing.T) {
	if r, _ := VerifyProjectId("1234"); r != VerifyProjectIdLooksOkay {
		t.Error(r)
	}
	if r, _ := VerifyProjectId("@#$%"); r != VerifyProjectIdInvalidCharacter {
		t.Error(r)
	}
	if r, _ := VerifyProjectId(strings.Repeat("x", ProjectIdMaxLength+1)); r != VerifyProjectIdTooLong {
		t.Error(r)
	}
}
