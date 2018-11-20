package oper_member

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/poc/oper"
	"github.com/watermint/toolbox/poc/oper/oper_cli"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestGroup_Tag(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}
	box := rice.MustFindBox("..")

	opers := []oper.Operation{
		&Member{},
		&Invite{},
		&List{},
	}

	ctx := &oper.Context{
		Logger: log,
		Box:    box,
		UI: &oper_cli.CUI{
			Out: os.Stdout,
			In:  os.Stdin,
		},
	}

	for _, o := range opers {
		op := oper.Operator{
			Context: ctx,
			Op:      o,
		}
		op.Init()
		log.Info("Operation",
			zap.String("title", op.Title()),
			zap.String("Desc", op.Desc()),
		)

		if op.IsGroup() {
			log.Info("Sub commands")
			for _, so := range op.SubOperators() {
				log.Info("Sub command",
					zap.String("Path", fmt.Sprintf("%s.%s", op.Tag(), so.Tag())),
				)
			}
		}
		if op.IsExecutable() {
			log.Info("Exec", zap.String("tag", op.Tag()))
			op.InjectLog()
			op.InjectContext()
			op.Executable().Exec()
		}

	}

}
