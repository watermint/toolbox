package file

import (
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dustin/go-humanize"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/dbsugar"
	"github.com/watermint/toolbox/infra/progress"
	"github.com/watermint/toolbox/integration/sdk"
	"path/filepath"
	"strings"
	"time"
)

type MoveContext struct {
	db              *dbsugar.DatabaseSugar
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
	FileByFile    bool
	BatchSize     int
	TokenFull     string
}

const (
	MOVE_BATCH_API_LIMIT      = 9500 // Actual limit is 10K, 500 is for buffer
	MOVE_BATCH_MAX_SIZE       = MOVE_BATCH_API_LIMIT
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
	)`

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
	)`

	move_create_table_shared_folder = `
	CREATE TABLE {{.TableName}} (
	  shared_folder_id                VARCHAR PRIMARY KEY,
	  name                            VARCHAR,
	  is_inside_team_folder           VARCHAR,
	  is_team_folder                  VARCHAR,
	  parent_shared_folder_id         VARCHAR,
	  path_lower                      VARCHAR
	)`

	move_create_table_operation_move = `
	CREATE TABLE {{.TableName}} (
	  operation_id                    VARCHAR PRIMARY KEY,
	  operation                       VARCHAR,
	  move_from                       VARCHAR,
	  move_to                         VARCHAR,
	  file_count                      INT
	)`

	move_create_table_operation_folder = `
	CREATE TABLE {{.TableName}} (
	  path_display                    VARCHAR PRIMARY KEY,
	  operation                       VARCHAR
	)`

	move_table_src_file          = "src_file"
	move_table_src_folder        = "src_folder"
	move_table_src_shared_folder = "src_shared_folder"
	move_table_dst_file          = "dst_file"
	move_table_dst_folder        = "dst_folder"
	move_table_dst_shared_folder = "dst_shared_folder"
	move_table_operation_move    = "opr_move"
	move_table_operation_folder  = "opr_folder"

	move_oper_file          = "file"
	move_oper_folder        = "folder"
	move_oper_create_folder = "create_folder"

	move_scan_src = 1
	move_scan_dst = 2
)

const (
	move_step_prepare_execution_plan = iota
	move_step_confirm_execution_plan
	move_step_prepare_scan
	move_step_scan_src_folders
	move_step_scan_src_shared_folders
	move_step_scan_dst_folders
	move_step_scan_dst_shared_folders
	move_step_validate_src
	move_step_validate_dst
	move_step_prepare_operation_plan
	move_step_execute_create_folder
	move_step_execute_operation_plan
	move_step_cleanup
)

type moveStepFunc func() error

