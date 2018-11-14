package cmd

import (
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"path/filepath"
	"strings"
)

type SharedLinkFilter struct {
	FilterByPath       string
	FilterByVisibility string
}

func (s *SharedLinkFilter) FlagConfig(f *flag.FlagSet) {
	descFilterPath := "Filter by file path (default: no filter)"
	f.StringVar(&s.FilterByPath, "filter-path", "", descFilterPath)

	descFilterVisibility := fmt.Sprintf(
		"Filter by visibility (default: no filter, {%s})",
		strings.Join(s.SupportedVisibility(), ", "),
	)
	f.StringVar(&s.FilterByVisibility, "filter-visibility", "", descFilterVisibility)
}

func (s *SharedLinkFilter) IsAcceptable(link *dbx_sharing.SharedLink) bool {
	linkBody := string(link.Link)
	result := true
	if s.FilterByVisibility != "" {
		visibility := gjson.Get(linkBody, "link_permissions.resolved_visibility.\\.tag").String()
		if visibility != s.FilterByVisibility {
			result = false
		}
	}

	if s.FilterByPath != "" {
		pathLower := gjson.Get(linkBody, "path_lower").String()
		filterPath := filepath.ToSlash(strings.ToLower(s.FilterByPath))

		if !strings.HasPrefix(pathLower, filterPath) {
			result = false
		}
	}
	return result
}

func (s *SharedLinkFilter) SupportedVisibility() []string {
	return []string{
		"public",
		"team_only",
		"password",
		"team_and_password",
		"shared_folder_only",
	}
}
