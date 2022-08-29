package sv_file_lock

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorUnexpectedResponseFormat = errors.New("unexpected response format")
)

type LockResult struct {
	Entry mo_file.Entry
	Error error
}

type Lock interface {
	Lock(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
	Unlock(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
	List(path mo_path.DropboxPath, onLockEntry func(entry *mo_file.LockInfo)) error

	// Batch lock paths. returns empty map if any error happened, thus entries never nil.
	LockBatch(path []mo_path.DropboxPath) (entries map[string]LockResult, err error)
	// Batch unlock paths. returns empty map if any error happened, thus entries never nil.
	UnlockBatch(path []mo_path.DropboxPath) (entries map[string]LockResult, err error)
}

func New(ctx dbx_client.Client) Lock {
	return &lockImpl{
		ctx: ctx,
	}
}

type lockImpl struct {
	ctx dbx_client.Client
}

func (z *lockImpl) List(path mo_path.DropboxPath, onLockEntry func(entry *mo_file.LockInfo)) error {
	svf := sv_file.NewFiles(z.ctx)
	return svf.ListEach(path, func(entry mo_file.Entry) {
		if fli := entry.LockInfo(); fli != nil {
			onLockEntry(fli)
		}
	}, sv_file.Recursive(true))
}

func (z *lockImpl) parseEntry(je es_json.Json) (entry mo_file.Entry, err error) {
	tag, found := je.FindString("\\.tag")
	if !found {
		return nil, ErrorUnexpectedResponseFormat
	}
	if tag == "success" {
		entry = &mo_file.Metadata{}
		err = je.FindModel("metadata", entry)
		return entry, err
	} else {
		reason, found := je.FindString("failure.\\.tag")
		if !found {
			return nil, ErrorUnexpectedResponseFormat
		}
		detail, found := je.FindString("failure." + reason + "\\.tag")
		if !found {
			return nil, errors.New(reason)
		}
		return nil, errors.New(reason + "/" + detail)
	}
}

func (z *lockImpl) lockOrUnlock(path mo_path.DropboxPath, endpoint string) (entry mo_file.Entry, err error) {
	l := z.ctx.Log().With(esl.String("path", path.Path()), esl.String("endpoint", endpoint))
	type PA struct {
		Path string `json:"path"`
	}
	type PB struct {
		Entries []PA `json:"entries"`
	}
	p := PB{
		[]PA{
			{path.Path()},
		},
	}
	res := z.ctx.Post(endpoint, api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	je, found := res.Success().Json().Find("entries.0")
	if !found {
		l.Debug("Entry array not found, or entry array has no entry", esl.String("response", res.Success().Json().RawString()))
		return nil, ErrorUnexpectedResponseFormat
	}
	return z.parseEntry(je)
}

func (z *lockImpl) Lock(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	return z.lockOrUnlock(path, "files/lock_file_batch")
}

func (z *lockImpl) Unlock(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	return z.lockOrUnlock(path, "files/unlock_file_batch")
}

func (z *lockImpl) lockOrUnlockBatch(paths []mo_path.DropboxPath, endpoint string) (entries map[string]LockResult, err error) {
	l := z.ctx.Log()
	type PA struct {
		Path string `json:"path"`
	}
	type PB struct {
		Entries []PA `json:"entries"`
	}
	pe := make([]PA, 0)
	for _, p := range paths {
		pe = append(pe, PA{Path: p.Path()})
	}
	p := PB{Entries: pe}
	entries = make(map[string]LockResult)

	res := z.ctx.Post(endpoint, api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	ja, found := res.Success().Json().FindArray("entries")
	if !found {
		l.Debug("`entries` not found")
		return entries, ErrorUnexpectedResponseFormat
	}
	if len(ja) != len(paths) {
		l.Debug("Length diff between `paths` and `entries`", esl.Int("lenEntries", len(ja)), esl.Int("lenPaths", len(paths)))
		return entries, ErrorUnexpectedResponseFormat
	}
	var lastErr error

	for i, p := range paths {
		entry, err := z.parseEntry(ja[i])
		entries[p.Path()] = LockResult{
			Entry: entry,
			Error: err,
		}
		if err != nil {
			lastErr = err
		}
	}
	return entries, lastErr
}

func (z *lockImpl) LockBatch(path []mo_path.DropboxPath) (entries map[string]LockResult, err error) {
	return z.lockOrUnlockBatch(path, "files/lock_file_batch")
}

func (z *lockImpl) UnlockBatch(path []mo_path.DropboxPath) (entries map[string]LockResult, err error) {
	return z.lockOrUnlockBatch(path, "files/unlock_file_batch")
}