func (m *MoveContext) Move() error {
	stepStart := move_step_prepare_execution_plan
	stepEnd := move_step_cleanup

	if m.Preflight {
		stepEnd = move_step_prepare_operation_plan
	}

	steps := make(map[int]moveStepFunc)
	title := make(map[int]string)

	title[move_step_prepare_execution_plan] = "Prepare execution plan"
	steps[move_step_prepare_execution_plan] = m.stepPrepareExecutionPlan

	title[move_step_confirm_execution_plan] = "Confirm execution plan"
	steps[move_step_confirm_execution_plan] = m.stepConfirmExecutionPlan

	title[move_step_prepare_scan] = "Preparing scan files and folders"
	steps[move_step_prepare_scan] = m.stepPrepareScan

	title[move_step_scan_src_folders] = "Scan source files and folders"
	steps[move_step_scan_src_folders] = m.stepScanSrcFolders

	title[move_step_scan_src_shared_folders] = "Scan source shared folders"
	steps[move_step_scan_src_shared_folders] = m.stepScanSrcSharedFolders

	title[move_step_scan_dst_folders] = "Scan dest files and folders"
	steps[move_step_scan_dst_folders] = m.stepScanDestFolder

	title[move_step_scan_dst_shared_folders] = "Scan dest shared folders"
	steps[move_step_scan_dst_shared_folders] = m.stepScanDestSharedFolders

	title[move_step_validate_src] = "Validate permissions of source files/folders"
	steps[move_step_validate_src] = m.stepValidateSrcPermissions

	title[move_step_validate_dst] = "Validate permissions of dest files/folders"
	steps[move_step_validate_dst] = m.stepValidateDestPermissions

	title[move_step_prepare_operation_plan] = "Prepare operation plan"
	steps[move_step_prepare_operation_plan] = m.stepPrepareOperationPlan

	title[move_step_execute_create_folder] = "Create folders in destination"
	steps[move_step_execute_create_folder] = m.stepExecuteCreateFolder

	title[move_step_execute_operation_plan] = "Execute operation plan"
	steps[move_step_execute_operation_plan] = m.stepExecuteOperationPlan

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
	if c != "/" && strings.HasSuffix(path, "/") {
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
		if n == "" {
			seelog.Warnf("File or folder not found for path[%s]", csp)
			return errors.New("file or folder not found")
		}
		m.cleanedSrcPaths = append(m.cleanedSrcPaths, csp)
		m.cleanedSrcBase = filepath.ToSlash(filepath.Dir(csp))

	} else {
		// Expand src if the path has the suffix "/".
		m.cleanedSrcBase = csp

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
		seelog.Warnf("No src files/folders found.")
		return errors.New("no src files or folders found")
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

func (m *MoveContext) stepPrepareScan() error {
	var err error

	m.dbFile = m.Infra.FileOnWorkPath("move.db")
	m.db = &dbsugar.DatabaseSugar{DataSourceName: m.dbFile}
	if err = m.db.Open(); err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", m.dbFile, err)
		return err
	}

	ddl := make(map[string]string)
	ddl[move_table_dst_file] = move_create_table_file
	ddl[move_table_dst_folder] = move_create_table_folder
	ddl[move_table_dst_shared_folder] = move_create_table_shared_folder
	ddl[move_table_src_file] = move_create_table_file
	ddl[move_table_src_folder] = move_create_table_folder
	ddl[move_table_src_shared_folder] = move_create_table_shared_folder
	ddl[move_table_operation_move] = move_create_table_operation_move
	ddl[move_table_operation_folder] = move_create_table_operation_folder

	for tn, d := range ddl {
		err = m.db.CreateTable(tn, d)
		if err != nil {
			seelog.Warnf("Failed to prepare table[%s]: error[%s]", tn, err)
			return err
		}
	}

	return nil
}

func (m *MoveContext) stepScanSrcFolders() error {
	pui := &progress.ProgressUI{}
	pui.Start(len(m.cleanedSrcPaths))
	defer pui.End()

	for _, s := range m.cleanedSrcPaths {
		err := m.scan(move_scan_src, s)
		if err != nil {
			seelog.Warnf("Failed to scan src files/folders : error[%s]", err)
			return err
		}
		pui.Incr()
	}

	m.scanReport(move_scan_src)

	return nil
}

func (m *MoveContext) stepScanDestFolder() error {
	seelog.Debugf("ScanDestFolder: cleanedDestPath[%s]", m.cleanedDestPath)
	if m.cleanedDestPath == "/" {
		for _, sp := range m.cleanedSrcPaths {
			n := filepath.Base(sp)
			p := filepath.ToSlash(filepath.Join("/", n))
			seelog.Debugf("ScanDestFolder: Path[%s] SrcPath[%s]", p, sp)

			err := m.scan(move_scan_dst, sp)
			if err != nil && strings.HasPrefix(err.Error(), "path/not_found") {
				seelog.Debugf("Skip scan destination folder")
				continue
			}
			if err != nil {
				seelog.Warnf("Failed to scan dest files/folders : error[%s]", err)
				return err
			}
		}

	} else {
		err := m.scan(move_scan_dst, m.cleanedDestPath)
		if err != nil && strings.HasPrefix(err.Error(), "path/not_found") {
			seelog.Debugf("Skip scan destination folder")
			return nil
		}
		if err != nil {
			seelog.Warnf("Failed to scan dest files/folders : error[%s]", err)
			return err
		}
	}
	m.scanReport(move_scan_dst)

	return nil
}

func (m *MoveContext) scanReport(scanTarget int) {
	var fCount int64
	var fSize uint64
	var folderCount int64

	err := m.db.QueryRow(
		"SELECT COUNT(file_id), SUM(size) FROM {{.TableName}}",
		m.tableNameFile(scanTarget),
	).Scan(
		&fCount,
		&fSize,
	)
	if err != nil {
		seelog.Debugf("Unable to retrieve file size/count : error[%s]", err)
		return
	}

	err = m.db.QueryRow(
		"SELECT COUNT(folder_id) FROM {{.TableName}}",
		m.tableNameFolder(scanTarget),
	).Scan(
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

func (m *MoveContext) scan(scanTarget int, path string) error {
	if path == "/" {

		return m.scanFolder(scanTarget, nil, nil)
	}

	client := files.New(m.dbxCfgFull)
	meta, err := client.GetMetadata(files.NewGetMetadataArg(path))
	if scanTarget == move_scan_dst &&
		err != nil &&
		strings.HasPrefix(err.Error(), "path/not_found") {
		seelog.Debugf("Skip scan destination folder")
		return nil
	}
	if err != nil {
		seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", path, err)
		return err
	}

	return m.scanDispatch(scanTarget, meta, nil)
}

func (m *MoveContext) scanDispatch(scanTarget int, meta files.IsMetadata, parentFolder *files.FolderMetadata) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		if scanTarget == move_scan_src {
			seelog.Debugf("Scan file: FileId[%s] PathDisplay[%s]", f.Id, f.PathDisplay)
		}
		return m.scanFile(scanTarget, f, parentFolder)

	case *files.FolderMetadata:
		seelog.Debugf("Scan folder: FolderId[%s] PathDisplay[%s]", f.Id, f.PathDisplay)
		return m.scanFolder(scanTarget, f, parentFolder)

	case *files.DeletedMetadata:
		seelog.Warnf("Deleted file cannot move [%s]", f.PathDisplay)
		return errors.New("deleted file cannot move")

	default:
		seelog.Debug("Unknown metadata type")
		return errors.New("unknown type of file/folder found")
	}
}

func (m *MoveContext) tableNameFile(scanTarget int) string {
	switch scanTarget {
	case move_scan_src:
		return move_table_src_file
	case move_scan_dst:
		return move_table_dst_file
	default:
		seelog.Errorf("Invalid scan target[%d]", scanTarget)
		return ""
	}
}

func (m *MoveContext) tableNameFolder(scanTarget int) string {
	switch scanTarget {
	case move_scan_src:
		return move_table_src_folder
	case move_scan_dst:
		return move_table_dst_folder
	default:
		seelog.Errorf("Invalid scan target[%d]", scanTarget)
		return ""
	}
}

func (m *MoveContext) tableNameSharedFolder(scanTarget int) string {
	switch scanTarget {
	case move_scan_src:
		return move_table_src_shared_folder
	case move_scan_dst:
		return move_table_dst_shared_folder
	default:
		seelog.Errorf("Invalid scan target[%d]", scanTarget)
		return ""
	}
}

func (m *MoveContext) scanFile(scanTarget int, meta *files.FileMetadata, parentFolder *files.FolderMetadata) error {
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
		`INSERT OR REPLACE INTO {{.TableName}} (
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
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.tableNameFile(scanTarget),
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

func (m *MoveContext) scanFolder(scanTarget int, meta, parentFolder *files.FolderMetadata) error {
	client := files.New(m.dbxCfgFull)

	path := ""
	if meta != nil {
		path = meta.PathDisplay
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
			`INSERT OR REPLACE INTO {{.TableName}} (
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
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			m.tableNameFolder(scanTarget),
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
	}

	listArg := files.NewListFolderArg(path)

	lf, err := client.ListFolder(listArg)
	if err != nil {
		seelog.Warnf("Unable to load folder[%s] : error[%s]", path, err)
		return err
	}
	more := true
	for more {
		for _, f := range lf.Entries {
			err = m.scanDispatch(scanTarget, f, meta)
			if err != nil {
				seelog.Warnf("Unable to prepare file/folder meta data[%s] : error[%s]", path, err)
				return err
			}
		}
		if lf.HasMore {
			lf, err = client.ListFolderContinue(files.NewListFolderContinueArg(lf.Cursor))
			if err != nil {
				seelog.Warnf("Unable to load folder (cont) [%s] : error[%s]", path, err)
				return err
			}
		}
		more = lf.HasMore
	}

	return nil
}

func (m *MoveContext) stepScanSrcSharedFolders() error {
	return m.scanSharedFolders(move_scan_src)
}

func (m *MoveContext) stepScanDestSharedFolders() error {
	return m.scanSharedFolders(move_scan_dst)
}

func (m *MoveContext) scanSharedFolders(scanTarget int) error {
	rows, err := m.db.Query(
		`SELECT DISTINCT sharing_shared_folder_id FROM {{.TableName}} WHERE sharing_shared_folder_id <> ""`,
		m.tableNameFolder(scanTarget),
	)
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
		m.scanSharedFolderInfo(scanTarget, sf)
	}

	return nil
}

