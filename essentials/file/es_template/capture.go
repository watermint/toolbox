package es_template

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

type Capture interface {
	// Capture create template from the path
	Capture(path es_filesystem.Path) (template Root, err error)
}

type CaptureOpts struct {
	HandlerSource HandlerCaptureSource
	HandlerTags   HandlerCaptureTags
}

type HandlerCaptureSource func(path es_filesystem.Path) (link string, err error)
type HandlerCaptureTags func(path es_filesystem.Path) (tags []string, err error)

func NewCapture(fs es_filesystem.FileSystem, opts CaptureOpts) Capture {
	return &capImpl{
		fs:   fs,
		opts: opts,
	}
}

type capImpl struct {
	fs   es_filesystem.FileSystem
	opts CaptureOpts
}

func (z capImpl) captureTags(path es_filesystem.Path) (tags []string, err error) {
	if z.opts.HandlerTags != nil {
		return z.opts.HandlerTags(path)
	}
	return []string{}, nil
}

func (z capImpl) annotateFile(entry es_filesystem.Entry, file File) (annotated File, err error) {
	file.Tags, err = z.captureTags(entry.Path())
	return file, err
}

func (z capImpl) annotateFolder(entry es_filesystem.Entry, folder Folder) (annotated Folder, err error) {
	folder.Tags, err = z.captureTags(entry.Path())
	return folder, err
}

func (z capImpl) captureFile(entry es_filesystem.Entry) (file File, err error) {
	if z.opts.HandlerSource != nil {
		var link string
		link, err = z.opts.HandlerSource(entry.Path())
		if err != nil {
			return File{}, err
		}
		return z.annotateFile(entry, File{
			Name:   entry.Name(),
			Source: link,
		})
	}

	return z.annotateFile(entry, File{
		Name: entry.Name(),
	})
}

func (z capImpl) captureFolderEntryAnnotated(entry es_filesystem.Entry) (folder Folder, err error) {
	folder.Folders, folder.Files, err = z.captureFolder(entry)
	if err != nil {
		return Folder{}, err
	}
	folder.Name = entry.Name()
	return z.annotateFolder(entry, folder)
}

func (z capImpl) captureFolder(entry es_filesystem.Entry) (folders []Folder, files []File, err error) {
	folders = make([]Folder, 0)
	files = make([]File, 0)
	entries, err := z.fs.List(entry.Path())
	for _, e := range entries {
		if e.IsFile() {
			fe, err := z.captureFile(e)
			if err != nil {
				return []Folder{}, []File{}, err
			}
			files = append(files, fe)
		}
		if e.IsFolder() {
			fe, err := z.captureFolderEntryAnnotated(e)
			if err != nil {
				return []Folder{}, []File{}, err
			}
			folders = append(folders, fe)
		}
	}
	return folders, files, nil
}

func (z capImpl) capturePath(path es_filesystem.Path) (folders []Folder, files []File, err error) {
	entry, err := z.fs.Info(path)
	if err != nil {
		return []Folder{}, []File{}, err
	}
	if entry.IsFile() {
		fileEntry, err := z.captureFile(entry)
		if err != nil {
			return []Folder{}, []File{}, err
		}
		return []Folder{}, []File{fileEntry}, nil
	}
	if entry.IsFolder() {
		return z.captureFolder(entry)
	}

	return []Folder{}, []File{}, nil
}

func (z capImpl) Capture(path es_filesystem.Path) (template Root, err error) {
	template.Folders, template.Files, err = z.capturePath(path)
	if err != nil {
		return Root{}, err
	}
	return template, nil
}
