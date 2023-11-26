package qt_messages

import (
	"errors"
	"flag"
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
	"io"
	"os"
	"sort"
	"strings"
)

var (
	ErrorCliArgSuggestFound = errors.New("cli.arg message might required")
)

func SuggestCliArgs(ctl app_control.Control, r rc_recipe.Recipe) error {
	l := ctl.Log()
	spec := rc_spec.New(r)
	existingCliArgs := ""
	if ctl.UI().Exists(spec.CliArgs()) {
		existingCliArgs = ctl.UI().Text(spec.CliArgs())
	}

	suggests := make([]string, 0)
	missings := make([]string, 0)
	optionals := make([]string, 0)
	suggestCount := 0
	for _, valName := range spec.ValueNames() {
		v := spec.Value(valName)
		vd := spec.ValueDefault(valName)
		valArg := "-" + ecase.ToLowerKebabCase(valName) + " "
		found := strings.Contains(existingCliArgs, valArg)

		switch vt := v.(type) {
		case *rc_value.ValueString:
			if vd != "" {
				if found {
					optionals = append(optionals, valArg)
				}
				continue
			} else {
				suggests = append(suggests, valArg+"VALUE")
			}

		case *rc_value.ValueMoUrlUrl:
			suggests = append(suggests, valArg+"URL")

		case *rc_value.ValueDaJsonInput:
			suggests = append(suggests, valArg+"/LOCAL/PATH/TO/INPUT.json")

		case *rc_value.ValueDaTextInput:
			suggests = append(suggests, valArg+"/LOCAL/PATH/TO/INPUT.txt")

		case *rc_value.ValueDaGridDataInput:
			suggests = append(suggests, valArg+"/LOCAL/PATH/TO/INPUT.csv")

		case *rc_value.ValueMoPathFileSystemPath:
			suggests = append(suggests, valArg+"/LOCAL/PATH/TO/PROCESS")

		case *rc_value.ValueMoPathDropboxPath:
			suggests = append(suggests, valArg+"/DROPBOX/PATH/TO/PROCESS")

		case *rc_value.ValueFdFileRowFeed:
			suggests = append(suggests, valArg+"/PATH/TO/DATA_FILE.csv")

		case *rc_value.ValueMoTimeTime:
			if vt.IsOptional() {
				continue
			}
			suggests = append(suggests, valArg+`\"2020-04-01 17:58:38\"`)

		default:
			l.Debug("Skip suggest", esl.Any("value", vt))
			continue
		}

		if !found {
			missings = append(missings, valArg)
			suggestCount++
		}
	}
	if suggestCount > 0 || len(optionals) > 0 {
		sort.Strings(suggests)

		msgCliArgs := spec.CliArgs()
		l.Error("cli.arg might required", esl.String("key", msgCliArgs.Key()))
		l.Error("cli.arg optional options", esl.Strings("optional", optionals))
		l.Error("cli.arg missing options", esl.Strings("missing", missings))
		l.Error("cli.arg existing cli.arg", esl.String("existing", existingCliArgs))
		l.Error("cli.arg suggested", esl.String("suggest", strings.Join(suggests, " ")))

		SuggestMessages(ctl, func(out io.Writer) {
			fmt.Fprintf(out, `"%s":"%s",`, msgCliArgs.Key(), strings.Join(suggests, " "))
			fmt.Fprintln(out)
		})

		return ErrorCliArgSuggestFound
	}
	return nil
}

func SuggestMessages(ctl app_control.Control, suggest func(out io.Writer)) {
	out := es_stdout.NewDefaultOut(ctl.Feature())
	fmt.Fprintln(out, "Please add those messages to message resource files:")
	fmt.Fprintln(out, "====================================================")
	suggest(out)
	fmt.Fprintln(out, "====================================================")
}

func VerifyMessages(ctl app_control.Control) error {
	cat := app_catalogue.Current()
	recipes := cat.Recipes()
	root := rc_group_impl.NewGroup()
	for _, r := range recipes {
		s := rc_spec.New(r)
		root.Add(s)
	}

	qui := app_ui.NewDiscard(ctl.Messages(), ctl.Log())
	verifyGroup(root, qui)

	missing := qt_msgusage.Record().Missing()
	if len(missing) > 0 {
		sort.Strings(missing)
		for _, k := range missing {
			ctl.Log().Error("Key missing", esl.String(k, ""))
		}

		SuggestMessages(ctl, func(out io.Writer) {
			for _, m := range missing {
				switch {
				case strings.HasSuffix(m, ".flag.peer"):
					fmt.Fprintf(out, `"%s":"Account alias",`, m)
				case strings.HasSuffix(m, ".flag.file"):
					fmt.Fprintf(out, `"%s":"Path to data file",`, m)
				case strings.HasSuffix(m, ".agreement"):
					fmt.Fprintf(out, `"%s":"This feature is in an early stage of development. This is not well tested. Please proceed by typing 'yes' to agree & enable this feature.",`, m)
				case strings.HasSuffix(m, ".disclaimer"):
					fmt.Fprintf(out, `"%s":"WARN: The early access feature is enabled.",`, m)
				case strings.HasSuffix(m, ".format.output_grid_data"):
					fmt.Fprintf(out, `"%s":"Output format",`, m)
				case strings.HasSuffix(m, ".input_grid_data.desc"):
					fmt.Fprintf(out, `"%s":"Input grid data file path. '-' for read from STDIN.",`, m)
				default:
					fmt.Fprintf(out, `"%s":"",`, m)
				}
				fmt.Fprintln(out)
			}
		})

		return errors.New("missing key found")
	}
	return nil
}

func verifyGroup(g rc_group.Group, ui app_ui.UI) {
	g.PrintUsage(ui, os.Args[0], app.BuildId)
	for _, sg := range g.SubGroups() {
		verifyGroup(sg, ui)
	}
	for _, r := range g.Recipes() {
		verifyRecipe(g, r, ui)
	}
}

func verifyRecipe(g rc_group.Group, r rc_recipe.Spec, ui app_ui.UI) {
	f := flag.NewFlagSet("", flag.ContinueOnError)

	r.SetFlags(f, ui)
	r.PrintUsage(ui)
}
