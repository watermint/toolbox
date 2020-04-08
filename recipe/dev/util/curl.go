package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tdewolff/parse/buffer"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/network/nw_capture"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"io"
	"os"
	"sort"
	"strings"
)

type Curl struct {
	Record     string
	BufferSize int
}

func (z *Curl) Preset() {
	z.BufferSize = 65536
}

func (z *Curl) Exec(c app_control.Control) error {
	l := c.Log()
	w := ut_io.NewDefaultOut(c.Feature().IsTest())
	bw := bufio.NewWriter(w)
	defer bw.Flush()
	var r io.Reader
	r = os.Stdin
	if z.Record != "" {
		r = buffer.NewReader([]byte(z.Record))
	}
	br := bufio.NewReaderSize(r, z.BufferSize)
	for {
		line, prefix, err := br.ReadLine()
		if prefix {
			l.Warn("Line is too long, terminate this operation")
			return nil
		}
		switch err {
		case nil:
			rec := &nw_capture.Record{}
			if pe := json.Unmarshal(line, rec); pe != nil {
				l.Error("Unable to unmarshal", zap.Error(err), zap.String("line", string(line)))
				return err
			}
			fmt.Fprintf(bw, "curl -D - -X POST %s \\\n", rec.Req.RequestUrl)
			reqKeys := make([]string, 0)
			for k := range rec.Req.RequestHeaders {
				reqKeys = append(reqKeys, k)
			}
			sort.Strings(reqKeys)
			for _, k := range reqKeys {
				fmt.Fprintf(bw, "     --header \"%s: %s\" \\\n", k, rec.Req.RequestHeaders[k])
			}
			fmt.Fprintf(bw, "     --data \"%s\"\n", strings.ReplaceAll(rec.Req.RequestParam, "\"", "\\\""))
			fmt.Fprintf(bw, "\n")
			fmt.Fprintf(bw, "HTTP/2 %d\n", rec.Res.ResponseCode)
			resKeys := make([]string, 0)
			for k := range rec.Res.ResponseHeaders {
				resKeys = append(resKeys, k)
			}
			sort.Strings(resKeys)
			for _, k := range resKeys {
				fmt.Fprintf(bw, "%s: %s\n", strings.ToLower(k), rec.Res.ResponseHeaders[k])
			}
			fmt.Fprintf(bw, "\n")
			if rec.Res.ResponseError != "" {
				fmt.Fprintln(bw, rec.Res.ResponseError)
			} else if rec.Res.ResponseBody != "" {
				fmt.Fprintln(bw, rec.Res.ResponseBody)
			} else {
				fmt.Fprintln(bw, string(rec.Res.ResponseJson))
			}
			fmt.Fprintf(bw, "\n\n\n")

		case io.EOF:
			return nil
		default:
			return err
		}
	}
}

func (z *Curl) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Curl{}, func(r rc_recipe.Recipe) {
		m := r.(*Curl)
		m.Record = `{"time":"2020-02-27T18:34:58.677+0900","req":{"method":"POST","url":"https://api.dropboxapi.com/2/files/create_folder_v2","param":"{\"path\":\"/watermint-toolbox-test/2020-02-27T18-34-37/file-sync-up/d-e-f\",\"autorename\":false}","headers":{"Authorization":"Bearer <secret>","Content-Type":"application/json"},"content_length":92},"res":{"code":429,"headers":{"Content-Disposition":"attachment; filename='error'","Content-Security-Policy":"sandbox; frame-ancestors 'none'","Content-Type":"application/json","Date":"Thu, 27 Feb 2020 09:34:58 GMT","Retry-After":"1","Server":"nginx","X-Content-Type-Options":"nosniff","X-Dropbox-Request-Id":"0b3e28a4804b2b99a4eab9c48e07536a","X-Frame-Options":"DENY"},"json":{"error_summary":"too_many_write_operations/.","error":{"reason":{".tag":"too_many_write_operations"}}},"content_length":108},"latency":857121385}`
	})
}
