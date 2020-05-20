package dc_readme

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_array"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"path/filepath"
	"sort"
	"strings"
)

func NewCommand(forPublish bool, commandPath string) dc_section.Section {
	return &Command{publish: forPublish, commandPath: commandPath}
}

type Command struct {
	publish            bool
	commandPath        string
	CommandHeader      app_msg.Message
	TableHeaderCommand app_msg.Message
	TableHeaderDesc    app_msg.Message
}

func (z Command) Title() app_msg.Message {
	return z.CommandHeader
}

func (z Command) commandName(spec rc_recipe.Spec) app_msg.Message {
	if z.publish {
		path := filepath.Join(z.commandPath, spec.SpecId()+".md")
		return app_msg.Raw(fmt.Sprintf("[%s](%s)", spec.CliPath(), path))
	} else {
		return app_msg.Raw(spec.CliPath())
	}
}

func (z Command) specListTable(ui app_ui.UI, header app_msg.Message, specs []rc_recipe.Spec) {
	ui.SubHeader(header)
	sort.Slice(specs, func(i, j int) bool {
		return strings.Compare(specs[i].CliPath(), specs[j].CliPath()) < 0
	})

	ui.WithTable("Commands", func(t app_ui.Table) {
		t.Header(z.TableHeaderCommand, z.TableHeaderDesc)
		for _, spec := range specs {
			t.Row(z.commandName(spec), spec.Title())
		}
	})
}

func (z Command) serviceKey(spec rc_recipe.Spec) string {
	specServices := spec.Services()
	sort.Strings(specServices)
	return strings.Join(specServices, "_")
}

func (z Command) recipes() []rc_recipe.Recipe {
	catalogue := app_catalogue.Current()
	recipes := catalogue.Recipes()
	available := make([]rc_recipe.Recipe, 0)

	for _, recipe := range recipes {
		spec := rc_spec.New(recipe)
		if spec.IsSecret() {
			continue
		}
		available = append(available, recipe)
	}
	return available
}

func (z Command) services() []string {
	recipes := z.recipes()
	services := make([]string, 0)
	for _, recipe := range recipes {
		spec := rc_spec.New(recipe)
		services = append(services, z.serviceKey(spec))
	}
	return es_array.NewByString(services...).Unique().Sort().AsStringArray()
}

func (z Command) specForService(services string) []rc_recipe.Spec {
	recipes := z.recipes()
	specs := make([]rc_recipe.Spec, 0)
	for _, recipe := range recipes {
		spec := rc_spec.New(recipe)
		key := z.serviceKey(spec)
		if key == services {
			specs = append(specs, spec)
		}
	}
	return specs
}

func (z Command) Body(ui app_ui.UI) {
	services := z.services()
	for _, svc := range services {
		suffix := svc
		if svc == "" {
			suffix = "utility"
		}
		header := app_msg.ObjMessage(&z, "services."+suffix)
		specs := z.specForService(svc)
		z.specListTable(ui, header, specs)
	}
}
