package main

import (
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_run"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRunbook(t *testing.T) {
	rbPath := filepath.Join(filepath.Dir(os.Args[0]), app_run.RunBookTestName)
	rb := &app_run.RunBook{
		Entry: []*app_run.RunEntry{
			{
				Args: []string{"dev", "echo", "-text", "Hey"},
			},
			{
				Args: []string{"dev", "echo", "-text", "Be quiet", "-quiet"},
			},
			{
				Args: []string{"dev", "echo", "-text", "Low memory", "-low-memory"},
			},
		},
	}
	rbContent, err := json.Marshal(rb)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(rbContent))

	if err = ioutil.WriteFile(rbPath, rbContent, 0644); err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(rbPath)

	if runBook, found := app_run.FindRunBook(true); found {
		bx := rice.MustFindBox("resources")
		web := rice.MustFindBox("web")
		runBook.Exec(bx, web)
	} else {
		t.Error("run book not found")
	}
}
