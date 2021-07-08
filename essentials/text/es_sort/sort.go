package es_sort

import (
	"bufio"
	"compress/zlib"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
)

const (
	// Default sort buffer memory limit: 50MiB
	defaultMemoryLimit = 50 * 1048576
)

var (
	ErrorAlreadyClosed = errors.New("the sorter resources already flushed and closed")
)

// Sorter Sort large text data
type Sorter interface {
	// WriteLine Write single line. The function will split into multiple lines if a line contains '\n'.
	WriteLine(line string) error

	// Flush the result into destination stream, then close all resources.
	Flush() error
}

// Compare returns an integer comparing two lines.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
type Compare func(x, y string) int

type sorterOpts struct {
	// Path to the temporary folder
	TempFolder string

	// Compress temporary files
	TempCompress bool

	// Target buffer memory size
	MemoryLimit int

	// Comparator
	Compare Compare

	// Remove duplicated lines
	Uniq bool

	// Logger
	Logger esl.Logger

	// Descending
	Desc bool
}

func (z sorterOpts) Apply(opts []SorterOpt) sorterOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type SorterOpt func(opts sorterOpts) sorterOpts

func Uniq(enabled bool) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.Uniq = enabled
		return opts
	}
}

// TempFolder Specify temporary folder. It's caller's responsibility to remove the temp folder if specified.
func TempFolder(path string) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.TempFolder = path
		return opts
	}
}
func TempCompress(enabled bool) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.TempCompress = enabled
		return opts
	}
}
func MemoryLimit(limit int) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.MemoryLimit = limit
		return opts
	}
}
func Comparator(comparator Compare) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.Compare = comparator
		return opts
	}
}
func Logger(logger esl.Logger) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.Logger = logger
		return opts
	}
}
func Desc(enabled bool) SorterOpt {
	return func(opts sorterOpts) sorterOpts {
		opts.Desc = enabled
		return opts
	}
}

func defaultOpts() sorterOpts {
	return sorterOpts{
		TempFolder:   "",
		TempCompress: false,
		MemoryLimit:  defaultMemoryLimit,
		Compare:      strings.Compare,
		Uniq:         false,
		Logger:       nil,
	}
}

func New(destStream io.WriteCloser, opts ...SorterOpt) Sorter {
	return &sorterImpl{
		buf:       nil,
		bufSize:   0,
		bufMutex:  sync.Mutex{},
		bufFile:   "",
		dest:      destStream,
		opts:      defaultOpts().Apply(opts),
		isFlushed: false,
	}
}

type sorterImpl struct {
	buf       []string
	bufSize   int
	bufMutex  sync.Mutex
	bufFile   string
	dest      io.WriteCloser
	opts      sorterOpts
	isFlushed bool
}

func (z *sorterImpl) log() esl.Logger {
	if z.opts.Logger != nil {
		return z.opts.Logger
	} else {
		return esl.Default()
	}
}

func (z *sorterImpl) createBufFileCompress(f func(w io.Writer) error) (path string, err error) {
	l := z.log()
	return z.createBufFilePlain(func(w io.Writer) error {
		zw, err := zlib.NewWriterLevel(w, zlib.BestCompression)
		if err != nil {
			l.Debug("Unable to create zlib writer", esl.Error(err))
			return err
		}

		defer func() {
			_ = zw.Flush()
			_ = zw.Close()
		}()

		return f(zw)
	})
}

func (z *sorterImpl) createBufFilePlain(f func(w io.Writer) error) (path string, err error) {
	l := z.log()
	bf, err := ioutil.TempFile(z.opts.TempFolder, "sort")
	if err != nil {
		l.Debug("Unable to create temp file", esl.Error(err))
		return "", err
	}
	path = bf.Name()
	l.Debug("Create buf file", esl.String("path", path))
	defer func() {
		_ = bf.Close()
	}()
	return path, f(bf)
}

func (z *sorterImpl) createBufFile(f func(w io.Writer) error) (path string, err error) {
	if z.opts.TempCompress {
		return z.createBufFileCompress(f)
	} else {
		return z.createBufFilePlain(f)
	}
}

func (z *sorterImpl) readBufFileWithWrapper(path string, wrapper func(r io.Reader) (rc io.ReadCloser, shouldClose bool, err error), f func(line string) error) error {
	l := z.log().With(esl.String("path", path))
	l.Debug("read plain buffer file")
	bf, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open the file", esl.Error(err))
		return err
	}

	defer func() {
		_ = bf.Close()
		rmErr := os.Remove(path)
		if rmErr != nil {
			l.Debug("Unable to remove the buf file", esl.Error(rmErr))
		}
	}()

	wr, sc, wErr := wrapper(bf)
	if wErr != nil {
		l.Debug("Unable to prepare wrapped reader", esl.Error(wErr))
		return wErr
	}
	defer func() {
		if sc {
			_ = wr.Close()
		}
	}()

	bs := bufio.NewScanner(wr)
	bs.Split(bufio.ScanLines)
	var readLines int64
	for bs.Scan() {
		readLines++
		fErr := f(bs.Text())
		if fErr != nil {
			l.Debug("Abort read due to an error", esl.Error(fErr))
			return fErr
		}
	}
	if scErr := bs.Err(); scErr != nil {
		l.Debug("Scanner error", esl.Error(scErr))
		return scErr
	}
	l.Debug("Read", esl.Int64("readLines", readLines))
	return nil
}

func (z *sorterImpl) readBufFileCompressed(path string, f func(line string) error) error {
	return z.readBufFileWithWrapper(
		path,
		func(r io.Reader) (rc io.ReadCloser, shouldClose bool, err error) {
			rc, err = zlib.NewReader(r)
			return rc, true, err
		},
		f,
	)
}

