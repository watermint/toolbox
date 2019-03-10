package cmd

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"path/filepath"
	"strings"
)

type SharedLinkFilter struct {
	FilterByPath       string
	FilterByVisibility string
}

func (z *SharedLinkFilter) FlagConfig(f *flag.FlagSet) {
	descFilterPath := "Filter by file path (default: no filter)"
	f.StringVar(&z.FilterByPath, "filter-path", "", descFilterPath)

	descFilterVisibility := fmt.Sprintf(
		"Filter by visibility (default: no filter, {%s})",
		strings.Join(z.SupportedVisibility(), ", "),
	)
	f.StringVar(&z.FilterByVisibility, "filter-visibility", "", descFilterVisibility)
}

func (z *SharedLinkFilter) IsAcceptable(link *dbx_sharing.SharedLink) bool {
	result := true
	if z.FilterByVisibility != "" {
		if link.PermissionResolvedVisibility != z.FilterByVisibility {
			result = false
		}
	}

	if z.FilterByPath != "" {
		filterPath := filepath.ToSlash(strings.ToLower(z.FilterByPath))

		if !strings.HasPrefix(link.PathLower, filterPath) {
			result = false
		}
	}
	return result
}

func (z *SharedLinkFilter) SupportedVisibility() []string {
	return []string{
		"public",
		"team_only",
		"password",
		"team_and_password",
		"shared_folder_only",
	}
}
