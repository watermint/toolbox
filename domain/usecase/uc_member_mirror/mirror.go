package uc_member_mirror

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Mirror interface {
	Mirror(srcEmail, dstEmail string) error
}

func New(ctxFileSrc, ctxFileDst api_context.Context) Mirror {
	return &mirrorImpl{
		ctxFileSrc: ctxFileSrc,
		ctxFileDst: ctxFileDst,
	}
}

type mirrorImpl struct {
	ctxFileSrc api_context.Context
	ctxFileDst api_context.Context
}

func (z *mirrorImpl) log() *zap.Logger {
	return z.ctxFileSrc.Log()
}

func (z *mirrorImpl) Mirror(srcEmail, dstEmail string) error {
	l := z.log().With(zap.String("srcEmail", srcEmail), zap.String("dstEmail", dstEmail))
	l.Debug("Start mirroring process")

	l.Debug("Lookup member profiles")
	srcProfile, err := sv_member.New(z.ctxFileSrc).ResolveByEmail(srcEmail)
	if err != nil {
		l.Error("Unable to lookup member", zap.String("lookupEmail", srcEmail), zap.Error(err))
		return err
	}
	dstProfile, err := sv_member.New(z.ctxFileDst).ResolveByEmail(dstEmail)
	if err != nil {
		l.Error("Unable to lookup member", zap.String("lookupEmail", dstEmail), zap.Error(err))
		return err
	}

	ctxFileSrcAsMember := z.ctxFileSrc.AsMemberId(srcProfile.TeamMemberId).WithPath(api_context.Namespace(srcProfile.MemberFolderId))
	ctxFileDstAsMember := z.ctxFileDst.AsMemberId(dstProfile.TeamMemberId).WithPath(api_context.Namespace(dstProfile.MemberFolderId))
	ucm := uc_file_mirror.New(ctxFileSrcAsMember, ctxFileDstAsMember)

	// shared folder list of src member (for excluding mirror)
	sharedFolders, err := sv_sharedfolder.New(z.ctxFileSrc.AsMemberId(srcProfile.TeamMemberId)).List()
	if err != nil {
		l.Error("Unable to list shared folders", zap.Error(err))
		return err
	}

	// prepare exclusion paths
	excludePaths := make([]string, 0)
	for _, sf := range sharedFolders {
		if sf.PathLower != "" {
			excludePaths = append(excludePaths, sf.PathLower)
		}
	}

	// ensure the given path `p` have descendant(s) of exclusion
	hasExclusionPath := func(p string) bool {
		hasExclusion := false
		for _, e := range excludePaths {
			r, err := filepath.Rel(p, e)
			if err != nil {
				l.Error("unable to determine file path", zap.Error(err))
			} else {
				// other side of the folder
				switch {
				case strings.HasPrefix(r, "../"):
				case r == "..":
				case r == ".":
				case strings.HasPrefix(r, "/"):
				default:
					hasExclusion = true
				}
			}
		}
		return hasExclusion
	}

	// ensure the given path `p` is a exclusion path or not
	isExclusionPath := func(p string) bool {
		for _, e := range excludePaths {
			if p == e {
				return true
			}
		}
		return false
	}

	// mirror path
	var mirrorPath func(p string) error
	mirrorPath = func(p string) error {
		ll := l.With(zap.String("p", p))
		path := mo_path.NewPath(p)

		// mirror path unless the path is or has exclusion path
		if !hasExclusionPath(p) && !isExclusionPath(p) {
			return ucm.Mirror(path, path)
		}

		ll.Debug("Given path have descendant(s) of exclusion")
		entries, err := sv_file.NewFiles(ctxFileSrcAsMember).List(path)
		if err != nil {
			ll.Error("Unable to list files", zap.Error(err))
			return err
		}

		var lastErr error = nil

		for _, entry := range entries {
			if folder, e := entry.Folder(); e {
				if isExclusionPath(folder.PathLower()) {
					ll.Debug("Skip shared folder", zap.Any("folder", folder))
				} else {
					// recurse into mirror path
					lastErr = mirrorPath(folder.PathLower())
				}
			}

			// copy files without check (no exclusion path for file)
			if file, e := entry.File(); e {
				fp := mo_path.NewPathDisplay(file.PathDisplay())
				if err := ucm.Mirror(fp, fp); err != nil {
					lastErr = err
				}
			}
		}

		return lastErr
	}

	l.Info("Start mirroring files")
	return mirrorPath("/")
}
