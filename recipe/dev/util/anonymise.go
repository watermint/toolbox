package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/time/ut_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"io"
	"net/url"
	"regexp"
	"strings"
)

const (
	jsonIndent = 2
)

type Anon struct {
}

func (z Anon) escapeConst(s string) (e string, completed bool) {
	prefixes := []string{
		"id:",
		"g:",
		"dbid:",
		"dbmid:",
		"dbtid:",
		"dbmsid:",
		"dbwsid:",
	}

	for _, p := range prefixes {
		if !strings.HasPrefix(s, p) {
			continue
		}
		return p + strings.Repeat("x", len(s)-len(p)), true
	}
	return s, false
}

func (z Anon) escapeEmail(s string) (e string, completed bool) {
	if dbx_util.RegexEmail.MatchString(s) {
		return regexp.MustCompile(`[^@^\.]`).ReplaceAllString(s, "x"), true
	}
	return s, false
}

func (z Anon) escapePath(s string) (e string, completed bool) {
	pre := regexp.MustCompile(`/(\w+/?)+`)
	if pre.MatchString(s) {
		return regexp.MustCompile(`[^/^\.]`).ReplaceAllString(s, "x"), true
	}
	return s, false
}

func (z Anon) escapeUrl(s string) (e string, completed bool) {
	u, err := url.Parse(s)
	if err != nil {
		// not an url
		return s, false
	}
	if u.Scheme == "" {
		return s, false
	}
	ap, _ := z.escapePath(u.Path)
	aq := strings.Repeat("x", len(u.RawQuery))
	af := strings.Repeat("x", len(u.Fragment))
	au := url.URL{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     ap,
		RawQuery: aq,
		Fragment: af,
	}
	return au.String(), true
}

func (z Anon) isDateTime(s string) bool {
	if _, v := ut_time.ParseTimestamp(s); v {
		return true
	}
	return false
}

func (z Anon) isTag(s string) bool {
	re := regexp.MustCompile(`^[a-z_]+$`)
	return re.MatchString(s)
}

func (z Anon) handleString(s string) string {
	if v := z.isDateTime(s); v {
		return s
	}
	if e, completed := z.escapeConst(s); completed {
		return e
	}
	if e, completed := z.escapeEmail(s); completed {
		return e
	}
	if e, completed := z.escapeUrl(s); completed {
		return e
	}
	if e, completed := z.escapePath(s); completed {
		return e
	}
	if z.isTag(s) {
		return s
	}
	return strings.Repeat("x", len(s))
}

func (z Anon) handleObject(o map[string]es_json.Json) interface{} {
	va := make(map[string]interface{})
	for k, v := range o {
		va[k] = z.handleValue(v)
	}
	return va
}

func (z Anon) handleArray(a []es_json.Json) interface{} {
	va := make([]interface{}, 0)
	for _, v := range a {
		va = append(va, z.handleValue(v))
	}
	return va
}

func (z Anon) handleValue(j es_json.Json) interface{} {
	l := esl.Default()
	if x, ok := j.String(); ok {
		y := z.handleString(x)
		l.Debug("Diff", esl.String("orig", x), esl.String("anon", y))
		return y
	}
	if x, ok := j.Number(); ok {
		if x.IsInt() {
			return x.Int64()
		} else {
			return x.Float64()
		}
	}
	if x, ok := j.Object(); ok {
		return z.handleObject(x)
	}
	if x, ok := j.Array(); ok {
		return z.handleArray(x)
	}
	if x, ok := j.Bool(); ok {
		return x
	}
	if ok := j.IsNull(); ok {
		return nil
	}

	l.Error("Unexpected value type", esl.Any("j", j))
	panic("unexpected value type")
}

func (z Anon) Anonymise(j es_json.Json) (json.RawMessage, error) {
	res := make(map[string]interface{})

	if resObj, found := j.FindObject("res"); !found {
		return nil, errors.New("response not found")
	} else {
		for k, v := range resObj {
			switch k {
			case "body":
				body, err := es_json.ParseString(v.RawString())
				if err != nil {
					return nil, err
				}
				res[k] = z.handleValue(body)
			case "headers":
				if headers, found := v.Object(); found {
					resHeaders := make(map[string]interface{})
					for hk, hv := range headers {
						switch hk {
						case "X-Dropbox-Request-Id":
							resHeaders[hk] = strings.Repeat("x", len(hv.RawString()))

						default:
							resHeaders[hk] = hv.Raw()
						}
					}
				}

			default:
				res[k] = v.Raw()
			}
		}
	}

	ajb, err := json.MarshalIndent(res, "", strings.Repeat(" ", jsonIndent))
	if err != nil {
		l := esl.Default()
		l.Error("Unable to marshal", esl.Error(err))
		return nil, err
	}
	return ajb, nil
}

type Anonymise struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	rc_recipe.RemarkConsole
	JobId mo_filter.Filter
	Path  mo_string.OptionalString
}

func (z *Anonymise) Preset() {
	z.JobId.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *Anonymise) anonymize(out io.Writer, lg app_job.LogFile) error {
	l := esl.Default()
	var buf bytes.Buffer
	if err := lg.CopyTo(&buf); err != nil {
		l.Debug("Unable to copy", esl.Error(err))
		return err
	}
	if buf.Len() < 1 {
		l.Debug("No content")
		return nil
	}
	chunks := strings.Split(buf.String(), "\n")
	for _, chunk := range chunks {
		if len(chunk) < 1 {
			continue
		}
		j, err := es_json.ParseString(chunk)
		if err != nil {
			l.Debug("unable to parse", esl.Error(err), esl.String("buf", chunk))
			return err
		}

		an := &Anon{}
		anonBody, err := an.Anonymise(j)
		if err != nil {
			l.Debug("unable to anonymize", esl.Error(err))
			return err
		}
		_, _ = out.Write(anonBody)
	}

	return nil
}

func (z *Anonymise) Exec(c app_control.Control) error {
	l := esl.Default()
	histories, err := app_job_impl.GetHistories(z.Path)
	if err != nil {
		return err
	}

	out := es_stdout.NewDefaultOut(c.Feature())
	for _, h := range histories {
		if !z.JobId.Accept(h.JobId()) {
			l.Debug("Skip", esl.String("jobId", h.JobId()))
			continue
		}
		logs, err := h.Logs()
		if err != nil {
			l.Debug("Unable to retrieve logs", esl.Error(err))
			return err
		}
		for _, lg := range logs {
			if lg.Type() != app_job.LogFileTypeCapture {
				l.Debug("Skip", esl.String("type", string(lg.Type())), esl.String("name", lg.Name()))
				continue
			}
			if err := z.anonymize(out, lg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Anonymise) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Anonymise{}, rc_recipe.NoCustomValues)
}
