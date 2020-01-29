package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type MonkeyWorker struct {
	Context api_context.Context
	Base    mo_path.DropboxPath
	File    mo_path.DropboxPath
	Sem     *semaphore.Weighted
}

func (z *MonkeyWorker) upload() error {
	l := z.Context.Log()
	tf, err := ioutil.TempFile("", "test-*.dat")
	if err != nil {
		l.Debug("Unable to create temp file", zap.Error(err))
		return err
	}
	defer os.Remove(tf.Name())
	td := make([]byte, rand.Intn(384)+1)
	rand.Read(td)
	tf.Write(td)

	entry, err := sv_file_content.NewUpload(z.Context).Overwrite(z.File, tf.Name())
	l.Info("Create or update", zap.Any("entry", entry), zap.Error(err))
	return err
}

func (z *MonkeyWorker) delete() error {
	l := z.Context.Log()
	entry, err := sv_file.NewFiles(z.Context).Remove(z.File)
	l.Info("Delete", zap.Any("entry", entry), zap.Error(err))
	return err
}

func (z *MonkeyWorker) Exec() error {
	defer z.Sem.Release(1)

	l := z.Context.Log().With(zap.String("file", z.File.Path()))
	entry, err := sv_file.NewFiles(z.Context).Resolve(z.File)

	l.Debug("Entry", zap.Any("entry", entry), zap.Error(err))

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
	Seconds      int
	Distribution int
	Path         mo_path.DropboxPath
	Peer         rc_conn.ConnUserFile
}

func (z *Monkey) Exec(c app_control.Control) error {
	if z.Distribution < 1 {
		return errors.New("distribution must be grater than 1")
	}
	files := make([]mo_path.DropboxPath, z.Distribution)
	for i := 0; i < z.Distribution; i++ {
		//file := fmt.Sprintf("test-%05d.dat", i/10)
		folder := fmt.Sprintf("test%d", i%10)
		files[i] = z.Path.ChildPath(folder)
	}
	sem := semaphore.NewWeighted(100)
	l := c.Log()
	l.Info("Monkey test start", zap.Int("Distribution", z.Distribution), zap.Int("Running time", z.Seconds))

	q := c.NewQueue()
	go func() {
		for {
			err := sem.Acquire(context.Background(), 1)
			if err != nil {
				l.Debug("Unable to acquire semaphore", zap.Error(err))
				return
			}

			f := files[rand.Intn(z.Distribution)]
			l.Debug("Enqueue file", zap.String("file", f.Path()))
			q.Enqueue(&MonkeyWorker{
				Context: z.Peer.Context(),
				Base:    z.Path,
				File:    f,
				Sem:     sem,
			})
		}
	}()

	time.Sleep(time.Duration(z.Seconds) * time.Second)
	return nil
}

func (z *Monkey) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Monkey{}, func(r rc_recipe.Recipe) {
		m := r.(*Monkey)
		m.Path = mo_path.NewDropboxPath(qt_recipe.TestTeamFolderName).ChildPath("dev-monkey")
		m.Seconds = 1
		m.Distribution = 1000
	})
}

func (z *Monkey) Preset() {
	z.Seconds = 10
	z.Distribution = 10000
}
