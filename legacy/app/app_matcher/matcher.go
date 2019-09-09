package app_matcher

import (
	"errors"
	"flag"
	app2 "github.com/watermint/toolbox/legacy/app"
	"regexp"
	"strings"
)

type Matcher struct {
	ExecContext    *app2.ExecContext
	optStartsWith  string
	optEndsWith    string
	optRegex       string
	optInteractive bool
	regex          *regexp.Regexp
	matcher        func(t string) bool
}

func (z *Matcher) FlagConfig(f *flag.FlagSet) {
	descStartsWith := z.ExecContext.Msg("app.common.matcher.flag.starts_with").T()
	f.StringVar(&z.optStartsWith, "starts-with", "", descStartsWith)

	descEndsWith := z.ExecContext.Msg("app.common.matcher.flag.ends_with").T()
	f.StringVar(&z.optEndsWith, "ends-with", "", descEndsWith)

	descRegex := z.ExecContext.Msg("app.common.matcher.flag.regex").T()
	f.StringVar(&z.optRegex, "regex", "", descRegex)

	descInteractive := z.ExecContext.Msg("app.common.matcher.flag.interactive").T()
	f.BoolVar(&z.optInteractive, "interactive", true, descInteractive)
}

func (z *Matcher) IsInteractive() bool {
	return z.optInteractive
}

func (z *Matcher) Init() (err error) {
	nf := 0
	if z.optStartsWith != "" {
		nf++
	}
	if z.optEndsWith != "" {
		nf++
	}
	if z.optRegex != "" {
		nf++
	}
	if nf == 0 {
		z.ExecContext.Msg("app.common.matcher.err.no_options").TellError()
		return errors.New("no options")
	}
	if nf > 1 {
		z.ExecContext.Msg("app.common.matcher.err.too_many_options").TellError()
		return errors.New("too many options")
	}
	switch {
	case z.optStartsWith != "":
		z.matcher = func(t string) bool {
			return strings.HasPrefix(t, z.optStartsWith)
		}

	case z.optEndsWith != "":
		z.matcher = func(t string) bool {
			return strings.HasPrefix(t, z.optEndsWith)
		}

	case z.optRegex != "":
		z.regex, err = regexp.Compile(z.optRegex)
		if err != nil {
			z.ExecContext.Msg("app.common.matcher.err.regex_compile").WithData(struct {
				Regex string
				Error string
			}{
				Regex: z.optRegex,
				Error: err.Error(),
			}).TellError()
			return err
		}

		z.matcher = func(t string) bool {
			return z.regex.MatchString(t)
		}
	}
	return nil
}

func (z *Matcher) Match(t string) bool {
	if z.matcher == nil {
		z.ExecContext.Log().Error("Matcher not initialized")
		return false
	}
	return z.matcher(t)
}
