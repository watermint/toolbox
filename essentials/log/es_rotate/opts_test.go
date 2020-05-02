package es_rotate

import "testing"

func TestRotateOpts_Apply(t *testing.T) {
	o := RotateOpts{}.Apply(
		BasePath("/somewhere"),
		BaseName("toolbox"),
		ChunkSize(1024),
		NumBackup(10),
		HookBeforeDelete(func(path string) {
		}),
	)
	if o.chunkSize != 1024 || o.numBackups != 10 || o.basePath != "/somewhere" || o.baseName != "toolbox" {
		t.Error(o)
	}
}
