package app_feature

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_config"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
	"os/user"
	"time"
)

type Feature interface {
	IsProduction() bool
	IsDebug() bool
	IsTest() bool
	IsQuiet() bool
	IsSecure() bool
	IsLowMemory() bool
	IsAutoOpen() bool
	UIFormat() string
	Config() app_config.Config
	OptInGet(oi OptIn) (f OptIn, found bool)
	OptInUpdate(oi OptIn) error
}

type OptIn interface {
	// The timestamp of opt-in, in ISO8601 format.
	// Empty when the user is not yet agreed.
	OptInTimestamp() string

	// Name of the user who opt'ed in.
	OptInUser() string

	// True when this feature enabled.
	OptInIsEnabled() bool

	// Name of the feature.
	OptInName(v OptIn) string

	// Description of the feature.
	OptInDescription(v OptIn) app_msg.Message

	// Message on enable this feature.
	OptInAgreement(v OptIn) app_msg.Message

	// Disclaimer message when the feature enabled runtime.
	OptInDisclaimer(v OptIn) app_msg.Message

	// Opt-in
	OptInCommit(enable bool)
}

func OptInFrom(v map[string]interface{}, oi OptIn) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, oi)
}

type OptInStatus struct {
	// The timestamp of opt-in, in ISO8601 format.
	Timestamp string `json:"timestamp"`

	// Name of the user who opt'ed in
	User string `json:"user"`

	// Opt-in status.
	Status bool `json:"status"`
}

func (z *OptInStatus) OptInCommit(enable bool) {
	usr, _ := user.Current()

	z.User = usr.Name
	z.Status = enable
	z.Timestamp = time.Now().Format(time.RFC3339)
}

func (z *OptInStatus) OptInName(v OptIn) string {
	return ut_reflect.Key(app.Pkg, v)
}

func (z *OptInStatus) OptInTimestamp() string {
	return z.Timestamp
}

func (z *OptInStatus) OptInUser() string {
	return z.User
}

func (z *OptInStatus) OptInIsEnabled() bool {
	return z.Status
}

func (z *OptInStatus) OptInAgreement(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "agreement")
}

func (z *OptInStatus) OptInDisclaimer(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "disclaimer")
}

func (z *OptInStatus) OptInDescription(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "desc")
}