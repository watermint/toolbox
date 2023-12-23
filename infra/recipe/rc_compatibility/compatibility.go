package rc_compatibility

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
	"time"
)

var (
	Definitions = CompatibilityDefinitions{}
)

type PathPair struct {
	Path []string `json:"path"`
	Name string   `json:"name"`
}

func LoadCompatibilityDefinition(data []byte) (cd CompatibilityDefinitions, err error) {
	err = json.Unmarshal(data, &cd)
	if err != nil {
		return cd, err
	}
	return
}

type PathChangeDefinition struct {
	Announcement        string     `json:"announcement,omitempty"`
	PruneAfterBuildDate string     `json:"prune_after_build_date,omitempty"`
	Current             PathPair   `json:"current"`
	FormerPaths         []PathPair `json:"former_paths"`
}

type CompatibilityDefinitions struct {
	PathChanges []PathChangeDefinition `json:"path_changes"`
}

func (z CompatibilityDefinitions) PathChangeFind(path []string, name string) (cd PathChangeDefinition, found bool) {
	for _, d := range z.PathChanges {
		if !reflect.DeepEqual(d.Current.Path, path) {
			continue
		}
		if d.Current.Name == name {
			return d, true
		}
	}
	return cd, false
}

func (z CompatibilityDefinitions) PathChangeFindAlive(path []string, name string) (cd PathChangeDefinition, found bool) {
	l := esl.Default()
	cd, found = z.PathChangeFind(path, name)
	if !found {
		return cd, false
	}
	if cd.PruneAfterBuildDate == "" {
		return cd, true
	}
	pruneAfter, err := time.Parse(time.RFC3339, cd.PruneAfterBuildDate)
	if err != nil {
		l.Error("Unable to parse prune after date", esl.Error(err))
		return cd, false
	}
	if pruneAfter.Before(time.Now()) {
		l.Debug("Prune after date", esl.String("pruneAfter", pruneAfter.String()))
		return cd, false
	}
	return cd, true
}

func (z CompatibilityDefinitions) PathChangeListAlive() (cds []PathChangeDefinition) {
	cds = make([]PathChangeDefinition, 0)
	return
}

func (z CompatibilityDefinitions) PathChangeListPruned() (cds []PathChangeDefinition) {
	cds = make([]PathChangeDefinition, 0)
	return
}
