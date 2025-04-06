package esl_rotate

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/watermint/toolbox/quality/infra/qt_file"
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
			if err := os.WriteFile(fp, []byte(name), 0644); err != nil {
				t.Error(err)
			}
			expectedFiles = append(expectedFiles, fp)
		}

		// new logs
		preserveFiles := make([]string, 0)
		for i := 0; i < numBackup; i++ {
			name := fmt.Sprintf("by_count.%04d.log", i+1000)
			fp := filepath.Join(path, name)
			if err := os.WriteFile(fp, []byte(name), 0644); err != nil {
				t.Error(err)
			}
			preserveFiles = append(preserveFiles, fp)
		}

		targetFiles, err := ro.PurgeTargets()
		if err != nil {
			t.Error(err)
		}

		// Sort all slices for consistent comparison
		sort.Strings(preserveFiles)
		sort.Strings(expectedFiles)
		sort.Strings(targetFiles)

		// Check for intersections using maps
		preserveMap := make(map[string]bool)
		for _, f := range preserveFiles {
			preserveMap[f] = true
		}

		expectedMap := make(map[string]bool)
		for _, f := range expectedFiles {
			expectedMap[f] = true
		}

		targetMap := make(map[string]bool)
		for _, f := range targetFiles {
			targetMap[f] = true
		}

		// Check preserve/target intersection
		for _, f := range targetFiles {
			if preserveMap[f] {
				t.Errorf("Target file %s should not be in preserve files", f)
			}
		}

		// Check expected/target intersection
		expectedCount := 0
		for _, f := range targetFiles {
			if expectedMap[f] {
				expectedCount++
			}
		}
		if expectedCount != numPurged {
			t.Errorf("Expected %d files to be purged, got %d", numPurged, expectedCount)
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
			if err := os.WriteFile(fp, data, 0644); err != nil {
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
			if err := os.WriteFile(fp, data, 0644); err != nil {
				t.Error(err)
			}
			preserveFiles = append(preserveFiles, fp)
		}

		targetFiles, err := ro.PurgeTargets()
		if err != nil {
			t.Error(err)
		}

		// Sort all slices for consistent comparison
		sort.Strings(preserveFiles)
		sort.Strings(expectedFiles)
		sort.Strings(targetFiles)

		// Check for intersections using maps
		preserveMap := make(map[string]bool)
		for _, f := range preserveFiles {
			preserveMap[f] = true
		}

		expectedMap := make(map[string]bool)
		for _, f := range expectedFiles {
			expectedMap[f] = true
		}

		targetMap := make(map[string]bool)
		for _, f := range targetFiles {
			targetMap[f] = true
		}

		// Check preserve/target intersection
		for _, f := range targetFiles {
			if preserveMap[f] {
				t.Errorf("Target file %s should not be in preserve files", f)
			}
		}

		// Check expected/target intersection
		expectedCount := 0
		for _, f := range targetFiles {
			if expectedMap[f] {
				expectedCount++
			}
		}
		if expectedCount != numPurge {
			t.Errorf("Expected %d files to be purged, got %d", numPurge, expectedCount)
		}
	})
}
