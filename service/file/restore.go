package file

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dustin/go-humanize"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/dbsugar"
	"github.com/watermint/toolbox/infra/progress"
	"path/filepath"
	"strings"
)

type RestoreContext struct {
	db          *dbsugar.DatabaseSugar
	dbFile      string
	cleanedPath string
	dbxCfgFull  dropbox.Config

	Infra     *infra.InfraOpts
	Path      string
	TokenFull string
	Preflight bool
}

const (
	restore_database_filename = "restore.db"

	restore_table_file             = "restore_file"
	restore_create_table_operation = `
	CREATE TABLE {{.TableName}} (
	  path_display                    VARCHAR PRIMARY KEY,
	  rev                             VARCHAR,
	  size                            INT8,
	  client_modified                 TIMESTAMP,
	  server_modified                 TIMESTAMP,
	  server_deleted                  TIMESTAMP
	)`
)

const (
	restore_step_prepare_execution_plan = iota
	restore_step_scan_path
	restore_step_restore
)

type restoreStepFunc func() error

func (r *RestoreContext) Restore() error {
	stepStart := restore_step_prepare_execution_plan
	stepEnd := restore_step_restore

	if r.Preflight {
		stepEnd = restore_step_scan_path
	}

	steps := make(map[int]restoreStepFunc)
	title := make(map[int]string)

	title[restore_step_prepare_execution_plan] = "Prepare execution plan"
	steps[restore_step_prepare_execution_plan] = r.restoreStepExecutionPlan

	title[restore_step_scan_path] = "Scan files and folders"
	steps[restore_step_scan_path] = r.restoreStepScanPath

	title[restore_step_restore] = "Restore"
	steps[restore_step_restore] = r.restoreStepRestore

	for s := stepStart; s <= stepEnd; s++ {
		seelog.Infof(
			"[Step %d of %d] %s",
			s+1,
			stepEnd+1,
			title[s],
		)

		if err := steps[s](); err != nil {
			seelog.Debugf("Abort step[%d] due to error[%s]", s+1, err)
			return err
		}
	}

	return nil
}

func (r *RestoreContext) restoreStepExecutionPlan() error {
	r.dbxCfgFull = dropbox.Config{Token: r.TokenFull}
	r.cleanedPath = filepath.ToSlash(filepath.Clean(r.Path))
	if !strings.HasPrefix(r.cleanedPath, "/") {
		r.cleanedPath = "/" + r.cleanedPath
	}
	if r.cleanedPath == "/" {
		r.cleanedPath = ""
	}

	r.dbFile = r.Infra.FileOnWorkPath(restore_database_filename)
	r.db = &dbsugar.DatabaseSugar{DataSourceName: r.dbFile}
	if err := r.db.Open(); err != nil {
		seelog.Warnf("Unable to open file path[%s] : error[%s]", r.dbFile, err)
		return err
	}

	ddl := make(map[string]string)
	ddl[restore_table_file] = restore_create_table_operation

	for tn, d := range ddl {
		if err := r.db.CreateTable(tn, d); err != nil {
			seelog.Warnf("Failed to prepare table[%s]: error[%s]", tn, err)
			return err
		}
	}

	return nil
}

func (r *RestoreContext) restoreStepScanPath() error {
	client := files.New(r.dbxCfgFull)
	arg := files.NewGetMetadataArg(r.cleanedPath)
	arg.IncludeDeleted = true
	meta, err := client.GetMetadata(arg)
	if err != nil {
		seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", r.cleanedPath, err)
		return err
	}
	err = r.scanDispatch(meta)
	if err != nil {
		seelog.Warnf("Failed to scan folders : error[%s]", err)
		return err
	}

	fileCount, fileSize, err := r.scanStats()
	if err != nil {
		return err
	}
	seelog.Infof("%s deleted file(s), total %s",
		humanize.Comma(fileCount),
		humanize.IBytes(fileSize),
	)

	return nil
}

func (r *RestoreContext) scanStats() (fileCount int64, fileSize uint64, err error) {
	err = r.db.QueryRow(
		`SELECT COUNT(path_display), SUM(size) FROM {{.TableName}}`,
		restore_table_file,
	).Scan(
		&fileCount,
		&fileSize,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve scan stats : error[%s]", err)
		return
	}
	return
}

