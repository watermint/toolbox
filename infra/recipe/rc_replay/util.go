package rc_replay

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"os"
)

var (
	ErrorPathNotFound = errors.New("replay path not found")
)

func ReplayPath(path mo_string.OptionalString) (string, error) {
	replayPath := ""
	if path.IsExists() {
		replayPath = path.Value()
	} else if rp := os.Getenv(app_definitions.EnvNameReplayPath); rp != "" {
		replayPath = rp
	}
	if replayPath == "" {
		return "", ErrorPathNotFound
	}
	replayPath, err := es_filepath.FormatPathWithPredefinedVariables(replayPath)
	if err != nil {
		return "", ErrorPathNotFound
	}
	return replayPath, nil
}
