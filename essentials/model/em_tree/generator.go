package em_tree

import (
	"encoding/base32"
	"github.com/watermint/toolbox/essentials/model/em_random"
	"math/rand"
	"time"
)

type Generator interface {
	Generate(opt ...Opt) (root Folder)

	Update(root Folder, r *rand.Rand)
}

func NewGenerator() Generator {
	return genImpl{}
}

type genImpl struct {
}

func (z genImpl) Update(root Folder, r *rand.Rand) {
	var update func(f Folder, r *rand.Rand) bool
	update = func(f Folder, r *rand.Rand) bool {
		switch r.Intn(20) {
		case 0: // rename self
			if root == f {
				break
			}
			f.Rename(NameByNodeId(r.Int63()))
			return true

		case 1: // delete one node
			descendants := f.Descendants()
			if len(descendants) > 0 {
				target := descendants[r.Intn(len(descendants))]
				f.Delete(target.Name())
				return true
			}

		case 2: // add a folder
			opts := Default().Apply([]Opt{NumNodes(4, 1, 10)})
			newFolder := NewGeneratedFolder(r.Int63(), 0, 0, 0, opts)
			f.Add(newFolder)
			return true

		case 3, 4: // rename a descendant
			descendants := f.Descendants()
			if len(descendants) > 0 {
				target := descendants[r.Intn(len(descendants))]
				target.Rename(NameByNodeId(r.Int63()))
				return true
			}

		case 5, 6, 7: // edit a file
			descendants := f.Descendants()
			for _, target := range descendants {
				if f, ok := target.(File); ok {
					f.UpdateContent(r.Int63(), r.Int63n(f.Size()+1)+f.Size()/4)
					return true
				}
			}

		default: // update descendant folder
			descendants := f.Descendants()
			for _, target := range descendants {
				if f, ok := target.(Folder); ok {
					if update(f, r) {
						return true
					}
				}
			}

		}

		// retry
		for {
			if update(f, r) {
				return true
			}
		}
	}

	update(root, r)
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
	return r.Int63n(opts.fileSizeRangeMax-opts.fileSizeRangeMin) + opts.fileSizeRangeMin
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

		// do not create empty folder
		if len(folder.Descendants()) > 0 {
			descendants = append(descendants, folder)
		}
	}

	return &folderNode{
		name:        NameByNodeId(nodeId),
		descendants: descendants,
	}
}
