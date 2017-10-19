package file

import (
	"github.com/watermint/toolbox/infra"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"database/sql"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"strings"
	"fmt"
	"path/filepath"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gosuri/uiprogress"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/watermint/toolbox/integration/sdk"
	"time"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
)

type MoveContext struct {
	db              *sql.DB
	dbFile          string
	dbAsync         *sql.DB
	dbAsyncFile     string
	cleanedSrcBase  string
	cleanedSrcPaths []string
	cleanedDestPath string

	Infra              *infra.InfraOpts
	SrcPath            string
	DestPath           string
	AutoRename         bool
	PromptConfirmation bool
}

const (
	MOVE_BATCH_SIZE = 100
)

func (m *MoveContext) Move(token string) error {
	totalSteps := 10

	seelog.Infof("[Step 1 of %d]: Prepare execution plan", totalSteps)
	err := m.preparePlan(token)
	if err != nil {
		seelog.Warnf("Suspend execution : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 2 of %d]: Confirm execution plan", totalSteps)
	cont := m.promptPlan()
	if !cont {
		seelog.Info("Execution cancelled.")
		return errors.New("execution cancelled")
	}

	seelog.Infof("[Step 3 of %d]: Preparing scan files and folders", totalSteps)
	err = m.prepareScan()
	if err != nil {
		seelog.Warnf("Preparation failed for scan : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 4 of %d]: Scan target files and folders", totalSteps)
	err = m.scanTarget(token)
	if err != nil {
		seelog.Warnf("Unable to scan target files and folders : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 5 of %d]: Scan sharing information", totalSteps)
	err = m.scanSharingInfo(token)
	if err != nil {
		seelog.Warnf("Unable to scan sharing information : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 6 of %d]: Validate permissions of files/folders", totalSteps)
	err = m.validatePermissions()
	if err != nil {
		return err
	}

	seelog.Infof("[Step 7 of %d]: Create destination folders", totalSteps)
	err = m.createFolders(token)
	if err != nil {
		seelog.Warnf("Unable to create folder(s) : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 8 of %d]: Move files", totalSteps)
	err = m.moveFiles(token)
	if err != nil {
		seelog.Warnf("Unable to move file(s) : error[%s]", err)
		return err
	}

	seelog.Infof("[Step 10 of %d]: Clean up folders of source folder", totalSteps)
	err = m.cleanupFolders(token)
	if err != nil {
		seelog.Warnf("Unable to clean up folder(s) : error[%s]", err)
		return err
	}

	return nil
}

func cleanPath(path string) (string, error) {
	c := filepath.Clean(path)
	c = filepath.ToSlash(c)
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

func (m *MoveContext) preparePlan(token string) error {
	var err error
	client := files.New(dropbox.Config{Token: token})

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
					m.cleanedSrcPaths = append(m.cleanedSrcPaths, filepath.Join(csp, n))
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

	return nil
}

func (m *MoveContext) promptPlan() bool {
	seelog.Info("Execution plan:")
	for _, sp := range m.cleanedSrcPaths {
		b := filepath.Base(sp)
		seelog.Infof("%s", sp)
		seelog.Infof("-> %s", filepath.Join(m.cleanedDestPath, b))
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
			return false
		}
	}

	return true
}

func (m *MoveContext) prepareScan() error {
	var err error

	m.dbFile = m.Infra.FileOnWorkPath("move.db")
	m.db, err = sql.Open("sqlite3", m.dbFile)

	if err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", m.dbFile, err)
		return err
	}

	m.dbAsyncFile = m.Infra.FileOnWorkPath("move-async.db")
	m.dbAsync, err = sql.Open("sqlite3", m.dbAsyncFile)

	if err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", m.dbFile, err)
		return err
	}


	q := `
	DROP TABLE IF EXISTS target_file
	`

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop table : error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE target_file (
	  file_id                         VARCHAR PRIMARY KEY,
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

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	q = `
	DROP TABLE IF EXISTS target_folder
	`

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop table : error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE target_folder (
	  folder_id                       VARCHAR PRIMARY KEY,
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

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	q = `
	DROP TABLE IF EXISTS target_shared_folder
	`

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop table : error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE target_shared_folder (
	  shared_folder_id                VARCHAR PRIMARY KEY,
	  name                            VARCHAR,
	  is_inside_team_folder           VARCHAR,
	  is_team_folder                  VARCHAR,
	  parent_shared_folder_id         VARCHAR,
	  path_lower                      VARCHAR
	)
	`

	_, err = m.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	q = `
	DROP TABLE IF EXISTS job_async
	`

	_, err = m.dbAsync.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop table : error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE job_async (
	  async_job_id    VARCHAR PRIMARY KEY,
	  job_timestamp   INT8
	)
	`

	_, err = m.dbAsync.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	return nil
}

func (m *MoveContext) scanTarget(token string) error {
	seelog.Flush()
	uip := uiprogress.New()
	uip.Start()
	defer uip.Stop()
	bar := uip.AddBar(len(m.cleanedSrcPaths))
	bar.PrependElapsed()
	bar.AppendCompleted()
	for _, s := range m.cleanedSrcPaths {
		err := m.scan(s, token)
		if err != nil {
			seelog.Warnf("Failed to scan target files/folders : error[%s]", err)
			return err
		}
		bar.Incr()
	}
	return nil
}

func (m *MoveContext) scan(src, token string) error {
	client := files.New(dropbox.Config{Token: token})
	meta, err := client.GetMetadata(files.NewGetMetadataArg(src))
	if err != nil {
		seelog.Warnf("Unable to load metadata for path[%s] : error[%s]", src, err)
		return err
	}

	return m.scanDispatch(meta, token)
}

func (m *MoveContext) scanDispatch(meta files.IsMetadata, token string) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		return m.scanFile(f)

	case *files.FolderMetadata:
		return m.scanFolder(f, token)

	case *files.DeletedMetadata:
		seelog.Warnf("Deleted file cannot move [%s]", f.PathDisplay)
		return errors.New("deleted file cannot move")

	default:
		seelog.Debug("Unknown metadata type")
		return errors.New("unknown type of file/folder found")
	}
}

func (m *MoveContext) scanFile(meta *files.FileMetadata) error {
	q := `
	INSERT OR REPLACE INTO target_file (
	  file_id,
	  name,
	  rev,
	  size,
	  path_lower,
	  path_display,
	  content_hash,
	  sharing_read_only,
	  sharing_parent_shared_folder_id
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	sharingReadOnly := false
	sharingParentSharedFolderId := ""

	if meta.SharingInfo != nil {
		sharingReadOnly = meta.SharingInfo.ReadOnly
		sharingParentSharedFolderId = meta.SharingInfo.ParentSharedFolderId
	}

	_, err := m.db.Exec(
		q,
		meta.Id,
		meta.Name,
		meta.Rev,
		meta.Size,
		meta.PathLower,
		meta.PathDisplay,
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

func (m *MoveContext) scanFolder(meta *files.FolderMetadata, token string) error {
	client := files.New(dropbox.Config{Token: token})

	q := `
	INSERT OR REPLACE INTO target_folder (
	  folder_id,
	  depth,
	  name,
	  path_lower,
	  path_display,
	  sharing_read_only,
	  sharing_parent_shared_folder_id,
	  sharing_shared_folder_id,
	  sharing_traverse_only,
	  sharing_no_access
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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

	_, err := m.db.Exec(
		q,
		meta.Id,
		len(strings.Split(meta.PathLower, "/")),
		meta.Name,
		meta.PathLower,
		meta.PathDisplay,
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
			err = m.scanDispatch(f, token)
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

func (m *MoveContext) scanSharingInfo(token string) error {
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
	seelog.Infof("%d shared folder(s) found in source path", len(sharedFolderIds))

	for _, sf := range sharedFolderIds {
		m.scanSharedFolderInfo(sf, token)
	}

	return nil
}

func (m *MoveContext) scanSharedFolderInfo(sharedFolderId, token string) error {
	client := sharing.New(dropbox.Config{Token: token})
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

	_, err = m.db.Exec(
		q,
		meta.SharedFolderId,
		meta.Name,
		meta.IsInsideTeamFolder,
		meta.IsTeamFolder,
		meta.ParentSharedFolderId,
		meta.PathLower,
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

func (m *MoveContext) validatePermissions() error {
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

func (m *MoveContext) createFolders(token string) error {
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

	client := files.New(dropbox.Config{Token: token})

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
		destPath := filepath.Join(m.cleanedDestPath, rel)
		seelog.Debugf("Create destionation folder[%s]", destPath)

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

func (m *MoveContext) moveFiles(token string) error {
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

	seelog.Infof("Moving %d files", cnt)

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

		destPath := filepath.Join(m.cleanedDestPath, rel)
		seelog.Debugf("Move file from [%s] to [%s]", pathDisplay, destPath)

		batch = append(batch, files.NewRelocationPath(pathDisplay, destPath))
		if len(batch) > MOVE_BATCH_SIZE {
			m.enqueueBatch(batch, token)
			batch = make([]*files.RelocationPath, 0)
		}
		bar.Incr()
	}

	if len(batch) > 0 {
		m.enqueueBatch(batch, token)
	}

	return nil
}

func (m *MoveContext) debugBatchMoveEntries(entries []*files.RelocationBatchResultData) {
	for _, e := range entries {
		switch f := e.Metadata.(type) {
		case *files.FileMetadata:
			seelog.Debugf("File moved into [%s] file_id[%s]", f.PathDisplay, f.Id)

		case *files.FolderMetadata:
			seelog.Warnf("This path should be file [%s] folder_id[%s]", f.PathDisplay, f.Id)

		case *files.DeletedMetadata:
			seelog.Warnf("This path should not be removed [%s]", f.PathDisplay)

		default:
			seelog.Warnf("Unknown metadata type found while moving file")
		}
	}
}

func (m *MoveContext) enqueueBatch(batch []*files.RelocationPath, token string) {
	arg := files.NewRelocationBatchArg(batch)
	arg.Autorename = false
	arg.AllowOwnershipTransfer = true
	arg.AllowSharedFolder = true

	res, err := sdk.ZMoveBatch(dropbox.Config{Token: token}, arg)
	if err != nil {
		seelog.Warnf("Unable to enqueue move task : error[%s]", err)
		return
	}

	seelog.Debugf("AsyncJobId[%s] Tag[%s]", res.AsyncJobId, res.Tag)
	client := files.New(dropbox.Config{Token: token})

	switch res.Tag {
	case "completed":
		seelog.Debugf("Move completed (%d relocations)", len(batch))

		if res.Complete.Entries != nil {
			m.debugBatchMoveEntries(res.Complete.Entries)
		}

	case "async_job_id":
		seelog.Debugf("Waiting task [%s]", res.AsyncJobId)
		for {
			time.Sleep(MOVE_BATCH_SIZE * time.Millisecond * 10)
			batchStatus, err := client.MoveBatchCheck(async.NewPollArg(res.AsyncJobId))
			if err != nil {
				seelog.Warnf("Unable to check job status[%s] : error[%s]", res.AsyncJobId, err)
			} else {
				switch res.Tag {
				case "in_progress":
					seelog.Debugf("async_job_id [%s] is still in progress", res.AsyncJobId)

				case "failed":
					seelog.Debugf("async_job_id [%s] failed with error [%s]", res.AsyncJobId, batchStatus.Failed.Tag)
					return

				case "completed":
					seelog.Debugf("async_job_id [%s] completed", res.AsyncJobId)
					if res.Complete != nil && res.Complete.Entries != nil {
						m.debugBatchMoveEntries(res.Complete.Entries)
					}
					return
				}
			}

		}
	}
}

func (m *MoveContext) cleanupFolders(token string) error {
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

	client := files.New(dropbox.Config{Token: token})

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
