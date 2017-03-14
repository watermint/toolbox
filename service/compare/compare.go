package compare

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"io"
	"os"
	"sync"
)

const (
	BLOCK_SIZE     = 4194304
	HASH_FOR_EMPTY = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

// Calculate File content hash
func ContentHash(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] error[%s]", path, err)
		return "", err
	}
	if info.Size() < 1 {
		return HASH_FOR_EMPTY, nil
	}

	f, err := os.Open(path)
	if err != nil {
		seelog.Warnf("Unable to open file [%s] error[%s]", path, err)
		return "", err
	}
	defer f.Close()

	var loadedBytes, totalBytes int64

	loadedBytes = 0
	totalBytes = info.Size()
	hashPerBlock := make([][32]byte, 0)

	for (totalBytes - loadedBytes) > 0 {
		r := io.LimitReader(f, BLOCK_SIZE)
		block := make([]byte, BLOCK_SIZE)
		readBytes, err := r.Read(block)
		if err == io.EOF {
			break
		}
		if err != nil {
			seelog.Warnf("Unable to load file [%s] error[%s]", path, err)
			return "", err
		}

		h := sha256.Sum256(block[:readBytes])
		hashPerBlock = append(hashPerBlock, h)
	}

	concatenated := make([]byte, 0)
	for _, h := range hashPerBlock {
		concatenated = append(concatenated, h[:]...)
	}
	h := sha256.Sum256(concatenated)
	return hex.EncodeToString(h[:]), nil
}

func Compare(infraOpts *infra.InfraOpts, token, localBasePath, dropboxBasePath string) error {
	var err error
	trav := Traverse{
		DropboxToken:    token,
		DropboxBasePath: dropboxBasePath,
		LocalBasePath:   localBasePath,
		InfraOpts:       infraOpts,
	}
	err = trav.Prepare()
	if err != nil {
		return err
	}
	defer trav.Close()

	seelog.Info("Start scanning local files")
	trav.ScanLocal()

	seelog.Info("Start scanning dropbox files")
	trav.ScanDropbox()

	wg := &sync.WaitGroup{}
	dtl := make(chan *CompareRowDropboxToLocal)
	ltd := make(chan *CompareRowLocalToDropbox)
	sah := make(chan *CompareRowSizeAndHash)

	go trav.CompareDropboxToLocal(dtl, wg)

	fmt.Println("*** Record: files not found in Local")
	for {
		row := <-dtl
		if row == nil {
			break
		}

		fmt.Printf("Path[%s] (lower:%s) Size[%d] Hash[%s] DropboxFileId[%s] DropboxRev[%s]\n",
			row.Path,
			row.PathLower,
			row.Size,
			row.ContentHash,
			row.DropboxFileId,
			row.DropboxRevision,
		)
	}

	go trav.CompareLocalToDropbox(ltd, wg)

	fmt.Println("*** Record: files not found in Dropbox")
	for {
		row := <-ltd
		if row == nil {
			break
		}
		fmt.Printf("Path[%s] (lower:%s) Size[%d] Hash[%s]\n",
			row.Path,
			row.PathLower,
			row.Size,
			row.ContentHash,
		)
	}

	go trav.CompareSizeAndHash(sah, wg)

	fmt.Println("*** Record: files size and/or hash not mached")
	for {
		row := <-sah
		if row == nil {
			break
		}

		fmt.Printf("Path[%s] (lower:%s) Size(Local:%d, Dropbox:%d), Hash(Local:%s, Dropbox:%s)\n",
			row.Path,
			row.PathLower,
			row.LocalSize,
			row.DropboxSize,
			row.LocalContentHash,
			row.DropboxContentHash,
		)
	}

	wg.Wait()

	return nil
}
