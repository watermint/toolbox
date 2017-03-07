package traverse

import (
	"database/sql"
	"github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/service/compare"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

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
	ch, err := compare.ContentHash(path)
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

type TraverseLocalFile struct {
	db        *sql.DB
	dbFile    string
	BasePath  string
	InfraOpts *infra.InfraOpts
}

func (t *TraverseLocalFile) Prepare() error {
	var err error
	t.dbFile = t.InfraOpts.FileOnWorkPath("traverselocal.db")
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

	return nil
}

func (t *TraverseLocalFile) Load(path string) error {
	seelog.Debugf("Loading path: path[%s]", path)
	lfi, err := NewLocalFileInfo(t.BasePath, path)
	if err != nil {
		seelog.Debugf("Unable to load path : path[%s] error[%s]", path, err)
		return err
	}
	return t.Insert(lfi)
}

func (t *TraverseLocalFile) Insert(fileInfo *LocalFileInfo) error {
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
		fileInfo.PathLower,
		fileInfo.Path,
		fileInfo.Size,
		fileInfo.ContentHash,
	)
	if err != nil {
		seelog.Warnf("Unable to insert/replace row: err[%s]", err)
		return err
	}

	return nil
}

func (t *TraverseLocalFile) Close() error {
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
func (t *TraverseLocalFile) Scan() error {
	return t.scanPath(t.BasePath)
}

func (t *TraverseLocalFile) scanPath(path string) error {
	seelog.Debugf("Scanning path: [%s]", path)
	info, err := os.Lstat(path)
	if err != nil {
		seelog.Warnf("Unable to acquire path information : path[%s] error[%s]", path, err)
		return err
	}
	if info.IsDir() {
		return t.scanDir(path)
	} else {
		return t.Load(path)
	}
}

func (t *TraverseLocalFile) scanDir(path string) error {
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
			err := t.scanDir(p)
			if err != nil {
				return err
			}
		} else {
			err := t.Load(p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *TraverseLocalFile) Retrieve(listener chan *LocalFileInfo, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	q := `
	SELECT path, path_lower, size, content_hash FROM traverselocalfile
	`

	seelog.Debug("Retrieve paths from local traverse results")
	rows, err := t.db.Query(q)
	if err != nil {
		seelog.Warnf("Unable to retrieve files which stored in internal database : error[%s]", err)
		return err
	}

	for rows.Next() {
		lfi := LocalFileInfo{}
		err = rows.Scan(&lfi.Path, &lfi.PathLower, &lfi.Size, &lfi.ContentHash)
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
