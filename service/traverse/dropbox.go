package traverse

import (
	"database/sql"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra"
	"os"
	"sync"
)

type DropboxFileInfo struct {
	DropboxFileId       string
	DropboxFileRevision string
	Path                string
	PathLower           string
	Size                int64
	ContentHash         string
}

type TraverseDropboxFile struct {
	db           *sql.DB
	config       dropbox.Config
	dbFile       string
	DropboxToken string
	BasePath     string
	InfraOpts    *infra.InfraOpts
}

func (t *TraverseDropboxFile) Prepare() error {
	t.config = dropbox.Config{
		Token: t.DropboxToken,
	}

	var err error
	t.dbFile = t.InfraOpts.FileOnWorkPath("traversedropbox.db")
	t.db, err = sql.Open("sqlite3", t.dbFile)

	q := `
	DROP TABLE IF EXISTS traversedropboxfile
	`
	_, err = t.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop existing table: error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE traversedropboxfile (
	  path_lower       VARCHAR PRIMARY KEY,
	  path             VARCHAR,
	  dropbox_file_id  VARCHAR,
	  dropbox_revision VARCHAR,
	  size             INT8,
	  content_hash     VARCHAR(32)
	)
	`
	_, err = t.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	return nil
}

func (t *TraverseDropboxFile) Close() error {
	if t.db == nil {
		return nil
	}
	err := t.db.Close()
	if err != nil {
		seelog.Errorf("Unable to close database: error[%s]", err)
		return err
	}
	err = os.Remove(t.dbFile)
	if err != nil {
		seelog.Warnf("Unable to remove database file : path[%s] error[%s]", t.dbFile, err)
		return err
	}
	return nil
}

func (t *TraverseDropboxFile) Scan() error {
	return t.scanPath(t.BasePath)
}

func (t *TraverseDropboxFile) loadFileMetadata(f *files.FileMetadata) error {
	q := `
	INSERT OR REPLACE INTO traversedropboxfile (
	  path_lower,
	  path,
	  dropbox_file_id,
	  dropbox_revision,
	  size,
	  content_hash
	) VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := t.db.Exec(
		q,
		f.PathLower,
		f.PathDisplay,
		f.Id,
		f.Rev,
		f.Size,
		f.ContentHash,
	)
	if err != nil {
		seelog.Warnf("Unable to insert/replace row : error[%s]", err)
		return err
	}

	return nil
}

func (t *TraverseDropboxFile) scanPath(path string) error {
	var meta files.IsMetadata
	var err error

	seelog.Debugf("Scanning path: [%s]", path)

	client := files.New(t.config)
	marg := files.NewGetMetadataArg(path)
	meta, err = client.GetMetadata(marg)
	if err != nil {
		seelog.Warnf("Unable to get meta data for path[%s] error[%s]", path, err)
		return err
	}

	return t.scanMeta(meta)
}

func (t *TraverseDropboxFile) scanMeta(meta files.IsMetadata) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		t.loadFileMetadata(f)

	case *files.FolderMetadata:
		t.scanFolder(f.PathLower)

	case *files.DeletedMetadata:
		seelog.Debugf("Ignore deleted file metadata: Path[%s]", f.PathLower)

	default:
		seelog.Debug("Ignore unknown metadata type")
	}
	return nil
}

func (t *TraverseDropboxFile) scanFolder(path string) error {
	seelog.Debugf("Scanning folder: [%s]", path)

	client := files.New(t.config)
	var entries []files.IsMetadata
	lfarg := files.NewListFolderArg(path)
	list, err := client.ListFolder(lfarg)
	if err != nil {
		seelog.Warnf("Unable to list_folder : path[%s] error[%s]", path, err)
		return err
	}

	entries = list.Entries
	for _, e := range entries {
		err := t.scanMeta(e)
		if err != nil {
			return err
		}
	}

	for list.HasMore {
		cont := files.NewListFolderContinueArg(list.Cursor)
		list, err = client.ListFolderContinue(cont)
		if err != nil {
			seelog.Warnf("Unable to list_folder_continue : path[%s] error[%s]", path, err)
			return err
		}
		entries = list.Entries
		for _, e := range entries {
			err := t.scanMeta(e)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TraverseDropboxFile) Retrieve(listener chan *DropboxFileInfo, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	q := `
	SELECT
	  path_lower,
	  path,
	  dropbox_file_id,
	  dropbox_revision,
	  size,
	  content_hash
	FROM
	  traversedropboxfile
  	`

	seelog.Debug("Retrieve paths from dropbox traverse results")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		dfi := DropboxFileInfo{}
		err = rows.Scan(
			&dfi.PathLower,
			&dfi.Path,
			&dfi.DropboxFileId,
			&dfi.DropboxFileRevision,
			&dfi.Size,
			&dfi.ContentHash,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}
		seelog.Debugf("Retrieved local traversed path: path[%s]", dfi.Path)
		listener <- &dfi
	}
	seelog.Debug("Finish retrieve dropbox traversed paths")
	listener <- nil
	return nil
}
