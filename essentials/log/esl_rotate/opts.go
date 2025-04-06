package esl_rotate

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/watermint/toolbox/essentials/file/es_gzip"
	"github.com/watermint/toolbox/essentials/log/esl"
)

const (
	UnlimitedBackups = -1
	UnlimitedQuota   = -1
	logFileExtension = ".log"
)

// Hook function that called when the log exceeds num backups.
// The file will be deleted after this function call.
type RotateHook func(path string)

// Rotate options
type RotateOpts struct {
	basePath string
	baseName string

	// Target size of single log file in bytes.
	chunkSize int64

	// Number of backups. No purge executed when this value is `UnlimitedBackups` (-1).
	numBackups int

	// Target storage quota for this logs.
	quota int64

	// Compress log file on rotate.
	compress   bool
	rotateHook RotateHook
}

// outInProgress tracks files that are currently being written to
var outInProgress sync.Map

func NewRotateOpts() RotateOpts {
	return RotateOpts{
		basePath:   "",
		baseName:   "",
		chunkSize:  0,
		numBackups: UnlimitedBackups,
		quota:      UnlimitedQuota,
		compress:   false,
		rotateHook: nil,
	}
}

func (z RotateOpts) IsCompress() bool {
	return z.compress
}

func (z RotateOpts) ChunkSize() int64 {
	if z.chunkSize <= 0 {
		return math.MaxInt64
	}
	return z.chunkSize
}

func (z RotateOpts) BasePath() string {
	return z.basePath
}

func (z RotateOpts) BaseName() string {
	return z.baseName
}

// Generate name of the current log file
func (z RotateOpts) CurrentName() string {
	suffix := fmt.Sprintf(".%16x%s", time.Now().UnixNano(), logFileExtension)
	return z.baseName + suffix
}

// Generate path to the current log file.
func (z RotateOpts) CurrentPath() string {
	return filepath.Join(z.basePath, z.CurrentName())
}

func (z RotateOpts) CurrentLogs() (entries []os.FileInfo, err error) {
	l := esl.ConsoleOnly()

	entries0, err := os.ReadDir(z.BasePath())
	if err != nil {
		l.Warn("Unable to read log directory", esl.String("path", z.BasePath()), esl.Error(err))
		return nil, err
	}
	entries = make([]os.FileInfo, 0)
	for _, entry := range entries0 {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		name := info.Name()
		if !strings.HasPrefix(name, z.BaseName()) {
			continue
		}
		if !strings.HasSuffix(name, logFileExtension) && !strings.HasSuffix(name, es_gzip.SuffixCompress) {
			continue
		}
		if _, ok := outInProgress.Load(filepath.Join(z.BasePath(), name)); ok {
			continue
		}
		entries = append(entries, info)
	}
	return entries, nil
}

func (z RotateOpts) targetsByCount(entries []os.FileInfo) []os.FileInfo {
	if z.numBackups == UnlimitedBackups || len(entries) < z.numBackups {
		return []os.FileInfo{}
	}

	numLogs := len(entries)
	numPurge := numLogs - z.numBackups
	if numPurge < 1 {
		return []os.FileInfo{}
	}

	// Sort entries by modification time (oldest first)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ModTime().Before(entries[j].ModTime())
	})

	// Return the oldest entries up to numPurge
	return entries[:numPurge]
}

func (z RotateOpts) targetsByQuota(entries []os.FileInfo) []os.FileInfo {
	if z.quota == UnlimitedQuota {
		return []os.FileInfo{}
	}

	// Sort entries by modification time (newest first)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ModTime().After(entries[j].ModTime())
	})

	var used int64
	var preserve []os.FileInfo

	// Keep the newest files that fit within the quota
	for _, entry := range entries {
		used += entry.Size()
		if used <= z.quota {
			preserve = append(preserve, entry)
		} else {
			break
		}
	}

	// Find files to purge (those not in preserve)
	purge := make([]os.FileInfo, 0)
	preserveMap := make(map[string]bool)
	for _, p := range preserve {
		preserveMap[p.Name()] = true
	}

	for _, entry := range entries {
		if !preserveMap[entry.Name()] {
			purge = append(purge, entry)
		}
	}

	return purge
}

func (z RotateOpts) PurgeTargets() (purge []string, err error) {
	logs, err := z.CurrentLogs()
	if err != nil {
		return nil, err
	}

	byCount := z.targetsByCount(logs)
	byQuota := z.targetsByQuota(logs)

	// Combine and deduplicate purge targets
	purgeMap := make(map[string]bool)
	for _, entry := range byCount {
		purgeMap[filepath.Join(z.BasePath(), entry.Name())] = true
	}
	for _, entry := range byQuota {
		purgeMap[filepath.Join(z.BasePath(), entry.Name())] = true
	}

	purge = make([]string, 0, len(purgeMap))
	for path := range purgeMap {
		purge = append(purge, path)
	}

	return purge, nil
}

// Apply all opts
func (z RotateOpts) Apply(opts ...RotateOpt) RotateOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		y, w := opts[0], opts[1:]
		return y(z).Apply(w...)
	}
}

type RotateOpt func(o RotateOpts) RotateOpts

// Compress the log file on rotate
func CompressEnabled(enabled bool) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.compress = enabled
		return o
	}
}

// Compress the log file on rotate
func Compress() RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.compress = true
		return o
	}
}

// Stay uncompressed the log file on rotate
func Uncompressed() RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.compress = false
		return o
	}
}

// Path to the log file
func BasePath(path string) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.basePath = path
		return o
	}
}

// Log file name without suffix
func BaseName(name string) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.baseName = name
		return o
	}
}

// Maximum size target for the single log file.
// Log file could exceed this size, but should not exceed too much.
func ChunkSize(size int64) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.chunkSize = size
		return o
	}
}

// Number of backups
func NumBackup(num int) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		if num != UnlimitedBackups && num < 0 {
			l := esl.ConsoleOnly()
			l.Warn("Invalid number of log backups", esl.Int("num", num))
			o.numBackups = 0
		} else {
			o.numBackups = num
		}
		return o
	}
}

func Quota(quota int64) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		if quota != UnlimitedQuota && quota < 0 {
			l := esl.ConsoleOnly()
			l.Warn("Invalid quota size", esl.Int64("quota", quota))
			o.quota = 0
		} else {
			o.quota = quota
		}
		return o
	}
}

// Hook function that called when just before the file deleted.
func HookBeforeDelete(hook RotateHook) RotateOpt {
	return func(o RotateOpts) RotateOpts {
		o.rotateHook = hook
		return o
	}
}
