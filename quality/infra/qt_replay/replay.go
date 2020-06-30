package qt_replay

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
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
		l.Warn("Test path not found")
		return nil, err
	}
	tp := filepath.Join(root, "test")
	rp := filepath.Join(tp, "replay", name)

	l.Debug("Loading replay", esl.String("path", rp))
	all, err := ioutil.ReadFile(rp)
	if err != nil {
		l.Debug("Unable to load", esl.Error(err))
		return nil, err
	}

	if strings.HasSuffix(name, ".gz") {
		cr, err := gzip.NewReader(bytes.NewReader(all))
		if err != nil {
			l.Debug("Unable to create gzip reader", esl.Error(err))
			return nil, err
		}
		cbr := new(bytes.Buffer)
		if _, err := io.Copy(cbr, cr); err != nil {
			l.Debug("unable to uncompress", esl.Error(err))
			return nil, err
		}
		all = cbr.Bytes()
	}

	// if an array of Response
	if err := json.Unmarshal(all, &rr); err == nil {
		return rr, nil
	}

	rr = make([]nw_replay.Response, 0)
	chunks := strings.Split(string(all), "\n")
	for _, chunk := range chunks {
		if len(strings.TrimSpace(chunk)) < 1 {
			continue
		}
		res := nw_replay.Response{}
		if err := json.Unmarshal([]byte(chunk), &res); err != nil {
			l.Debug("Unable to unmarshal", esl.Error(err))
			return nil, err
		}
		rr = append(rr, res)
	}

	if len(rr) < 1 {
		l.Error("No replay loaded")
		return nil, errors.New("no replay loaded")
	}

	l.Debug("Replay loaded", esl.Int("numRecords", len(rr)))
	return rr, nil
}
