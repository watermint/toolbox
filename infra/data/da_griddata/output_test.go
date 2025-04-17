package da_griddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainGridDataFormatter_Format(t *testing.T) {
	f := PlainGridDataFormatter{}
	v := f.Format(123, 0, 0)
	assert.Equal(t, 123, v)
	v2 := f.Format("abc", 1, 2)
	assert.Equal(t, "abc", v2)
}

func TestDefaultGridDataFormatter(t *testing.T) {
	v := DefaultGridDataFormatter.Format(456, 0, 0)
	assert.Equal(t, 456, v)
}
