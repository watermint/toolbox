package esl_rotate

import (
	"crypto/rand"
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_array_deprecated"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRotateOpts_Apply(t *testing.T) {
	o := NewRotateOpts().Apply(
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

func TestRotateOpts_PurgeTargets_ByCount(t *testing.T) {
	qt_file.TestWithTestFolder(t, "by_count", false, func(path string) {
		numBackup := 10
		numPurged := 5
		ro := NewRotateOpts().Apply(
			BasePath(path),
			BaseName("by_count"),
			NumBackup(numBackup),
		)

		// old logs
		expectedFiles := make([]string, 0)
		for i := 0; i < numPurged; i++ {
			name := fmt.Sprintf("by_count.%04d.log", i)
			fp := filepath.Join(path, name)
			if err := ioutil.WriteFile(fp, []byte(name), 0644); err != nil {
				t.Error(err)
			}
			expectedFiles = append(expectedFiles, fp)
		}

		// new logs
		preserveFiles := make([]string, 0)
		for i := 0; i < numBackup; i++ {
			name := fmt.Sprintf("by_count.%04d.log", i+1000)
			fp := filepath.Join(path, name)
			if err := ioutil.WriteFile(fp, []byte(name), 0644); err != nil {
				t.Error(err)
			}
			preserveFiles = append(preserveFiles, fp)
		}

		targetFiles, err := ro.PurgeTargets()
		if err != nil {
			t.Error(err)
		}

		pa := es_array_deprecated.NewByString(preserveFiles...)
		ea := es_array_deprecated.NewByString(expectedFiles...)
		ta := es_array_deprecated.NewByString(targetFiles...)

		// Preserve/Target
		if cf := pa.Intersection(ta); cf.Size() != 0 {
			t.Error(cf)
		}

		// Purge/Target
		if cf := ea.Intersection(ta); cf.Size() != numPurged {
			t.Error(ea)
		}
	})
}

func TestRotateOpts_PurgeTargets_ByQuota(t *testing.T) {
	qt_file.TestWithTestFolder(t, "by_quota", false, func(path string) {
		quota := 1000
		chunk := 100
		numFiles := quota / chunk
		numPurge := numFiles / 2
		ro := NewRotateOpts().Apply(
			BasePath(path),
			BaseName("by_quota"),
			Quota(int64(quota)),
		)

		// old logs
		expectedFiles := make([]string, 0)
		for i := 0; i < numPurge; i++ {
			name := fmt.Sprintf("by_quota.%04d.log", i)
			data := make([]byte, chunk)
			if _, err := rand.Read(data); err != nil {
				t.Error(err)
			}
			fp := filepath.Join(path, name)
			if err := ioutil.WriteFile(fp, data, 0644); err != nil {
				t.Error(err)
			}
			expectedFiles = append(expectedFiles, fp)
		}

		// new logs
		preserveFiles := make([]string, 0)
		for i := 0; i < numFiles; i++ {
			name := fmt.Sprintf("by_quota.%04d.log", i+1000)
			data := make([]byte, chunk)
			if _, err := rand.Read(data); err != nil {
				t.Error(err)
			}
			fp := filepath.Join(path, name)
			if err := ioutil.WriteFile(fp, data, 0644); err != nil {
				t.Error(err)
			}
			preserveFiles = append(preserveFiles, fp)
		}

		targetFiles, err := ro.PurgeTargets()
		if err != nil {
			t.Error(err)
		}

		pa := es_array_deprecated.NewByString(preserveFiles...)
		ea := es_array_deprecated.NewByString(expectedFiles...)
		ta := es_array_deprecated.NewByString(targetFiles...)

		// Preserve/Target
		if cf := pa.Intersection(ta); cf.Size() != 0 {
			t.Error(cf)
		}

		// Purge/Target
		if cf := ea.Intersection(ta); cf.Size() != numPurge {
			t.Error(ea)
		}
	})
}
