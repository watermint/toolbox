package app_ui

import (
	"encoding/json"
	"github.com/watermint/toolbox/experimental/app_msg"
	"github.com/watermint/toolbox/experimental/app_msg_container"
	"github.com/watermint/toolbox/experimental/app_root"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type WebUILog struct {
	Tag     string
	Message string
	Cols    []string
	Link    string
}

type LinkForLocalFile func(path string) string

type Web struct {
	mc  app_msg_container.Container
	out io.Writer
	llf LinkForLocalFile
}

const (
	WebTagHeader       = "header"
	WebTagInfo         = "info"
	WebTagTableStart   = "table_start"
	WebTagTableHeader  = "table_header"
	WebTagTableRow     = "table_row"
	WebTagTableFinish  = "table_finish"
	WebTagError        = "error"
	WebTagBreak        = "break"
	WebTagAskCont      = "ask_cont"
	WebTagAskText      = "ask_text"
	WebTagAskSecure    = "ask_secure"
	WebTagArtifactXlsx = "artifact_xslx"
	WebTagArtifactCsv  = "artifact_csv"
	WebTagArtifactJson = "artifact_json"
)

func (z *Web) uiLog(w *WebUILog) {
	l := app_root.Log()
	j, err := json.Marshal(w)
	if err != nil {
		l.Error("Unable to marshal JSON", zap.Error(err))
		return
	}
	_, err = z.out.Write(j)
	if err != nil {
		l.Warn("Unable to write UI message", zap.Error(err))
		return
	}
	_, err = z.out.Write([]byte("\n"))
	if err != nil {
		l.Warn("Unable to write UI message", zap.Error(err))
		return
	}
}

func (z *Web) Header(key string, p ...app_msg.Param) {
	z.uiLog(&WebUILog{
		Tag:     WebTagHeader,
		Message: z.Text(key, p...),
	})
}

func (z *Web) Info(key string, p ...app_msg.Param) {
	z.uiLog(&WebUILog{
		Tag:     WebTagInfo,
		Message: z.Text(key, p...),
	})
}

func (z *Web) InfoTable(border bool) Table {
	t := &WebTable{
		mc: z.mc,
		w:  z,
	}
	z.uiLog(&WebUILog{
		Tag: WebTagTableStart,
	})
	return t
}

func (z *Web) Error(key string, p ...app_msg.Param) {
	z.uiLog(&WebUILog{
		Tag:     WebTagError,
		Message: z.Text(key, p...),
	})
}

func (z *Web) Break() {
	z.uiLog(&WebUILog{
		Tag: WebTagBreak,
	})
}

func (z *Web) Text(key string, p ...app_msg.Param) string {
	return z.mc.Compile(app_msg.M(key, p...))
}

func (z *Web) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	panic("not supported")
}

func (z *Web) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	panic("not supported")
}

func (z *Web) AskSecure(key string, p ...app_msg.Param) (secure string, cancel bool) {
	panic("not supported")
}

func (z *Web) OpenArtifact(path string) {
	l := app_root.Log()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		l.Warn("Unable to read path", zap.Error(err), zap.String("path", path))
		return
	}

	for _, f := range files {
		e := filepath.Ext(f.Name())
		switch strings.ToLower(e) {
		case ".xlsx":
			z.uiLog(&WebUILog{
				Tag:  WebTagArtifactXlsx,
				Link: z.llf(filepath.Join(path, f.Name())),
			})

		case ".csv":
			z.uiLog(&WebUILog{
				Tag:  WebTagArtifactCsv,
				Link: z.llf(filepath.Join(path, f.Name())),
			})

		case ".json":
			z.uiLog(&WebUILog{
				Tag:  WebTagArtifactJson,
				Link: z.llf(filepath.Join(path, f.Name())),
			})

		default:
			l.Debug("Unsupported extension", zap.String("name", f.Name()))
		}
	}
}

func (z *Web) IsConsole() bool {
	return false
}

func (z *Web) IsWeb() bool {
	return true
}

type WebTable struct {
	mc app_msg_container.Container
	w  *Web
}

func (z *WebTable) Header(h ...app_msg.Message) {
	cols := make([]string, 0)
	for _, c := range h {
		cols = append(cols, z.mc.Compile(c))
	}
	z.w.uiLog(&WebUILog{
		Tag:  WebTagTableHeader,
		Cols: cols,
	})
}

func (z *WebTable) HeaderRaw(h ...string) {
	cols := make([]string, 0)
	for _, c := range h {
		cols = append(cols, c)
	}
	z.w.uiLog(&WebUILog{
		Tag:  WebTagTableHeader,
		Cols: cols,
	})
}

func (z *WebTable) Row(m ...app_msg.Message) {
	cols := make([]string, 0)
	for _, c := range m {
		cols = append(cols, z.mc.Compile(c))
	}
	z.w.uiLog(&WebUILog{
		Tag:  WebTagTableRow,
		Cols: cols,
	})
}

func (z *WebTable) RowRaw(m ...string) {
	cols := make([]string, 0)
	for _, c := range m {
		cols = append(cols, c)
	}
	z.w.uiLog(&WebUILog{
		Tag:  WebTagTableRow,
		Cols: cols,
	})
}

func (z *WebTable) Flush() {
	z.w.uiLog(&WebUILog{
		Tag: WebTagTableFinish,
	})
}
