package file

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dustin/go-humanize"
	"github.com/gosuri/uiprogress"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/sdk"
	"path/filepath"
	"strings"
	"time"
)

type MoveContext struct {
	db              *sql.DB
	dbFile          string
	cleanedSrcBase  string
	cleanedSrcPaths []string
	cleanedDestPath string
	dbxCfgFull      dropbox.Config

	Infra         *infra.InfraOpts
	SrcPath       string
	DestPath      string
	Preflight     bool
	PreflightAnon bool
	BatchSize     int
	TokenFull     string
}

const (
	MOVE_BATCH_MAX_SIZE       = 9000
	MOVE_BATCH_RETRY_INTERVAL = 30
	MOVE_BATCH_CHECK_INTERVAL = 3

	move_create_table_file = `
	CREATE TABLE {{.TableName}} (
	  file_id                         VARCHAR PRIMARY KEY,
	  parent_folder_id                VARCHAR,
	  name                            VARCHAR,
	  rev                             VARCHAR,
	  size                            INT8,
	  path_lower                      VARCHAR,
	  path_display                    VARCHAR,
	  content_hash                    VARCHAR(32),
	  sharing_read_only               BOOL,
	  sharing_parent_shared_folder_id VARCHAR
	)
	`

	move_create_table_folder = `
	CREATE TABLE {{.TableName}} (
	  folder_id                       VARCHAR PRIMARY KEY,
	  parent_folder_id                VARCHAR,
	  depth                           INT,
	  name                            VARCHAR,
	  path_lower                      VARCHAR,
	  path_display                    VARCHAR,
	  sharing_read_only               BOOL,
	  sharing_parent_shared_folder_id VARCHAR,
	  sharing_shared_folder_id        VARCHAR,
	  sharing_traverse_only           BOOL,
	  sharing_no_access               BOOL
	)
	`

	move_create_table_shared_folder = `
	CREATE TABLE {{.TableName}} (
	  shared_folder_id                VARCHAR PRIMARY KEY,
	  name                            VARCHAR,
	  is_inside_team_folder           VARCHAR,
	  is_team_folder                  VARCHAR,
	  parent_shared_folder_id         VARCHAR,
	  path_lower                      VARCHAR
	)
	`

	move_table_target_file          = "target_file"
	move_table_target_folder        = "target_folder"
	move_table_target_shared_folder = "target_shared_folder"
	move_table_dest_file            = "dest_file"
	move_table_dest_folder          = "dest_folder"
	move_table_dest_shared_folder   = "dest_shared_folder"
)

const (
	move_step_prepare_execution_plan = iota
	move_step_confirm_execution_plan
	move_step_prepare_scan
	move_step_scan_target_folders
	move_step_scan_target_shared_folders
	move_step_validate
	move_step_create_destination_folders
	move_step_move_files
	move_step_cleanup
)

type moveStepFunc func() error

