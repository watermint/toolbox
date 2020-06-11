package es_demo

import "testing"

func TestTermImpl_Write(t *testing.T) {
	ti := &termImpl{}
	ti.Write([]byte("watermint toolbox 69.4.123"))
}
