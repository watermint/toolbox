package recipe

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/skratchdot/open-golang/open"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_conn_impl"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_control_impl"
	"github.com/watermint/toolbox/experimental/app_control_launcher"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_recipe_group"
	"github.com/watermint/toolbox/experimental/app_template"
	"github.com/watermint/toolbox/experimental/app_ui"
	"github.com/watermint/toolbox/experimental/app_user"
	"github.com/watermint/toolbox/experimental/app_vo"
	"github.com/watermint/toolbox/experimental/app_vo_impl"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	webPort                = 7800
	webUserHashCookie      = "user_hash"
	webPathRoot            = "/"
	webPathLogin           = "/login"
	webPathHome            = "/home"
	webPathConnectStart    = "/connect/start"
	webPathConnectAuth     = "/connect/auth"
	webPathConnectFinish   = "/connect/finish"
	webPathRun             = "/run"
	webPathJob             = "/job"
	webPathJobArtifact     = "/artifact"
	webPathForbidden       = "/error/forbidden"
	webPathServerError     = "/error/server"
	webPathAuthFailed      = "/error/auth_failed"
	webPathCommandNotFound = "/error/command_not_found"
)

type webJobRun struct {
	name      string
	jobId     string
	recipe    app_recipe.Recipe
	vo        app_vo.ValueObject
	uc        app_control.Control
	uiLogFile *os.File
}

type WebVO struct {
	Port int
}

func (z *WebVO) Validate(t app_vo.Validator) {
}

type Web struct {
}

func (z *Web) Console() {
}

func (z *Web) Requirement() app_vo.ValueObject {
	return &WebVO{
		Port: webPort,
	}
}

func (z *Web) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	wvo := vo.(*WebVO)

	l := k.Log()
	repo, err := app_user.SingleUserRepository(k.Control().Workspace())
	if err != nil {
		return err
	}
	rur := repo.(app_user.RootUserRepository)
	ru := rur.RootUser()

	if k.Control().IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	hfs := k.Control().(app_control.ControlHttpFileSystem)
	htp := hfs.Template()
	htr := htp.(render.HTMLRender)
	cl := k.Control().(app_control_launcher.ControlLauncher)

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))
	//g.StaticFS("/assets", hfs.HttpFileSystem())
	g.HTMLRender = htr

	baseUrl := fmt.Sprintf("http://localhost:%d", wvo.Port)
	jobChan := make(chan *webJobRun)

	wh := &WebHandler{
		control:     k.Control(),
		Template:    htp,
		Launcher:    cl,
		BaseUrl:     baseUrl,
		authForUser: make(map[string]api_auth.Web),
		jobChan:     jobChan,
	}
	wh.Setup(g)

	go z.jobRunner(k.Control(), jobChan)

	loginUrl := baseUrl + webPathLogin + "/" + ru.UserHash()

	k.Log().Info("Login url", zap.String("url", loginUrl))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))

	time.Sleep(2 * time.Second)

	open.Start(loginUrl)

	return nil
}

func (z *Web) jobRunner(ctl app_control.Control, jc <-chan *webJobRun) {
	for job := range jc {
		l := ctl.Log().With(zap.String("name", job.name), zap.String("jobId", job.jobId))
		l.Debug("Start a new job")
		k := app_kitchen.NewKitchen(job.uc, job.vo)
		err := job.recipe.Exec(k)
		if err != nil {
			l.Error("Unable to finish the job", zap.Error(err))
		}
		l.Debug("Closing log file")
		job.uiLogFile.Close()

		l.Debug("Job spin down")
		job.uc.Down()

		l.Debug("The job finished")
	}
}

type WebHandler struct {
	control     app_control.Control
	Template    app_template.Template
	Launcher    app_control_launcher.ControlLauncher
	BaseUrl     string
	Root        *app_recipe_group.Group
	authForUser map[string]api_auth.Web
	jobChan     chan *webJobRun
}

