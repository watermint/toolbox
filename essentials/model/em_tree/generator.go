package em_tree

import (
	"encoding/base32"
	"github.com/watermint/toolbox/essentials/model/em_random"
	"math/rand"
	"time"
)

type Generator interface {
	Generate(opt ...Opt) (root Folder)

	Update(root Node, numTries int, opt ...Opt)
}

func NewGenerator() Generator {
	return genImpl{}
}

type genImpl struct {
}

func (z genImpl) Update(root Node, numTries int, opt ...Opt) {

}

func (z genImpl) Generate(opt ...Opt) (root Folder) {
	opts := Default().Apply(opt)
	return NewGeneratedFolder(opts.seed, 0, 0, 0, opts)
}

func NameByNodeId(nodeId int64) string {
	raw := make([]byte, 6)
	r := rand.New(rand.NewSource(nodeId))
	r.Read(raw)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(raw)
}

func SizeByNodeId(nodeId int64, opts Opts) int64 {
	r := rand.New(rand.NewSource(nodeId))
	lambda := float64(opts.fileSizeLambda)
	min := float64(opts.fileSizeRangeMin)
	max := float64(opts.fileSizeRangeMax)
	return int64(em_random.PoissonWithRange(r, lambda, min, max))
}

func TimeByNodeId(nodeId int64, opts Opts) time.Time {
	r := rand.New(rand.NewSource(nodeId))
	return time.Unix(r.Int63n(opts.fileDateRangeMax.Unix()-opts.fileDateRangeMin.Unix())+opts.fileDateRangeMin.Unix(), 0)
}

func NewGeneratedFile(nodeId int64, opts Opts) File {
	return &fileNode{
		name:  NameByNodeId(nodeId),
		size:  SizeByNodeId(nodeId, opts),
		mtime: TimeByNodeId(nodeId, opts),
	}
}

func NewGeneratedFolder(nodeId int64, size int64, depth, nodes int, opts Opts) Folder {
	r := rand.New(rand.NewSource(nodeId))
	descendants := make([]Node, 0)
	ratio := r.Float32()
	numNodes := int(em_random.PoissonWithRange(r, float64(opts.numDescendantLambda), float64(opts.numDescendantRangeMin), float64(opts.numDescendantRangeMax)))

	if opts.numNodesRangeMax < nodes+numNodes {
		numNodes = opts.numNodesRangeMax - nodes
	}
	numFiles := int(float32(numNodes) * ratio)
	numFolders := int(float32(numNodes) * (1 - ratio))
	if opts.depthRangeMax <= depth+1 {
		numFolders = 0
	}
	if numFolders < 1 && nodes+numNodes < opts.numNodesRangeMin {
		numFiles = opts.numNodesRangeMin - numFolders
	}
	var folderSize int64

	for i := 0; i < numFiles; i++ {
		file := NewGeneratedFile(r.Int63(), opts)
		folderSize += file.Size()
		descendants = append(descendants, file)
	}
	for i := 0; i < numFolders; i++ {
		folder := NewGeneratedFolder(r.Int63(), folderSize+size, depth+1, nodes+numNodes, opts)
		size += SumFileSize(folder)
		nodes += SumNumNode(folder)
		descendants = append(descendants, folder)
	}

	return &folderNode{
		name:        NameByNodeId(nodeId),
		descendants: descendants,
	}
}
