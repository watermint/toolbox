package gui

import (
	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"net/http"
)

type Server struct {
	ctl               app_control.Control
	expectedUserAgent string
}

func (z *Server) withSession(g *gin.Context, f func()) {
	if g.Request.UserAgent() != z.expectedUserAgent {
		g.HTML(
			http.StatusOK,
			"error",
			gin.H{
				"Header": "No session",
				"Detail": "Please close all `watermint toolbox` application, then start it.",
			},
		)
		return
	}

	f()
}

func (z *Server) home(g *gin.Context) {
	z.withSession(g, func() {
		menu := make([]map[string]string, 0)
		menu = append(menu, map[string]string{
			"Uri":         "/catalogue",
			"Title":       "Commands",
			"Description": "Show available commands",
		})

		g.HTML(
			http.StatusOK,
			"home",
			gin.H{
				"Menu": menu,
			},
		)
	})
}

func (z *Server) catalogue(g *gin.Context) {
	z.withSession(g, func() {
		ui := z.ctl.UI()

		catRecipes := app_catalogue.Current().Recipes()
		cat := make([]map[string]string, 0)
		for _, r := range catRecipes {
			s := rc_spec.New(r)
			if s.IsSecret() {
				continue
			}
			cat = append(cat, map[string]string{
				"Title":       s.CliPath(),
				"Description": ui.Text(s.Title()),
				"Uri":         "/command/" + Command{Command: s.CliPath()}.EncodeCommandUrl(),
			})
		}

		g.HTML(
			http.StatusOK,
			"catalogue",
			gin.H{
				"Commands": cat,
			},
		)
	})
}

func (z *Server) command(g *gin.Context) {
	z.withSession(g, func() {
		ui := z.ctl.UI()

		cmd := &Command{}
		if err := g.ShouldBindUri(cmd); err != nil {
			g.HTML(
				http.StatusOK,
				"error",
				gin.H{},
			)
			return
		}
		cmdCliPath, err := cmd.DecodeCommandName()
		if err != nil {
			g.HTML(
				http.StatusOK,
				"error",
				gin.H{},
			)
			return
		}
		spec := app_catalogue.Current().RecipeSpec(cmdCliPath)

		cmdRequiredValues := make([]map[string]interface{}, 0)
		cmdOptionalValues := make([]map[string]interface{}, 0)
		for _, valName := range spec.ValueNames() {
			valDefault := spec.ValueDefault(valName)
			valDefinition := map[string]interface{}{
				"Name":    valName,
				"Desc":    ui.Text(spec.ValueDesc(valName)),
				"Default": spec.ValueDefault(valName),
			}
			if valDefault == "" {
				cmdRequiredValues = append(cmdRequiredValues, valDefinition)
			} else {
				cmdOptionalValues = append(cmdOptionalValues, valDefinition)
			}
		}

		g.HTML(
			http.StatusOK,
			"command",
			gin.H{
				"Command":        cmdCliPath,
				"CommandTitle":   ui.Text(spec.Title()),
				"CommandDesc":    ui.TextOrEmpty(spec.Desc()),
				"RequiredValues": cmdRequiredValues,
				"OptionalValues": cmdOptionalValues,
			},
		)
	})
}
