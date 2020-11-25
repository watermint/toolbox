package dc_supplemental

import "github.com/watermint/toolbox/infra/doc/dc_section"

var (
	Docs = []dc_section.Document{
		&PathVariable{},
		&ExperimentalFeature{},
		&Troubleshooting{},
		&DropboxBusiness{},
	}
)
