package mo_file

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"html"
)

type Match struct {
	Raw              json.RawMessage
	EntryTag         string `path:"metadata.metadata.\\.tag" json:"tag"`
	EntryName        string `path:"metadata.metadata.name" json:"name"`
	EntryPathDisplay string `path:"metadata.metadata.path_display" json:"path_display"`
	EntryPathLower   string `path:"metadata.metadata.path_lower" json:"path_lower"`
}

func (z *Match) Concrete() *ConcreteEntry {
	ce := &ConcreteEntry{}
	if err := api_parser.ParseModelPathRaw(ce, z.Raw, "metadata.metadata"); err != nil {
		app_root.Log().Debug("Unable to parse json", zap.Error(err), zap.ByteString("raw", z.Raw))
		return ce
	}
	ce.Raw = json.RawMessage(gjson.ParseBytes(z.Raw).Get("metadata.metadata").Raw)
	return ce
}

func (z *Match) HighlightHtml() string {
	hls := gjson.ParseBytes(z.Raw).Get("highlight_spans")
	if !hls.Exists() || !hls.IsArray() {
		return ""
	}

	body := ""
	for _, span := range hls.Array() {
		s := span.Get("highlight_str").String()
		h := span.Get("is_highlighted").Bool()

		e := html.EscapeString(s)
		if h {
			body = body + "<b>" + e + "</b>"
		} else {
			body = body + e
		}
	}
	return body
}

func (z *Match) Highlighted() (mh *MatchHighlighted) {
	return &MatchHighlighted{
		Raw:              z.Raw,
		EntryTag:         z.EntryTag,
		EntryName:        z.EntryName,
		EntryPathDisplay: z.EntryPathDisplay,
		EntryPathLower:   z.EntryPathLower,
		HighlightHtml:    z.HighlightHtml(),
	}
}

type MatchHighlighted struct {
	Raw              json.RawMessage
	EntryTag         string `json:"tag"`
	EntryName        string `json:"name"`
	EntryPathDisplay string `json:"path_display"`
	EntryPathLower   string `json:"path_lower"`
	HighlightHtml    string `json:"highlight_html"`
}