func (m *MoveContext) scanSharedFolderInfo(scanTarget int, sharedFolderId string) error {
	client := sharing.New(m.dbxCfgFull)
	meta, err := client.GetFolderMetadata(sharing.NewGetMetadataArgs(sharedFolderId))
	if err != nil {
		seelog.Warnf("Unable to load shared_folder metadata [%s] : error[%s]", sharedFolderId, err)
		return err
	}

	name := meta.Name
	pathLower := meta.PathLower

	if m.Preflight && m.PreflightAnon {
		name = ""
		pathLower = ""
	}

	_, err = m.db.Exec(
		`INSERT OR REPLACE INTO {{.TableName}} (
		  shared_folder_id,
		  name,
		  is_inside_team_folder,
		  is_team_folder,
		  parent_shared_folder_id,
		  path_lower
		) VALUES (?, ?, ?, ?, ?, ?)`,
		m.tableNameSharedFolder(scanTarget),
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
	cnt := 0
	err := m.db.QueryRow(
		`SELECT COUNT(DISTINCT sharing_shared_folder_id)
		FROM  {{.TableName}}
		WHERE sharing_shared_folder_id <> ""`,
		move_table_src_folder,
	).Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count shared folders : error[%s]", err)
		return 0, err
	}

	return cnt, nil
}

