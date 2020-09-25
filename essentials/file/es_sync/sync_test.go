package es_sync

import (
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_copier"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/model/em_file_random"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSyncImpl_Sync(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
	)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
	missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	if len(missingSources) > 0 {
		t.Error(missingSources)
	}
	if len(missingTargets) > 0 {
		t.Error(missingTargets)
	}
	if len(typeDiffs) > 0 {
		t.Error(typeDiffs)
	}
	if len(fileDiffs) > 0 {
		t.Error(es_json.ToJsonString(fileDiffs))
	}
}

func TestSyncImpl_SingleFile(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{})
	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
	)
	err := syncer.Sync(es_filesystem_model.NewPath("/a/x"), es_filesystem_model.NewPath("/w"))
	if err != nil {
		t.Error(err)
	}
	x := em_file.ResolvePath(tree1, "/a/x")
	if x == nil {
		t.Error(x)
	}

	w := em_file.ResolvePath(tree2, "/w/x")
	if w == nil {
		t.Error(w)
	}
}

func TestSyncImpl_ReplaceFolderByFile(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	tree1a := em_file.ResolvePath(tree1, "/a")
	tree1a.(em_file.Folder).Delete("c")
	if c := em_file.ResolvePath(tree1, "/a/c"); c != nil {
		t.Error(c)
	}
	tree1ac := em_file_random.NewGeneratedFile(rand.Int63(), em_file_random.Default())
	tree1ac.Rename("c")
	tree1a.(em_file.Folder).Add(tree1ac)

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(true),
	)
	em_file.Display(esl.Default(), tree1)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
	missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	if len(missingSources) > 0 {
		t.Error(missingSources)
	}
	if len(missingTargets) > 0 {
		t.Error(missingTargets)
	}
	if len(typeDiffs) > 0 {
		t.Error(typeDiffs)
	}
	if len(fileDiffs) > 0 {
		t.Error(es_json.ToJsonString(fileDiffs))
	}
}

func TestSyncImpl_ReplaceFileByFolder(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	tree1a := em_file.ResolvePath(tree1, "/a")
	tree1a.(em_file.Folder).Delete("x")
	if c := em_file.ResolvePath(tree1, "/a/x"); c != nil {
		t.Error(c)
	}
	tree1ax := em_file.NewFolder("x", []em_file.Node{})
	tree1a.(em_file.Folder).Add(tree1ax)

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(true),
	)
	em_file.Display(esl.Default(), tree1)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
	missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	if len(missingSources) > 0 {
		t.Error(missingSources)
	}
	if len(missingTargets) > 0 {
		t.Error(missingTargets)
	}
	if len(typeDiffs) > 0 {
		t.Error(typeDiffs)
	}
	if len(fileDiffs) > 0 {
		t.Error(es_json.ToJsonString(fileDiffs))
	}
}

func TestSyncImpl_Filter(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)
	filter := mo_filter.New("")
	filter.SetOptions(mo_filter.NewTestNameFilter("x"))
	if !filter.IsEnabled() {
		t.Error(filter)
	}

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(true),
		WithNameFilter(filter),
	)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	em_file.Display(esl.Default(), tree2)

	x := em_file.ResolvePath(tree2, "/a/y")
	if x != nil {
		t.Error(x)
	}
}

func TestSyncImpl_FileEdit(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	tree1ax := em_file.ResolvePath(tree1, "/a/x")
	tree1ax.(em_file.File).UpdateContent(rand.Int63(), 20)

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(true),
	)
	em_file.Display(esl.Default(), tree1)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
	missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	if len(missingSources) > 0 {
		t.Error(missingSources)
	}
	if len(missingTargets) > 0 {
		t.Error(missingTargets)
	}
	if len(typeDiffs) > 0 {
		t.Error(typeDiffs)
	}
	if len(fileDiffs) > 0 {
		t.Error(es_json.ToJsonString(fileDiffs))
	}
}

