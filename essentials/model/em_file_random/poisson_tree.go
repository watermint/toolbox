package em_file_random

import (
	"github.com/watermint/toolbox/essentials/model/em_file"
	"math/rand"
)

type poissonTreeImpl struct {
}

func (z poissonTreeImpl) electFolderFromExisting(r *rand.Rand, base em_file.Folder, depth int, opts Opts) em_file.Folder {
	descendants := base.Descendants()
	folders := make([]em_file.Folder, 0)
	for _, d := range descendants {
		if folder, ok := d.(em_file.Folder); ok {
			folders = append(folders, folder)
		}
	}

	// Elect base folder if no sub folders found
	if len(folders) < 1 {
		return base
	}

	x := rand.Intn(10)
	switch {
	case x < 2:
		return folders[r.Intn(len(folders))]

	default:
		if depth < opts.depthRangeMax {
			return z.electFolderFromExisting(r, folders[r.Intn(len(folders))], depth+1, opts)
		} else {
			return folders[r.Intn(len(folders))]
		}
	}
}

func (z poissonTreeImpl) electFolderNewFolder(r *rand.Rand, base em_file.Folder, depth int, opts Opts) em_file.Folder {
	if depth <= opts.depthRangeMax {
		if opts.maxFoldersInFolder < base.NumFolders() {
			return z.electFolderFromExisting(r, base, depth, opts)
		}

		folder := em_file.NewFolder(NameByNodeId(r.Int63()), []em_file.Node{})
		base.Add(folder)
		return folder
	}

	return z.electFolderFromExisting(r, base, depth, opts)
}

func (z poissonTreeImpl) electFolder(r *rand.Rand, base em_file.Folder, depth int, opts Opts) em_file.Folder {
	x := rand.Intn(48)
	switch {
	case x < 2:
		folder := z.electFolderFromExisting(r, base, depth, opts)
		return z.electFolderNewFolder(r, folder, depth, opts)

	case x < 3:
		return z.electFolderNewFolder(r, base, depth, opts)

	default:
		return z.electFolderFromExisting(r, base, depth, opts)
	}
}

func (z poissonTreeImpl) addFile(nodeId int64, root em_file.Folder, opts Opts) {
	r := rand.New(rand.NewSource(nodeId))
	file := NewGeneratedFile(r.Int63(), opts)

	folder := z.electFolder(r, root, 0, opts)
	folder.Add(file)
}

func (z poissonTreeImpl) newTreeFolder(nodeId int64, size int64, depth, files int, opts Opts) em_file.Folder {
	r := rand.New(rand.NewSource(nodeId))
	descendants := make([]em_file.Node, 0)
	ratio := r.Float32()
	numNodes := r.Intn(opts.maxFilesInFolder + opts.maxFoldersInFolder)

	numFiles := int(float32(numNodes) * ratio)
	numFolders := int(float32(numNodes) * (1 - ratio))
	if opts.depthRangeMax <= depth+1 {
		numFolders = 0
	}
	if numFolders < 1 {
		if files+numFiles < opts.numFiles {
			numFiles = opts.numFiles - numFolders
		}
	}
	if opts.numFiles < files+numFiles {
		numFiles = opts.numFiles - files
	}

	var folderSize int64

	for i := 0; i < numFiles; i++ {
		file := NewGeneratedFile(r.Int63(), opts)
		folderSize += file.Size()
		descendants = append(descendants, file)
	}
	for i := 0; i < numFolders; i++ {
		folder := z.newTreeFolder(r.Int63(), folderSize+size, depth+1, files+numNodes, opts)
		size += em_file.SumFileSize(folder)
		files += em_file.SumNumNode(folder)

		// do not create empty folder
		if len(folder.Descendants()) > 0 {
			descendants = append(descendants, folder)
		}
	}

	return em_file.NewFolder(NameByNodeId(nodeId), descendants)
}

func (z poissonTreeImpl) Update(root em_file.Folder, r *rand.Rand) {
	var update func(f em_file.Folder, r *rand.Rand) bool
	update = func(f em_file.Folder, r *rand.Rand) bool {
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
			z.electFolderNewFolder(r, f, 0, Default())
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
				if f, ok := target.(em_file.File); ok {
					f.UpdateContent(r.Int63(), r.Int63n(f.Size()+1)+f.Size()/4)
					return true
				}
			}

		default: // update descendant folder
			descendants := f.Descendants()
			for _, target := range descendants {
				if f, ok := target.(em_file.Folder); ok {
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

func (z poissonTreeImpl) Generate(opt ...Opt) (root em_file.Folder) {
	opts := Default().Apply(opt)
	root = em_file.NewFolder("", []em_file.Node{})
	r := rand.New(rand.NewSource(opts.seed))

	for i := 0; i < opts.numFiles; i++ {
		z.addFile(r.Int63(), root, opts)
	}
	return
}
