package es_rotate

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_array"
	"github.com/watermint/toolbox/essentials/collections/es_value"
	"github.com/watermint/toolbox/essentials/file/es_gzip"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	UnlimitedBackups = -1
	UnlimitedQuota   = -1
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
	l := es_log.Default()

	entries0, err := ioutil.ReadDir(z.BasePath())
	if err != nil {
		l.Warn("Unable to read log directory", es_log.String("path", z.BasePath()), es_log.Error(err))
		return nil, err
	}
	entries = make([]os.FileInfo, 0)
	for _, entry := range entries0 {
		name := entry.Name()
		if !strings.HasPrefix(name, z.BaseName()) {
			continue
		}
		if !strings.HasSuffix(name, logFileExtension) && !strings.HasSuffix(name, es_gzip.SuffixCompress) {
			continue
		}
		if _, ok := outInProgress.Load(filepath.Join(z.BasePath(), name)); ok {
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (z RotateOpts) targetsByCount(entries []os.FileInfo) (purge es_array.Array) {
	if z.numBackups == UnlimitedBackups || len(entries) < z.numBackups {
		return es_array.Empty()
	}

	numLogs := len(entries)
	numPurge := numLogs - z.numBackups
	if numPurge < 1 {
		return es_array.Empty()
	}

	return es_array.NewByFileInfo(entries...).
		Sort().
		Left(numPurge)
}

func (z RotateOpts) targetsByQuota(entries []os.FileInfo) (purge es_array.Array) {
	if z.quota == UnlimitedQuota {
		return es_array.Empty()
	}

	var used int64
	all := es_array.NewByFileInfo(entries...)
	preserve := all.Sort().RightWhile(
		func(v es_value.Value) bool {
			fi := v.AsInterface().(os.FileInfo)
			used += fi.Size()
			return used <= z.quota
		},
	)
	return all.Diff(preserve)
}

func (z RotateOpts) PurgeTargets() (purge []string, err error) {
	logs, err := z.CurrentLogs()
	if err != nil {
		return nil, err
	}

	byCount := z.targetsByCount(logs)
	byQuota := z.targetsByQuota(logs)

	purge = byCount.Union(byQuota).Map(func(v es_value.Value) es_value.Value {
		fi := v.AsInterface().(os.FileInfo)
		return es_value.New(filepath.Join(z.BasePath(), fi.Name()))
	}).AsStringArray()
	return
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
			l := es_log.Default()
			l.Warn("Invalid number of log backups", es_log.Int("num", num))
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
			l := es_log.Default()
			l.Warn("Invalid quota size", es_log.Int64("quota", quota))
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
