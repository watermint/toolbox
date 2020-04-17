package dev

import (
	"bufio"
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_worker_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"sort"
)

type AsyncWorker struct {
	// job context
	group *mo_group.Group

	// recipe's context
	ctl  app_control.Control
	conn dbx_context.Context
	rep  rp_model.RowReport
}

func (z *AsyncWorker) Exec() error {
	l := z.ctl.Log()
	l.Debug("Scan group (Multi thread)", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Group", z.group))

	msv := sv_group_member.New(z.conn, z.group)
	members, err := msv.List()
	if err != nil {
		return err
	}
	for _, m := range members {
		row := mo_group_member.NewGroupMember(z.group, m)
		z.rep.Row(row)
	}
	return nil
}

type Async struct {
	rc_recipe.RemarkSecret
	RunConcurrently bool
	Peer            dbx_conn.ConnBusinessInfo
	Rows            rp_model.RowReport
}

func (z *Async) Preset() {
	z.Rows.SetModel(
		&mo_group_member.GroupMember{},
		rp_model.HiddenColumns(
			"group_id",
			"account_id",
			"team_member_id",
		),
	)
}

func (z *Async) Exec(c app_control.Control) error {
	ctxInfo := z.Peer.Context()

	gsv := sv_group.New(ctxInfo)
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	err = z.Rows.Open()
	if err != nil {
		return err
	}

	q := c.NewQueue()

	// Launch additional routines (because only single routine running when the recipe
	// run through test
	qq := q.(*rc_worker_impl.Queue)
	qq.Launch(4)

	for _, group := range groups {
		if z.RunConcurrently {
			w := &AsyncWorker{
				group: group,
				ctl:   c,
				conn:  ctxInfo,
				rep:   z.Rows,
			}
			q.Enqueue(w)
		} else {
			c.Log().Debug("Scan group (Single thread)", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Group", group))
			msv := sv_group_member.New(ctxInfo, group)
			members, err := msv.List()
			if err != nil {
				return err
			}
			for _, m := range members {
				row := mo_group_member.NewGroupMember(group, m)
				z.Rows.Row(row)
			}
		}
	}
	q.Wait()

	return nil
}

func (z *Async) Test(c app_control.Control) error {
	l := c.Log()

	var singleTheadReport, multiThreadReport string

	// Concurrent
	{
		l.Info("Running multi thread operation")
		cf := c.(app_control_launcher.ControlFork)
		cc, err := cf.Fork("async")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(cc, &Async{}, func(r rc_recipe.Recipe) {
			ar := r.(*Async)
			ar.RunConcurrently = true
		})
		if err != nil {
			return err
		}
		multiThreadReport = filepath.Join(cc.Workspace().Report(), "rows.csv")
	}

	// Single thread
	{
		l.Info("Running single-thread operation")
		cf := c.(app_control_launcher.ControlFork)
		cc, err := cf.Fork("single-thread")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(cc, &Async{}, func(r rc_recipe.Recipe) {
			ar := r.(*Async)
			ar.RunConcurrently = true
		})
		if err != nil {
			return err
		}
		singleTheadReport = filepath.Join(cc.Workspace().Report(), "rows.csv")
	}

	singleReport := make([]string, 0)
	{
		f, err := os.Open(singleTheadReport)
		if err != nil {
			return err
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			singleReport = append(singleReport, s.Text())
		}
	}
	sort.Strings(singleReport)

	concurrentReport := make([]string, 0)
	{
		f, err := os.Open(multiThreadReport)
		if err != nil {
			return err
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			concurrentReport = append(concurrentReport, s.Text())
		}
	}
	sort.Strings(concurrentReport)

	d := cmp.Diff(singleReport, concurrentReport)
	l.Debug("Diff")
	if d != "" {
		l.Error("Diff found", zap.String("diff", d))
		return errors.New("diff found")
	}
	return nil
}
