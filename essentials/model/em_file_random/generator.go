package em_file_random

import (
	"encoding/base32"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"math/rand"
	"time"
)

type Generator interface {
	Generate(opt ...Opt) (root em_file.Folder)

	Update(root em_file.Folder, r *rand.Rand)
}

func NewPoissonTree() Generator {
	return poissonTreeImpl{}
}

func NameByNodeId(nodeId int64) string {
	raw := make([]byte, 6)
	r := rand.New(rand.NewSource(nodeId))
	r.Read(raw)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(raw)
}

func SizeByNodeId(nodeId int64, opts Opts) int64 {
	r := rand.New(rand.NewSource(nodeId))
	return r.Int63n(opts.fileSizeRangeMax-opts.fileSizeRangeMin) + opts.fileSizeRangeMin
}

func TimeByNodeId(nodeId int64, opts Opts) time.Time {
	r := rand.New(rand.NewSource(nodeId))
	return time.Unix(r.Int63n(opts.fileDateRangeMax.Unix()-opts.fileDateRangeMin.Unix())+opts.fileDateRangeMin.Unix(), 0)
}

func NewGeneratedFile(nodeId int64, opts Opts) em_file.File {
	return em_file.NewFile(
		NameByNodeId(nodeId),
		SizeByNodeId(nodeId, opts),
		TimeByNodeId(nodeId, opts),
		nodeId,
	)
}