func (m *MoveContext) stepValidateSrcPermissions() error {
	return m.validatePermissions(move_scan_src)
}

func (m *MoveContext) stepValidateDestPermissions() error {
	return m.validatePermissions(move_scan_dst)
}

func (m *MoveContext) validatePermissions(scanTarget int) error {
	cntReadOnly := 0
	cntTraverseOnly := 0
	cntNoAccess := 0

	// Validate files
	err := m.db.QueryRow(
		`SELECT IFNULL(SUM(sharing_read_only), 0) FROM {{.TableName}}`,
		m.tableNameFile(scanTarget),
	).Scan(
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
	err = m.db.QueryRow(
		`SELECT
	  		IFNULL(SUM(sharing_read_only), 0),
	  		IFNULL(SUM(sharing_traverse_only), 0),
	  		IFNULL(SUM(sharing_no_access), 0)
		FROM {{.TableName}}`,
		m.tableNameFolder(scanTarget),
	).Scan(
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

func (m *MoveContext) moveBatch(batch []*files.RelocationPath, pui *progress.ProgressUI) error {
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
				pui.Incr()
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
					pui.Incr()
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
						pui.Incr()
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
	cnt := 0
	err := m.db.QueryRow(
		`SELECT COUNT(path_display) FROM {{.TableName}}`,
		move_table_src_folder,
	).Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count folders : error[%s]", err)
		return err
	}

	if cnt == 0 {
		seelog.Debugf("No folder cleanup required")
		return nil
	}

	seelog.Infof("Remove %d folder(s) in source", cnt)

	rows, err := m.db.Query(
		`SELECT path_display FROM {{.TableName}} ORDER BY depth DESC`,
		move_table_src_folder,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve folder info : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	pui := &progress.ProgressUI{}
	pui.Start(cnt)
	defer pui.End()

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
		if err != nil && strings.HasPrefix(err.Error(), "path/not_found") {
			seelog.Debugf("Path[%s] already moved", pathDisplay)
			pui.Incr()
			continue
		}
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

		pui.Incr()
	}

	return nil
}

func (m *MoveContext) stepPrepareOperationPlan() error {
	baseFolderId := ""
	fileCnt, movable, err := m.operPlanSrcFolder(baseFolderId)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan: error[%s]", err)
		return err
	}

	if movable {
		seelog.Debugf("OprFolder: move FolderId[%s] FileCount[%d]", baseFolderId, fileCnt)
		m.operPlanMoveFolder(baseFolderId, fileCnt)
	}

	return nil
}

func (m *MoveContext) operPlanSrcFolder(folderId string) (fileCnt int, folderMovable bool, err error) {
	folderMovable = false
	seelog.Debugf("OperPlan: Start plan for FolderId[%s]", folderId)

	// Validate file count
	err = m.db.QueryRow(
		`SELECT COUNT(file_id) FROM {{.TableName}} WHERE parent_folder_id = ?`,
		move_table_src_file,
		folderId,
	).Scan(
		&fileCnt,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
		return 0, false, err
	}

	seelog.Debugf("OperPlanFolder: %d files found under folder[%s]", fileCnt, folderId)

	if MOVE_BATCH_API_LIMIT < fileCnt {
		seelog.Debugf("OperPlanFolder: too many files for move in one operation.")

		// Plan move files if folderMovable == false
		err = m.operPlanMoveFiles(folderId)
		return fileCnt, false, err
	}

	// Validate shared folder
	sharedFolderId := ""
	if folderId != "" {
		err = m.db.QueryRow(
			`SELECT sharing_shared_folder_id FROM {{.TableName}} WHERE folder_id = ?`,
			move_table_src_folder,
			folderId,
		).Scan(
			&sharedFolderId,
		)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			return fileCnt, false, err
		}
		seelog.Debugf("OperPlan: Ensure shared folder[%s]", sharedFolderId)
	}

	// Validate dest folders
	var srcPathDisplay string
	var destFolderExists bool
	destFolderExists = false
	destFolderPath := ""

	if folderId != "" {
		err = m.db.QueryRow(
			`SELECT path_display FROM {{.TableName}} WHERE folder_id = ?`,
			move_table_src_folder,
			folderId,
		).Scan(
			&srcPathDisplay,
		)
		destFolderPath, err = m.destPathFromSrcPath(srcPathDisplay)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			return fileCnt, false, err
		}

		seelog.Debugf("OperPlan: Compare path srcPathDisplay[%s] destPathDisplay[%s]", srcPathDisplay, destFolderPath)

		var pd int
		err = m.db.QueryRow(
			`SELECT COUNT(path_display) FROM {{.TableName}} WHERE path_display = ?`,
			move_table_dst_folder,
			destFolderPath,
		).Scan(
			&pd,
		)
		seelog.Debugf("OperPlan: Ensure dest path exist[%d] for folderId[%s]", pd, folderId)
		if pd > 0 {
			destFolderExists = true
		}
	}

	// Validate child folders
	row, err := m.db.Query(
		`SELECT folder_id, path_display FROM {{.TableName}} WHERE parent_folder_id = ?`,
		move_table_src_folder,
		folderId,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
		return fileCnt, false, err
	}

	childFolders := make(map[string]string)
	childMovable := make(map[string]bool)
	childFileCnt := make(map[string]int)

	for row.Next() {
		var childFolderId, pathDisplay string
		err = row.Scan(
			&childFolderId,
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			return fileCnt, false, err
		}
		seelog.Debugf("OperPlanFolder: Scan for oper plan: FolderId[%s] PathDisplay[%s]", childFolderId, pathDisplay)
		childFolders[childFolderId] = pathDisplay
	}

	allChildMovable := true
	for childFolderId, _ := range childFolders {
		childCnt, movable, err := m.operPlanSrcFolder(childFolderId)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			return fileCnt, false, err
		}

		if !movable {
			seelog.Debugf("OperPlanFolder: Cannot move child folder FolderId[%s]", childFolderId)
			allChildMovable = false
		}

		seelog.Debugf("OperPlan: ChildFolderId[%s] movable[%t] FileCount[%d]", childFolderId, movable, childCnt)

		childMovable[childFolderId] = movable
		childFileCnt[childFolderId] = childCnt
		fileCnt += childCnt
	}

	if allChildMovable &&
		fileCnt < MOVE_BATCH_API_LIMIT &&
		sharedFolderId == "" &&
		folderId != "" &&
		!destFolderExists &&
		!m.FileByFile {

		seelog.Debugf("OperPlan: Folder[%s] movable. AllChildMovable[%t] FileCount[%d] SharedFolderId[%s]", folderId, allChildMovable, fileCnt, sharedFolderId)
		return fileCnt, true, nil
	}

	if destFolderPath != "" && !destFolderExists {
		seelog.Debugf("OperPlan: Create Folder DestPath[%s]", destFolderPath)
		m.operPlanCreateFolder(destFolderPath)
	}

	for childFolderId, movable := range childMovable {
		seelog.Debugf("OperPlan: ChildFolderId[%s] Movable[%t]", childFolderId, movable)
		if movable {
			seelog.Debugf("OperPlan: Move child ChildFolderId[%s] Movable[%t]", childFolderId, movable)
			m.operPlanMoveFolder(childFolderId, childFileCnt[childFolderId])
		}
	}

	m.operPlanMoveFiles(folderId)

	return fileCnt, false, nil
}

