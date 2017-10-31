package file

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dustin/go-humanize"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/dbsugar"
	"github.com/watermint/toolbox/infra/progress"
	"path/filepath"
	"strings"
	"time"
)

type RestoreContext struct {
	db              *dbsugar.DatabaseSugar
	dbFile          string
	cleanedBasePath string
	scanBasePath    string
	dbxCfgFull      dropbox.Config

	Infra           *infra.InfraOpts
	BasePath        string
	TokenFull       string
	Preflight       bool
	FilterTimeAfter *time.Time
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
	restore_step_prepare_resources = iota
	restore_step_find_base
	restore_step_scan_path
	restore_step_restore
)

type restoreStepFunc func() error

func (r *RestoreContext) Restore() error {
	stepStart := restore_step_prepare_resources
	stepEnd := restore_step_restore

	if r.Preflight {
		stepEnd = restore_step_scan_path
	}

	steps := make(map[int]restoreStepFunc)
	title := make(map[int]string)

	title[restore_step_prepare_resources] = "Prepare resources"
	steps[restore_step_prepare_resources] = r.restoreStepPrepareResources

	title[restore_step_find_base] = "Find base path of scan"
	steps[restore_step_find_base] = r.restoreStepFindBase

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

func (r *RestoreContext) restoreStepPrepareResources() error {
	r.dbxCfgFull = dropbox.Config{Token: r.TokenFull}
	r.cleanedBasePath = filepath.ToSlash(filepath.Clean(r.BasePath))
	if !strings.HasPrefix(r.cleanedBasePath, "/") {
		r.cleanedBasePath = "/" + r.cleanedBasePath
	}
	if r.cleanedBasePath == "/" {
		r.cleanedBasePath = ""
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

func (r *RestoreContext) restoreStepFindBase() error {
	client := files.New(r.dbxCfgFull)
	sp := r.cleanedBasePath
	for 0 < len(sp) {
		seelog.Debugf("FindBase: GetMetadata[%s]", sp)
		arg := files.NewGetMetadataArg(sp)
		arg.IncludeDeleted = true
		meta, err := client.GetMetadata(arg)
		if err != nil {
			seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", r.cleanedBasePath, err)
			return err
		}
		switch f := meta.(type) {
		case *files.DeletedMetadata:
			pp := filepath.ToSlash(filepath.Dir(sp))
			if pp == "/" {
				seelog.Debugf("Scan from root")
				r.scanBasePath = ""
				return nil
			}
			seelog.Debugf("Path[%s] found as deleted. Dig into parent [%s]", sp, pp)
			sp = pp

		case *files.FolderMetadata:
			seelog.Debugf("Scan root path: [%s]", f.PathDisplay)
			r.scanBasePath = f.PathDisplay
			return nil

		default:
			seelog.Warnf("Unexpected metadata type: path[%s]", sp)
			return errors.New("unexpected metadata type")
		}
	}

	return nil
}

func (r *RestoreContext) restoreStepScanPath() error {
	r.scanFolder(r.scanBasePath, true)

	fileCount, fileSize, err := r.scanStats()
	if err != nil {
		return err
	}
	seelog.Infof("   Scan result: %s deleted file(s), total %s",
		humanize.Comma(fileCount),
		humanize.IBytes(fileSize),
	)
	fileCount, fileSize, err = r.scanFilteredStats()
	if err != nil {
		return err
	}
	seelog.Infof("Restore target: %s deleted file(s), total %s",
		humanize.Comma(fileCount),
		humanize.IBytes(fileSize),
	)

	return nil
}

func (r *RestoreContext) scanStats() (fileCount int64, fileSize uint64, err error) {
	err = r.db.QueryRow(
		`SELECT COUNT(path_display), IFNULL(SUM(size), 0) FROM {{.TableName}}`,
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

func (r *RestoreContext) scanFilteredStats() (fileCount int64, fileSize uint64, err error) {
	rows, err := r.filteredQuery("COUNT(path_display), IFNULL(SUM(size), 0)")
	if err != nil {
		seelog.Warnf("Unable to retrieve scan stats : error[%s]", err)
		return
	}
	if !rows.Next() {
		seelog.Warnf("Unable to retrieve scan stats : error[%s]", err)
		err = errors.New("no row found")
		return
	}
	rows.Scan(
		&fileCount,
		&fileSize,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve scan stats : error[%s]", err)
		return
	}
	return
}

func (r *RestoreContext) scanFolder(pathDisplay string, recursive bool) error {
	seelog.Debugf("ScanFolder: Path[%s]", pathDisplay)
	client := files.New(r.dbxCfgFull)
	arg := files.NewListFolderArg(pathDisplay)
	arg.Recursive = recursive
	arg.IncludeDeleted = true
	arg.IncludeMountedFolders = false

	lf, err := client.ListFolder(arg)
	if err != nil {
		seelog.Warnf("Unable to load folder[%s] : error[%s]", pathDisplay, err)
		return err
	}
	if len(lf.Entries) == 1 {
		switch f := lf.Entries[0].(type) {
		case *files.FolderMetadata:
			if f.PathDisplay == pathDisplay {
				seelog.Debugf("Retry with recursive=false: Path[%s]", pathDisplay)
				return r.scanFolder(pathDisplay, false)
			}
		}
	}

	more := true
	for more {
		for _, f := range lf.Entries {
			switch g := f.(type) {
			case *files.FileMetadata:
				// ignore

			case *files.FolderMetadata:
				seelog.Tracef("ScanDispatch: FolderId[%s] PathDisplay[%s]", g.Id, g.PathDisplay)
				if !recursive && g.PathDisplay != pathDisplay {
					err = r.scanFolder(g.PathDisplay, false)
					if err != nil {
						seelog.Warnf("Unable to prepare file/folder meta data[%s] : error[%s]", err)
						return err
					}
				}

			case *files.DeletedMetadata:
				seelog.Tracef("ScanDispatch: Deleted[%s]", g.PathDisplay)
				err = r.scanDeleted(g)
				if err != nil {
					seelog.Warnf("Unable to prepare file/folder meta data[%s] : error[%s]", err)
					return err
				}

			default:
				return errors.New("unexpected metadata type")
			}
		}
		if lf.HasMore {
			lf, err = client.ListFolderContinue(files.NewListFolderContinueArg(lf.Cursor))
			if err != nil {
				seelog.Warnf("Unable to load folder (cont) [%s] : error[%s]", pathDisplay, err)
				return err
			}
		}
		more = lf.HasMore
	}
	return nil
}

func (r *RestoreContext) scanDeleted(meta *files.DeletedMetadata) error {
	seelog.Tracef("ScanDeleted: Path[%s]", meta.PathDisplay)

	if !strings.HasPrefix(meta.PathDisplay, r.cleanedBasePath) {
		seelog.Tracef("Skip: Base[%s] Path[%s]", r.cleanedBasePath, meta.PathDisplay)
		return nil
	}

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
	fileCount, _, err := r.scanFilteredStats()
	if err != nil {
		return err
	}
	if fileCount < 1 {
		seelog.Infof("No restore target files")
		return nil
	}

	pui := progress.ProgressUI{Infra: r.Infra}
	pui.Start(int(fileCount))

	rows, err := r.filteredQuery("path_display, rev")
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

func (r *RestoreContext) filteredQuery(columns string) (rows *sql.Rows, err error) {
	if r.FilterTimeAfter != nil {
		rows, err = r.db.Query(
			fmt.Sprintf("SELECT %s FROM {{.TableName}} WHERE ? <= server_deleted", columns),
			restore_table_file,
			r.FilterTimeAfter,
		)
	} else {
		rows, err = r.db.Query(
			fmt.Sprintf("SELECT %s FROM {{.TableName}}", columns),
			restore_table_file,
		)
	}
	return
}