func (m *MoveContext) Move() error {
	stepStart := move_step_prepare_execution_plan
	stepEnd := move_step_cleanup

	if m.Preflight {
		stepEnd = move_step_validate
	}

	steps := make(map[int]moveStepFunc)
	title := make(map[int]string)

	title[move_step_prepare_execution_plan] = "Prepare execution plan"
	steps[move_step_prepare_execution_plan] = m.stepPrepareExecutionPlan

	title[move_step_confirm_execution_plan] = "Confirm execution plan"
	steps[move_step_confirm_execution_plan] = m.stepConfirmExecutionPlan

	title[move_step_prepare_scan] = "Preparing scan files and folders"
	steps[move_step_prepare_scan] = m.stepPrepareScan

	title[move_step_scan_target_folders] = "Scan target files and folders"
	steps[move_step_scan_target_folders] = m.stepScanTargetFolders

	title[move_step_scan_target_shared_folders] = "Scan sharing information"
	steps[move_step_scan_target_shared_folders] = m.stepScanTargetSharedFolders

	title[move_step_validate] = "Validate permissions of files/folders"
	steps[move_step_validate] = m.stepValidatePermissions

	title[move_step_create_destination_folders] = "Create destination folders"
	steps[move_step_create_destination_folders] = m.stepCreateDestFolders

	title[move_step_move_files] = "Move file(s)"
	steps[move_step_move_files] = m.stepMoveFiles

	title[move_step_cleanup] = "Clean up folders of source folder"
	steps[move_step_cleanup] = m.stepCleanupSourceFolders

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

func cleanPath(path string) (string, error) {
	c := filepath.ToSlash(filepath.Clean(path))
	if !strings.HasPrefix(c, "/") {
		c = "/" + c
	}
	if strings.HasSuffix(path, "/") {
		c = c + "/"
	}
	return c, nil
}

func nameInMetadata(meta files.IsMetadata) string {
	switch f := meta.(type) {
	case *files.FileMetadata:
		return f.Name
	case *files.FolderMetadata:
		return f.Name
	case *files.DeletedMetadata:
		return ""

	default:
		seelog.Debug("Unknown metadata type")
		return ""
	}
}

func (m *MoveContext) stepPrepareExecutionPlan() error {
	var err error
	m.dbxCfgFull = dropbox.Config{Token: m.TokenFull}
	client := files.New(m.dbxCfgFull)

	if MOVE_BATCH_MAX_SIZE < m.BatchSize {
		seelog.Infof("Batch size reduced to avoid hitting API upper limit: %d -> %d", m.BatchSize, MOVE_BATCH_MAX_SIZE)
		m.BatchSize = MOVE_BATCH_MAX_SIZE
	}

	// Prepare dest path
	m.cleanedDestPath, err = cleanPath(m.DestPath)
	if err != nil {
		seelog.Warnf("Unable to clean dest path[%s] : error[%s]", m.DestPath, err)
		return err
	}

	// Prepare src path
	csp, err := cleanPath(m.SrcPath)
	m.cleanedSrcBase = csp
	m.cleanedSrcPaths = make([]string, 0)
	if err != nil {
		seelog.Warnf("Unable to clean src path[%s] : error[%s]", m.SrcPath, err)
		return err
	}

	if !strings.HasSuffix(csp, "/") {
		// Ensure the file/folder exist
		meta, err := client.GetMetadata(files.NewGetMetadataArg(csp))
		if err != nil {
			seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", csp, err)
			return err
		}
		n := nameInMetadata(meta)
		if n != "" {
			seelog.Warnf("File or folder not found for path[%s]", csp)
			return errors.New("file or folder not found")
		}
		m.cleanedSrcPaths = append(m.cleanedSrcPaths, csp)

	} else {
		// Expand src if the path has the suffix "/".

		listArg := files.NewListFolderArg(csp)
		lf, err := client.ListFolder(listArg)
		if err != nil {
			seelog.Warnf("Unable to load folder[%s] : error[%s]", csp, err)
			return err
		}
		more := true
		for more {
			for _, f := range lf.Entries {
				n := nameInMetadata(f)
				if n != "" {
					m.cleanedSrcPaths = append(m.cleanedSrcPaths, filepath.ToSlash(filepath.Join(csp, n)))
				}
			}
			if lf.HasMore {
				lf, err = client.ListFolderContinue(files.NewListFolderContinueArg(lf.Cursor))
				if err != nil {
					seelog.Warnf("Unable to load folder (cont) [%s] : error[%s]", csp, err)
					return err
				}
			}
			more = lf.HasMore
		}
	}

	if len(m.cleanedSrcPaths) < 1 {
		seelog.Warnf("No target files/folders found.")
		return errors.New("no target files or folders found")
	}

	return nil
}

func (m *MoveContext) stepConfirmExecutionPlan() error {

	seelog.Info("Execution plan:")
	for _, sp := range m.cleanedSrcPaths {
		b := filepath.Base(sp)
		seelog.Infof("%s", sp)
		seelog.Infof("-> %s", filepath.ToSlash(filepath.Join(m.cleanedDestPath, b)))
	}

	if m.Preflight {
		seelog.Info("Skip confirmation for preflight mode")
		return nil
	}

	phrase := "move"
	cancel := "cancel"
	code := ""
	seelog.Flush()

	for code != phrase {
		fmt.Printf("Query: please confirm execution plan and, type [%s]. To cancel execution, type [%s].\n", phrase, cancel)

		if _, err := fmt.Scan(&code); err != nil {
			seelog.Errorf("Input error (%s), try again.", err)
			continue
		}
		trim := strings.TrimSpace(code)
		if len(trim) < 1 {
			seelog.Errorf("Input error, try again.")
			continue
		}
		if trim == cancel {
			return errors.New("operation cancelled by user")
		}
	}

	return nil
}

func (m *MoveContext) createTable(tableName, ddlTmpl string) error {
	dropTable := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := m.db.Exec(dropTable)
	if err != nil {
		seelog.Warnf("Unable to drop table [%s] : error[%s]", tableName, err)
		return err
	}

	ddl, err := util.CompileTemplate(ddlTmpl, struct {
		TableName string
	}{
		TableName: tableName,
	})

	if err != nil {
		seelog.Warnf("Unable to compile TableName[%s], DDL[%s] : error[%s]", tableName, ddlTmpl, err)
		return err
	}

	_, err = m.db.Exec(ddl)
	if err != nil {
		seelog.Warnf("Unable to create table [%s] : error[%s]", tableName, err)
		return err
	}
	return nil
}

func (m *MoveContext) stepPrepareScan() error {
	var err error

	m.dbFile = m.Infra.FileOnWorkPath("move.db")
	m.db, err = sql.Open("sqlite3", m.dbFile)

	if err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", m.dbFile, err)
		return err
	}

	ddl := make(map[string]string)
	ddl[move_table_dest_file] = move_create_table_file
	ddl[move_table_dest_folder] = move_create_table_folder
	ddl[move_table_dest_shared_folder] = move_create_table_shared_folder
	ddl[move_table_target_file] = move_create_table_file
	ddl[move_table_target_folder] = move_create_table_folder
	ddl[move_table_target_shared_folder] = move_create_table_shared_folder

	for tn, d := range ddl {
		err = m.createTable(tn, d)
		if err != nil {
			seelog.Warnf("Failed to prepare table[%s]: error[%s]", tn, err)
			return err
		}
	}

	return nil
}

