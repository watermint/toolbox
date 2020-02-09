package main

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	run([]string{os.Args[0], "dev", "echo", "-text", "Hey"}, true)
}

func TestRunbook(t *testing.T) {
	rbPath := filepath.Join(filepath.Dir(os.Args[0]), app_workflow.RunBookTestName)
	rb := &app_workflow.RunBook{
		Version: "1",
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
