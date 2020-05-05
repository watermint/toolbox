package main

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}
	run([]string{os.Args[0], "dev", "echo", "-text", "Hey"}, true)
}

func TestRunbook(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}

	rbPath := filepath.Join(filepath.Dir(os.Args[0]), app_workflow.RunBookTestName)
	rb := &app_workflow.RunBook{
		Version: 1,
		Steps: []*app_workflow.RunStep{
			{
				Name: "echo-hello",
				Args: []string{"dev", "echo", "-text", "Hello"},
			},
			{
				Name: "echo-world",
				Args: []string{"dev", "echo", "-text", "World"},
			},
		},
		Workers: []*app_workflow.RunWorker{
			{
				Name: "parallelA",
				Steps: []*app_workflow.RunStep{
					{
						Name: "echo-para1",
						Args: []string{"dev", "echo", "-text", "Parallel A1"},
					},
					{
						Name: "echo-para2",
						Args: []string{"dev", "echo", "-text", "Parallel A2"},
					},
				},
			},
			{
				Name: "parallelB",
				Steps: []*app_workflow.RunStep{
					{
						Name: "echo-para1",
						Args: []string{"dev", "echo", "-text", "Parallel B1"},
					},
					{
						Name: "echo-para2",
						Args: []string{"dev", "echo", "-text", "Parallel B2"},
					},
				},
			},
		},
	}
	rbContent, err := json.Marshal(rb)
	if err != nil {
		t.Error(err)
		return
	}

	if err = ioutil.WriteFile(rbPath, rbContent, 0644); err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(rbPath)

	run([]string{os.Args[0]}, true)
}

func TestRunbookLoop(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}

	p, err := ioutil.TempDir("", "loop")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(p)
	rbPath := filepath.Join(p, "loop.runbook")
	rb := &app_workflow.RunBook{
		Version: 1,
		Steps: []*app_workflow.RunStep{
			{
				Name: "echo-hello",
				Args: []string{"dev", "echo", "-text", "Hello"},
			},
		},
	}
	rbContent, err := json.Marshal(rb)
	if err != nil {
		t.Error(err)
	}

	if err = ioutil.WriteFile(rbPath, rbContent, 0644); err != nil {
		t.Error(err)
		return
	}

	run([]string{os.Args[0],
		"job", "loop",
		"-runbook-path", rbPath,
		"-until", time.Now().Add(2 * 1000 * time.Millisecond).Format(time.RFC3339),
		"-interval-seconds", "1",
	}, true)
}
