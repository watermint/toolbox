package app_ui

import (
	"bufio"
	"fmt"
	"github.com/watermint/toolbox/app/app_msg"
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

func (z *CUI) Tell(msg app_msg.UIMessage) {
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellError(msg app_msg.UIMessage) {
	fmt.Fprint(z.Out, "ERR: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellDone(msg app_msg.UIMessage) {
	fmt.Fprint(z.Out, "DONE: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellSuccess(msg app_msg.UIMessage) {
	fmt.Fprint(z.Out, "SUCCESS: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellFailure(msg app_msg.UIMessage) {
	fmt.Fprint(z.Out, "FAILURE: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) AskRetry(msg app_msg.UIMessage) bool {
	fmt.Fprintln(z.Out, msg.Text())
	fmt.Fprintln(z.Out, "Retry? (y/n)")
	return z.YesNo()
}

func (z *CUI) AskText(msg app_msg.UIMessage) string {
	fmt.Fprintln(z.Out, msg.Text())
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
		fmt.Fprintln(z.Out, msg.Text())
	}
}
