package es_response_impl

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const (
	ReadBufferSize = 10 * 1048576 // 10MiB
	ReadChunkSize  = 64 * 1024    // 64KiB
)

var (
	ErrorInvalidBufferState = errors.New("invalid buffer state")
)

func readJitterWait() {
	wms := rand.Intn(50) + 10
	time.Sleep(time.Duration(wms) * time.Millisecond)
}

func Read(ctx es_context.Context, resBody io.ReadCloser) (es_response.Body, error) {
	if body, err := read(ctx, resBody, ReadBufferSize, ReadChunkSize); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func read(ctx es_context.Context, resBody io.ReadCloser, readBufSize, readChunkSize int) (body es_response.Body, err error) {
	l := ctx.Log().With(
		es_log.Int("readBufSize", readBufSize),
		es_log.Int("readChunkSize", readChunkSize))
	// Subtract single chunk size from read buf size, for not to exceed read buffer.
	readBufSize = readBufSize - readChunkSize
	if readBufSize < 1 {
		l.Error("Read buffer size is smaller than chunk size.")
		return nil, ErrorInvalidBufferState
	}

	var bodyBuf bytes.Buffer
	bodyReader := nw_bandwidth.WrapReader(resBody)
	defer func() {
		// close response body, due to it's a duty of a caller.
		if err := resBody.Close(); err != nil {
			l.Debug("Unable to close response body", es_log.Error(err))
		}
	}()

	// ready content into the buffer
	for {
		if bodyBuf.Len() > readBufSize {
			l.Debug("Body size exceeds read buffer size", es_log.Int("readBytes", bodyBuf.Len()))
			break
		}
		n, err := io.CopyN(&bodyBuf, bodyReader, int64(readChunkSize))
		switch err {
		case io.EOF:
			return newMemoryBody(ctx, bodyBuf.Bytes()), nil

		case nil:
			if n == 0 {
				readJitterWait()
			}
			continue

		default:
			l.Debug("Body read error",
				es_log.Int("readBytes", bodyBuf.Len()),
				es_log.Error(err))
			return nil, err
		}
	}

	// Create file
	bodyFile, err := ioutil.TempFile("", ctx.ClientHash())
	if err != nil {
		l.Debug("Unable to create file", es_log.Error(err))
		return nil, err
	}

	// Keep the file safer
	if err := bodyFile.Chmod(0600); err != nil {
		l.Debug("Unable to change file mode", es_log.Error(err))
		return nil, err
	}

	cleanupOnError := func() {
		if err := bodyFile.Close(); err != nil {
			l.Debug("Error on closing body file", es_log.Error(err))
		}
		if err := os.Remove(bodyFile.Name()); err != nil {
			l.Debug("Error on removing the file", es_log.Error(err))
		}
	}

	// Flush buffer to the file
	readBodyBufSize := int64(bodyBuf.Len())
	fileBuf := bufio.NewWriter(bodyFile)
	fileBytes, err := io.Copy(fileBuf, &bodyBuf)
	if err != nil {
		l.Debug("Unable to write read body buffer", es_log.Error(err))
		cleanupOnError()
		return nil, err
	}

	if fileBytes != readBodyBufSize {
		l.Debug("Buffer content mismatch",
			es_log.Int64("writtenToFile", fileBytes),
			es_log.Int64("bodyBufferSize", readBodyBufSize))
		cleanupOnError()
		return nil, ErrorInvalidBufferState
	}

	// Read body & write content into the file
	for {
		n, err := io.CopyN(fileBuf, bodyReader, int64(readChunkSize))
		fileBytes += n

		switch err {
		case io.EOF:
			if err := fileBuf.Flush(); err != nil {
				l.Debug("Unable to flush content into the file",
					es_log.Int64("readBytes", fileBytes),
					es_log.Error(err))
				cleanupOnError()
				return nil, err
			}
			if err := bodyFile.Close(); err != nil {
				l.Debug("Unable to close file", es_log.Error(err))
				cleanupOnError()
				return nil, err
			}
			return newFileBody(ctx, bodyFile.Name(), fileBytes), nil

		case nil:
			if n == 0 {
				readJitterWait()
			}
			continue

		default:
			l.Debug("Body read error",
				es_log.Int64("readBytes", fileBytes),
				es_log.Error(err))
			return nil, err
		}
	}
}
