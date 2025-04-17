package da_griddata

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type dummyFormatter struct{}

func (d dummyFormatter) Format(v interface{}, col, row int) interface{} {
	return v
}

func TestNewOutput_Basic(t *testing.T) {
	out := NewOutput(nil, "testOut")
	assert.Equal(t, "testOut", out.Name())
	assert.Equal(t, "testOut", out.Debug().(map[string]interface{})["Name"])
}

func TestGdOutput_SetFormatter(t *testing.T) {
	out := NewOutput(nil, "testOut")
	out.SetFormatter(OutputTypeCsv, dummyFormatter{})
	// No panic, formatter set
}

func TestGdOutput_FilePath(t *testing.T) {
	out := NewOutput(nil, "testOut")
	assert.Equal(t, "", out.FilePath())
}

func TestGdOutput_CaptureRestore(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.csv")
	out := NewOutput(nil, "testOut").(*gdOutput)
	out.filePath = filePath
	out.outputType = OutputTypeJson
	cap := out.Capture()
	b, err := json.Marshal(cap)
	assert.Nil(t, err)
	j := es_json.MustParse(b)
	out2 := NewOutput(nil, "testOut").(*gdOutput)
	err = out2.Restore(j)
	assert.Nil(t, err)
	assert.Equal(t, filePath, out2.filePath)
	assert.Equal(t, OutputTypeJson, out2.outputType)
}

func TestGdOutput_Restore_Error(t *testing.T) {
	out := NewOutput(nil, "testOut").(*gdOutput)
	b, err := json.Marshal("not-an-object")
	assert.Nil(t, err)
	j := es_json.MustParse(b)
	err = out.Restore(j)
	assert.Equal(t, ErrorValueRestoreFailed, err)
}
