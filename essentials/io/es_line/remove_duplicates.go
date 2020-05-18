package es_line

import "io"

func NewRemoveRedundantLinesWriter(out io.Writer) io.Writer {
	return &RemoveRedundantLinesWriter{out: out}
}

type RemoveRedundantLinesWriter struct {
	out   io.Writer
	last0 byte
	last1 byte
}

func (z *RemoveRedundantLinesWriter) Write(p []byte) (n int, err error) {
	t := make([]byte, 0)
	for _, b := range p {
		if z.last0 == '\n' && z.last1 == '\n' && b == '\n' {
			continue
		}
		t = append(t, b)
		z.last0 = z.last1
		z.last1 = b
	}
	return z.out.Write(t)
}
