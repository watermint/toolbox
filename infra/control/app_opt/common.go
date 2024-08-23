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
	AutoOpen bool `json:"auto_open,omitempty"`

	// Limit bandwidth to downloading/uploading contents
	BandwidthKb int `json:"bandwidth_kb,omitempty"`

	// Set concurrency of worker execution
	Concurrency int `json:"concurrency,omitempty"`

	// Enable debug mode
	Debug bool `json:"debug,omitempty"`

	// Enable verbose mode
	Verbose bool `json:"verbose,omitempty"`

	// Enable experiments
	Experiment string `json:"experiment,omitempty"`

	// Language
	Lang mo_string.SelectString `json:"lang,omitempty"`

	// Memory budget
	BudgetMemory mo_string.SelectString `json:"budget_memory,omitempty"`

	// Storage budget
	BudgetStorage mo_string.SelectString `json:"budget_storage,omitempty"`

	// Job data
	RetainJobData mo_string.SelectString `json:"retain_job_data,omitempty"`

	// Set output format
	Output mo_string.SelectString `json:"output,omitempty"`

	// Set output format filter
	OutputFilter mo_string.OptionalString `json:"output_filter,omitempty"`

	// Explicitly set proxy the hostname and the port number
	Proxy mo_string.OptionalString `json:"proxy,omitempty"`

	// Path to auth database
	AuthDatabase mo_string.OptionalString `json:"auth_database,omitempty"`

	// Quiet mode
	Quiet bool `json:"quiet,omitempty"`

	// Do not store token in the file
	Secure bool `json:"secure,omitempty"`

	// Skip logging
	SkipLogging bool `json:"skip_logging,omitempty"`

	// Extra parameters
	Extra mo_string.OptionalString `json:"extra,omitempty"`

	// loaded extra options value
	extraCache *ExtraOpts `json:"extra_cache,omitempty"`

	// Specify workspace path
	Workspace mo_string.OptionalString `json:"workspace,omitempty"`
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
