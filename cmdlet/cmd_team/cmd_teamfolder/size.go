package cmd_teamfolder

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type CmdTeamTeamFolderSize struct {
	*cmdlet.SimpleCommandlet
	optDepth     int
	optCachePath string
	report       report.Factory
}

func (CmdTeamTeamFolderSize) Name() string {
	return "size"
}

func (CmdTeamTeamFolderSize) Desc() string {
	return "Calculate size of team folder"
}

func (CmdTeamTeamFolderSize) Usage() string {
	return ""
}

func (z *CmdTeamTeamFolderSize) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descOptDepth := "Depth directories deep"
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)

	descUseCached := "Use cached information, or create cache if not exist"
	f.StringVar(&z.optCachePath, "cache", "", descUseCached)
}

func (z *CmdTeamTeamFolderSize) Exec(args []string) {
	z.report.Init(z.Log())
	defer z.report.Close()

	nsz := &NamespaceSizes{}
	nsz.Init()
	if z.optCachePath != "" && z.isCacheFilesAvailable() {
		z.Log().Info("Calculating size from cache")
		z.loadFromCache(nsz)
	} else {
		z.Log().Info("Retrieve data from API, then calculating size")
		z.loadFromApi(nsz)
	}

	z.Log().Info("Reporting result")
	for _, sz := range nsz.Sizes {
		z.report.Report(sz)
	}
}

func (z *CmdTeamTeamFolderSize) loadFromCache(nsz *NamespaceSizes) {
	if f, err := os.Open(z.cachePathFile()); err != nil {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				z.Log().Warn(
					"Unable to read file",
					zap.String("file", z.cachePathFile()),
					zap.Error(err),
				)
				break
			}

			nsf := &dbx_namespace.NamespaceFile{}
			if err = json.Unmarshal(line, nsf); err != nil {
				z.Log().Warn(
					"Unable to parse line",
					zap.String("file", z.cachePathFile()),
					zap.String("line", string(line)),
					zap.Error(err),
				)
				break
			}

			nsz.OnFile(nsf)
		}
		f.Close()
	}

	if f, err := os.Open(z.cachePathFolder()); err != nil {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				z.Log().Warn(
					"Unable to read file",
					zap.String("file", z.cachePathFolder()),
					zap.Error(err),
				)
				break
			}

			nsf := &dbx_namespace.NamespaceFolder{}
			if err = json.Unmarshal(line, nsf); err != nil {
				z.Log().Warn(
					"Unable to parse line",
					zap.String("file", z.cachePathFolder()),
					zap.String("line", string(line)),
					zap.Error(err),
				)
				break
			}

			nsz.OnFolder(nsf)
		}
		f.Close()
	}

}

func (z *CmdTeamTeamFolderSize) loadFromApi(nsz *NamespaceSizes) {
	apiFile, err := z.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		z.DefaultErrorHandler(ea)
		return
	}
	cacheWriter := report.Factory{}
	cacheWriter.ReportPath = z.optCachePath
	cacheWriter.ReportFormat = "json"
	cacheWriter.DefaultWriter = ioutil.Discard
	cacheWriter.Init(z.Log())

	namespaceFile := dbx_namespace.ListNamespaceFile{}
	namespaceFile.AsAdminId = admin.TeamMemberId
	namespaceFile.OnError = z.DefaultErrorHandler
	namespaceFile.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
		if namespace.NamespaceType != "team_folder" {
			z.Log().Debug(
				"Skip non `team_folder` namespace",
				zap.String("namespace_type", namespace.NamespaceType),
				zap.String("namespace_id", namespace.NamespaceId),
				zap.String("name", namespace.Name),
			)
			return false
		}
		z.Log().Info("Scanning folder",
			zap.String("namespace_type", namespace.NamespaceType),
			zap.String("namespace_id", namespace.NamespaceId),
			zap.String("name", namespace.Name),
		)
		return true
	}
	namespaceFile.OnFolder = func(folder *dbx_namespace.NamespaceFolder) bool {
		nsz.OnFolder(folder)
		cacheWriter.Report(folder)
		return true
	}
	namespaceFile.OnFile = func(file *dbx_namespace.NamespaceFile) bool {
		nsz.OnFile(file)
		cacheWriter.Report(file)
		return true
	}
	namespaceFile.OnDelete = func(deleted *dbx_namespace.NamespaceDeleted) bool {
		nsz.OnDelete(deleted)
		cacheWriter.Report(deleted)
		return true
	}
	namespaceFile.List(apiFile)
}

func (z *CmdTeamTeamFolderSize) cachePathFile() string {
	return filepath.Join(z.optCachePath, "NamespaceFile.json")
}
func (z *CmdTeamTeamFolderSize) cachePathFolder() string {
	return filepath.Join(z.optCachePath, "NamespaceFolder.json")
}

func (z *CmdTeamTeamFolderSize) isCacheFilesAvailable() bool {
	avail := false
	if _, err := os.Stat(z.cachePathFile()); !os.IsNotExist(err) {
		avail = true
	}
	if _, err := os.Stat(z.cachePathFolder()); !os.IsNotExist(err) {
		avail = true
	}
	return avail
}

type NamespaceSize struct {
	Namespace       *dbx_namespace.Namespace `json:"namespace"`
	Path            string                   `json:"path"`
	FileCount       int64                    `json:"file_count"`
	FolderCount     int64                    `json:"folder_count"`
	DescendantCount int64                    `json:"descendant_count"`
	Size            int64                    `json:"size"`
}

type NamespaceSizes struct {
	Sizes map[string]*NamespaceSize
	Depth int
}

func (z *NamespaceSizes) Init() {
	z.Sizes = make(map[string]*NamespaceSize)
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
	if len(pathElems) >= z.Depth+1 {
		dp = strings.Join(pathElems[:z.Depth+1], "/")
	} else {
		dp = strings.Join(pathElems, "/")
	}
	if dp == "" {
		return namespace.NamespaceId + "/"
	}

	return namespace.NamespaceId + dp
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
