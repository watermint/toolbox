package app_ui

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewDefaultCUI() UI {
	return &CUI{
		Out: os.Stdout,
		In:  os.Stdin,
	}
}

type CUI struct {
	Out io.Writer
	In  io.Reader
}

func (z *CUI) DebugMode(debug bool) {
}

func (z *CUI) YesNo(msg string) bool {
	br := bufio.NewReader(z.In)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			return false
		}
		ans := strings.ToLower(strings.TrimSpace(string(line)))
		switch ans {
		case "y":
			return true
		case "yes":
			return true
		case "n":
			return false
		case "no":
			return false
		}

		// ask again
		fmt.Fprintf(z.Out, msg)
	}
}

func (z *CUI) Tell(msg UIMessage) {
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) TellError(msg UIMessage) {
	fmt.Fprint(z.Out, "ERR: ")
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) TellSuccess(msg UIMessage) {
	fmt.Fprint(z.Out, "SUCCESS: ")
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) TellFailure(msg UIMessage) {
	fmt.Fprint(z.Out, "FAILURE: ")
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) AskRetry(msg UIMessage) bool {
	m := "Retry? (y/n)"
	fmt.Fprintln(z.Out, msg.T())
	fmt.Fprintln(z.Out, m)
	return z.YesNo(m)
}

func (z *CUI) AskText(msg UIMessage) string {
	fmt.Fprintln(z.Out, msg.T())
	br := bufio.NewReader(z.In)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			return ""
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			return text
		}

		// ask again
		fmt.Fprintln(z.Out, msg.T())
	}
}

func (z *CUI) AskConfirm(msg UIMessage) bool {
	m := "? (y/n)"
	fmt.Fprintln(z.Out, msg.T())
	fmt.Fprintln(z.Out, m)
	return z.YesNo(m)
}
