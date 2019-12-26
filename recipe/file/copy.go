package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Copy struct {
	Peer rc_conn.ConnUserFile
	Src  mo_path.DropboxPath
	Dst  mo_path.DropboxPath
}

func (z *Copy) Preset() {
}

func (z *Copy) Console() {
}

func (z *Copy) Exec(k rc_kitchen.Kitchen) error {
	uc := uc_file_relocation.New(z.Peer.Context())
	return uc.Copy(z.Src, z.Dst)
}

func (z *Copy) Test(c app_control.Control) error {
	return qt_endtoend.ScenarioTest()
}
