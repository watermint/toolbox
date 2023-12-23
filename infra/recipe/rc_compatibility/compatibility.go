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

func IsAlive(pruneAfterBuildDate string) bool {
	if pruneAfterBuildDate == "" {
		return true
	}
	pruneAfter, err := time.Parse(time.RFC3339, pruneAfterBuildDate)
	if err != nil {
		l := esl.Default()
		l.Error("Unable to parse prune after date", esl.Error(err))
		return false
	}
	return pruneAfter.After(time.Now())
}

func LoadCompatibilityDefinition(data []byte) (cd CompatibilityDefinitions, err error) {
	err = json.Unmarshal(data, &cd)
	if err != nil {
		return cd, err
	}
	return
}

type CompatibilityDefinitions struct {
	PathChanges []PathChangeDefinition `json:"path_change"`
	Prune       []PruneDefinition      `json:"prune"`
}

func (z CompatibilityDefinitions) FindPathChange(path []string, name string) (cd PathChangeDefinition, found bool) {
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

func (z CompatibilityDefinitions) FindAlivePathChange(path []string, name string) (cd PathChangeDefinition, found bool) {
	l := esl.Default()
	cd, found = z.FindPathChange(path, name)
	if !found {
		return cd, false
	}
	if IsAlive(cd.PruneAfterBuildDate) {
		return cd, true
	} else {
		l.Debug("Prune after date", esl.String("pruneAfter", cd.PruneAfterBuildDate))
		return cd, false
	}
}

func (z CompatibilityDefinitions) ListAlivePathChange() (cds []PathChangeDefinition) {
	cds = make([]PathChangeDefinition, 0)
	for _, d := range z.PathChanges {
		if IsAlive(d.PruneAfterBuildDate) {
			cds = append(cds, d)
		}
	}
	return
}

func (z CompatibilityDefinitions) ListPrunedPathChange() (cds []PathChangeDefinition) {
	cds = make([]PathChangeDefinition, 0)
	for _, d := range z.PathChanges {
		if !IsAlive(d.PruneAfterBuildDate) {
			cds = append(cds, d)
		}
	}
	return
}

func (z CompatibilityDefinitions) FindPrune(path []string, name string) (cd PruneDefinition, found bool) {
	for _, d := range z.Prune {
		if !reflect.DeepEqual(d.Current.Path, path) {
			continue
		}
		if d.Current.Name == name {
			return d, true
		}
	}
	return cd, false
}

func (z CompatibilityDefinitions) FindAlivePrune(path []string, name string) (cd PruneDefinition, found bool) {
	l := esl.Default()
	cd, found = z.FindPrune(path, name)
	if !found {
		return cd, false
	}
	if IsAlive(cd.PruneAfterBuildDate) {
		return cd, true
	} else {
		l.Debug("Prune after date", esl.String("pruneAfter", cd.PruneAfterBuildDate))
		return cd, false
	}
}

func (z CompatibilityDefinitions) ListAlivePrune() (cds []PruneDefinition) {
	cds = make([]PruneDefinition, 0)
	for _, cd := range z.Prune {
		if IsAlive(cd.PruneAfterBuildDate) {
			cds = append(cds, cd)
		}
	}
	return
}

func (z CompatibilityDefinitions) ListPrunedPrune() (cds []PruneDefinition) {
	cds = make([]PruneDefinition, 0)
	for _, cd := range z.Prune {
		if !IsAlive(cd.PruneAfterBuildDate) {
			cds = append(cds, cd)
		}
	}
	return
}
