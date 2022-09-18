package es_template

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"path/filepath"
)

type HandlerApplyTagAdd func(path es_filesystem.Path, tags []string) error
type HandlerApplyPutFile func(path es_filesystem.Path, f io.ReadSeeker) error
type OnCreateFolder func(path es_filesystem.Path)
type ApplyOpts struct {
	HandlerTagAdd  HandlerApplyTagAdd
	HandlerPutFile HandlerApplyPutFile
	OnCreateFolder OnCreateFolder
}

type Apply interface {
	// Apply template to the path
	Apply(path es_filesystem.Path, template Root) (err error)
}

func NewApply(fs es_filesystem.FileSystem, opts ApplyOpts) Apply {
	return &applyImpl{
		fs: fs,
		ao: opts,
	}
}

type applyImpl struct {
	fs es_filesystem.FileSystem
	ao ApplyOpts
}

func (z applyImpl) applyName(entryName string) (name string, err error) {
	return es_filepath.FormatPathWithPredefinedVariables(entryName)
}

func (z applyImpl) putFileAndTags(path es_filesystem.Path, entry File, f io.ReadSeeker) error {
	if z.ao.HandlerPutFile == nil {
		return nil
	}
	if err := z.ao.HandlerPutFile(path, f); err != nil {
		return err
	}
	if z.ao.HandlerTagAdd != nil && 0 < len(entry.Tags) {
		return z.ao.HandlerTagAdd(path, entry.Tags)
	}
	return nil
}

func (z applyImpl) putFileFromSource(path es_filesystem.Path, entry File) (err error) {
	l := esl.Default().With(esl.String("path", path.Path()), esl.String("source", entry.Source))
	d, err := os.MkdirTemp("", "putFile")
	if err != nil {
		l.Debug("Unable to create temporary folder", esl.Error(err))
		return err
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()

	filePath := filepath.Join(d, entry.Name)
	err = es_download.Download(l, entry.Source, filePath)
	if err != nil {
		l.Debug("Unable to download the file", esl.Error(err))
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		l.Debug("Unable to open the downloaded file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	return z.putFileAndTags(path, entry, f)
}

func (z applyImpl) putFile(path es_filesystem.Path, entry File) (err error) {
	l := esl.Default().With(esl.String("path", path.Path()))
	// skip put file
	if z.ao.HandlerPutFile == nil {
		l.Debug("No put file handler, skip")
		return nil
	}

	if entry.Source != "" {
		l.Debug("Put file from source")
		return z.putFileFromSource(path, entry)
	}
	if entry.Content != "" {
		l.Debug("Put file content directly")
		return z.putFileAndTags(path, entry, bytes.NewReader([]byte(entry.Content)))
	}

	l.Debug("No content found. skip")
	return nil
}

func (z applyImpl) applyFolder(path es_filesystem.Path, entry Folder) (err error) {
	l := esl.Default().With(esl.String("path", path.Path()))
	if z.ao.OnCreateFolder != nil {
		z.ao.OnCreateFolder(path)
	}
	_, fsErr := z.fs.CreateFolder(path)
	if fsErr != nil {
		if !fsErr.IsConflict() {
			l.Debug("Unable to create folder", esl.Error(err))
			return fsErr
		} else {
			l.Debug("The folder already exists")
			// fall through
		}
	}
	if z.ao.HandlerTagAdd != nil && 0 < len(entry.Tags) {
		l.Debug("Adding tags", esl.Strings("tags", entry.Tags))
		if err := z.ao.HandlerTagAdd(path, entry.Tags); err != nil {
			return err
		}
	}

	l.Debug("Apply descendants")
	return z.applyFilesAndFolders(path, entry.Folders, entry.Files)
}

func (z applyImpl) applyFilesAndFolders(path es_filesystem.Path, folders []Folder, files []File) error {
	for _, e := range folders {
		folderName, err := z.applyName(e.Name)
		if err != nil {
			return err
		}
		folderPath := path.Descendant(folderName)
		if afErr := z.applyFolder(folderPath, e); afErr != nil {
			return afErr
		}
	}
	for _, e := range files {
		fileName, err := z.applyName(e.Name)
		if err != nil {
			return err
		}
		filePath := path.Descendant(fileName)
		if afErr := z.putFile(filePath, e); afErr != nil {
			return afErr
		}
	}
	return nil
}

func (z applyImpl) Apply(path es_filesystem.Path, template Root) (err error) {
	return z.applyFilesAndFolders(path, template.Folders, template.Files)
}
