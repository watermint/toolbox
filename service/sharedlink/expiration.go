package sharedlink

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/integration/business"
	"github.com/watermint/toolbox/integration/sdk"
	"sync"
	"time"
)

type SharedLinkReceiverData interface {
}

type SharedLinkReceiverContent struct {
	SharedLink *sharing.SharedLinkMetadata
	Dropbox    *dropbox.Config
}
type SharedLinkReceiverEOF struct {
}

type UpdateSharedLinkExpireContext struct {
	TargetUser string
	Expiration time.Duration
	Overwrite  bool
}

func (u *UpdateSharedLinkExpireContext) IsTargetUser(m *team.TeamMemberInfo) bool {
	if u.TargetUser == "" {
		return true
	}
	return u.TargetUser == m.Profile.Email
}

func (u *UpdateSharedLinkExpireContext) ExpireTime() time.Time {
	return sdk.RebaseTimeForAPI(time.Now().Add(u.Expiration))
}

func (u *UpdateSharedLinkExpireContext) ShouldUpdate(s *sharing.SharedLinkMetadata) (bool, *time.Time) {
	seelog.Tracef("Validate sharedLink: id[%s] expire[%s]", s.Id, s.Expires.String())

	expire := u.ExpireTime()
	if s.Expires.IsZero() {
		return true, &expire
	}
	if u.Overwrite && s.Expires.After(expire) {
		return true, &expire
	}
	return false, nil
}

func UpdateSharedLinkForTeam(token string, expiration UpdateSharedLinkExpireContext) {
	queue := make(chan *team.TeamMemberInfo)
	receiver := make(chan SharedLinkReceiverData)
	wg := &sync.WaitGroup{}

	go UpdateSharedLinkExpire(receiver, expiration, wg)
	go business.LoadTeamMembersQueue(token, queue)

	for m := range queue {
		if m == nil {
			break
		}
		if !expiration.IsTargetUser(m) {
			seelog.Infof("Skip user [%s]", m.Profile.Email)
			continue
		}
		seelog.Infof("Loading shared link for member[%s]", m.Profile.Email)

		c := dropbox.Config{
			Token:      token,
			AsMemberID: m.Profile.TeamMemberId,
		}

		ListAllSharedLinks(&c, receiver)
	}
	receiver <- SharedLinkReceiverEOF{}
	wg.Wait()
}

func UpdateSharedLinkExpire(receiver chan SharedLinkReceiverData, expire UpdateSharedLinkExpireContext, wg *sync.WaitGroup) {
	wg.Add(1)
	for q := range receiver {
		switch s := q.(type) {
		case SharedLinkReceiverContent:
			if u, newExpire := expire.ShouldUpdate(s.SharedLink); u {
				UpdateExpiration(&s, newExpire)
			}

		case SharedLinkReceiverEOF:
			wg.Done()
			return
		}
	}
}

func UpdateExpiration(s *SharedLinkReceiverContent, newExpire *time.Time) {
	seelog.Infof("Update shared link[%s] expiration from [%s] to [%s]", s.SharedLink.Id, s.SharedLink.Expires.String(), newExpire.String())

	client := sharing.New(*s.Dropbox)
	settings := sharing.NewSharedLinkSettings()
	settings.Expires = *newExpire

	args := sharing.NewModifySharedLinkSettingsArgs(s.SharedLink.Url, settings)
	_, err := client.ModifySharedLinkSettings(args)
	if err != nil {
		seelog.Warnf("Unable to update expiration for id[%s] url[%s] error[%s]", s.SharedLink.Id, s.SharedLink.Url, err)
	}
}

func ListAllSharedLinks(dropbox *dropbox.Config, receiver chan SharedLinkReceiverData) error {
	cursor := ""
	client := sharing.New(*dropbox)

	for {
		arg := sharing.NewListSharedLinksArg()
		arg.Cursor = cursor
		seelog.Tracef("ListSharedLink with cursor[%s]", cursor)
		result, err := client.ListSharedLinks(arg)

		if err != nil {
			seelog.Warnf("Unable to load shared links: %s", err)
			return err
		}

		var entries []sharing.IsSharedLinkMetadata
		seelog.Tracef("[%d] links found, has_more[%t]", len(entries), result.HasMore)
		for _, l := range entries {
			switch meta := l.(type) {
			case *sharing.SharedLinkMetadata:
				receiver <- SharedLinkReceiverContent{
					SharedLink: meta,
					Dropbox:    dropbox,
				}

			default:
				seelog.Warnf("Unable to load shared link metadata: %s")
				return errors.New("Unable to extract metadata")
			}
		}

		if !result.HasMore {
			return nil
		}
		cursor = result.Cursor
	}
}
