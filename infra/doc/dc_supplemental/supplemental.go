package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
)

func Docs(media dc_index.MediaType) []dc_section.Document {
	return []dc_section.Document{
		&PathVariable{},
		&ExperimentalFeature{},
		&Troubleshooting{},
		NewDropboxBusiness(media),
		NewDocSpecChange(),
		&ReportingOptions{},
	}
}
