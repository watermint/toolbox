package rp_artifact

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Artifact interface {
	// File name
	Name() string

	// File path
	Path() string
}

// Retrieve artifacts from the Job.
// Returns empty array when any errors occurred.
func Artifacts(job app_workspace.Job) (artifacts []Artifact) {
	l := esl.Default()
	artifacts = make([]Artifact, 0)
	path := job.Report()
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		l.Debug("Unable to read directory", esl.Error(err))
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		switch ext {
		case ".csv", ".json", ".xlsx":
			artifacts = append(artifacts, newAft(filepath.Join(path, entry.Name()), entry.Name()))
		}
	}
	return
}

func newAft(path string, name string) Artifact {
	return &aftImpl{
		name: name,
		path: path,
	}
}

type aftImpl struct {
	name string
	path string
}

func (z aftImpl) Name() string {
	return z.name
}

func (z aftImpl) Path() string {
	return z.path
}
