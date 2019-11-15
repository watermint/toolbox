package app_ui

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

type WebUILog struct {
	Tag       string
	Message   string
	TableId   string
	TableCols []string
	Link      string
	LinkLabel string
}

type LinkForLocalFile func(path string) string

func NewWeb(baseUI UI, out io.Writer, llf LinkForLocalFile) UI {
	return &Web{
		baseUI: baseUI,
		out:    out,
		llf:    llf,
	}
}

type Web struct {
	baseUI UI
	out    io.Writer
	llf    LinkForLocalFile
	mutex  sync.Mutex
}

const (
	WebTagHeader         = "header"
	WebTagInfo           = "info"
	WebTagTableStart     = "table_start"
	WebTagTableHeader    = "table_header"
	WebTagTableRow       = "table_row"
	WebTagTableFinish    = "table_finish"
	WebTagError          = "error"
	WebTagBreak          = "break"
	WebTagAskCont        = "ask_cont"
	WebTagAskText        = "ask_text"
	WebTagAskSecure      = "ask_secure"
	WebTagArtifactHeader = "artifact_header"
	WebTagArtifactXlsx   = "artifact_xlsx"
	WebTagArtifactCsv    = "artifact_csv"
	WebTagArtifactJson   = "artifact_json"
	WebTagResultSuccess  = "result_success"
	WebTagResultFailure  = "result_failure"
	WebTagRefresh        = "refresh"
)

func (z *Web) uiLog(w *WebUILog) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

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

func (z *Web) InfoM(m app_msg.Message) {
	z.Info(m.Key(), m.Params()...)
}

func (z *Web) ErrorM(m app_msg.Message) {
	z.Error(m.Key(), m.Params()...)
}

func (z *Web) Success(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagResultSuccess,
		Message: z.Text(key, p...),
	})
}

func (z *Web) Failure(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagResultFailure,
		Message: z.Text(key, p...),
	})
}

func (z *Web) Header(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagHeader,
		Message: z.Text(key, p...),
	})
}

func (z *Web) Info(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagInfo,
		Message: z.Text(key, p...),
	})
}

func (z *Web) InfoTable(name string) Table {
	t := &WebTable{
		baseUI:    z.baseUI,
		tableName: name,
		w:         z,
	}
	return t
}

func (z *Web) Error(key string, p ...app_msg.P) {
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

func (z *Web) Text(key string, p ...app_msg.P) string {
	return z.baseUI.Text(key, p...)
}

func (z *Web) TextOrEmpty(key string, p ...app_msg.P) string {
	return z.baseUI.TextOrEmpty(key, p...)
}

func (z *Web) AskCont(key string, p ...app_msg.P) (cont bool, cancel bool) {
	panic("not supported")
}

func (z *Web) AskText(key string, p ...app_msg.P) (text string, cancel bool) {
	panic("not supported")
}

func (z *Web) AskSecure(key string, p ...app_msg.P) (secure string, cancel bool) {
	panic("not supported")
}

func (z *Web) OpenArtifact(path string) {
	l := app_root.Log()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		l.Warn("Unable to read path", zap.Error(err), zap.String("path", path))
		return
	}
	z.uiLog(&WebUILog{
		Tag: WebTagArtifactHeader,
	})

	for _, f := range files {
		e := filepath.Ext(f.Name())
		switch strings.ToLower(e) {
		case ".xlsx":
			z.uiLog(&WebUILog{
				Tag:       WebTagArtifactXlsx,
				Link:      z.llf(filepath.Join(path, f.Name())),
				LinkLabel: f.Name(),
			})

		case ".csv":
			z.uiLog(&WebUILog{
				Tag:       WebTagArtifactCsv,
				Link:      z.llf(filepath.Join(path, f.Name())),
				LinkLabel: f.Name(),
			})

		case ".json":
			z.uiLog(&WebUILog{
				Tag:       WebTagArtifactJson,
				Link:      z.llf(filepath.Join(path, f.Name())),
				LinkLabel: f.Name(),
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
	baseUI    UI
	tableName string
	w         *Web
}

func (z *WebTable) Header(h ...app_msg.Message) {
	cols := make([]string, 0)
	for _, c := range h {
		cols = append(cols, z.baseUI.Text(c.Key(), c.Params()...))
	}
	z.w.uiLog(&WebUILog{
		Tag:       WebTagTableHeader,
		TableId:   z.tableName,
		TableCols: cols,
	})
}

func (z *WebTable) HeaderRaw(h ...string) {
	cols := make([]string, 0)
	for _, c := range h {
		cols = append(cols, c)
	}
	z.w.uiLog(&WebUILog{
		Tag:       WebTagTableHeader,
		TableId:   z.tableName,
		TableCols: cols,
	})
}

func (z *WebTable) Row(m ...app_msg.Message) {
	cols := make([]string, 0)
	for _, c := range m {
		cols = append(cols, z.baseUI.Text(c.Key(), c.Params()...))
	}
	z.w.uiLog(&WebUILog{
		Tag:       WebTagTableRow,
		TableId:   z.tableName,
		TableCols: cols,
	})
}

func (z *WebTable) RowRaw(m ...string) {
	cols := make([]string, 0)
	for _, c := range m {
		cols = append(cols, c)
	}
	z.w.uiLog(&WebUILog{
		Tag:       WebTagTableRow,
		TableId:   z.tableName,
		TableCols: cols,
	})
}

func (z *WebTable) Flush() {
}