func (m *MoveContext) stepScanTargetFolders() error {
	seelog.Flush()
	uip := uiprogress.New()
	uip.Start()
	defer uip.Stop()
	bar := uip.AddBar(len(m.cleanedSrcPaths))
	bar.PrependElapsed()
	bar.AppendCompleted()
	for _, s := range m.cleanedSrcPaths {
		err := m.scan(s)
		if err != nil {
			seelog.Warnf("Failed to scan target files/folders : error[%s]", err)
			return err
		}
		bar.Incr()
	}

	m.scanReport()

	return nil
}

func (m *MoveContext) scanReport() {
	var fCount int64
	var fSize uint64
	var folderCount int64

	q := `
	SELECT COUNT(file_id), SUM(size) FROM target_file
	`

	row := m.db.QueryRow(q)
	err := row.Scan(
		&fCount,
		&fSize,
	)
	if err != nil {
		seelog.Debugf("Unable to retrieve file size/count : error[%s]", err)
		return
	}

	q = `
	SELECT COUNT(folder_id) FROM target_folder
	`

	row = m.db.QueryRow(q)
	err = row.Scan(
		&folderCount,
	)
	if err != nil {
		seelog.Debugf("Unable to retrieve folder count : error[%s]", err)
		return
	}

	seelog.Infof(
		"Found: %s folders, %s files, total %s",
		humanize.Comma(folderCount),
		humanize.Comma(fCount),
		humanize.IBytes(fSize),
	)
}

func (m *MoveContext) scan(src string) error {
	client := files.New(m.dbxCfgFull)
	meta, err := client.GetMetadata(files.NewGetMetadataArg(src))
	if err != nil {
		seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", src, err)
		return err
	}

	return m.scanDispatch(meta, nil)
}

