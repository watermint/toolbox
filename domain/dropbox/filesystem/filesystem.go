package filesystem

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

const (
	FileSystemTypeDropbox  = "dropbox"
	ApiComplexityThreshold = 10_000
)

func NewFileSystem(ctx dbx_context.Context) es_filesystem.FileSystem {
	return &dbxFs{
		ctx: ctx,
	}
}

type dbxFs struct {
	ctx dbx_context.Context
}

func (z dbxFs) OperationalComplexity(entries []es_filesystem.Entry) (complexity int64) {
	if x := len(entries); x > ApiComplexityThreshold {
		return int64(x)
	} else {
		return 1
	}
}

func (z dbxFs) List(path es_filesystem.Path) (entries []es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("path", path.AsData()))
	dbxPath, err := ToDropboxPath(path)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	dbxEntries, dbxErr := sv_file.NewFiles(z.ctx).List(dbxPath)
	if dbxErr != nil {
		l.Debug("Unable to list entries", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	entries = make([]es_filesystem.Entry, 0)
	for _, dbxEntry := range dbxEntries {
		entries = append(entries, NewEntry(dbxEntry))
	}
	return entries, nil
}

func (z dbxFs) Info(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("path", path.AsData()))
	dbxPath, err := ToDropboxPath(path)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	dbxEntry, dbxErr := sv_file.NewFiles(z.ctx).Resolve(dbxPath)
	if dbxErr != nil {
		l.Debug("unable to retrieve entry", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	return NewEntry(dbxEntry), nil
}

func (z dbxFs) Delete(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("path", path.AsData()))

	var opDelete func(dbxPath mo_path.DropboxPath) es_filesystem.FileSystemError
	opDelete = func(dbxPath mo_path.DropboxPath) es_filesystem.FileSystemError {
		_, dbxErr := sv_file.NewFiles(z.ctx).Remove(dbxPath)
		remoteErr := dbx_error.NewErrors(dbxErr)
		switch {
		case remoteErr == nil:
			l.Debug("Successfully deleted")
			return nil

		case remoteErr.IsTooManyFiles():
			l.Debug("Too many files, recurse into descendants")
			remoteEntries, err := sv_file.NewFiles(z.ctx).List(dbxPath)
			if err != nil {
				l.Debug("Unable to retrieve entries", esl.Error(err))
				return NewError(err)
			}

			var lastErr es_filesystem.FileSystemError
			for _, remoteEntry := range remoteEntries {
				err = opDelete(mo_path.NewDropboxPath(remoteEntry.PathDisplay()))
				if err != nil {
					l.Debug("Unable to remove", esl.Any("descendantEntry", remoteEntry.Concrete()), esl.Error(err))
					lastErr = NewError(err)
				}
			}

			if lastErr != nil {
				return lastErr
			}
			l.Debug("Descendant removed without error, try remove the folder")
			_, err = sv_file.NewFiles(z.ctx).Remove(dbxPath)
			if err != nil {
				l.Debug("Unable to remove", esl.Error(err))
				return NewError(err)
			}
			return nil

		default:
			l.Debug("Unable to remove", esl.Error(err))
			return err
		}
	}

	dbxPath0, err := ToDropboxPath(path)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return err
	}

	return opDelete(dbxPath0)
}

func (z dbxFs) CreateFolder(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("path", path.AsData()))
	dbxPath, err := ToDropboxPath(path)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return err
	}
	_, dbxErr := sv_file_folder.New(z.ctx).Create(dbxPath)
	if dbxErr != nil {
		return NewError(dbxErr)
	}
	return nil
}

func (z dbxFs) Entry(data es_filesystem.EntryData) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log()
	if data.FileSystemType != FileSystemTypeDropbox {
		l.Debug("Invalid file system type", esl.String("actualType", data.FileSystemType))
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}
	if data.Attributes == nil {
		l.Debug("Attributes was nil")
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}
	if raw, ok := data.Attributes["raw"]; ok {
		if rawMap, ok := raw.(map[string]interface{}); ok {
			rawJson, jsErr := json.Marshal(rawMap)
			if jsErr != nil {
				l.Debug("Unable to serialize", esl.Error(jsErr))
				return nil, NewError(ErrorInvalidEntryDataFormat)
			}

			meta := &mo_file.Metadata{}
			if err := api_parser.ParseModelRaw(meta, rawJson); err != nil {
				l.Debug("Unable to parse as a model", esl.Error(err))
				return nil, NewError(ErrorInvalidEntryDataFormat)
			}
			return NewEntry(meta), nil
		}
		if rawJson, ok := raw.(json.RawMessage); ok {
			meta := &mo_file.Metadata{}
			if err := api_parser.ParseModelRaw(meta, rawJson); err != nil {
				l.Debug("Unable to parse as a model", esl.Error(err))
				return nil, NewError(ErrorInvalidEntryDataFormat)
			}
			return NewEntry(meta), nil
		}
		l.Debug("Attributes found, but invalid type")
	}
	return nil, NewError(ErrorInvalidEntryDataFormat)
}

func (z dbxFs) Path(data es_filesystem.PathData) (path es_filesystem.Path, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeDropbox {
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}
	return NewPath(data.EntryShard.ShardId, mo_path.NewDropboxPath(data.EntryPath)), nil
}

func (z dbxFs) Shard(data es_filesystem.ShardData) (namespace es_filesystem.Shard, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeDropbox {
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}
	return data, nil
}

func (z dbxFs) FileSystemType() string {
	return FileSystemTypeDropbox
}
