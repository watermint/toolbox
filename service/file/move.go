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

	move_table_src_file          = "src_file"
	move_table_src_folder        = "src_folder"
	move_table_src_shared_folder = "src_shared_folder"
	move_table_dst_file          = "dst_file"
	move_table_dst_folder        = "dst_folder"
	move_table_dst_shared_folder = "dst_shared_folder"

	move_scan_src = 1
	move_scan_dst = 2
)

const (
	move_step_prepare_execution_plan = iota
	move_step_confirm_execution_plan
	move_step_prepare_scan
	move_step_scan_src_folders
	move_step_scan_src_shared_folders
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

	title[move_step_scan_src_folders] = "Scan source files and folders"
	steps[move_step_scan_src_folders] = m.stepScanSrcFolders

	title[move_step_scan_src_shared_folders] = "Scan source shared folders"
	steps[move_step_scan_src_shared_folders] = m.stepScanSrcSharedFolders

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
		err := m.scan(s)
		if err != nil {
			seelog.Warnf("Failed to scan src files/folders : error[%s]", err)
			return err
		}
		pui.Incr()
	}

	m.scanReport()

	return nil
}

func (m *MoveContext) scanReport() {
	var fCount int64
	var fSize uint64
	var folderCount int64

	err := m.db.QueryRow(
		"SELECT COUNT(file_id), SUM(size) FROM {{.TableName}}",
		move_table_src_file,
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
		move_table_src_folder,
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
	INSERT OR REPLACE INTO {{.TableName}} (
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
		move_table_src_file,
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
		move_table_src_folder,
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

func (m *MoveContext) stepScanSrcSharedFolders() error {
	rows, err := m.db.Query(
		`SELECT DISTINCT sharing_shared_folder_id FROM {{.TableName}} WHERE sharing_shared_folder_id <> ""`,
		move_table_src_folder,
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
		move_table_src_shared_folder,
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

func (m *MoveContext) stepValidatePermissions() error {
	cntReadOnly := 0
	cntTraverseOnly := 0
	cntNoAccess := 0

	// Validate files
	err := m.db.QueryRow(
		`SELECT IFNULL(SUM(sharing_read_only), 0) FROM {{.TableName}}`,
		move_table_src_file,
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
		move_table_src_folder,
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

func (m *MoveContext) stepCreateDestFolders() error {
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
		seelog.Debugf("No create folder required")
		return nil
	}

	seelog.Infof("Create %d folder(s) in destination", cnt)

	rows, err := m.db.Query(
		`SELECT path_display FROM {{.TableName}} ORDER BY depth`,
		move_table_src_folder,
	)
	if err != nil {
		seelog.Warnf("Unable to retrieve shared_folder info : error[%s]", err)
		return err
	}

	client := files.New(m.dbxCfgFull)

	pui := progress.ProgressUI{}
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

		pui.Incr()
	}

	return nil
}

func (m *MoveContext) stepMoveFiles() error {
	cnt := 0
	err := m.db.QueryRow(
		`SELECT COUNT(path_display) FROM {{.TableName}}`,
		move_table_src_file,
	).Scan(
		&cnt,
	)
	if err != nil {
		seelog.Warnf("Unable to count file(s) : error[%s]", err)
		return err
	}

	seelog.Infof("Move %d files into destination", cnt)

	pui := &progress.ProgressUI{}
	pui.Start(cnt)
	defer pui.End()

	rows, err := m.db.Query(
		`SELECT path_display FROM {{.TableName}}`,
		move_table_src_file,
	)
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
			m.moveBatch(batch, pui)
			batch = make([]*files.RelocationPath, 0)
		}
	}

	if 0 < len(batch) {
		m.moveBatch(batch, pui)
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
