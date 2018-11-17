package oper_cli

import (
	"bufio"
	"fmt"
	"github.com/watermint/toolbox/poc/oper/oper_msg"
	"io"
	"strings"
)

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

func (z *CUI) Tell(msg oper_msg.UIMessage) {
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellError(msg oper_msg.UIMessage) {
	fmt.Fprint(z.Out, "ERR: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellDone(msg oper_msg.UIMessage) {
	fmt.Fprint(z.Out, "DONE: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellSuccess(msg oper_msg.UIMessage) {
	fmt.Fprint(z.Out, "SUCCESS: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellFailure(msg oper_msg.UIMessage) {
	fmt.Fprint(z.Out, "FAILURE: ")
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) TellProgress(msg oper_msg.UIMessage) {
	fmt.Fprintln(z.Out, msg.Text())
}

func (z *CUI) AskRetry(msg oper_msg.UIMessage) bool {
	fmt.Fprintln(z.Out, msg.Text())
	fmt.Fprintln(z.Out, "Retry? (y/n)")
	return z.YesNo()
}

func (z *CUI) AskWarn(msg oper_msg.UIMessage) bool {
	fmt.Fprint(z.Out, "WARN: ")
	fmt.Fprintln(z.Out, msg.Text())
	fmt.Fprintln(z.Out, "Continue? (y/n)")
	return z.YesNo()
}

func (z *CUI) AskText(msg oper_msg.UIMessage) string {
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