func (z *sorterImpl) readBufFilePlain(path string, f func(line string) error) error {
	return z.readBufFileWithWrapper(
		path,
		func(r io.Reader) (rc io.ReadCloser, shouldClose bool, err error) {
			return ioutil.NopCloser(r), false, nil
		},
		f,
	)
}

func (z *sorterImpl) readBufFile(path string, f func(line string) error) error {
	if z.opts.TempCompress {
		return z.readBufFileCompressed(path, f)
	} else {
		return z.readBufFilePlain(path, f)
	}
}

func (z *sorterImpl) clearBuffer(path string) {
	l := z.log()
	z.bufFile = path
	z.buf = make([]string, 0)
	z.bufSize = 0
	l.Debug("Buf file created", esl.String("path", path))
}

func (z *sorterImpl) rotate(flush func(w io.Writer) error) error {
	l := z.log()
	l.Debug("rotate buffer", esl.Int("lines", len(z.buf)))
	path, err := z.createBufFile(func(w io.Writer) error {
		return flush(w)
	})

	if err != nil {
		l.Debug("Unable to create buffer file", esl.Error(err))
		return err
	}

	z.clearBuffer(path)
	return nil
}

func (z *sorterImpl) rotateSimple() error {
	return z.rotate(z.simpleFlush)
}

func (z *sorterImpl) rotateMerge() error {
	return z.rotate(z.mergeFlush)
}

func (z *sorterImpl) compare(x, y string) int {
	if z.opts.Desc {
		return z.opts.Compare(x, y) * -1
	} else {
		return z.opts.Compare(x, y)
	}
}

func (z *sorterImpl) simpleFlush(w io.Writer) error {
	bw := bufio.NewWriter(w)
	l := z.log()
	l.Debug("Simple flush", esl.Int("lines", len(z.buf)))
	sort.Slice(z.buf, func(i, j int) bool {
		return z.compare(z.buf[i], z.buf[j]) < 0
	})
	for _, line := range z.buf {
		_, wErr := bw.WriteString(line + "\n")
		if wErr != nil {
			l.Debug("Unable to write", esl.Error(wErr))
			return wErr
		}
	}
	l.Debug("Wrote", esl.Int("lines", len(z.buf)))
	return bw.Flush()
}

func (z *sorterImpl) mergeFlush(w io.Writer) error {
	bw := bufio.NewWriter(w)
	bufLines := len(z.buf)
	l := z.log()
	l.Debug("Merge flush", esl.Int("lines", bufLines))
	bufIndex := 0
	var bufFileIndex int64
	sort.Slice(z.buf, func(i, j int) bool {
		return z.compare(z.buf[i], z.buf[j]) < 0
	})

	rbErr := z.readBufFile(z.bufFile, func(line string) error {
		bufFileIndex++
		if bufLines <= bufIndex {
			_, wErr := bw.WriteString(line + "\n")
			if wErr != nil {
				l.Debug("Write error", esl.Error(wErr))
			}
			return wErr
		}

		for bufIndex < bufLines {
			cmp := z.compare(z.buf[bufIndex], line)
			// Skip if duplicate found
			if cmp == 0 && z.opts.Uniq {
				bufIndex++
				continue
			}
			// end loop when buf[bufIndex] grater than the current line
			if 0 < cmp {
				break
			}
			if _, wErr := bw.WriteString(z.buf[bufIndex] + "\n"); wErr != nil {
				l.Debug("Write error", esl.Error(wErr))
				return wErr
			}
			bufIndex++
		}
		_, wErr := bw.WriteString(line + "\n")
		return wErr
	})
	for ; bufIndex < bufLines; bufIndex++ {
		if _, wErr := bw.WriteString(z.buf[bufIndex] + "\n"); wErr != nil {
			l.Debug("Write error", esl.Error(wErr))
			return wErr
		}
	}
	_ = bw.Flush()

	l.Debug("Merge finished", esl.Error(rbErr), esl.Int64("bufFileLines", bufFileIndex), esl.Int("bufMemLines", bufLines))
	return rbErr
}

func (z *sorterImpl) flushBuffer() error {
	l := z.log()
	l.Debug("Flush buffer")

	if z.bufFile == "" {
		return z.rotateSimple()
	} else {
		return z.rotateMerge()
	}
}

func (z *sorterImpl) writeEachLine(line string) error {
	lineSize := len(line)
	if z.opts.MemoryLimit < lineSize+z.bufSize {
		if err := z.flushBuffer(); err != nil {
			return err
		}
	}

	// skip if duplicate found in the buf
	if z.opts.Uniq {
		for _, ln := range z.buf {
			if ln == line {
				return nil
			}
		}
	}

	z.buf = append(z.buf, line)
	z.bufSize += lineSize
	return nil
}

func (z *sorterImpl) WriteLine(line string) error {
	z.bufMutex.Lock()
	defer z.bufMutex.Unlock()

	if z.isFlushed {
		return ErrorAlreadyClosed
	}

	for _, l := range strings.Split(line, "\n") {
		if err := z.writeEachLine(l); err != nil {
			return err
		}
	}
	return nil
}

func (z *sorterImpl) Flush() error {
	l := z.log()
	z.bufMutex.Lock()
	defer z.bufMutex.Unlock()

	if z.isFlushed {
		l.Debug("Ignore flush (already flushed)")
		return nil
	}

	if z.bufFile == "" {
		if fErr := z.simpleFlush(z.dest); fErr != nil {
			l.Debug("Unable to flush", esl.Error(fErr))
			return fErr
		}
	} else {
		if mErr := z.mergeFlush(z.dest); mErr != nil {
			l.Debug("Unable to merge & flush", esl.Error(mErr))
			return mErr
		}
	}

	_ = z.dest.Close()
	z.isFlushed = true
	return nil
}
