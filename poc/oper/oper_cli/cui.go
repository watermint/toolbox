package oper_cli

import (
	"bufio"
	"github.com/watermint/toolbox/poc/oper/oper_i18n"
	"github.com/watermint/toolbox/poc/oper/oper_ui"
	"golang.org/x/text/language"
	"io"
	"os"
)

type CUI struct {
	out  *bufio.Writer
	in   *bufio.Reader
	Out  io.Writer
	In   io.Reader
	Lang language.Tag
}

func (z *CUI) TellSuccess(msg oper_i18n.UIMessage) {
	z.out.WriteString("SUCCESS: ")
	z.Tell(msg)
}

func (z *CUI) TellFailure(msg oper_i18n.UIMessage) {
	z.out.WriteString("FAILURE: ")
	z.Tell(msg)
}

func (z *CUI) Init() {
	z.out = bufio.NewWriter(z.Out)
	z.in = bufio.NewReader(z.In)
}

func (z *CUI) Tell(msg oper_i18n.UIMessage) {
	z.out.WriteString(msg.Message(z.Lang))
	z.out.WriteString("\n")
}

func (z *CUI) TellDone(msg oper_i18n.UIMessage) {
	z.Tell(msg)
}

func (z *CUI) TellError(msg oper_i18n.UIMessage) {
	z.out.WriteString("ERR: ")
	z.out.WriteString(msg.Message(z.Lang))
	z.out.WriteString("\n")
}

func (z *CUI) TellTable(tbl oper_ui.UITable) {
	panic("implement me")
}

func (z *CUI) TellProgress(msg oper_i18n.UIMessage) {
	z.Tell(msg)
}

func (z *CUI) AskRetry(msg oper_i18n.UIMessage) bool {
	panic("implement me")
}

func (z *CUI) AskWarn(msg oper_i18n.UIMessage) bool {
	panic("implement me")
}

func (z *CUI) AskOptions(title oper_i18n.UIMessage, opts map[string]oper_i18n.UIMessage) string {
	panic("implement me")
}

func (z *CUI) AskInputFile(msg oper_i18n.UIMessage) *os.File {
	z.out.WriteString("Ask[File]")
	z.Tell(msg)
	z.out.WriteString("\n")

	for {
		z.out.WriteString("Filename:\n")
		pathByte, _, err := z.in.ReadLine()
		if err == io.EOF {
			return nil
		}
		path := string(pathByte)
		if s, err := os.Stat(path); err != nil {
			z.out.WriteString("ERR: " + err.Error() + "\n")
			continue
		} else if s.IsDir() {
			z.out.WriteString("ERR: " + path + " is a directory")
			continue
		}

		if f, err := os.Open(path); err != nil {
			z.out.WriteString("ERR: " + err.Error() + "\n")
			return nil
		} else {
			return f
		}
	}
}

func (z *CUI) AskOutputFile(msg oper_i18n.UIMessage, filename string, tmpFilePath string) {
	panic("implement me")
}

func (z *CUI) AskText(msg oper_i18n.UIMessage) string {
	panic("implement me")
}
