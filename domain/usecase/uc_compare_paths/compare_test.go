package uc_compare_paths

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"testing"
)

func TestEndToEndCompareImpl_Diff(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		tf1Path := qt_api.ToolboxTestSuiteFolder.ChildPath("compare/compare1")
		tf2Path := qt_api.ToolboxTestSuiteFolder.ChildPath("compare/compare2")

		testFolder1, err := sv_file.NewFiles(ctx).Resolve(tf1Path)
		if err != nil {
			ctx.Log().Warn("Test folder1 not found", zap.Error(err))
			return
		}
		tf1, e := testFolder1.Folder()
		if !e {
			ctx.Log().Warn("Test folder1 is not a shared folder")
			return
		}

		testFolder2, err := sv_file.NewFiles(ctx).Resolve(tf2Path)
		if err != nil {
			ctx.Log().Warn("Test folder2 not found", zap.Error(err))
			return
		}
		tf2, e := testFolder2.Folder()
		if !e {
			ctx.Log().Warn("Test folder2 is not a shared folder")
			return
		}

		// Compare 1 to 1, should not have diff (with path root)
		ctx.Log().Info("Compare 1 to 1 (with Path Root)")
		{
			ucc := New(
				ctx.WithPath(api_context.Namespace(tf1.SharedFolderId())),
				ctx.WithPath(api_context.Namespace(tf1.SharedFolderId())),
				app_ui.NewDummy(),
			)
			count, err := ucc.Diff(mo_path.NewDropboxPath(""), mo_path.NewDropboxPath(""), func(diff mo_file_diff.Diff) error {
				t.Error(diff)
				return nil
			})
			if count > 0 || err != nil {
				t.Error(count, err)
			}
		}

		// Compare 1 to 1, should not have diff (with path args)
		ctx.Log().Info("Compare 1 to 1 (with Path Args)")
		{
			ucc := New(
				ctx,
				ctx,
				app_ui.NewDummy(),
			)
			count, err := ucc.Diff(tf1Path, tf1Path, func(diff mo_file_diff.Diff) error {
				t.Error(diff)
				return nil
			})
			if count > 0 || err != nil {
				t.Error(count, err)
			}
		}

		// Compare 2 to 2, should not have diff
		ctx.Log().Info("Compare 2 to 2")
		{
			ucc := New(
				ctx.WithPath(api_context.Namespace(tf2.SharedFolderId())),
				ctx.WithPath(api_context.Namespace(tf2.SharedFolderId())),
				app_ui.NewDummy(),
			)
			count, err := ucc.Diff(mo_path.NewDropboxPath(""), mo_path.NewDropboxPath(""), func(diff mo_file_diff.Diff) error {
				t.Error(diff)
				return nil
			})
			if count > 0 || err != nil {
				t.Error(count, err)
			}
		}

		// Compare 1 to 2, should have two diffs (with path root)
		ctx.Log().Info("Compare 1 to 2 (path root)")
		{
			ucc := New(
				ctx.WithPath(api_context.Namespace(tf1.SharedFolderId())),
				ctx.WithPath(api_context.Namespace(tf2.SharedFolderId())),
				app_ui.NewDummy(),
			)
			count, err := ucc.Diff(mo_path.NewDropboxPath(""), mo_path.NewDropboxPath(""), func(diff mo_file_diff.Diff) error {
				ctx.Log().Debug("diff", zap.Any("diff", diff))
				return nil
			})
			if count != 2 || err != nil {
				t.Error(count, err)
			}
		}

		// Compare 2 to 1, should have two diffs (with paths)
		ctx.Log().Info("Compare 2 to 1 (path args)")
		{
			ucc := New(
				ctx,
				ctx,
				app_ui.NewDummy(),
			)
			count, err := ucc.Diff(tf2Path, tf1Path, func(diff mo_file_diff.Diff) error {
				ctx.Log().Debug("diff", zap.Any("diff", diff))
				return nil
			})
			if count != 2 || err != nil {
				t.Error(count, err)
			}
		}
	})
}

// Mock test

func TestCompareImpl_Diff(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ctx := api_context_impl.NewMock(ctl)
		sv := New(ctx, ctx, ctl.UI())
		_, err := sv.Diff(qt_recipe.NewTestDropboxFolderPath(), qt_recipe.NewTestDropboxFolderPath(), func(diff mo_file_diff.Diff) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
