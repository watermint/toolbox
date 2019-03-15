package dbx_size

import (
	"bufio"
	"encoding/json"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_namespace"
	"github.com/watermint/toolbox/model/dbx_profile"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type NamespaceSize struct {
	Namespace       *dbx_namespace.Namespace `json:"namespace"`
	Path            string                   `json:"path"`
	FileCount       int64                    `json:"file_count"`
	FolderCount     int64                    `json:"folder_count"`
	DescendantCount int64                    `json:"descendant_count"`
	Size            int64                    `json:"size"`
}

type NamespaceSizes struct {
	Sizes                  map[string]*NamespaceSize
	ec                     *app.ExecContext
	OnError                func(err error) bool
	OptDepth               int
	OptCachePath           string
	OptIncludeAppFolder    bool
	OptIncludeMemberFolder bool
	OptIncludeTeamFolder   bool
	OptIncludeSharedFolder bool
}

func (z *NamespaceSizes) Init(ec *app.ExecContext) {
	z.Sizes = make(map[string]*NamespaceSize)
	z.ec = ec
}

func (z *NamespaceSizes) parent(path string) string {
	pe := strings.Split(path, "/")
	le := len(pe)
	if le <= 2 {
		return pe[0] + "/"
	}
	return strings.Join(pe[:le-1], "/")
}

func (z *NamespaceSizes) increment(namespace *dbx_namespace.Namespace, path string, folder, file, size int64) {
	if sz, ok := z.Sizes[path]; ok {
		sz.DescendantCount += folder + file
		sz.FolderCount += folder
		sz.FileCount += file
		sz.Size += size
	} else {
		sz = &NamespaceSize{
			Namespace:       namespace,
			Path:            path,
			DescendantCount: folder + file,
			FolderCount:     folder,
			FileCount:       file,
			Size:            size,
		}
		z.Sizes[path] = sz
	}

	pp := z.parent(path)
	if pp != path && path != "/" {
		z.increment(namespace, pp, folder, file, size)
	}
}

func (z *NamespaceSizes) Dir(path string) string {
	pathElems := strings.Split(path, "/")
	if len(pathElems) <= 1 {
		return "/"
	} else {
		pathElems = pathElems[:len(pathElems)-1]
		return strings.Join(pathElems, "/")
	}
}

func (z *NamespaceSizes) Path(namespace *dbx_namespace.Namespace, path string) string {
	pathElems := strings.Split(path, "/")
	dp := ""
	if len(pathElems) >= z.OptDepth+1 {
		dp = strings.Join(pathElems[:z.OptDepth+1], "/")
	} else {
		dp = strings.Join(pathElems, "/")
	}
	if dp == "" {
		return "ns:" + namespace.NamespaceId + "/"
	}

	return "ns:" + namespace.NamespaceId + dp
}

func (z *NamespaceSizes) OnFolder(folder *dbx_namespace.NamespaceFolder) bool {
	z.increment(folder.Namespace, z.Path(folder.Namespace, z.Dir(folder.Folder.PathLower)), 1, 0, 0)
	return true
}

func (z *NamespaceSizes) OnFile(file *dbx_namespace.NamespaceFile) bool {
	z.increment(file.Namespace, z.Path(file.Namespace, z.Dir(file.File.PathLower)), 0, 1, file.File.Size)
	return true
}

func (z *NamespaceSizes) OnDelete(deleted *dbx_namespace.NamespaceDeleted) bool {
	// nop
	return true
}

func (z *NamespaceSizes) Load(c *dbx_api.Context) bool {
	if z.OptCachePath != "" && z.isCacheFilesAvailable() {
		z.ec.Log().Info("Calculating size from cache")
		return z.LoadFromCache()
	} else {
		z.ec.Log().Info("Retrieve data from API, then calculating size")
		return z.LoadFromApi(c)
	}
}

func (z *NamespaceSizes) LoadFromCache() bool {
	if f, err := os.Open(z.cachePathFile()); err != nil {
		z.ec.Log().Warn(
			"Skip loading",
			zap.String("file", z.cachePathFile()),
			zap.Error(err),
		)
	} else {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				z.ec.Log().Warn(
					"Unable to read file",
					zap.String("file", z.cachePathFile()),
					zap.Error(err),
				)
				break
			}

			nsf := &dbx_namespace.NamespaceFile{}
			if err = json.Unmarshal(line, nsf); err != nil {
				z.ec.Log().Warn(
					"Unable to parse line",
					zap.String("file", z.cachePathFile()),
					zap.String("line", string(line)),
					zap.Error(err),
				)
				break
			}

			z.OnFile(nsf)
		}
		f.Close()
	}

	if f, err := os.Open(z.cachePathFolder()); err != nil {
		z.ec.Log().Warn(
			"Skip loading",
			zap.String("file", z.cachePathFolder()),
			zap.Error(err),
		)
	} else {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				z.ec.Log().Warn(
					"Unable to read file",
					zap.String("file", z.cachePathFolder()),
					zap.Error(err),
				)
				break
			}

			nsf := &dbx_namespace.NamespaceFolder{}
			if err = json.Unmarshal(line, nsf); err != nil {
				z.ec.Log().Warn(
					"Unable to parse line",
					zap.String("file", z.cachePathFolder()),
					zap.String("line", string(line)),
					zap.Error(err),
				)
				break
			}

			z.OnFolder(nsf)
		}
		f.Close()
	}

	return true
}

