package app_opt

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"io/ioutil"
	"runtime"
)

const (
	OutputNone           = "none"
	OutputText           = "text"
	OutputMarkdown       = "markdown"
	OutputJson           = "json"
	BudgetLow            = "low"
	BudgetNormal         = "normal"
	BudgetUnlimited      = "unlimited"
	LangAuto             = "auto"
	LangEnglish          = "en"
	LangJapanese         = "ja"
	RetainJobDataDefault = "default"
	RetainJobDataOnError = "on_error"
	RetainJobDataNone    = "none"
)

type ExtraOpts struct {
	AppKeys     map[string]string `json:"app_keys,omitempty"`
	Experiments []string          `json:"experiments,omitempty"`
}

func (z ExtraOpts) AppKey(key string) (value string, found bool) {
	if z.AppKeys == nil {
		return "", false
	}
	value, found = z.AppKeys[key]
	return
}

func (z ExtraOpts) HasExperiment(key string) bool {
	for _, experiment := range z.Experiments {
		if experiment == key {
			return true
		}
	}
	return false
}

type CommonOpts struct {
	// Automatically open the artifact folder, after successful execution
	AutoOpen bool

	// Limit bandwidth to downloading/uploading contents
	BandwidthKb int

	// Set concurrency of worker execution
	Concurrency int

	// Enable debug mode
	Debug bool

	// Enable verbose mode
	Verbose bool

	// Enable experiments
	Experiment string

	// Language
	Lang mo_string.SelectString

	// Memory budget
	BudgetMemory mo_string.SelectString

	// Storage budget
	BudgetStorage mo_string.SelectString

	// Job data
	RetainJobData mo_string.SelectString

	// Set output format
	Output mo_string.SelectString

	// Explicitly set proxy the hostname and the port number
	Proxy mo_string.OptionalString

	// Path to auth database
	AuthDatabase mo_string.OptionalString

	// Quiet mode
	Quiet bool

	// Do not store token in the file
	Secure bool

	// Skip logging
	SkipLogging bool

	// Extra parameters
	Extra mo_string.OptionalString

	// loaded extra options value
	extraCache *ExtraOpts

	// Specify workspace path
	Workspace mo_string.OptionalString
}

// ExtraLoad Load extra opts
func (z *CommonOpts) ExtraLoad() error {
	if z.Extra == nil || !z.Extra.IsExists() {
		return nil
	}

	data, err := ioutil.ReadFile(z.Extra.Value())
	if err != nil {
		return err
	}

	e := &ExtraOpts{}
	if err := json.Unmarshal(data, e); err != nil {
		return err
	}

	z.extraCache = e
	return nil
}

func (z *CommonOpts) ShouldDeleteJobData(err error) bool {
	switch z.RetainJobData.Value() {
	case RetainJobDataNone:
		// remove job data on exit
		return true

	case RetainJobDataOnError:
		if err != nil {
			return true
		}
	}

	return false
}

func (z *CommonOpts) ExtraOpts() ExtraOpts {
	if z.extraCache == nil {
		return ExtraOpts{}
	}
	return *z.extraCache
}

func (z *CommonOpts) Preset() {
	z.AutoOpen = false
	z.BandwidthKb = 0
	z.Concurrency = runtime.NumCPU()
	z.Debug = false
	z.Verbose = false
	z.SkipLogging = false
	z.Experiment = ""
	z.BudgetMemory.SetOptions(BudgetNormal, BudgetLow, BudgetNormal)
	z.BudgetStorage.SetOptions(string(app_budget.DefaultBudget), app_budget.StorageBudgets...)
	z.RetainJobData.SetOptions(RetainJobDataDefault, RetainJobDataDefault, RetainJobDataOnError, RetainJobDataNone)
	z.Lang.SetOptions(LangAuto, LangAuto, LangEnglish, LangJapanese)
	z.Output.SetOptions(OutputText, OutputText, OutputMarkdown, OutputJson, OutputNone)
	z.Proxy = mo_string.NewOptional("")
	z.AuthDatabase = mo_string.NewOptional("")
	z.Quiet = false
	z.Secure = false
	z.Workspace = mo_string.NewOptional("")
	z.Extra = mo_string.NewOptional("")
}

func Default() CommonOpts {
	com := CommonOpts{}
	com.Workspace = mo_string.NewOptional("")
	com.Proxy = mo_string.NewOptional("")
	com.BudgetMemory = mo_string.NewSelect()
	com.BudgetStorage = mo_string.NewSelect()
	com.RetainJobData = mo_string.NewSelect()
	com.Lang = mo_string.NewSelect()
	com.Output = mo_string.NewSelect()
	com.Extra = mo_string.NewOptional("")
	com.Preset()
	return com
}
