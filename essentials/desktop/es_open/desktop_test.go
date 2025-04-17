package es_open

import (
    "os"
    "testing"
)

func TestCurrentDesktop(t *testing.T) {
    if !testing.Verbose() {
        t.Skip("Skip test")
    }
    d := CurrentDesktop()
    p, err := os.MkdirTemp("", "desktop")
    if err != nil {
        t.Error(err)
        return
    }
    err = d.Open(p)
    if err == nil {
        t.Log("success")
    } else {
        t.Log("open failure or unsupported", err)
    }
}
