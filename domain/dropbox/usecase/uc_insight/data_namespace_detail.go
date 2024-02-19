package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type NamespaceDetail struct {
	// primary keys
	NamespaceId string `path:"shared_folder_id" gorm:"primaryKey"`

	// attributes
	FileId               string `path:"file_id" gorm:"index"`
	Name                 string `path:"name"`
	PolicyManageAccess   string `path:"policy.acl_update_policy.\\.tag"`
	PolicyMember         string `path:"policy.member_policy.\\.tag"`
	PolicyMemberResolved string `path:"policy.resolved_member_policy.\\.tag"`
	PolicySharedLink     string `path:"policy.shared_link_policy.\\.tag"`
	PolicyViewerInfo     string `path:"policy.viewer_info_policy.\\.tag"`
	AccessInheritance    string `path:"access_inheritance.\\.tag"`
	IsInsideTeamFolder   bool   `path:"is_inside_team_folder"`
	IsTeamFolder         bool   `path:"is_team_folder"`
	ParentNamespaceId    string `path:"parent_shared_folder_id" gorm:"index"`
	OwnerTeamId          string `path:"owner_team.id" gorm:"index"`
	OwnerTeamName        string `path:"owner_team.name"`
	IsSameTeam           bool   `gorm:"index"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewNamespaceDetail(s *mo_sharedfolder.SharedFolder) (ns *NamespaceDetail, err error) {
	ns = &NamespaceDetail{
		NamespaceId:          s.SharedFolderId,
		Name:                 s.Name,
		PolicyManageAccess:   s.PolicyManageAccess,
		PolicyMemberResolved: s.PolicyMember,
		PolicyMember:         s.PolicyMemberFolder,
		PolicySharedLink:     s.PolicySharedLink,
		PolicyViewerInfo:     s.PolicyViewerInfo,
		AccessInheritance:    s.AccessInheritance,
		IsInsideTeamFolder:   s.IsInsideTeamFolder,
		IsTeamFolder:         s.IsTeamFolder,
		ParentNamespaceId:    s.ParentSharedFolderId,
		OwnerTeamId:          s.OwnerTeamId,
		OwnerTeamName:        s.OwnerTeamName,
		Raw:                  s.Raw,
	}

	return ns, nil
}

type NamespaceDetailParam struct {
	NamespaceId string `path:"namespace_id" json:"namespace_id"`
	IsRetry     bool   `path:"is_retry" json:"is_retry"`
}

type NamespaceDetailError struct {
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	ApiError
}

func (z NamespaceDetailError) ToParam() interface{} {
	return &NamespaceDetailParam{
		NamespaceId: z.NamespaceId,
		IsRetry:     true,
	}
}

func (z tsImpl) scanNamespaceDetail(param *NamespaceDetailParam, stage eq_sequence.Stage, admin *mo_profile.Profile, team *mo_team.Info) (err error) {
	l := z.ctl.Log().With(esl.String("namespaceId", param.NamespaceId))
	isSuccessOrRetriable := func(err error) bool {
		dbxErr := dbx_error.NewErrors(err)
		if dbxErr == nil {
			return true
		}
		return dbxErr.HasPrefix("invalid_id") || dbxErr.Path().IsNotFound()
	}

	onError := func(err error) error {
		z.db.Save(&NamespaceDetailError{
			NamespaceId: param.NamespaceId,
			ApiError:    ApiErrorFromError(err),
		})
		return err
	}
	client := z.client.AsAdminId(admin.TeamMemberId)
	ns, err := sv_sharedfolder.New(client).Resolve(param.NamespaceId)
	if !isSuccessOrRetriable(err) {
		l.Debug("Unable to resolve namespace", esl.Error(err))
		return onError(err)
	}
	n, err := NewNamespaceDetail(ns)
	if err != nil {
		l.Debug("Unable to retrieve namespace detail", esl.Error(err))
		return onError(err)
	}
	m, err := sv_file.NewFiles(client).Resolve(mo_path.NewDropboxPath("ns:" + param.NamespaceId))
	switch {
	case err == nil:
		// fall through
	case isSuccessOrRetriable(err):
		l.Debug("Unable to resolve namespace", esl.Error(err))
		if param.IsRetry {
			z.db.Delete(&NamespaceDetailError{}, "namespace_id = ?", param.NamespaceId)
		}
		return nil
	default:
		l.Debug("Unable to resolve namespace", esl.Error(err))
		return onError(err)
	}

	ce := m.Concrete()

	n.FileId = ce.Id
	n.IsSameTeam = team.TeamId == n.OwnerTeamId
	z.db.Save(n)
	if z.db.Error != nil {
		l.Debug("Unable to save namespace detail", esl.Error(z.db.Error))
		return z.db.Error
	}

	if n.ParentNamespaceId == "" {

		z.db.Save(&NamespaceEntry{
			NamespaceId:              param.NamespaceId,
			FileId:                   ce.Id,
			ParentFolderId:           "",
			EntryType:                "folder",
			Name:                     ce.Name,
			Size:                     0,
			Rev:                      "",
			IsDownloadable:           false,
			HasExplicitSharedMembers: false,
			ClientModified:           "",
			ServerModified:           "",
			ContentHash:              "",
			PathLower:                ce.PathLower,
			PathDisplay:              ce.PathDisplay,
			EntryNamespaceId:         param.NamespaceId,
			ParentNamespaceId:        ce.ParentSharedFolderId,
			Updated:                  0,
			Raw:                      nil,
		})
	}

	if param.IsRetry {
		z.db.Delete(&NamespaceDetailError{}, "namespace_id = ?", param.NamespaceId)
	}

	return nil
}