func (m *MoveContext) operPlanCreateFolder(pathDisplay string) error {
	_, err := m.db.Exec(
		`INSERT OR REPLACE INTO {{.TableName}} (
		  path_display,
		  operation
		) VALUES (?, ?)`,
		move_table_operation_folder,
		pathDisplay,
		move_oper_create_folder,
	)
	if err != nil {
		seelog.Warnf("Unable to plan create folder : error[%s]", err)
		return err
	}
	return nil
}

func (m *MoveContext) operPlanMoveFolder(folderId string, fileCnt int) (err error) {
	if folderId == "" {
		return nil
	}

	var pathDisplay string
	err = m.db.QueryRow(
		`SELECT path_display FROM {{.TableName}} WHERE folder_id = ?`,
		move_table_src_folder,
		folderId,
	).Scan(
		&pathDisplay,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan FolderId[%s] : error[%s]", folderId, err)
		return
	}

	destPath, err := m.destPathFromSrcPath(pathDisplay)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan FolderId[%s] DisplayPath[%s]: error[%s]", pathDisplay, err)
		return
	}

	_, err = m.db.Exec(
		`INSERT OR REPLACE INTO {{.TableName}} (
		  operation_id,
		  operation,
		  move_from,
		  move_to,
		  file_count
		) VALUES (?, ?, ?, ?, ?)`,
		move_table_operation_move,
		folderId,
		move_oper_folder,
		pathDisplay,
		destPath,
		fileCnt,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan FolderId[%s] DisplayPath[%s]: error[%s]", pathDisplay, err)
		return
	}
	return
}

