package mo_file_size

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
)

const (
	ApiComplexityThreshold = 10_000
)

type Size struct {
	Path            string `json:"path"`
	CountFile       int64  `json:"count_file"`
	CountFolder     int64  `json:"count_folder"`
	CountDescendant int64  `json:"count_descendant"`
	Size            int64  `json:"size"`
	ApiComplexity   int64  `json:"api_complexity"`
}

func (z Size) Eval(entries []mo_file.Entry) Size {
	if n := len(entries); n <= ApiComplexityThreshold {
		z.ApiComplexity = int64(n)
	} else {
		z.ApiComplexity = 1
	}

	for _, entry := range entries {
		z.CountDescendant++
		if f, e := entry.File(); e {
			z.CountFile++
			z.Size += f.Size
		} else if _, e := entry.Folder(); e {
			z.CountFolder++
		}
	}
	return z
}

func (z Size) Plus(path string, s Size) Size {
	return Size{
		Path:            path,
		CountFile:       z.CountFile + s.CountFile,
		CountFolder:     z.CountFolder + s.CountFolder,
		CountDescendant: z.CountDescendant + s.CountDescendant,
		ApiComplexity:   z.ApiComplexity + s.ApiComplexity,
		Size:            z.Size + s.Size,
	}
}

type NamespaceSize struct {
	NamespaceName     string `json:"namespace_name"`
	NamespaceId       string `json:"namespace_id"`
	NamespaceType     string `json:"namespace_type"`
	OwnerTeamMemberId string `json:"owner_team_member_id"`
	Path              string `json:"path"`
	CountFile         int64  `json:"count_file"`
	CountFolder       int64  `json:"count_folder"`
	CountDescendant   int64  `json:"count_descendant"`
	Size              int64  `json:"size"`
	ApiComplexity     int64  `json:"api_complexity"`
}

func NewNamespaceSize(namespace *mo_namespace.Namespace, size Size) *NamespaceSize {
	return &NamespaceSize{
		NamespaceName:     namespace.Name,
		NamespaceId:       namespace.NamespaceId,
		NamespaceType:     namespace.NamespaceType,
		OwnerTeamMemberId: namespace.TeamMemberId,
		Path:              size.Path,
		CountFile:         size.CountFile,
		CountFolder:       size.CountFolder,
		CountDescendant:   size.CountDescendant,
		Size:              size.Size,
		ApiComplexity:     size.ApiComplexity,
	}
}
