package compare

import (
	"database/sql"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Traverse struct {
	db              *sql.DB
	dropboxConfig   dropbox.Config
	dbFile          string
	DropboxToken    string
	DropboxBasePath string
	InfraOpts       *infra.InfraOpts
	LocalBasePath   string
}

type DropboxFileInfo struct {
	DropboxFileId       string
	DropboxFileRevision string
	Path                string
	PathLower           string
	Size                int64
	ContentHash         string
}

type LocalFileInfo struct {
	Path        string
	PathLower   string
	Size        int64
	ContentHash string
}

func NewLocalFileInfo(basePath, path string) (*LocalFileInfo, error) {
	rel, err := filepath.Rel(basePath, path)
	if err != nil {
		seelog.Warnf("Unable to compute relative path : path[%s], error[%s]", path, err)
		return nil, err
	}
	inf, err := os.Lstat(path)
	if err != nil {
		seelog.Warnf("Unable to acquire lstat: path[%s] error[%s]", path, err)
		return nil, err
	}
	ch, err := ContentHash(path)
	if err != nil {
		seelog.Debugf("Unable to compute hash: path[%s] erorr[%s]", path, err)
		return nil, err
	}
	p := filepath.ToSlash(filepath.Clean(rel))
	lfi := LocalFileInfo{
		Path:        p,
		PathLower:   strings.ToLower(p),
		Size:        inf.Size(),
		ContentHash: ch,
	}
	return &lfi, nil
}

func (t *Traverse) normalizeKeyPath(path string) string {
	path = strings.ToLower(path)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Normalize Unicode: macOS normalize file names by NFD (with some exception).
	path = t.normalizeUnicodePath(path)
	return path
}

func (t *Traverse) normalizeUnicodePath(path string) string {
	return string(norm.NFC.Bytes([]byte(path)))
}

func (t *Traverse) Prepare() error {
	var err error
	t.dropboxConfig = dropbox.Config{
		Token: t.DropboxToken,
	}
	t.dbFile = t.InfraOpts.FileOnWorkPath("traverse.db")
	t.db, err = sql.Open("sqlite3", t.dbFile)
	if err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", t.dbFile, err)
		return err
	}

	q := `
	DROP TABLE IF EXISTS traverselocalfile
	`
	_, err = t.db.Exec(q)
	if err != nil {
		seelog.Errorf("Unable to drop table: %s", err)
		return err
	}

	q = `
	CREATE TABLE traverselocalfile (
	  path_lower   VARCHAR PRIMARY KEY,
	  path         VARCHAR,
	  size         INT8,
	  content_hash VARCHAR(32)
	)
	`
	_, err = t.db.Exec(q)
	if err != nil {
		seelog.Errorf("Unable to create table: %s", err)
		return nil
	}

	q = `
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

func (t *Traverse) ScanDropbox() error {
	return t.scanDropboxPath(t.DropboxBasePath)
}

func (t *Traverse) loadDropboxFileMetadata(f *files.FileMetadata) error {
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
		t.normalizeKeyPath(f.PathLower),
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

func (t *Traverse) scanDropboxPath(path string) error {
	var meta files.IsMetadata
	var err error

	seelog.Debugf("Scanning path: [%s]", path)

	if path == "" {
		return t.scanDropboxFolder(path)
	} else {

		client := files.New(t.dropboxConfig)
		marg := files.NewGetMetadataArg(path)
		meta, err = client.GetMetadata(marg)
		if err != nil {
			seelog.Warnf("Unable to get meta data for path[%s] error[%s]", path, err)
			return err
		}

		return t.scanDropboxMeta(meta)
	}
}

func (t *Traverse) scanDropboxMeta(meta files.IsMetadata) error {
	switch f := meta.(type) {
	case *files.FileMetadata:
		t.loadDropboxFileMetadata(f)

	case *files.FolderMetadata:
		t.scanDropboxFolder(f.PathLower)

	case *files.DeletedMetadata:
		seelog.Debugf("Ignore deleted file metadata: Path[%s]", f.PathLower)

	default:
		seelog.Debug("Ignore unknown metadata type")
	}
	return nil
}

func (t *Traverse) scanDropboxFolder(path string) error {
	seelog.Debugf("Scanning folder: [%s]", path)

	client := files.New(t.dropboxConfig)
	var entries []files.IsMetadata
	lfarg := files.NewListFolderArg(path)
	list, err := client.ListFolder(lfarg)
	if err != nil {
		seelog.Warnf("Unable to list_folder : path[%s] error[%s]", path, err)
		return err
	}

	entries = list.Entries
	for _, e := range entries {
		err := t.scanDropboxMeta(e)
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
			err := t.scanDropboxMeta(e)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Traverse) RetrieveDropbox(listener chan *DropboxFileInfo, wg *sync.WaitGroup) error {
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
	ORDER BY
	  path_lower
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

func (t *Traverse) SummaryDropbox() (count, size int64, err error) {
	seelog.Debug("Summarise")

	q := `
	SELECT
	  COUNT(path_lower),
	  SUM(size)
	FROM
	  traversedropboxfile
	`

	row := t.db.QueryRow(q)
	err = row.Scan(
		&count,
		&size,
	)
	if err != nil {
		seelog.Warnf("Unable to summrise : error[%s]", err)
		return 0, 0, err
	}

	return
}

func (t *Traverse) FetchDropbox(path string) (*DropboxFileInfo, error) {
	seelog.Debugf("Fetch path[%s]", path)

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
	WHERE
	  path_lower = ?
  	`

	dfi := DropboxFileInfo{}
	row := t.db.QueryRow(q, t.normalizeKeyPath(path))
	err := row.Scan(
		&dfi.PathLower,
		&dfi.Path,
		&dfi.DropboxFileId,
		&dfi.DropboxFileRevision,
		&dfi.Size,
		&dfi.ContentHash,
	)
	if err != nil {
		seelog.Debugf("Query failed for path[%s] error[%s]", path, err)
		return nil, err
	} else {
		return &dfi, nil
	}
}