func (z *WebHandler) auth(user app_user.User, uc app_control.Control) api_auth.Web {
	if a, ok := z.authForUser[user.UserHash()]; ok {
		return a
	} else {
		a = api_auth_impl.NewWeb(uc)
		z.authForUser[user.UserHash()] = a
		return a
	}
}

func (z *WebHandler) setupUrls(g *gin.Engine) {
	g.Use()

	g.GET(webPathLogin+"/:user_hash", z.Login)

	g.GET(webPathForbidden, z.Forbidden)
	g.GET(webPathServerError, z.ServerError)
	g.GET(webPathCommandNotFound, z.CommandNotFound)
	g.GET(webPathAuthFailed, z.AuthFailed)

	g.NoRoute(z.NotFound)
	z.Template.Define("error", "layout/layout.html", "pages/error.html")

	g.GET(webPathHome, z.Home)
	g.GET(webPathHome+"/:command", z.Home)
	g.POST(webPathHome+"/:command", z.Home)
	g.GET(webPathHome+"/:command/:tokenType", z.Home)
	g.GET("param", z.renderRecipeParam)
	z.Template.Define("home-catalogue", "layout/layout.html", "pages/catalogue.html")
	z.Template.Define("home-recipe-conn", "layout/layout.html", "pages/recipe_conn.html")
	z.Template.Define("home-recipe-param", "layout/layout.html", "pages/recipe_param.html")
	z.Template.Define("home-recipe-run", "layout/layout.html", "pages/recipe_run.html")

	g.POST(webPathRun+"/:command", z.Run)
	g.POST(webPathJob+"/:command/:jobId", z.Job)
	z.Template.Define(webPathJob, "layout/layout.html", "pages/job.html")

	g.GET(webPathJobArtifact+"/:command/:jobId/:artifactName", z.Artifact)

	g.GET(webPathRoot, z.Instruction)
	z.Template.Define(webPathRoot, "layout/layout.html", "pages/home.html")

	g.GET(webPathConnectStart, z.connectStart)
	g.GET(webPathConnectAuth, z.connectAuth)
	g.GET(webPathConnectFinish, z.connectFinish)
	z.Template.Define(webPathConnectFinish, "layout/layout.html", "pages/recipe_conn_finish.html")
}

func (z *WebHandler) setupCatalogue() {
	recipes := z.Launcher.Catalogue()
	z.Root = app_recipe_group.NewGroup([]string{}, "")
	for _, r := range recipes {
		_, ok := r.(app_recipe.SecretRecipe)
		if ok {
			continue
		}
		_, ok = r.(app_recipe.ConsoleRecipe)
		if ok {
			continue
		}

		z.Root.Add(r)
	}
}

func (z *WebHandler) findRecipe(cmd string) (grp *app_recipe_group.Group, rcp app_recipe.Recipe, err error) {
	cmdPath := strings.Split(cmd, "-")
	_, grp, rcp, _, err = z.Root.Select(cmdPath)

	if cmd == "" {
		grp = z.Root
		rcp = nil
		err = nil
	}
	return
}

func (z *WebHandler) recipeRequirements(rcp app_recipe.Recipe) (conns map[string]string, paramTypes map[string]string, paramDefaults map[string]interface{}) {
	conns = make(map[string]string)
	paramTypes = make(map[string]string)
	paramDefaults = make(map[string]interface{})

	var vo interface{} = rcp.Requirement()
	vc := app_vo_impl.NewValueContainer(vo)

	for k, v := range vc.Values {
		if d, ok := v.(bool); ok {
			paramTypes[k] = "bool"
			paramDefaults[k] = d
		} else if _, ok := v.(app_conn.ConnBusinessInfo); ok {
			conns[k] = api_auth.DropboxTokenBusinessInfo
		} else if _, ok := v.(app_conn.ConnBusinessFile); ok {
			conns[k] = api_auth.DropboxTokenBusinessFile
		} else if _, ok := v.(app_conn.ConnBusinessAudit); ok {
			conns[k] = api_auth.DropboxTokenBusinessAudit
		} else if _, ok := v.(app_conn.ConnBusinessMgmt); ok {
			conns[k] = api_auth.DropboxTokenBusinessManagement
		} else if _, ok := v.(app_conn.ConnUserFile); ok {
			conns[k] = api_auth.DropboxTokenFull
		}
	}
	return
}