func TestSyncImpl_SyncRandom(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	l := esl.Default()
	seed := time.Now().UnixNano()
	l.Debug("Random test with seed", esl.Int64("seed", seed))

	r := rand.New(rand.NewSource(seed))

	tree1 := em_file_random.NewPoissonTree().Generate(em_file_random.NumFiles(100))
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	for i := 0; i < 3; i++ {
		l.Info("Sync try", esl.Int("tries", i))
		seq := eq_sequence.New()
		conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)
		fs1 := es_filesystem_model.NewFileSystem(tree1)
		fs2 := es_filesystem_model.NewFileSystem(tree2)

		syncer := New(
			esl.Default(),
			seq,
			fs1,
			fs2,
			conn,
			SyncOverwrite(true),
			SyncDelete(true),
		)
		err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
		if err != nil {
			t.Error(seed, i, err)
		}
		folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
		missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
		if err != nil {
			t.Error(seed, i, err)
		}
		if len(missingSources) > 0 {
			t.Error(seed, i, es_json.ToJsonString(missingSources))
		}
		if len(missingTargets) > 0 {
			t.Error(seed, i, es_json.ToJsonString(missingTargets))
		}
		if len(typeDiffs) > 0 {
			t.Error(seed, i, es_json.ToJsonString(typeDiffs))
		}
		if len(fileDiffs) > 0 {
			t.Error(seed, i, es_json.ToJsonString(fileDiffs))
		}

		for j := 0; j < 10; j++ {
			em_file_random.NewPoissonTree().Update(tree1, r)
		}
	}
}

func TestSyncImpl_SyncRandomReduceCreateFolder(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	l := esl.Default()
	seed := time.Now().UnixNano()
	l.Debug("Random test with seed", esl.Int64("seed", seed))

	r := rand.New(rand.NewSource(seed))

	tree1 := em_file_random.NewPoissonTree().Generate(em_file_random.NumFiles(100))
	em_file.DeleteEmptyFolders(tree1)
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	for i := 0; i < 3; i++ {
		l.Info("Sync try", esl.Int("tries", i))
		seq := eq_sequence.New()
		conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)
		fs1 := es_filesystem_model.NewFileSystem(tree1)
		fs2 := es_filesystem_model.NewFileSystem(tree2)

		syncer := New(
			esl.Default(),
			seq,
			fs1,
			fs2,
			conn,
			SyncOverwrite(true),
			SyncDelete(true),
			OptimizePreventCreateFolder(true),
		)
		err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
		if err != nil {
			t.Error(seed, i, err)
		}
		folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
		missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
		if err != nil {
			t.Error(seed, i, err)
		}
		if len(missingSources) > 0 {
			t.Error(seed, i, es_json.ToJsonString(missingSources))
		}
		if len(missingTargets) > 0 {
			t.Error(seed, i, es_json.ToJsonString(missingTargets))
		}
		if len(typeDiffs) > 0 {
			t.Error(seed, i, es_json.ToJsonString(typeDiffs))
		}
		if len(fileDiffs) > 0 {
			t.Error(seed, i, es_json.ToJsonString(fileDiffs))
		}

		for j := 0; j < 10; j++ {
			em_file_random.NewPoissonTree().Update(tree1, r)
		}
		em_file.DeleteEmptyFolders(tree1)
	}
}

func BenchmarkSyncImpl_SyncRandomTest(b *testing.B) {
	l := esl.Default()
	masterSeed := time.Now().UnixNano()
	l.Debug("Random test with seed", esl.Int64("seed", masterSeed))
	masterRand := rand.New(rand.NewSource(masterSeed))
	wg := sync.WaitGroup{}

	bench := func(runner int) {
		seed := masterRand.Int63()
		l.Debug("Random test with seed", esl.Int64("seed", seed))

		r := rand.New(rand.NewSource(seed))

		tree1 := em_file_random.NewPoissonTree().Generate()
		tree2 := em_file.NewFolder("root", []em_file.Node{})

		for i := 0; i < b.N; i++ {
			l.Info("Sync try", esl.Int("tries", i), esl.Int("runner", runner))
			seq := eq_sequence.New()
			conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)
			fs1 := es_filesystem_model.NewFileSystem(tree1)
			fs2 := es_filesystem_model.NewFileSystem(tree2)

			syncer := New(
				esl.Default(),
				seq,
				fs1,
				fs2,
				conn,
				SyncOverwrite(true),
				SyncDelete(true),
			)
			err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
			if err != nil {
				b.Error(seed, i, err)
			}
			folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
			missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
			if err != nil {
				b.Error(seed, i, err)
			}
			if len(missingSources) > 0 {
				b.Error(seed, i, es_json.ToJsonString(missingSources))
			}
			if len(missingTargets) > 0 {
				b.Error(seed, i, es_json.ToJsonString(missingTargets))
			}
			if len(typeDiffs) > 0 {
				b.Error(seed, i, es_json.ToJsonString(typeDiffs))
			}
			if len(fileDiffs) > 0 {
				b.Error(seed, i, es_json.ToJsonString(fileDiffs))
			}

			em_file_random.NewPoissonTree().Update(tree1, r)
		}
		wg.Done()
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go bench(i)
	}
	wg.Wait()
}
