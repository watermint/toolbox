package mo_sharedlink

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/infra/api_parser"
)

type SharedLink interface {
	SharedLinkId() string
	LinkTag() string
	LinkUrl() string
	LinkName() string
	LinkExpires() string
	LinkVisibility() string
	LinkPathLower() string
	File() (file *File, ok bool)
	Folder() (folder *Folder, ok bool)
}

type Metadata struct {
	Raw        json.RawMessage
	Id         string `path:"id"`
	Tag        string `path:"\\.tag"`
	Url        string `path:"url"`
	Name       string `path:"name"`
	Expires    string `path:"expires"`
	PathLower  string `path:"path_lower"`
	Visibility string `path:"link_permissions.resolved_visibility.\\.tag"`
}

func (z *Metadata) LinkTag() string {
	return z.Tag
}

func (z *Metadata) File() (file *File, ok bool) {
	if z.Tag == "file" {
		file := &File{}
		if err := api_parser.ParseModelRaw(file, z.Raw); err != nil {
			return nil, false
		}
		return file, true
	}
	return nil, false
}

func (z *Metadata) Folder() (folder *Folder, ok bool) {
	if z.Tag == "folder" {
		folder := &Folder{}
		if err := api_parser.ParseModelRaw(folder, z.Raw); err != nil {
			return nil, false
		}
		return folder, true
	}
	return nil, false
}

func (z *Metadata) SharedLinkId() string {
	return z.Id
}

func (z *Metadata) LinkUrl() string {
	return z.Url
}

func (z *Metadata) LinkName() string {
	return z.Name
}

func (z *Metadata) LinkExpires() string {
	return z.Expires
}

func (z *Metadata) LinkPathLower() string {
	return z.PathLower
}

func (z *Metadata) LinkVisibility() string {
	return z.Visibility
}

type File struct {
	Raw            json.RawMessage
	Id             string `path:"id"`
	Tag            string `path:"\\.tag"`
	Url            string `path:"url"`
	Name           string `path:"name"`
	ClientModified string `path:"client_modified"`
	ServerModified string `path:"server_modified"`
	Revision       string `path:"rev"`
	Expires        string `path:"expires"`
	PathLower      string `path:"path_lower"`
	Size           int    `path:"size"`
	Visibility     string `path:"link_permissions.resolved_visibility.\\.tag"`
}

func (z *File) SharedLinkId() string {
	return z.Id
}

func (z *File) LinkTag() string {
	return z.Tag
}

func (z *File) LinkUrl() string {
	return z.Url
}

func (z *File) LinkName() string {
	return z.Name
}

func (z *File) LinkExpires() string {
	return z.Expires
}

func (z *File) LinkPathLower() string {
	return z.PathLower
}

func (z *File) LinkVisibility() string {
	return z.LinkVisibility()
}

func (z *File) File() (file *File, ok bool) {
	return z, true
}

func (z *File) Folder() (folder *Folder, ok bool) {
	return nil, false
}

type Folder struct {
	Raw        json.RawMessage
	Id         string `path:"id"`
	Tag        string `path:"\\.tag"`
	Url        string `path:"url"`
	Name       string `path:"name"`
	Expires    string `path:"expires"`
	PathLower  string `path:"path_lower"`
	Visibility string `path:"link_permissions.resolved_visibility.\\.tag"`
}

func (z *Folder) SharedLinkId() string {
	return z.Id
}

func (z *Folder) LinkTag() string {
	return z.Tag
}

func (z *Folder) LinkUrl() string {
	return z.Url
}

func (z *Folder) LinkName() string {
	return z.Name
}

func (z *Folder) LinkExpires() string {
	return z.Expires
}

func (z *Folder) LinkVisibility() string {
	return z.Visibility
}

func (z *Folder) LinkPathLower() string {
	return z.PathLower
}

func (z *Folder) File() (file *File, ok bool) {
	return nil, false
}

func (z *Folder) Folder() (folder *Folder, ok bool) {
	return z, true
}
