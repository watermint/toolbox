package dev

import (
	"bufio"
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/recipe/rc_worker_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"sort"
)

type AsyncVO struct {
	RunConcurrently bool
	Peer            rc_conn.ConnBusinessInfo
}

const (
	reportAsyncConcurrent = "concurrent"
	reportAsyncSingle     = "single_thread"
)

func (z *AsyncVO) reportName() string {
	if z.RunConcurrently {
		return reportAsyncConcurrent
	} else {
		return reportAsyncSingle
	}
}

type AsyncWorker struct {
	// job context
	group *mo_group.Group

	// recipe's context
	ctl  app_control.Control
	conn api_context.Context
	rep  rp_model.Report
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
}

func (z *Async) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportAsyncConcurrent, &mo_group_member.GroupMember{}),
		rp_spec_impl.Spec(reportAsyncSingle, &mo_group_member.GroupMember{}),
	}
}

func (z *Async) Hidden() {
}

func (z *Async) Requirement() rc_vo.ValueObject {
	return &AsyncVO{}
}

func (z *Async) Exec(k rc_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*AsyncVO)
	connInfo, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	gsv := sv_group.New(connInfo)
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(lvo.reportName())
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()

	// Launch additional routines (because only single routine running when the recipe
	// run through test
	qq := q.(*rc_worker_impl.Queue)
	qq.Launch(4)

	for _, group := range groups {
		if lvo.RunConcurrently {
			w := &AsyncWorker{
				group: group,
				ctl:   k.Control(),
				conn:  connInfo,
				rep:   rep,
			}
			q.Enqueue(w)
		} else {
			k.Log().Debug("Scan group (Single thread)", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Group", group))
			msv := sv_group_member.New(connInfo, group)
			members, err := msv.List()
			if err != nil {
				return err
			}
			for _, m := range members {
				row := mo_group_member.NewGroupMember(group, m)
				rep.Row(row)
			}
		}
	}
	q.Wait()

	return nil
}

func (z *Async) Test(c app_control.Control) error {
	lvo := &AsyncVO{}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}

	l := c.Log()

	// Non concurrent operation:
	l.Info("Running single thread operation")
	lvo.RunConcurrently = false
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	singleReportPath := filepath.Join(c.Workspace().Report(), lvo.reportName()+".csv")

	// Concurrent operation:
	l.Info("Running multi thread operation")
	lvo.RunConcurrently = true
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	concurrentReportPath := filepath.Join(c.Workspace().Report(), lvo.reportName()+".csv")

	singleReport := make([]string, 0)
	{
		f, err := os.Open(singleReportPath)
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
		f, err := os.Open(concurrentReportPath)
		if err != nil {
			return err
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			concurrentReport = append(concurrentReport, s.Text())
		}
	}
	sort.Strings(concurrentReport)

	var err error
	if len(singleReport) != len(concurrentReport) {
		l.Error("Size mismatch")
		err = errors.New("report size mismatch")
	}

	l.Info("Compare single to concurrent",
		zap.Int("singleRecords", len(singleReport)),
		zap.Int("concurrentRecords", len(concurrentReport)),
	)
	for i, single := range singleReport {
		if len(concurrentReport) < i {
			break
		}
		concurrent := concurrentReport[i]
		if concurrent != single {
			l.Error("Line diff",
				zap.Int("Line", i),
				zap.String("Single", single),
				zap.String("Concurrent", concurrent),
			)
			err = errors.New("line diff found")
		}
	}

	l.Info("Compare concurrent to single")
	for i, concurrent := range concurrentReport {
		if len(singleReport) < i {
			break
		}
		single := singleReport[i]
		if concurrent != single {
			l.Error("Line diff",
				zap.Int("Line", i),
				zap.String("Single", single),
				zap.String("Concurrent", concurrent),
			)
			err = errors.New("line diff found")
		}
	}

	return err
}
