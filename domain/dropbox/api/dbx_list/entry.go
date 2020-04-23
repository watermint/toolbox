package dbx_list

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

type listEntryImpl struct {
	entry gjson.Result
}

func (z listEntryImpl) Json() (res gjson.Result, err error) {
	return z.entry, nil
}

func (z listEntryImpl) Model(v interface{}) error {
	return api_parser.ParseModel(v, z.entry)
}

func (z listEntryImpl) ModelWithPath(v interface{}, path string) error {
	return api_parser.ParseModel(v, z.entry.Get(path))
}
