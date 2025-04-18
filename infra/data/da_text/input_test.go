package da_text_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestNewTextInput_SetFilePathAndFilePath(t *testing.T) {
	tmpDir := t.TempDir()
	resTestFilePath := filepath.Join(tmpDir, "test.txt")
	input := da_text.NewTextInput("testInput", nil)
	input.SetFilePath(resTestFilePath)
	assert.Equal(t, resTestFilePath, input.FilePath())
}

func TestNewTextInput_Debug(t *testing.T) {
	tmpDir := t.TempDir()
	resTestFilePath := filepath.Join(tmpDir, "test.txt")
	input := da_text.NewTextInput("testInput", nil)
	input.SetFilePath(resTestFilePath)
	dbg := input.Debug()
	m, ok := dbg.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "testInput", m["Name"])
	assert.Equal(t, resTestFilePath, m["FilePath"])
}

func TestTxInput_Open_FileNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	resNotFoundPath := filepath.Join(tmpDir, "notfound.txt")
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		input := da_text.NewTextInput("testInput", nil)
		input.SetFilePath(resNotFoundPath)
		err := input.Open(ctl)
		assert.NotNil(t, err)
		assert.True(t, os.IsNotExist(err))
	})
}

func TestTxInput_Content_FileNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	resNotFoundPath := filepath.Join(tmpDir, "notfound.txt")
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		input := da_text.NewTextInput("testInput", nil)
		input.SetFilePath(resNotFoundPath)
		_ = input.Open(ctl)
		_, err := input.Content()
		assert.NotNil(t, err)
		assert.True(t, os.IsNotExist(err))
	})
}
