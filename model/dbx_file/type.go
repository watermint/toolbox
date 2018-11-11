package dbx_file

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
)

type Folder struct {
	Name          string          `json:"name"`
	FolderId      string          `json:"id"`
	PathLower     string          `json:"path_lower"`
	PathDisplay   string          `json:"path_display"`
	SharingInfo   json.RawMessage `json:"sharing_info"`
	PropertyGroup json.RawMessage `json:"property_group"`
}

type File struct {
	Name                     string          `json:"name"`
	FileId                   string          `json:"id"`
	ClientModified           string          `json:"client_modified"`
	ServerModified           string          `json:"server_modified"`
	Revision                 string          `json:"rev"`
	Size                     int64           `json:"size"`
	PathLower                string          `json:"path_lower,omitempty"`
	PathDisplay              string          `json:"path_display,omitempty"`
	MediaInfo                json.RawMessage `json:"media_info,omitempty"`
	SymlinkInfo              json.RawMessage `json:"symlink_info,omitempty"`
	SharingInfo              json.RawMessage `json:"sharing_info,omitempty"`
	PropertyGroups           json.RawMessage `json:"property_groups,omitempty"`
	HasExplicitSharedMembers bool            `json:"has_explicit_shared_members,omitempty"`
	ContentHash              string          `json:"content_hash,omitempty"`
}

type Deleted struct {
	Name        string `json:"name"`
	PathLower   string `json:"path_lower"`
	PathDisplay string `json:"path_display"`
}

type EntryParser struct {
	Logger   *zap.Logger
	log      *zap.Logger
	OnError  func(annotation dbx_api.ErrorAnnotation) bool
	OnFolder func(folder *Folder) bool
	OnFile   func(file *File) bool
	OnDelete func(deleted *Deleted) bool
}

func (p *EntryParser) Parse(entry gjson.Result) bool {
	p.log = p.Logger.With(zap.String("parser", "EntryParser"))

	tag := entry.Get(dbx_api.ResJsonDotTag)
	if !tag.Exists() {
		return dbx_api.ParserError(
			"`.tag` not found in the entry",
			entry.Str,
			p.log,
			p.OnError,
		)
	}

	switch tag.String() {
	case "file":
		return p.parseFile(entry)

	case "folder":
		return p.parseFolder(entry)

	case "deleted":
		return p.parseDeleted(entry)

	default:
		return dbx_api.ParserError(
			"unknown `.tag` value found in the entry",
			entry.Str,
			p.log.With(zap.String("tag", tag.String())),
			p.OnError,
		)
	}

	return true
}

func (p *EntryParser) parseFile(entry gjson.Result) bool {
	f := &File{}
	if err := json.Unmarshal([]byte(entry.Raw), f); err != nil {
		dbx_api.ParserError(
			"unable to unmarshal entry",
			entry.Str,
			p.log.With(zap.Error(err)),
			p.OnError,
		)
		return false
	}
	return p.OnFile(f)
}
func (p *EntryParser) parseFolder(entry gjson.Result) bool {
	f := &Folder{}
	if err := json.Unmarshal([]byte(entry.Raw), f); err != nil {
		dbx_api.ParserError(
			"unable to unmarshal entry",
			entry.Str,
			p.log.With(zap.Error(err)),
			p.OnError,
		)
		return false
	}
	return p.OnFolder(f)
}
func (p *EntryParser) parseDeleted(entry gjson.Result) bool {
	d := &Deleted{}
	if err := json.Unmarshal([]byte(entry.Raw), d); err != nil {
		dbx_api.ParserError(
			"unable to unmarshal entry",
			entry.Str,
			p.log.With(zap.Error(err)),
			p.OnError,
		)
		return false
	}
	return p.OnDelete(d)
}