func (r *RestoreContext) scanDispatch(meta files.IsMetadata) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		// ignore
		return nil

	case *files.FolderMetadata:
		seelog.Debugf("ScanDispatch: FolderId[%s] PathDisplay[%s]", f.Id, f.PathDisplay)
		return r.scanFolder(f)

	case *files.DeletedMetadata:
		seelog.Debugf("ScanDispatch: Deleted[%s]", f.PathDisplay)
		return r.scanDeleted(f)
	}
	return errors.New("unexpected metadata type")
}

func (r *RestoreContext) scanFolder(meta *files.FolderMetadata) error {
	seelog.Debugf("ScanFolder: Path[%s]", meta.PathDisplay)
	client := files.New(r.dbxCfgFull)
	arg := files.NewListFolderArg(meta.PathDisplay)
	arg.IncludeDeleted = true
	more := true

	lf, err := client.ListFolder(arg)
	if err != nil {
		seelog.Warnf("Unable to load folder[%s] : error[%s]", meta.PathDisplay, err)
		return err
	}
	for more {
		for _, f := range lf.Entries {
			err = r.scanDispatch(f)
			if err != nil {
				seelog.Warnf("Unable to prepare file/folder meta data[%s] : error[%s]", err)
				return err
			}
		}
		if lf.HasMore {
			lf, err = client.ListFolderContinue(files.NewListFolderContinueArg(lf.Cursor))
			if err != nil {
				seelog.Warnf("Unable to load folder (cont) [%s] : error[%s]", meta.PathDisplay, err)
				return err
			}
		}
		more = lf.HasMore
	}
	return nil
}

func (r *RestoreContext) scanDeleted(meta *files.DeletedMetadata) error {
	seelog.Debugf("ScanDeleted: Path[%s]", meta.PathDisplay)
	client := files.New(r.dbxCfgFull)

	arg := files.NewListRevisionsArg(meta.PathDisplay)
	arg.Limit = 1

	res, err := client.ListRevisions(arg)
	if err != nil && strings.HasPrefix(err.Error(), "path/not_file") {
		seelog.Debugf("Skip for deleted folder[%s]", meta.PathDisplay)
		return nil
	}
	if err != nil {
		seelog.Warnf("Unable to list revisions for path[%s] : error[%s]", meta.PathDisplay, err)
		return err
	}
	if res.Entries == nil || len(res.Entries) < 1 {
		seelog.Warnf("Skip: Revision not found for path[%s]", meta.PathDisplay)
		return nil
	}

	rev := res.Entries[0]
	seelog.Debugf("ScanDeleted: Path[%s] Revision[%s] Size[%d]", rev.PathDisplay, rev.Rev, rev.Size)
	_, err = r.db.Exec(
		`INSERT INTO {{.TableName}} (
		  path_display,
		  rev,
		  size,
		  client_modified,
		  server_modified,
		  server_deleted
		) VALUES (?, ?, ?, ?, ?, ?)`,
		restore_table_file,
		rev.PathDisplay,
		rev.Rev,
		rev.Size,
		rev.ClientModified,
		rev.ServerModified,
		res.ServerDeleted,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare restore data : error[%s]", err)
		return err
	}
	return nil
}

func (r *RestoreContext) restoreStepRestore() error {
	seelog.Debugf("Restore Start")
	fileCount, _, err := r.scanStats()
	if err != nil {
		return err
	}

	pui := progress.ProgressUI{Infra: r.Infra}
	pui.Start(int(fileCount))

	rows, err := r.db.Query(
		`SELECT path_display, rev FROM {{.TableName}}`,
		restore_table_file,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve scan data : error[%s]", err)
		return err
	}
	client := files.New(r.dbxCfgFull)

	for rows.Next() {
		var pathDisplay, rev string
		err = rows.Scan(&pathDisplay, &rev)
		if err != nil {
			seelog.Warnf("Unable to retrieve path/rev from scan data: error[%s]", err)
			return err
		}

		seelog.Debugf("Restore: Path[%s] Rev[%s]", pathDisplay, rev)
		arg := files.NewRestoreArg(pathDisplay, rev)
		_, err := client.Restore(arg)
		if err != nil {
			seelog.Warnf("Unable to restore file[%s] : error[%s]", pathDisplay, err)
			return err
		}
		pui.Incr()
	}
	pui.End()
	seelog.Debugf("Restore Finished")

	return nil
}