func (z *NamespaceSizes) LoadFromApi(c *dbx_api.Context) bool {
	admin, err := dbx_profile.AuthenticatedAdmin(c)
	if err != nil {
		return z.OnError(err)
	}
	cacheWriter := app_report.Factory{}
	cacheWriter.ReportPath = z.OptCachePath
	cacheWriter.ReportFormat = "json"
	cacheWriter.DefaultWriter = ioutil.Discard
	cacheWriter.Init(z.ec)

	nsl := dbx_namespace.ListNamespaceFile{}
	nsl.AsAdminId = admin.TeamMemberId
	nsl.OptIncludeTeamFolder = true
	nsl.OnError = z.OnError
	nsl.OptIncludeMemberFolder = z.OptIncludeMemberFolder
	nsl.OptIncludeAppFolder = z.OptIncludeAppFolder
	nsl.OptIncludeSharedFolder = z.OptIncludeSharedFolder
	nsl.OptIncludeTeamFolder = z.OptIncludeTeamFolder
	nsl.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
		z.ec.Log().Info("Scanning folder",
			zap.String("namespace_type", namespace.NamespaceType),
			zap.String("namespace_id", namespace.NamespaceId),
			zap.String("name", namespace.Name),
		)
		return true
	}
	nsl.OnFolder = func(folder *dbx_namespace.NamespaceFolder) bool {
		z.OnFolder(folder)
		cacheWriter.Report(folder)
		return true
	}
	nsl.OnFile = func(file *dbx_namespace.NamespaceFile) bool {
		z.OnFile(file)
		cacheWriter.Report(file)
		return true
	}
	nsl.OnDelete = func(deleted *dbx_namespace.NamespaceDeleted) bool {
		z.OnDelete(deleted)
		cacheWriter.Report(deleted)
		return true
	}
	return nsl.List(c)
}

func (z *NamespaceSizes) cachePathFile() string {
	return filepath.Join(z.OptCachePath, "NamespaceFile.json")
}
func (z *NamespaceSizes) cachePathFolder() string {
	return filepath.Join(z.OptCachePath, "NamespaceFolder.json")
}

func (z *NamespaceSizes) isCacheFilesAvailable() bool {
	avail := false
	if _, err := os.Stat(z.cachePathFile()); !os.IsNotExist(err) {
		avail = true
	}
	if _, err := os.Stat(z.cachePathFolder()); !os.IsNotExist(err) {
		avail = true
	}
	return avail
}
