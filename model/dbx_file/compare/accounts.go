package compare

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_file"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type BetweenAccounts struct {
	ExecContext       *app.ExecContext
	LeftApi           *dbx_api.Context
	LeftAccountAlias  string
	LeftAsMemberId    string
	LeftAsAdminId     string
	LeftPath          string
	LeftPathRoot      interface{}
	RightApi          *dbx_api.Context
	RightAccountAlias string
	RightAsMemberId   string
	RightAsAdminId    string
	RightPath         string
	RightPathRoot     interface{}
	OnDiff            func(diff Diff)
}

type Diff struct {
	DiffType  string `json:"diff_type"`
	LeftPath  string `json:"left_path,omitempty"`
	LeftKind  string `json:"left_kind,omitempty"`
	LeftSize  *int64 `json:"left_size,omitempty"`
	LeftHash  string `json:"left_hash,omitempty"`
	RightPath string `json:"right_path,omitempty"`
	RightKind string `json:"right_kind,omitempty"`
	RightSize *int64 `json:"right_size,omitempty"`
	RightHash string `json:"right_hash,omitempty"`
}

func (z *BetweenAccounts) handleError(err error) bool {
	z.ExecContext.Log().Debug("ignore errors", zap.Error(err))
	return true
}

func (z *BetweenAccounts) ignoreDeleted(deleted *dbx_file.Deleted) bool {
	return true
}

func (z *BetweenAccounts) compareLevel(path string) int {
	diffCount := 0
	leftFiles := make(map[string]*dbx_file.File)
	leftFolders := make(map[string]*dbx_file.Folder)
	rightFiles := make(map[string]*dbx_file.File)
	rightFolders := make(map[string]*dbx_file.Folder)

	lp := dbx_file.ListFolder{
		AsMemberId:                      z.LeftAsMemberId,
		AsAdminId:                       z.LeftAsAdminId,
		PathRoot:                        z.LeftPathRoot,
		IncludeHasExplicitSharedMembers: false,
		IncludeDeleted:                  false,
		IncludeMediaInfo:                false,
		IncludeMountedFolders:           true,
		OnError:                         z.handleError,
		OnDelete:                        z.ignoreDeleted,
		OnFolder: func(folder *dbx_file.Folder) bool {
			leftFolders[strings.ToLower(folder.Name)] = folder
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			leftFiles[strings.ToLower(file.Name)] = file
			return true
		},
	}
	leftPath := filepath.ToSlash(filepath.Join(z.LeftPath, path))
	z.ExecContext.Log().Debug("Scanning left path", zap.String("path", leftPath))
	lp.List(z.LeftApi, leftPath)

	rp := dbx_file.ListFolder{
		AsMemberId:                      z.RightAsMemberId,
		AsAdminId:                       z.RightAsAdminId,
		PathRoot:                        z.RightPathRoot,
		IncludeHasExplicitSharedMembers: false,
		IncludeDeleted:                  false,
		IncludeMediaInfo:                false,
		IncludeMountedFolders:           true,
		OnError:                         z.handleError,
		OnDelete:                        z.ignoreDeleted,
		OnFolder: func(folder *dbx_file.Folder) bool {
			rightFolders[strings.ToLower(folder.Name)] = folder
			return true
		},
		OnFile: func(file *dbx_file.File) bool {
			rightFiles[strings.ToLower(file.Name)] = file
			return true
		},
	}
	rightPath := filepath.ToSlash(filepath.Join(z.RightPath, path))
	z.ExecContext.Log().Debug("Scanning right path", zap.String("path", rightPath))
	rp.List(z.RightApi, rightPath)

	// compare files left to right
	for _, lf := range leftFiles {
		if rf, e := rightFiles[strings.ToLower(lf.Name)]; e {
			if lf.ContentHash != rf.ContentHash {
				z.OnDiff(Diff{
					DiffType:  "file_content_diff",
					LeftPath:  lf.PathDisplay,
					LeftKind:  "file",
					LeftSize:  &lf.Size,
					LeftHash:  lf.ContentHash,
					RightPath: rf.PathDisplay,
					RightKind: "file",
					RightSize: &rf.Size,
					RightHash: rf.ContentHash,
				})
				diffCount++
			}
		} else {
			z.OnDiff(Diff{
				DiffType: "right_file_missing",
				LeftPath: lf.PathDisplay,
				LeftKind: "file",
				LeftSize: &lf.Size,
				LeftHash: lf.ContentHash,
			})
			diffCount++
		}
	}

	// compare files right to left
	for _, rf := range rightFiles {
		if _, e := leftFiles[strings.ToLower(rf.Name)]; !e {
			z.OnDiff(Diff{
				DiffType:  "left_file_missing",
				RightPath: rf.PathDisplay,
				RightKind: "file",
				RightSize: &rf.Size,
				RightHash: rf.ContentHash,
			})
			diffCount++
		}
	}

	// compare folders left to right
	for _, lf := range leftFolders {
		if _, e := rightFolders[strings.ToLower(lf.Name)]; e {
			// proceed to ancestors
			pd, err := filepath.Rel(strings.ToLower(z.LeftPath), lf.PathLower)
			if err != nil {
				z.ExecContext.Log().Warn("unable to calculate relative path", zap.String("leftPathBase", z.LeftPath), zap.String("leftPath", lf.PathLower), zap.Error(err))
				continue
			}
			if strings.HasPrefix(pd, "..") {
				z.ExecContext.Log().Error("invalid relative path", zap.String("pd", pd), zap.String("zLeftPath", z.LeftPath), zap.String("lfPathLower", lf.PathLower))
				continue
			}

			z.ExecContext.Log().Debug("proceed into ancestors", zap.String("path", pd))
			diffCount += z.compareLevel(pd)

		} else {
			z.OnDiff(Diff{
				DiffType: "right_folder_missing",
				LeftPath: lf.PathDisplay,
				LeftKind: "folder",
			})
			diffCount++
		}
	}

	// compare folders right to left
	for _, rf := range rightFolders {
		if _, e := leftFolders[strings.ToLower(rf.Name)]; !e {
			z.OnDiff(Diff{
				DiffType:  "left_folder_missing",
				RightPath: rf.PathDisplay,
				RightKind: "folder",
			})
			diffCount++
		}
	}
	return diffCount
}

func (z *BetweenAccounts) Compare() {
	// start
	z.ExecContext.Msg("dbx_file.compare.progress.start").Tell()

	// compare
	diffs := z.compareLevel("")

	// report result
	z.ExecContext.Msg("dbx_file.compare.progress.done").WithData(struct {
		DiffCount int
	}{
		DiffCount: diffs,
	}).TellSuccess()
}