func (m *MoveContext) scanDispatch(meta files.IsMetadata, parentFolder *files.FolderMetadata) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		return m.scanFile(f, parentFolder)

	case *files.FolderMetadata:
		return m.scanFolder(f, parentFolder)

	case *files.DeletedMetadata:
		seelog.Warnf("Deleted file cannot move [%s]", f.PathDisplay)
		return errors.New("deleted file cannot move")

	default:
		seelog.Debug("Unknown metadata type")
		return errors.New("unknown type of file/folder found")
	}
}

func (m *MoveContext) scanFile(meta *files.FileMetadata, parentFolder *files.FolderMetadata) error {
	q := `
	INSERT OR REPLACE INTO target_file (
	  file_id,
	  parent_folder_id,
	  name,
	  rev,
	  size,
	  path_lower,
	  path_display,
	  content_hash,
	  sharing_read_only,
	  sharing_parent_shared_folder_id
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	sharingReadOnly := false
	sharingParentSharedFolderId := ""

	if meta.SharingInfo != nil {
		sharingReadOnly = meta.SharingInfo.ReadOnly
		sharingParentSharedFolderId = meta.SharingInfo.ParentSharedFolderId
	}

	name := meta.Name
	pathLower := meta.PathLower
	pathDisplay := meta.PathDisplay

	if m.Preflight && m.PreflightAnon {
		name = ""
		pathLower = ""
		pathDisplay = ""
	}

	parentFolderId := ""
	if parentFolder != nil {
		parentFolderId = parentFolder.Id
	}

	_, err := m.db.Exec(
		q,
		meta.Id,
		parentFolderId,
		name,
		meta.Rev,
		meta.Size,
		pathLower,
		pathDisplay,
		meta.ContentHash,
		sharingReadOnly,
		sharingParentSharedFolderId,
	)

	if err != nil {
		seelog.Warnf("Unable to prepare target file meta data[%s] : error[%s]", meta.PathDisplay, err)
		return err
	}

	return nil
}

func (m *MoveContext) scanFolder(meta, parentFolder *files.FolderMetadata) error {
	client := files.New(m.dbxCfgFull)

	q := `
	INSERT OR REPLACE INTO target_folder (
	  folder_id,
	  parent_folder_id,
	  depth,
	  name,
	  path_lower,
	  path_display,
	  sharing_read_only,
	  sharing_parent_shared_folder_id,
	  sharing_shared_folder_id,
	  sharing_traverse_only,
	  sharing_no_access
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	sharingReadOnly := false
	sharingParentSharedFolderId := ""
	sharingSharedFolderId := ""
	sharingTraverseOnly := false
	sharingNoAccess := false

	if meta.SharingInfo != nil {
		sharingReadOnly = meta.SharingInfo.ReadOnly
		sharingParentSharedFolderId = meta.SharingInfo.ParentSharedFolderId
		sharingSharedFolderId = meta.SharingInfo.SharedFolderId
		sharingTraverseOnly = meta.SharingInfo.TraverseOnly
		sharingNoAccess = meta.SharingInfo.NoAccess
	}

	name := meta.Name
	pathLower := meta.PathLower
	pathDisplay := meta.PathDisplay

	if m.Preflight && m.PreflightAnon {
		name = ""
		pathLower = ""
		pathDisplay = ""
	}

	parentFolderId := ""
	if parentFolder != nil {
		parentFolderId = parentFolder.Id
	}

	_, err := m.db.Exec(
		q,
		meta.Id,
		parentFolderId,
		len(strings.Split(meta.PathLower, "/")),
		name,
		pathLower,
		pathDisplay,
		sharingReadOnly,
		sharingParentSharedFolderId,
		sharingSharedFolderId,
		sharingTraverseOnly,
		sharingNoAccess,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare target folder meta data[%s] : error[%s]", meta.PathDisplay, err)
		return err
	}

	listArg := files.NewListFolderArg(meta.PathLower)

	lf, err := client.ListFolder(listArg)
	if err != nil {
		seelog.Warnf("Unable to load folder[%s] : error[%s]", meta.PathDisplay, err)
		return err
	}
	more := true
	for more {
		for _, f := range lf.Entries {
			err = m.scanDispatch(f, meta)
			if err != nil {
				seelog.Warnf("Unable to prepare file/folder meta data[%s] : error[%s]", meta.PathDisplay, err)
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

func (m *MoveContext) stepScanTargetSharedFolders() error {
	q := `
	SELECT DISTINCT sharing_shared_folder_id FROM target_folder WHERE sharing_shared_folder_id <> ""
	`
	rows, err := m.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve shared_folder info : error[%s]", err)
		return err
	}

	sharedFolderIds := make([]string, 0)

	for rows.Next() {
		sharedFolderId := ""
		err = rows.Scan(
			&sharedFolderId,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}

		sharedFolderIds = append(sharedFolderIds, sharedFolderId)
	}
	rows.Close()

	if len(sharedFolderIds) < 1 {
		return nil
	}
	seelog.Infof("%s shared folder(s) found in source path",
		humanize.Comma(int64(len(sharedFolderIds))))

	for _, sf := range sharedFolderIds {
		m.scanSharedFolderInfo(sf)
	}

	return nil
}

func (m *MoveContext) scanSharedFolderInfo(sharedFolderId string) error {
	client := sharing.New(m.dbxCfgFull)
	meta, err := client.GetFolderMetadata(sharing.NewGetMetadataArgs(sharedFolderId))
	if err != nil {
		seelog.Warnf("Unable to load shared_folder metadata [%s] : error[%s]", sharedFolderId, err)
		return err
	}

	q := `
	INSERT OR REPLACE INTO target_shared_folder (
	  shared_folder_id,
	  name,
	  is_inside_team_folder,
	  is_team_folder,
	  parent_shared_folder_id,
	  path_lower
	) VALUES (?, ?, ?, ?, ?, ?)
	`

	name := meta.Name
	pathLower := meta.PathLower

	if m.Preflight && m.PreflightAnon {
		name = ""
		pathLower = ""
	}

	_, err = m.db.Exec(
		q,
		meta.SharedFolderId,
		name,
		meta.IsInsideTeamFolder,
		meta.IsTeamFolder,
		meta.ParentSharedFolderId,
		pathLower,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare shared folder info : error[%s]", err)
		return err
	}

	return nil
}

func (m *MoveContext) numberOfSharedFoldersInSrc() (int, error) {
	q := `
	SELECT COUNT(DISTINCT sharing_shared_folder_id) FROM target_folder WHERE sharing_shared_folder_id <> ""
	`

	cnt := 0
	row := m.db.QueryRow(q)
	err := row.Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count shared folders : error[%s]", err)
		return 0, err
	}

	return cnt, nil
}

func (m *MoveContext) stepValidatePermissions() error {
	cntReadOnly := 0
	cntTraverseOnly := 0
	cntNoAccess := 0

	// Validate files
	q := `
	SELECT
	  SUM(sharing_read_only)
	FROM
	  target_file
	`

	row := m.db.QueryRow(q)
	err := row.Scan(
		&cntReadOnly,
	)
	if err != nil {
		seelog.Warnf("Unable to validate target files : error[%s]", err)
		return err
	}

	if cntReadOnly > 0 {
		seelog.Errorf("%d file(s) are read only. Cannot move those files.", cntReadOnly)
		return errors.New("file(s) are read only")
	}

	// Validate folders
	q = `
	SELECT
	  SUM(sharing_read_only),
	  SUM(sharing_traverse_only),
	  SUM(sharing_no_access)
	FROM
	  target_folder
	`

	row = m.db.QueryRow(q)
	err = row.Scan(
		&cntReadOnly,
		&cntTraverseOnly,
		&cntNoAccess,
	)
	if err != nil {
		seelog.Warnf("Unable to validate target files : error[%s]", err)
		return err
	}

	if cntReadOnly > 0 {
		seelog.Errorf("%d folder(s) are read only. Cannot move those folder.", cntReadOnly)
		return errors.New("folder(s) are read only")
	}
	if cntTraverseOnly > 0 {
		seelog.Errorf("%d folder(s) are traverse only. Cannot move those folder.", cntReadOnly)
		return errors.New("folder(s) are traverse only")
	}
	if cntNoAccess > 0 {
		seelog.Errorf("%d folder(s) are no access. Cannot move those folder.", cntReadOnly)
		return errors.New("folder(s) are no access")
	}

	return nil
}

func (m *MoveContext) stepCreateDestFolders() error {
	q := `
	SELECT COUNT(path_display) FROM target_folder
	`

	cnt := 0
	row := m.db.QueryRow(q)
	err := row.Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count folders : error[%s]", err)
		return err
	}

	seelog.Infof("Create %d folder(s) in destination", cnt)

	q = `
	SELECT path_display FROM target_folder ORDER BY depth
	`

	rows, err := m.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve shared_folder info : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	seelog.Flush()
	uip := uiprogress.New()
	uip.Start()
	defer uip.Stop()
	bar := uip.AddBar(cnt)
	bar.PrependElapsed()
	bar.AppendCompleted()

	for rows.Next() {
		pathDisplay := ""
		err = rows.Scan(
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve folder name : error[%s]", err)
			return err
		}

		rel, err := filepath.Rel(m.cleanedSrcBase, pathDisplay)
		if err != nil {
			seelog.Warnf("Unable to calculate relative path[%s] : error[%s]", pathDisplay, err)
			return err
		}
		destPath := filepath.ToSlash(filepath.Join(m.cleanedDestPath, rel))
		seelog.Debugf("Create destination folder[%s]", destPath)

		arg := files.NewCreateFolderArg(destPath)
		arg.Autorename = false
		res, err := client.CreateFolderV2(arg)
		if err != nil {
			if strings.HasPrefix(err.Error(), "path/conflict") {
				seelog.Debugf("Folder [%s] already exists. Skip", destPath)
			} else {
				seelog.Warnf("Unable to create folder[%s] : error[%s]", destPath, err)
				return err
			}
		} else {
			seelog.Debugf("Destination folder created[%s] folder_id[%s]", destPath, res.Metadata.Id)
		}

		bar.Incr()
	}

	return nil
}

func (m *MoveContext) stepMoveFiles() error {
	q := `
	SELECT COUNT(path_display) FROM target_file
	`

	cnt := 0
	row := m.db.QueryRow(q)
	err := row.Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count file(s) : error[%s]", err)
		return err
	}

	seelog.Infof("Move %d files into destination", cnt)

	seelog.Flush()
	uip := uiprogress.New()
	uip.Start()
	defer uip.Stop()
	bar := uip.AddBar(cnt)
	bar.PrependElapsed()
	bar.AppendCompleted()

	q = `
	SELECT path_display FROM target_file
	`

	rows, err := m.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve shared_folder info : error[%s]", err)
		return err
	}

	batch := make([]*files.RelocationPath, 0)

	for rows.Next() {
		pathDisplay := ""
		err = rows.Scan(
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve folder name : error[%s]", err)
			return err
		}

		rel, err := filepath.Rel(m.cleanedSrcBase, pathDisplay)
		if err != nil {
			seelog.Warnf("Unable to calculate relative path[%s] : error[%s]", pathDisplay, err)
			return err
		}

		destPath := filepath.ToSlash(filepath.Join(m.cleanedDestPath, rel))
		seelog.Debugf("Moving file from [%s] to [%s]", pathDisplay, destPath)

		batch = append(batch, files.NewRelocationPath(pathDisplay, destPath))
		if m.BatchSize <= len(batch) {
			m.moveBatch(batch, bar)
			batch = make([]*files.RelocationPath, 0)
		}
	}

	if 0 < len(batch) {
		m.moveBatch(batch, bar)
	}

	return nil
}

