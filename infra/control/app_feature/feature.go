package app_feature

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os/user"
	"time"
)

type Feature interface {
	IsProduction() bool
	IsDebug() bool
	IsVerbose() bool
	IsTest() bool
	IsTestWithMock() bool
	IsTestWithSeqReplay() (replay []nw_replay.Response, enabled bool)
	IsTestWithReplay() (replay kv_storage.Storage, enabled bool)
	IsQuiet() bool
	IsSecure() bool
	IsAutoOpen() bool
	IsTransient() bool
	IsSkipLogging() bool

	// Experiment Is the experiment enabled or not.
	Experiment(name string) bool

	// UIFormat UI format
	UIFormat() string

	// Concurrency configuration.
	Concurrency() int

	// Home Toolbox home path. Returns empty if a user doesn't specify the path.
	Home() string

	// BudgetMemory Budget for memory usage
	BudgetMemory() app_budget.Budget

	// BudgetStorage Budget for storage usage
	BudgetStorage() app_budget.Budget

	// KvsEngine Key-value storage engine type
	KvsEngine() kv_storage.KvsEngine

	// OptInGet Retrieve feature
	OptInGet(oi OptIn) (f OptIn, found bool)

	// OptInUpdate Update opt-in feature
	OptInUpdate(oi OptIn) error

	// AsTest With test mode
	AsTest(useMock bool) Feature

	// AsSeqReplayTest With sequential replay test
	AsSeqReplayTest(replay []nw_replay.Response) Feature

	// AsReplayTest With replay test
	AsReplayTest(replays kv_storage.Storage) Feature

	// AsQuiet With quiet mode, but this will not guarantee UI/log are converted into quiet mode.
	AsQuiet() Feature

	// ConsoleLogLevel Console log level
	ConsoleLogLevel() esl.Level

	// Extra options
	Extra() app_opt.ExtraOpts

	// PathAuthRepository path to auth repository
	PathAuthRepository() string

	// IsDefaultPathAuthRepository returns true when the repository
	IsDefaultPathAuthRepository() bool
}

type OptIn interface {
	// The timestamp of opt-in, in ISO8601 format.
	// Empty when the user is not yet agreed.
	OptInTimestamp() string

	// Name of the user who opt'ed in.
	OptInUser() string

	// True when this feature enabled.
	OptInIsEnabled() bool

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

	switch {
	case usr.Name != "":
		z.User = usr.Name
	case usr.Username != "":
		z.User = usr.Username
	default:
		z.User = "unknown"
	}
	z.Status = enable
	z.Timestamp = time.Now().Format(time.RFC3339)
}

func (z OptInStatus) OptInTimestamp() string {
	return z.Timestamp
}

func (z OptInStatus) OptInUser() string {
	return z.User
}

func (z OptInStatus) OptInIsEnabled() bool {
	return z.Status
}

func OptInName(v OptIn) string {
	return es_reflect.Key(v)
}

func OptInAgreement(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "agreement")
}

func OptInDisclaimer(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "disclaimer")
}

func OptInDescription(v OptIn) app_msg.Message {
	return app_msg.ObjMessage(v, "desc")
}
