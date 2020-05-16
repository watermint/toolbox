package qt_replay

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func LoadReplay(name string) (rr []nw_replay.Response, err error) {
	l := esl.Default().With(esl.String("name", name))
	root, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Error("Test path not found")
		return nil, err
	}
	tp := filepath.Join(root, "test")
	rp := filepath.Join(tp, "replay", name)

	l.Debug("Loading replay", esl.String("path", rp))
	b, err := ioutil.ReadFile(rp)
	if err != nil {
		l.Debug("Unable to load", esl.Error(err))
		return nil, err
	}

	if strings.HasSuffix(name, ".gz") {
		cr, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			l.Debug("Unable to create gzip reader", esl.Error(err))
			return nil, err
		}
		cbr := new(bytes.Buffer)
		if _, err := io.Copy(cbr, cr); err != nil {
			l.Debug("unable to uncompress", esl.Error(err))
			return nil, err
		}
		b = cbr.Bytes()
	}

	if err := json.Unmarshal(b, &rr); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return nil, err
	}

	l.Debug("Replay loaded", esl.Int("numRecords", len(rr)))
	return rr, nil
}