func (m *MoveContext) moveBatch(batch []*files.RelocationPath, bar *uiprogress.Bar) error {
	retry := false
	for {
		if retry {
			seelog.Debugf("Re-trying to move %d files", len(batch))
		} else {
			seelog.Debugf("Trying to move %d files", len(batch))
		}
		arg := files.NewRelocationBatchArg(batch)
		arg.AllowSharedFolder = true
		arg.AllowOwnershipTransfer = true
		arg.Autorename = false
		mbRes, err := sdk.ZMoveBatch(m.dbxCfgFull, arg)
		if err != nil {
			seelog.Warnf("Unable to call `move_batch` : error[%s]", err)
			return err
		}

		asyncJobId := ""
		switch mbRes.Tag {
		case "complete":
			seelog.Debugf("Move completed for %d files", len(batch))
			for i := 0; i < len(batch); i++ {
				bar.Incr()
			}
			return nil

		case "async_job_id":
			asyncJobId = mbRes.AsyncJobId
			seelog.Debugf("`async_job_id` issued [%s]", asyncJobId)
			break

		default:
			seelog.Warnf("Unknown tag for /files/move_batch: [%s]", mbRes.Tag)
			return errors.New("unknown tag")
		}

		for {
			ckRes, err := sdk.ZMoveBatchCheck(m.dbxCfgFull, async.NewPollArg(asyncJobId))
			if err != nil {
				seelog.Warnf("Unable to call `/files/move_batch_check` : error[%s]", err)
				return err
			}

			switch ckRes.Tag {
			case "complete":
				seelog.Debugf("Move completed by async job for %d files", len(batch))
				for i := 0; i < len(batch); i++ {
					bar.Incr()
				}
				return nil

			case "in_progress":
				seelog.Debugf("`async_job_id`[%s] is in progress", asyncJobId)
				time.Sleep(MOVE_BATCH_CHECK_INTERVAL * time.Second)

			case "failed":
				if strings.HasPrefix(ckRes.Failed.Tag, "too_many_write_operations") {
					seelog.Debugf("`too_many_write_operations`: Wait for seconds")
					time.Sleep(MOVE_BATCH_RETRY_INTERVAL * time.Second)
					retry = true

				} else {
					seelog.Warnf("Unable to move %d files: error [%s]", len(batch), ckRes.Failed.Tag)
					for i := 0; i < len(batch); i++ {
						bar.Incr()
					}

					return errors.New(ckRes.Failed.Tag)
				}
			}

			if retry {
				break
			}
		}
	}
}

