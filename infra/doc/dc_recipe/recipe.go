package dc_recipe

type Recipe struct {
	Name            string            `json:"name"`
	Title           string            `json:"title"`
	Desc            string            `json:"desc"`
	Remarks         string            `json:"remarks"`
	Path            string            `json:"path"`
	CliArgs         string            `json:"cli_args"`
	CliNote         string            `json:"cli_note"`
	ConnUsePersonal bool              `json:"conn_use_personal"`
	ConnUseBusiness bool              `json:"conn_use_business"`
	ConnScopes      map[string]string `json:"conn_scopes"`
	Services        []string          `json:"services"`
	IsSecret        bool              `json:"is_secret"`
	IsConsole       bool              `json:"is_console"`
	IsExperimental  bool              `json:"is_experimental"`
	IsIrreversible  bool              `json:"is_irreversible"`
	IsTransient     bool              `json:"is_transient"`
	Reports         []*Report         `json:"reports"`
	Feeds           []*Feed           `json:"feeds"`
	Values          []*Value          `json:"values"`
}
