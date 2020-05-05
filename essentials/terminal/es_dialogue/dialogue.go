package es_dialogue

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"strings"
)

type Dialogue interface {
	// Ask weak confirmation to proceed the operation.
	// Prompt message could be like "hit any key to proceed".
	AskProceed(p Prompt)

	// Ask sticky confirmation to proceed the operation.
	AskCont(p Prompt, v VerifyCont) (c bool)

	// Ask text
	AskText(p Prompt, v VerifyText) (t string, cancel bool)

	// Ask secure text (text should not echo back)
	// Always read from os.Stdin
	AskSecure(p Prompt) (t string, cancel bool)
}

type Prompt func()

// Return result in cont.
// Return valid = false when given t is invalid.
// The value cont is not used when valid == false.
type VerifyCont func(t string) (cont, valid bool)

func YesNoCont(t string) (cont, valid bool) {
	a := strings.ToLower(strings.TrimSpace(t))
	switch a {
	case "y", "yes":
		return true, true
	case "n", "no":
		return false, true
	default:
		return false, false
	}
}

// Ensure the text is valid format or not.
type VerifyText func(t string) (s string, valid bool)

func NonEmptyText(t string) (s string, valid bool) {
	s = strings.TrimSpace(t)
	return s, s != ""
}

func New(w io.Writer) Dialogue {
	return &dlgImpl{
		in: os.Stdin,
		wr: w,
	}
}

type dlgImpl struct {
	in *os.File
	wr io.Writer
}

// This func does not rely on es_terminfo.IsInTerminal() for testing purpose.
func (z dlgImpl) isInTerminal() bool {
	return terminal.IsTerminal(int(z.in.Fd()))
}

func (z dlgImpl) AskProceed(p Prompt) {
	// Ignore when os.Stdin is pipe
	if !z.isInTerminal() {
		return
	}
	p()
	_, _, _ = keyboard.GetSingleKey()
}

func (z dlgImpl) AskCont(p Prompt, v VerifyCont) (c bool) {
	l := es_log.Default()
	r := bufio.NewReader(z.in)
	for {
		// prompt
		p()

		// read
		ln, _, err := r.ReadLine()
		ls := string(ln)
		if co, va := v(ls); va {
			return co
		}
		// return on error including io.EOF
		if err != nil {
			l.Debug("Error or cancelled", es_log.Error(err))
			return false
		}
	}
}

func (z dlgImpl) AskText(p Prompt, v VerifyText) (t string, cancel bool) {
	l := es_log.Default()
	r := bufio.NewReader(z.in)
	for {
		// prompt
		p()

		// read
		ln, _, err := r.ReadLine()
		ls := string(ln)
		if vt, va := v(ls); va {
			return vt, false
		}

		// return on error including io.EOF
		if err != nil {
			l.Debug("Error or cancelled", es_log.Error(err))
			return "", true
		}
	}
}

func (z dlgImpl) AskSecure(p Prompt) (t string, cancel bool) {
	// Always cancel when stdin is pipe
	if !z.isInTerminal() {
		return "", true
	}

	l := es_log.Default()
	for {
		// prompt
		p()

		// read password
		s, err := terminal.ReadPassword(int(z.in.Fd()))
		if err != nil {
			l.Debug("Unable to read password", es_log.Error(err))
			return "", true
		}
		_, _ = fmt.Fprintln(z.wr)
		return string(s), false
	}
}

func DenyAll() Dialogue {
	return &denyImpl{}
}

type denyImpl struct {
}

func (z denyImpl) AskProceed(p Prompt) {
}

func (z denyImpl) AskCont(p Prompt, v VerifyCont) (c bool) {
	return false
}

func (z denyImpl) AskText(p Prompt, v VerifyText) (t string, cancel bool) {
	return "", true
}

func (z denyImpl) AskSecure(p Prompt) (t string, cancel bool) {
	return "", true
}