func (z *WebHandler) Setup(g *gin.Engine) {
	z.setupCatalogue()
	z.setupUrls(g)
}

func (z *WebHandler) Login(g *gin.Context) {
	hash := g.Param("user_hash")
	repo, err := app_user.SingleUserRepository(z.control.Workspace())
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
	}
	_, err = repo.Resolve(hash)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
	}

	g.SetCookie(
		webUserHashCookie,
		hash,
		86400,
		"/",
		"localhost",
		false,
		true,
	)
	g.Redirect(http.StatusTemporaryRedirect, webPathHome)
}

func (z *WebHandler) NotFound(g *gin.Context) {
	g.HTML(
		http.StatusNotFound,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.notfound.header"),
			"Detail": z.control.UI().Text("web.error.notfound.detail"),
		},
	)
}

func (z *WebHandler) ServerError(g *gin.Context) {
	g.HTML(
		http.StatusInternalServerError,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.server.header"),
			"Detail": z.control.UI().Text("web.error.server.detail"),
		},
	)
}
func (z *WebHandler) CommandNotFound(g *gin.Context) {
	g.HTML(
		http.StatusBadRequest,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.command_not_found.header"),
			"Detail": z.control.UI().Text("web.error.command_not_found.detail"),
		},
	)
}

func (z *WebHandler) AuthFailed(g *gin.Context) {
	g.HTML(
		http.StatusOK,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.auth_failed.header"),
			"Detail": z.control.UI().Text("web.error.auth_failed.detail"),
		},
	)
}

func (z *WebHandler) Forbidden(g *gin.Context) {
	g.HTML(
		http.StatusForbidden,
		webPathServerError,
		gin.H{
			"Header": z.control.UI().Text("web.error.forbidden.header"),
			"Detail": z.control.UI().Text("web.error.forbidden.detail"),
		},
	)
}

func (z *WebHandler) Instruction(g *gin.Context) {
	g.HTML(
		http.StatusOK,
		webPathRoot,
		gin.H{},
	)
	//g.Redirect(http.StatusTemporaryRedirect, "https://github.com/watermint/toolbox")
}

func (z *WebHandler) connectStart(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		//cmd := g.Query("command")
		tokenType := g.Query("tokenType")
		redirectUrl := z.BaseUrl + webPathConnectAuth
		_, url, err := z.auth(user, uc).New(tokenType, redirectUrl)
		if err != nil {
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}
		g.Redirect(http.StatusTemporaryRedirect, url)
	})
}

func (z *WebHandler) connectAuth(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		state := g.Query("state")
		code := g.Query("code")

		_, _, err := z.auth(user, uc).Auth(state, code)
		if err != nil {
			g.Redirect(http.StatusTemporaryRedirect, webPathAuthFailed)
		} else {
			g.Redirect(http.StatusTemporaryRedirect, webPathConnectFinish)
		}
	})
}

func (z *WebHandler) connectFinish(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		g.HTML(
			http.StatusOK,
			webPathConnectFinish,
			gin.H{},
		)
	})
}

func (z *WebHandler) Home(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		cmd := g.Param("command")
		grp, rcp, err := z.findRecipe(cmd)

		switch {
		case err != nil:
			g.Redirect(http.StatusTemporaryRedirect, webPathCommandNotFound)

		case rcp != nil:
			// TODO: Breadcrumb list
			z.renderRecipeConn(g, cmd, rcp, user, uc)

		case grp != nil:
			// TODO: Breadcrumb list
			z.renderCatalogue(g, cmd, grp)
		}
	})
}

