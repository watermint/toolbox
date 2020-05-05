package test

import (
	"context"
	"fmt"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MonkeyWorker struct {
	Context dbx_context.Context
	Base    mo_path.DropboxPath
	Name    string
	Sem     *semaphore.Weighted
}

func (z *MonkeyWorker) upload() error {
	l := z.Context.Log()
	tf, err := ioutil.TempFile("", "monkey")
	if err != nil {
		l.Debug("Unable to create temp file", es_log.Error(err))
		return err
	}
	td := make([]byte, rand.Intn(384)+1)
	rand.Read(td)
	tf.Write(td)
	tf.Close()
	path := filepath.Join(filepath.Dir(tf.Name()), z.Name)
	os.Rename(tf.Name(), path)
	defer os.Remove(path)

	entry, err := sv_file_content.NewUpload(z.Context).Overwrite(z.Base, path)
	l.Info("Create or update", es_log.Any("entry", entry), es_log.Error(err))
	return nil
}

func (z *MonkeyWorker) delete() error {
	l := z.Context.Log()
	entry, err := sv_file.NewFiles(z.Context).Remove(z.Base.ChildPath(z.Name))
	l.Info("Delete", es_log.Any("entry", entry), es_log.Error(err))
	return nil
}

func (z *MonkeyWorker) Exec() error {
	defer z.Sem.Release(1)

	l := z.Context.Log().With(es_log.String("base", z.Base.Path()), es_log.String("name", z.Name))
	entry, err := sv_file.NewFiles(z.Context).Resolve(z.Base.ChildPath(z.Name))

	l.Debug("Entry", es_log.Any("entry", entry), es_log.Error(err))

	// Create if the file not found
	if err != nil {
		return z.upload()
	}

	// Probability; Update : Delete = 9 : 1
	isDelete := rand.Float32() < .1
	if isDelete {
		return z.delete()
	} else {
		return z.upload()
	}
}

type Monkey struct {
	rc_recipe.RemarkSecret
	Seconds      mo_int.RangeInt
	Distribution mo_int.RangeInt
	Path         mo_path.DropboxPath
	Peer         dbx_conn.ConnUserFile
	Extension    string
}

func (z *Monkey) Exec(c app_control.Control) error {
	exts := make([]string, 0)
	for _, e := range strings.Split(z.Extension, ",") {
		exts = append(exts, strings.TrimSpace(e))
	}
	extNum := len(exts)

	files := make([]string, z.Distribution.Value())
	folders := make([]mo_path.DropboxPath, z.Distribution.Value())
	for i := 0; i < z.Distribution.Value(); i++ {
		e := exts[rand.Intn(extNum)]
		files[i] = fmt.Sprintf("test-%05d.%s", i/10, e)
		folders[i] = z.Path.ChildPath(fmt.Sprintf("test%d", i%10))
	}
	sem := semaphore.NewWeighted(100)
	l := c.Log()
	l.Info("Monkey test start", es_log.Int("Distribution", z.Distribution.Value()), es_log.Int("Running time", z.Seconds.Value()))

	q := c.NewQueue()
	go func() {
		for {
			err := sem.Acquire(context.Background(), 1)
			if err != nil {
				l.Debug("Unable to acquire semaphore", es_log.Error(err))
				return
			}

			index := rand.Intn(z.Distribution.Value())
			file := files[index]
			folder := folders[index]

			l.Debug("Enqueue file", es_log.String("path", folder.Path()), es_log.String("name", file))
			q.Enqueue(&MonkeyWorker{
				Context: z.Peer.Context(),
				Base:    folder,
				Name:    file,
				Sem:     sem,
			})
		}
	}()

	time.Sleep(time.Duration(z.Seconds.Value()) * 1000 * time.Millisecond)
	return nil
}

func (z *Monkey) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Monkey{}, func(r rc_recipe.Recipe) {
		m := r.(*Monkey)
		m.Path = qt_recipe.NewTestDropboxFolderPath("dev-monkey")
		m.Seconds.SetValue(1)
		m.Distribution.SetValue(1000)
	})
}

func (z *Monkey) Preset() {
	z.Seconds.SetRange(1, 86400, 10)
	z.Distribution.SetRange(1, math.MaxInt32, 10000)
	z.Extension = "jpg,pdf,xlsx,docx,pptx,zip,png,txt,bak,csv,mov,mp4,html,gif,lzh,bmp,wmi,ini,ai,psd"
}