func (t *Traverse) LoadLocal(path string) error {
	seelog.Debugf("Loading path: path[%s]", path)
	lfi, err := NewLocalFileInfo(t.LocalBasePath, path)
	if err != nil {
		seelog.Debugf("Unable to load path : path[%s] error[%s]", path, err)
		return err
	}
	return t.InsertLocal(lfi)
}

func (t *Traverse) InsertLocal(fileInfo *LocalFileInfo) error {
	q := `
	INSERT OR REPLACE INTO traverselocalfile (
	  path_lower,
	  path,
	  size,
	  content_hash
	) VALUES (?, ?, ?, ?)
	`
	_, err := t.db.Exec(
		q,
		t.normalizeKeyPath(fileInfo.PathLower),
		t.normalizeUnicodePath(fileInfo.Path),
		fileInfo.Size,
		fileInfo.ContentHash,
	)
	if err != nil {
		seelog.Warnf("Unable to insert/replace row: err[%s]", err)
		return err
	}

	return nil
}

func (t *Traverse) Close() error {
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

// Scan from base path.
func (t *Traverse) ScanLocal() error {
	return t.scanLocalPath(t.LocalBasePath)
}

func (t *Traverse) scanLocalPath(path string) error {
	seelog.Debugf("Scanning path: [%s]", path)
	info, err := os.Lstat(path)
	if err != nil {
		seelog.Warnf("Unable to acquire path information : path[%s] error[%s]", path, err)
		return err
	}
	if info.IsDir() {
		return t.scanLocalDir(path)
	} else {
		return t.LoadLocal(path)
	}
}

func (t *Traverse) scanLocalDir(path string) error {
	seelog.Debugf("Scanning directory: [%s]", path)
	list, err := ioutil.ReadDir(path)
	if err != nil {
		seelog.Warnf("Unable to list files of directory : path[%s] error[%s]", path, err)
		return err
	}
	for _, f := range list {
		p := filepath.Join(path, f.Name())
		seelog.Debugf("Directory entry[%s] isDir[%t] size[%d]", p, f.IsDir(), f.Size())
		if f.IsDir() {
			err := t.scanLocalDir(p)
			if err != nil {
				return err
			}
		} else {
			err := t.LoadLocal(p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Traverse) RetrieveLocal(listener chan *LocalFileInfo, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	q := `
	SELECT
	  path_lower,
	  path,
	  size,
	  content_hash
	FROM
	  traverselocalfile
	ORDER BY
	  path_lower
	`

	seelog.Debug("Retrieve paths from local traverse results")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		lfi := LocalFileInfo{}
		err = rows.Scan(
			&lfi.PathLower,
			&lfi.Path,
			&lfi.Size,
			&lfi.ContentHash,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}
		seelog.Debugf("Retrieved local traversed path: path[%s]", lfi.Path)
		listener <- &lfi
	}
	seelog.Debug("Finish retrieve local traversed paths")
	listener <- nil
	return nil
}

func (t *Traverse) FetchLocal(path string) (*LocalFileInfo, error) {
	seelog.Debugf("Fetch path[%s]", path)

	q := `
	SELECT
	  path_lower,
	  path,
	  size,
	  content_hash
	FROM
	  traverselocalfile
	WHERE
	  path_lower = ?
	`

	lfi := LocalFileInfo{}
	row := t.db.QueryRow(q, t.normalizeKeyPath(path))
	err := row.Scan(
		&lfi.PathLower,
		&lfi.Path,
		&lfi.Size,
		&lfi.ContentHash,
	)
	if err != nil {
		seelog.Debugf("Query failed for path[%s] error[%s]", path, err)
		return nil, err
	} else {
		return &lfi, nil
	}
}

func (t *Traverse) SummaryLocal() (count, size int64, err error) {
	seelog.Debug("Summarise")

	q := `
	SELECT
	  COUNT(path_lower),
	  SUM(size)
	FROM
	  traverselocalfile
	`

	row := t.db.QueryRow(q)
	err = row.Scan(
		&count,
		&size,
	)
	if err != nil {
		seelog.Warnf("Unable to summrise : error[%s]", err)
		return 0, 0, err
	}

	return
}

type CompareRowLocalToDropbox struct {
	PathLower   string
	Path        string
	Size        int64
	ContentHash string
}

type CompareRowDropboxToLocal struct {
	PathLower       string
	Path            string
	DropboxFileId   string
	DropboxRevision string
	Size            int64
	ContentHash     string
}

type CompareRowSizeAndHash struct {
	PathLower          string
	Path               string
	LocalSize          int64
	DropboxSize        int64
	LocalContentHash   string
	DropboxContentHash string
}

func (t *Traverse) CompareLocalToDropbox(listener chan *CompareRowLocalToDropbox, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	q := `
	SELECT
	  path_lower,
	  path,
	  size,
	  content_hash
	FROM
	  traverselocalfile
        WHERE
          path_lower NOT IN (SELECT path_lower FROM traversedropboxfile)
        ORDER BY
          path_lower
	`

	seelog.Debug("Compare Local to Dropbox")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		row := CompareRowLocalToDropbox{}
		err = rows.Scan(
			&row.PathLower,
			&row.Path,
			&row.Size,
			&row.ContentHash,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}
		seelog.Debugf("Retrieved diff row: path[%s]", row.Path)
		listener <- &row
	}
	seelog.Debug("Finish diff rows")
	listener <- nil
	return nil
}

func (t *Traverse) CompareDropboxToLocal(listener chan *CompareRowDropboxToLocal, wg *sync.WaitGroup) error {
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
	WHERE
	  path_lower NOT IN (SELECT path_lower FROM traverselocalfile)
	ORDER BY
	  path_lower
	`

	seelog.Debug("Compare Dropbox to Local")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		row := CompareRowDropboxToLocal{}
		err = rows.Scan(
			&row.PathLower,
			&row.Path,
			&row.DropboxFileId,
			&row.DropboxRevision,
			&row.Size,
			&row.ContentHash,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}
		seelog.Debugf("Retrieved diff row: path[%s]", row.Path)
		listener <- &row
	}
	seelog.Debug("Finish diff rows")
	listener <- nil
	return nil
}

func (t *Traverse) CompareSizeAndHash(listener chan *CompareRowSizeAndHash, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	q := `
	SELECT
	  l.path_lower,
	  l.path,
	  l.size,
	  d.size,
	  l.content_hash,
	  d.content_hash
	FROM
	  traverselocalfile   l,
	  traversedropboxfile d
	WHERE
              l.path_lower = d.path_lower
          AND (l.size <> d.size OR l.content_hash <> d.content_hash)
	`

	seelog.Debug("Compare size and/or hash")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		row := CompareRowSizeAndHash{}
		err = rows.Scan(
			&row.PathLower,
			&row.Path,
			&row.LocalSize,
			&row.DropboxSize,
			&row.LocalContentHash,
			&row.DropboxContentHash,
		)
		if err != nil {
			seelog.Warnf("Unable to retrieve row : error[%s]", err)
			return err
		}
		seelog.Debugf("Retrieved diff row: path[%s]", row.Path)
		listener <- &row
	}
	seelog.Debug("Finish diff rows")
	listener <- nil
	return nil
}