func (z *WebHandler) Run(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		cmd := g.Param("command")
		l := z.control.Log().With(zap.String("cmd", cmd))
		_, rcp, err := z.findRecipe(cmd)
		if rcp == nil || err != nil {
			l.Debug("Invalid run request", zap.String("Cmd", cmd))
			g.Redirect(http.StatusTemporaryRedirect, webPathCommandNotFound)
			return
		}
		selectedConns := g.PostFormMap("Conn")

		var vo interface{} = rcp.Requirement()
		vc := app_vo_impl.NewValueContainer(vo)

		for k, v := range vc.Values {
			if d, ok := v.(bool); ok {
				vc.Values[k] = d
			} else if _, ok := v.(app_conn.ConnBusinessInfo); ok {
				if pn, ok := selectedConns[k]; ok {
					vc.Values[k] = &app_conn_impl.ConnBusinessInfo{
						PeerName: pn,
					}
				} else {
					l.Debug("Unable to find required conn", zap.String("key", k))
				}
			} else if _, ok := v.(app_conn.ConnBusinessFile); ok {
				if pn, ok := selectedConns[k]; ok {
					vc.Values[k] = &app_conn_impl.ConnBusinessFile{
						PeerName: pn,
					}
				} else {
					l.Debug("Unable to find required conn", zap.String("key", k))
				}
			} else if _, ok := v.(app_conn.ConnBusinessAudit); ok {
				if pn, ok := selectedConns[k]; ok {
					vc.Values[k] = &app_conn_impl.ConnBusinessAudit{
						PeerName: pn,
					}
				} else {
					l.Debug("Unable to find required conn", zap.String("key", k))
				}
			} else if _, ok := v.(app_conn.ConnBusinessMgmt); ok {
				if pn, ok := selectedConns[k]; ok {
					vc.Values[k] = &app_conn_impl.ConnBusinessMgmt{
						PeerName: pn,
					}
				} else {
					l.Debug("Unable to find required conn", zap.String("key", k))
				}
			} else if _, ok := v.(app_conn.ConnUserFile); ok {
				if pn, ok := selectedConns[k]; ok {
					vc.Values[k] = &app_conn_impl.ConnUserFile{
						PeerName: pn,
					}
				} else {
					l.Debug("Unable to find required conn", zap.String("key", k))
				}
			}
		}

		vc.Apply(vo)

		wuiLogPath := filepath.Join(uc.Workspace().Log(), "webui.log")
		l.Debug("Create web ui log file", zap.String("path", wuiLogPath))
		wuiLog, err := os.Create(wuiLogPath)
		if err != nil {
			l.Debug("Unable to create web ui log file", zap.String("path", wuiLogPath), zap.Error(err))
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}

		linkForLocalFile := func(path string) string {
			rel, err := filepath.Rel(uc.Workspace().Job(), path)
			if err != nil {
				l.Warn("Unable to calc rel path", zap.Error(err))
				return ""
			}
			p := base64.URLEncoding.EncodeToString([]byte(rel))
			return fmt.Sprintf("%s/%s/%s/%s", webPathJobArtifact, cmd, uc.Workspace().JobId(), p)
		}

		wui := app_ui.NewWeb(uc.UI(), wuiLog, linkForLocalFile)
		if muc, ok := uc.(*app_control_impl.Multi); !ok {
			l.Debug("Control was not expected impl.")
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		} else {
			juc := muc.WithNewUI(wui)
			wj := &webJobRun{
				name:      cmd,
				jobId:     juc.Workspace().JobId(),
				recipe:    rcp,
				vo:        vo.(app_vo.ValueObject),
				uc:        juc,
				uiLogFile: wuiLog,
			}
			l.Debug("Enqueue Job", zap.String("name", cmd), zap.String("jobId", wj.jobId))
			z.jobChan <- wj

			g.Redirect(
				http.StatusTemporaryRedirect,
				webPathJob+"/"+cmd+"/"+wj.jobId,
			)
		}

	})
}

