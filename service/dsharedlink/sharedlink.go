package dsharedlink

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/integration/business"
	"reflect"
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
	return time.Now().Round(time.Second).UTC().Add(u.Expiration)
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
	members, err := business.TeamMembers(token)
	if err != nil {
		seelog.Errorf("Unable to load team members: error[%s]", err)
		return
	}

	receiver := make(chan SharedLinkReceiverData)
	wg := &sync.WaitGroup{}

	go UpdateSharedLinkExpire(receiver, expiration, wg)

	for _, m := range members {
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
		result, err := client.ListSharedLinks(arg)

		if err != nil {
			seelog.Warnf("Unable to load shared links: %s", err)
			return err
		}

		for _, l := range result.Links {
			meta, err := extractSharedLinkMetadata(l)
			if err != nil {
				seelog.Warnf("Unable to load shared link metadata: %s", err)
				return err
			}

			receiver <- SharedLinkReceiverContent{
				SharedLink: meta,
				Dropbox:    dropbox,
			}
		}

		if !result.HasMore {
			return nil
		}
		cursor = result.Cursor
	}
}

// Workaround: There are no function to retrieve SharedLinkMetadata from IsSharedLinkMetadata.
func extractSharedLinkMetadata(link sharing.IsSharedLinkMetadata) (*sharing.SharedLinkMetadata, error) {
	v := reflect.ValueOf(link).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		tf := t.Field(i)

		if tf.Name == "SharedLinkMetadata" {
			val := vf.Interface()
			switch s := val.(type) {
			case sharing.SharedLinkMetadata:
				return &s, nil

			default:
				seelog.Warnf("Unknown type for 'SharedLinkMetadata' (%s)", vf.Type().Name())
			}
		}
	}
	return nil, errors.New("Unable to load 'SharedLinkMetadata'")
}
