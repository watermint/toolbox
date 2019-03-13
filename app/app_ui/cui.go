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

func (z *CUI) YesNo() bool {
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
		fmt.Fprintln(z.Out, "Retry? (y/n)")
	}
}

func (z *CUI) Tell(msg UIMessage) {
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) TellError(msg UIMessage) {
	fmt.Fprint(z.Out, "ERR: ")
	fmt.Fprintln(z.Out, msg.T())
}

func (z *CUI) TellDone(msg UIMessage) {
	fmt.Fprint(z.Out, "DONE: ")
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
	fmt.Fprintln(z.Out, msg.T())
	fmt.Fprintln(z.Out, "Retry? (y/n)")
	return z.YesNo()
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