func (z *WebHandler) Job(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		cmd := g.Param("command")
		jobId := g.Param("jobId")
		l := z.control.Log().With(zap.String("jobId", jobId), zap.String("cmd", cmd))

		jobPath := filepath.Join(user.Workspace().UserHome(), "jobs", jobId)
		logPath := filepath.Join(jobPath, "logs", "webui.log")
		logFile, err := os.Open(logPath)
		if err != nil {
			l.Debug("Unable to open file", zap.Error(err), zap.String("logPath", logPath))
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}
		defer logFile.Close()

		logs := make([]*app_ui.WebUILog, 0)

		s := bufio.NewScanner(logFile)
		for s.Scan() {
			line := s.Bytes()
			wl := &app_ui.WebUILog{}
			if err = json.Unmarshal(line, wl); err != nil {
				l.Warn("Unable to unmarshal line, skip", zap.Error(err), zap.String("line", s.Text()))
			} else {
				logs = append(logs, wl)
			}
		}
		if s.Err() != nil {
			l.Warn("There is an error on read log", zap.Error(err))
		}

		g.HTML(
			http.StatusOK,
			webPathJob,
			gin.H{
				"Recipe": cmd,
				"Logs":   logs,
			},
		)
	})
}

func (z *WebHandler) Artifact(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User, uc app_control.Control) {
		jobId := g.Param("jobId")
		artifactName := g.Param("artifactName")

		l := z.control.Log().With(zap.String("jobId", jobId), zap.String("artifactName", artifactName))

		rel, err := base64.URLEncoding.DecodeString(artifactName)
		if err != nil {
			l.Warn("Unable to decode artifact name", zap.Error(err))
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}
		relPath := string(rel)

		jobPath := filepath.Join(uc.Workspace().Home(), "jobs", jobId)
		path := filepath.Join(jobPath, relPath)
		l.Debug("Artifact path", zap.String("path", path), zap.String("jobPath", jobPath))
		if !strings.HasPrefix(filepath.Clean(path), jobPath) {
			l.Warn("Look like malicious path", zap.String("path", filepath.Clean(path)), zap.String("jobPath", jobPath))
			g.Data(http.StatusNoContent, "application/octet-stream", []byte{})
			return
		}

		contentType := "application/octet-stream"
		switch strings.ToLower(filepath.Ext(relPath)) {
		case ".xlsx":
			contentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
		case ".csv":
			contentType = "text/csv"
		case ".json":
			contentType = "application/json"
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			l.Warn("Unable to load binary", zap.Error(err))
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}

		fileName := filepath.Base(relPath)

		g.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		g.Data(http.StatusOK, contentType, data)
	})
}

func (z *WebHandler) renderRecipeConn(g *gin.Context, cmd string, rcp app_recipe.Recipe, user app_user.User, uc app_control.Control) {
	l := z.control.Log().With(zap.String("cmd", cmd))
	reqConns, reqParams, _ := z.recipeRequirements(rcp)
	selectedConns := g.PostFormMap("Conn")
	orderedConnName := make([]string, 0)
	var nextConnName, nextConnType = "", ""

	for k := range reqConns {
		orderedConnName = append(orderedConnName, k)
	}
	sort.Strings(orderedConnName)

	for _, k := range orderedConnName {
		if v, ok := selectedConns[k]; !ok || v == "" {
			nextConnName = k
			nextConnType = reqConns[k]
		}
	}

	if nextConnName == "" { // no more required conns
		if len(reqParams) > 0 {
			// TODO forward to param rendering
			z.renderRecipeParam(g)
		} else {
			// TODO forward to confirm & run
			z.renderRecipeRun(g, cmd, rcp, user, uc)
		}
		return
	}

	existingConns, err := z.auth(user, uc).List(nextConnType)
	if err != nil {
		l.Debug("Unable to list connections", zap.Error(err))
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
		return
	}
	listConns := make([]string, 0)
	connDesc := make(map[string]string)
	connSuppl := make(map[string]string)

	for _, e := range existingConns {
		listConns = append(listConns, e.PeerName)
		connDesc[e.PeerName] = e.Description
		connSuppl[e.PeerName] = e.Supplemental
	}
	sort.Strings(listConns)

	g.HTML(
		http.StatusOK,
		"home-recipe-conn",
		gin.H{
			"Recipe":                cmd,
			"ExistingConns":         listConns,
			"ExistingConnDesc":      connDesc,
			"ExistingConnSuppl":     connSuppl,
			"SelectedConns":         selectedConns,
			"CurrentConn":           nextConnName,
			"CurrentConnType":       nextConnType,
			"CurrentConnTypeHeader": z.control.UI().Text("web.conn." + nextConnType + ".header"),
			"CurrentConnTypeDetail": z.control.UI().Text("web.conn." + nextConnType + ".detail"),
		},
	)
}

