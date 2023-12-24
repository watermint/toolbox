package rc_compatibility

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/exp/slices"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	Definitions = CompatibilityDefinitions{}
)

type PathPair struct {
	Path []string `json:"path"`
	Name string   `json:"name"`
}

func (z PathPair) CliPath() string {
	return strings.Join(append(z.Path, z.Name), " ")
}

func (z PathPair) IsValid() bool {
	return z.Name != ""
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

func LoadOrNewCompatibilityDefinition(path string) (cds CompatibilityDefinitions, err error) {
	if _, err := os.Lstat(path); os.IsNotExist(err) {
		cds = CompatibilityDefinitions{
			Prune:       make([]PruneDefinition, 0),
			PathChanges: make([]PathChangeDefinition, 0),
		}
	} else {
		cdsBody, err := os.ReadFile(path)
		if err != nil {
			return cds, err
		}
		cds, err = ParseCompatibilityDefinition(cdsBody)
		if err != nil {
			return cds, err
		}
	}
	return cds, nil
}

func SaveCompatibilityDefinition(path string, cds CompatibilityDefinitions, compact bool) (err error) {
	var cdsNewBody []byte
	slices.SortFunc(cds.Prune, func(a, b PruneDefinition) int {
		return strings.Compare(a.Current.CliPath(), b.Current.CliPath())
	})
	slices.SortFunc(cds.PathChanges, func(a, b PathChangeDefinition) int {
		return strings.Compare(a.Current.CliPath(), b.Current.CliPath())
	})
	if compact {
		cdsNewBody, err = json.Marshal(cds)
		if err != nil {
			return err
		}
	} else {
		cdsNewBody, err = json.MarshalIndent(cds, "", "  ")
		if err != nil {
			return err
		}
	}
	if err := os.WriteFile(path, cdsNewBody, 0644); err != nil {
		return err
	}
	return nil
}

func ParseCompatibilityDefinition(data []byte) (cd CompatibilityDefinitions, err error) {
	err = json.Unmarshal(data, &cd)
	if err != nil {
		return cd, err
	}

	for _, pc := range cd.PathChanges {
		if !pc.Current.IsValid() {
			l := esl.Default()
			l.Error("Invalid path change", esl.Any("pathChange", pc))
			panic("invalid path change")
		}
		for _, fp := range pc.FormerPaths {
			if !fp.IsValid() {
				l := esl.Default()
				l.Error("Invalid former path", esl.Any("pathChange", pc))
				panic("invalid former path")
			}
		}
	}
	for _, p := range cd.Prune {
		if !p.Current.IsValid() {
			l := esl.Default()
			l.Error("Invalid prune", esl.Any("prune", p))
			panic("invalid prune")
		}
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

func (z CompatibilityDefinitions) FindPlannedPrune(path []string, name string) (cd PruneDefinition, found bool) {
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

func (z CompatibilityDefinitions) FindPrunedPrune(path []string, name string) (cd PruneDefinition, found bool) {
	l := esl.Default()
	cd, found = z.FindPrune(path, name)
	if !found {
		return cd, false
	}
	if IsAlive(cd.PruneAfterBuildDate) {
		return cd, false
	} else {
		l.Debug("Prune after date", esl.String("pruneAfter", cd.PruneAfterBuildDate))
		return cd, true
	}
}
func (z CompatibilityDefinitions) ListPlannedPrune() (cds []PruneDefinition) {
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
