package app_template_impl

import (
	"bytes"
	"github.com/gin-gonic/gin/render"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_template"
	"html/template"
	"io/ioutil"
	"net/http"
)

func NewDev(fs http.FileSystem, ctl app_control.Control) app_template.Template {
	return &DevFileSystem{
		resNames: make(map[string][]string),
		fs:       fs,
		ctl:      ctl,
	}
}

// Dynamic loading
type DevFileSystem struct {
	resNames map[string][]string
	fs       http.FileSystem
	ctl      app_control.Control
}

func (z *DevFileSystem) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: z.tmpl(name),
		Data:     data,
	}
}

func (z *DevFileSystem) Define(name string, resNames ...string) error {
	l := z.ctl.Log().With(esl.String("name", name))
	l.Debug("Loading", esl.Strings("resources", resNames))
	for _, r := range resNames {
		f, err := z.fs.Open(r)
		if err != nil {
			l.Error("Unable to open resource", esl.Error(err))
			return err
		}
		f.Close()
	}
	z.resNames[name] = resNames
	return nil
}

func (z *DevFileSystem) tmpl(name string) *template.Template {
	l := z.ctl.Log().With(esl.String("name", name))
	t := template.New("render")
	rs, ok := z.resNames[name]
	if !ok {
		l.Error("Resource unavailable")
		return nil
	}
	for _, r := range rs {
		ll := l.With(esl.String("resource", r))
		f, err := z.fs.Open(r)
		if err != nil {
			ll.Error("Unable to open resource", esl.Error(err))
			return nil
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			ll.Error("Unable to read resource", esl.Error(err))
			return nil
		}
		f.Close()
		t, err = t.Parse(string(b))
		if err != nil {
			ll.Error("Unable to parse template", esl.Error(err))
			return nil
		}
	}
	return t
}

func (z *DevFileSystem) Render(name string, d ...app_template.D) string {
	l := z.ctl.Log().With(esl.String("name", name))
	t := z.tmpl(name)
	if t == nil {
		return ""
	}

	data := make(map[string]interface{})
	for _, di := range d {
		for k, v := range di {
			data[k] = v
		}
	}

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	if err != nil {
		l.Error("Unable to execute template", esl.Error(err))
		return ""
	}

	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	h, err := m.String("text/html", doc.String())
	if err != nil {
		l.Warn("Unable to minify result", esl.Error(err))
		return doc.String()
	}
	return h
}
