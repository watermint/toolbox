package sv_event

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/calendar/model/mo_event"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Event interface {
	ListEach(h func(event *mo_event.Event), calendarId string, opts ListOpts) error
}

type ListOpt func(o ListOpts) ListOpts

type ListOpts struct {
	MaxAttendees           *int   `url:"maxAttendees,omitempty"`
	MaxResults             *int   `url:"maxResults,omitempty"`
	OrderBy                string `url:"orderBy,omitempty"`
	PageToken              string `url:"pageToken,omitempty"`
	Q                      string `url:"q,omitempty"`
	SharedExtendedProperty string `url:"sharedExtendedProperty,omitempty"`
	ShowDeleted            *bool  `url:"showDeleted,omitempty"`
	ShowHiddenInvitations  *bool  `url:"showHiddenInvitations,omitempty"`
	SingleEvents           *bool  `url:"singleEvents,omitempty"`
	TimeMax                string `url:"timeMax,omitempty"`
	TimeMin                string `url:"timeMin,omitempty"`
	TimeZone               string `url:"timeZone,omitempty"`
	UpdatedMin             string `url:"updatedMin,omitempty"`
}

func TimeMax(v string) ListOpt {
	return func(o ListOpts) ListOpts {
		o.TimeMax = v
		return o
	}
}
func TimeMin(v string) ListOpt {
	return func(o ListOpts) ListOpts {
		o.TimeMin = v
		return o
	}
}
func Query(v string) ListOpt {
	return func(o ListOpts) ListOpts {
		o.Q = v
		return o
	}
}

func New(ctx goog_client.Client) Event {
	return &eventImpl{
		ctx: ctx,
	}
}

type eventImpl struct {
	ctx goog_client.Client
}

func (z eventImpl) ListEach(h func(event *mo_event.Event), calendarId string, opts ListOpts) error {
	endPoint := "calendars/" + calendarId + "/events"
	listEach := func(opts ListOpts) (nextPageToken string, err error) {
		res := z.ctx.Get(endPoint, api_request.Query(&opts))
		if err, f := res.Failure(); f {
			return "", err
		}
		j, err := res.Success().AsJson()
		if err != nil {
			return "", err
		}
		err = j.FindArrayEach("items", func(e es_json.Json) error {
			event := &mo_event.Event{}
			if err := e.Model(event); err != nil {
				return err
			}
			h(event)
			return nil
		})
		if err != nil {
			return "", err
		}
		nextPageToken, found := j.FindString("nextPageToken")
		if found {
			return nextPageToken, nil
		} else {
			return "", nil
		}
	}

	pageToken, err := listEach(opts)
	if err != nil {
		return err
	}

	for pageToken != "" {
		pageToken, err = listEach(ListOpts{
			PageToken: pageToken,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
