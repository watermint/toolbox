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

func (z *Web) AskCont(m app_msg.Message) (cont bool, cancel bool) {
	return false, true
}

func (z *Web) AskText(m app_msg.Message) (text string, cancel bool) {
	return "", true
}

func (z *Web) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	return "", true
}

func (z *Web) Header(m app_msg.Message) {
	z.HeaderK(m.Key(), m.Params()...)
}

func (z *Web) Text(m app_msg.Message) string {
	return z.TextK(m.Key(), m.Params()...)
}

func (z *Web) TextOrEmpty(m app_msg.Message) string {
	return z.TextOrEmptyK(m.Key(), m.Params()...)
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

func (z *Web) Info(m app_msg.Message) {
	z.InfoK(m.Key(), m.Params()...)
}

func (z *Web) Error(m app_msg.Message) {
	z.ErrorK(m.Key(), m.Params()...)
}

func (z *Web) Success(m app_msg.Message) {
	z.uiLog(&WebUILog{
		Tag:     WebTagResultSuccess,
		Message: z.Text(m),
	})
}

func (z *Web) SuccessK(key string, p ...app_msg.P) {
	z.Success(app_msg.M(key, p...))
}

func (z *Web) Failure(m app_msg.Message) {
	z.uiLog(&WebUILog{
		Tag:     WebTagResultFailure,
		Message: z.Text(m),
	})
}

func (z *Web) FailureK(key string, p ...app_msg.P) {
	z.Failure(app_msg.M(key, p...))
}

func (z *Web) HeaderK(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagHeader,
		Message: z.TextK(key, p...),
	})
}

func (z *Web) InfoK(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagInfo,
		Message: z.TextK(key, p...),
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

func (z *Web) ErrorK(key string, p ...app_msg.P) {
	z.uiLog(&WebUILog{
		Tag:     WebTagError,
		Message: z.TextK(key, p...),
	})
}

func (z *Web) Break() {
	z.uiLog(&WebUILog{
		Tag: WebTagBreak,
	})
}

func (z *Web) TextK(key string, p ...app_msg.P) string {
	return z.baseUI.TextK(key, p...)
}

func (z *Web) TextOrEmptyK(key string, p ...app_msg.P) string {
	return z.baseUI.TextOrEmptyK(key, p...)
}

func (z *Web) AskContK(key string, p ...app_msg.P) (cont bool, cancel bool) {
	panic("not supported")
}

func (z *Web) AskTextK(key string, p ...app_msg.P) (text string, cancel bool) {
	panic("not supported")
}

func (z *Web) AskSecureK(key string, p ...app_msg.P) (secure string, cancel bool) {
	panic("not supported")
}

func (z *Web) OpenArtifact(path string, autoOpen bool) {
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
		cols = append(cols, z.baseUI.TextK(c.Key(), c.Params()...))
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
		cols = append(cols, z.baseUI.TextK(c.Key(), c.Params()...))
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