func (m *MoveContext) operPlanMoveFiles(parentFolderId string) error {
	row, err := m.db.Query(
		`SELECT file_id, path_display FROM {{.TableName}} WHERE parent_folder_id = ?`,
		move_table_src_file,
		parentFolderId,
	)
	if err != nil {
		seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
		return err
	}

	oper := make(map[string]string)
	for row.Next() {
		var fileId, pathDisplay string
		err = row.Scan(
			&fileId,
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			return err
		}

		seelog.Debugf("File: Scan for oper plan: FolderId[%s] PathDisplay[%s]", fileId, pathDisplay)

		oper[fileId] = pathDisplay
	}

	for fileId, pathDisplay := range oper {
		destPath, err := m.destPathFromSrcPath(pathDisplay)
		if err != nil {
			seelog.Warnf("Unable to calculate dest path[%s] : error[%s]", pathDisplay, err)
			continue
		}

		_, err = m.db.Exec(
			`INSERT OR REPLACE INTO {{.TableName}} (
			  operation_id,
			  operation,
			  move_from,
			  move_to,
			  file_count
			) VALUES (?, ?, ?, ?, 1)`,
			move_table_operation_move,
			fileId,
			move_oper_file,
			pathDisplay,
			destPath,
		)
		if err != nil {
			seelog.Warnf("Unable to prepare operation plan : error[%s]", err)
			continue
		}
	}

	return nil
}

func (m *MoveContext) destPathFromSrcPath(srcDisplayPath string) (destPath string, err error) {
	rel, err := filepath.Rel(m.cleanedSrcBase, srcDisplayPath)
	if err != nil {
		seelog.Warnf("Unable to calculate relative path[%s] : error[%s]", srcDisplayPath, err)
		return
	}

	destPath = filepath.ToSlash(filepath.Join(m.cleanedDestPath, rel))
	return
}