func (m *MoveContext) stepCleanupSourceFolders() error {
	q := `
	SELECT COUNT(path_display) FROM target_folder
	`

	cnt := 0
	row := m.db.QueryRow(q)
	err := row.Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count folders : error[%s]", err)
		return err
	}

	seelog.Infof("Remove %d folder(s) in source", cnt)

	q = `
	SELECT path_display FROM target_folder ORDER BY depth DESC
	`

	rows, err := m.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve folder info : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	seelog.Flush()
	uip := uiprogress.New()
	uip.Start()
	defer uip.Stop()
	bar := uip.AddBar(cnt)
	bar.PrependElapsed()
	bar.AppendCompleted()

	for rows.Next() {
		pathDisplay := ""
		err = rows.Scan(
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve folder name : error[%s]", err)
			return err
		}

		lf, err := client.ListFolder(files.NewListFolderArg(pathDisplay))
		if err != nil {
			seelog.Warnf("Unable to list folder[%s] : error[%s]", pathDisplay, err)
			return err
		}
		if len(lf.Entries) > 0 {
			seelog.Warnf("Unable to remove folder[%s]. File exists.", pathDisplay)
			return errors.New("unable to remove folder")
		}

		arg := files.NewDeleteArg(pathDisplay)
		_, err = client.DeleteV2(arg)
		if err != nil {
			seelog.Warnf("Unable to remove folder[%s] : error[%s]", pathDisplay, err)
			return err
		}

		bar.Incr()
	}

	return nil
}