func (z *WebHandler) renderRecipeParam(g *gin.Context) {

	g.HTML(
		http.StatusOK,
		"home-recipe-param",
		gin.H{
			"Recipe": "sharedfolder-list",
		},
	)
}

func (z *WebHandler) renderRecipeRun(g *gin.Context, cmd string, rcp app_recipe.Recipe, user app_user.User, uc app_control.Control) {
	l := z.control.Log().With(zap.String("cmd", cmd))
	reqConns, _, _ := z.recipeRequirements(rcp)

	selectedConns := g.PostFormMap("Conn")
	listConns := make([]string, 0)
	connDesc := make(map[string]string)
	connSuppl := make(map[string]string)

	for connName, tokenType := range reqConns {
		listConns = append(listConns, connName)
		conns, err := z.auth(user, uc).List(tokenType)
		if err != nil {
			l.Debug("Unable to list connections", zap.Error(err))
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}
		for _, c := range conns {
			if c.PeerName == selectedConns[connName] {
				connDesc[connName] = c.Description
				connSuppl[connName] = c.Supplemental
				break
			}
		}
	}
	sort.Strings(listConns)

	g.HTML(
		http.StatusOK,
		"home-recipe-run",
		gin.H{
			"Recipe":        cmd,
			"Conns":         listConns,
			"ConnsSelected": selectedConns,
			"ConnDesc":      connDesc,
			"ConnSuppl":     connSuppl,
		},
	)
}

func (z *WebHandler) renderCatalogue(g *gin.Context, cmd string, grp *app_recipe_group.Group) {
	cmds := make([]string, 0)
	dict := make(map[string]gin.H)
	jobs := make([]gin.H, 0)

	for _, g := range grp.SubGroups {
		if g.IsSecret() {
			continue
		}

		path := make([]string, 0)
		path = append(path, grp.Path...)
		path = append(path, g.Name)

		dict[g.Name] = gin.H{
			"Title":       g.Name,
			"Description": z.control.UI().Text(grp.CommandDesc(g.Name).Key()),
			"Uri":         webPathHome + "/" + strings.Join(path, "-"),
		}
	}
	for name := range grp.Recipes {
		path := make([]string, 0)
		path = append(path, grp.Path...)
		path = append(path, name)

		dict[name] = gin.H{
			"Title":       name,
			"Description": z.control.UI().Text(grp.CommandDesc(name).Key()),
			"Uri":         webPathHome + "/" + strings.Join(path, "-"),
		}
	}

	for k := range dict {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)
	for _, c := range cmds {
		jobs = append(jobs, dict[c])
	}

	g.HTML(
		http.StatusOK,
		"home-catalogue",
		gin.H{
			"Detail": cmd,
			"Jobs":   jobs,
		},
	)
}

func (z *WebHandler) withUser(g *gin.Context, f func(g *gin.Context, user app_user.User, uc app_control.Control)) {
	l := z.control.Log()

	hash, err := g.Cookie(webUserHashCookie)
	if err != nil {
		l.Debug("No cookie access")
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
		return
	}
	l.With(zap.String("UserHash", hash))
	repo, err := app_user.SingleUserRepository(z.control.Workspace())
	if err != nil {
		l.Debug("Unable to prepare user repo", zap.Error(err))
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
		return
	}
	user, err := repo.Resolve(hash)
	if err != nil {
		l.Debug("Unable to resolve user by hash", zap.Error(err))
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
		return
	}
	uc, err := z.control.(app_control_launcher.ControlLauncher).NewControl(user.Workspace())
	if err != nil {
		l.Debug("Unable to create new control for the user", zap.Error(err))
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
		return
	}

	f(g, user, uc)
}