func (m *MoveContext) stepExecuteOperationPlan() error {
	var fCount int

	err := m.db.QueryRow(
		"SELECT COUNT(file_id) FROM {{.TableName}}",
		move_table_src_file,
	).Scan(
		&fCount,
	)
	if err != nil {
		seelog.Debugf("Unable to retrieve file size/count : error[%s]", err)
		return err
	}

	pui := &progress.ProgressUI{}
	if fCount < 1 {
		pui.Start(1)
	} else {
		pui.Start(fCount)
	}
	defer pui.End()

	err = m.executeOperFolders(pui)
	if err != nil {
		seelog.Warnf("Unable to execute operation : error[%s]", err)
		return err
	}
	err = m.executeOperFiles(pui)
	if err != nil {
		seelog.Warnf("Unable to execute operation : error[%s]", err)
		return err
	}

	return nil
}

func (m *MoveContext) stepExecuteCreateFolder() error {
	rows, err := m.db.Query(
		`SELECT
		  path_display
		FROM {{.TableName}}
		WHERE operation = ?
		`,
		move_table_operation_folder,
		move_oper_create_folder,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve operation plan : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	for rows.Next() {
		pathDisplay := ""
		err := rows.Scan(
			&pathDisplay,
		)
		if err != nil {
			seelog.Warnf("Unable to retrive operaiton plan : error[%s]", err)
			return err
		}

		arg := files.NewCreateFolderArg(pathDisplay)
		arg.Autorename = false
		_, err = client.CreateFolderV2(arg)
		if err != nil {
			seelog.Warnf("Failed to create folder[%s] : error[%s]", pathDisplay, err)
			return err
		}
	}
	return nil
}

func (m *MoveContext) executeOperFolders(pui *progress.ProgressUI) error {
	rows, err := m.db.Query(
		`SELECT
		  move_from,
		  move_to,
		  file_count
		FROM
		  {{.TableName}}
		WHERE
		  operation = ?`,
		move_table_operation_move,
		move_oper_folder,
	)
	if err != nil {
		seelog.Warnf("Unable to execute operation for folder : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	for rows.Next() {
		var moveFrom, moveTo string
		var cnt int
		err = rows.Scan(
			&moveFrom,
			&moveTo,
			&cnt,
		)
		if err != nil {
			seelog.Warnf("Unable to execute operation for folder : error[%s]", err)
			return err
		}

		arg := files.NewRelocationArg(moveFrom, moveTo)
		arg.Autorename = false
		arg.AllowOwnershipTransfer = true
		arg.AllowSharedFolder = true

		seelog.Debugf("ExecOper: move from[%s] to[%s]", moveFrom, moveTo)

		_, err := client.MoveV2(arg)
		if err != nil {
			seelog.Warnf("Failed to move folder from[%s] to [%s] : error[%s]", moveFrom, moveTo, err)
			return err
		}

		for i := 0; i < cnt; i++ {
			pui.Incr()
		}
	}

	return nil
}

func (m *MoveContext) executeOperFiles(pui *progress.ProgressUI) error {
	rows, err := m.db.Query(
		`SELECT
		  move_from,
		  move_to
		FROM
		  {{.TableName}}
		WHERE
		  operation = ?`,
		move_table_operation_move,
		move_oper_file,
	)
	if err != nil {
		seelog.Warnf("Unable to execute operation for folder : error[%s]", err)
		return err
	}

	batch := make([]*files.RelocationPath, 0)
	for rows.Next() {
		var moveFrom, moveTo string

		rows.Scan(
			&moveFrom,
			&moveTo,
		)

		seelog.Debugf("ExecOper: move from[%s] to[%s]", moveFrom, moveTo)

		batch = append(batch, files.NewRelocationPath(moveFrom, moveTo))
		if m.BatchSize <= len(batch) {
			err = m.moveBatch(batch, pui)
			if err != nil {
				seelog.Warnf("Failed to move batch file move : error[%s]", err)
				return err
			}
			batch = make([]*files.RelocationPath, 0)
		}
	}
	if 0 < len(batch) {
		err = m.moveBatch(batch, pui)
		if err != nil {
			seelog.Warnf("Failed to move batch file move : error[%s]", err)
			return err
		}
	}
	return nil
}
