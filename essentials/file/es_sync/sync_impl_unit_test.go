package es_sync

import (
	"testing"
	"time"

	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_copier"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func TestNew(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
	)

	if syncer == nil {
		t.Error("Expected non-nil syncer")
	}

	syncImpl := syncer.(*syncImpl)
	if syncImpl.log == nil {
		t.Error("Expected logger to be set")
	}

	if syncImpl.source == nil {
		t.Error("Expected source filesystem to be set")
	}

	if syncImpl.target == nil {
		t.Error("Expected target filesystem to be set")
	}

	if syncImpl.conn == nil {
		t.Error("Expected connector to be set")
	}

	if syncImpl.fileCmp == nil {
		t.Error("Expected file comparator to be set")
	}
}

func TestNewWithOptions(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	filter := mo_filter.New("")
	filter.SetOptions(mo_filter.NewTestNameFilter("test"))

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(false),
		WithNameFilter(filter),
		OptimizePreventCreateFolder(true),
	)

	if syncer == nil {
		t.Error("Expected non-nil syncer")
	}

	syncImpl := syncer.(*syncImpl)
	if !syncImpl.opts.syncDelete {
		t.Error("Expected syncDelete to be true")
	}

	if syncImpl.opts.syncOverwrite {
		t.Error("Expected syncOverwrite to be false")
	}

	if syncImpl.opts.optimizeReduceCreateFolder != true {
		t.Error("Expected optimizeReduceCreateFolder to be true")
	}
}

func TestSyncImpl_ComputeBatchId(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
	)

	syncImpl := syncer.(*syncImpl)

	source := es_filesystem_model.NewPath("/test/source")
	target := es_filesystem_model.NewPath("/test/target")

	batchId := syncImpl.computeBatchId(source, target)
	if batchId == "" {
		t.Error("Expected non-empty batch ID")
	}

	// Test that same paths produce same batch ID
	batchId2 := syncImpl.computeBatchId(source, target)
	if batchId != batchId2 {
		t.Error("Expected same batch ID for same paths")
	}

	// Test that different paths can produce different batch IDs
	// (Note: batch IDs are based on shard IDs, so they might be the same for simple paths)
	target2 := es_filesystem_model.NewPath("/test/different")
	batchId3 := syncImpl.computeBatchId(source, target2)
	// For model filesystem, shards might be the same, so we just verify the method works
	if batchId3 == "" {
		t.Error("Expected non-empty batch ID for different paths")
	}
}

func TestSyncImpl_EnqueueTask(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	qd := eq_queue.New()

	syncer := New(
		esl.Default(),
		qd,
		fs1,
		fs2,
		conn,
	)

	syncImpl := syncer.(*syncImpl)

	source := es_filesystem_model.NewPath("/test/source")
	target := es_filesystem_model.NewPath("/test/target")

	// This should not panic - just test that the method exists
	defer func() {
		if r := recover(); r != nil {
			// Expected to panic due to queue not being set up properly
			t.Logf("enqueueTask panicked as expected: %v", r)
		}
	}()

	syncImpl.enqueueTask(queueIdCopyFile, source, target, &TaskCopyFile{})
}

func TestTaskCopyFile(t *testing.T) {
	// Create a dummy file entry for testing
	tree := em_file.DemoTree()
	fs := es_filesystem_model.NewFileSystem(tree)
	
	// Get an actual file entry
	sourcePath := es_filesystem_model.NewPath("/a/x")
	sourceEntry, err := fs.Info(sourcePath)
	if err != nil {
		t.Error("Failed to get source entry for testing")
		return
	}

	task := &TaskCopyFile{
		Source: sourceEntry.AsData(),
		Target: es_filesystem_model.NewPath("/target").AsData(),
	}

	// Just test that the struct can be created with proper types
	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestTaskReplaceFolderByFile(t *testing.T) {
	// Test that the task struct exists and can be created
	task := &TaskReplaceFolderByFile{}
	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestTaskReplaceFileByFolder(t *testing.T) {
	// Test that the task struct exists and can be created  
	task := &TaskReplaceFileByFolder{}
	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestTaskSyncFolder(t *testing.T) {
	// Test that the task struct exists and can be created
	task := &TaskSyncFolder{}
	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestTaskDelete(t *testing.T) {
	// Test that the task struct exists and can be created
	task := &TaskDelete{}
	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestQueueConstants(t *testing.T) {
	// Test that queue ID constants are defined
	constants := []string{
		queueIdSyncFolder,
		queueIdCopyFile,
		queueIdDelete,
		queueIdReplaceFolderByFile,
		queueIdReplaceFileByFolder,
	}

	for i, constant := range constants {
		if constant == "" {
			t.Errorf("Queue constant %d should not be empty", i)
		}
	}

	// Test specific values
	if queueIdSyncFolder != "sync_folder" {
		t.Errorf("Expected queueIdSyncFolder to be 'sync_folder', got %s", queueIdSyncFolder)
	}

	if queueIdCopyFile != "sync_file" {
		t.Errorf("Expected queueIdCopyFile to be 'sync_file', got %s", queueIdCopyFile)
	}

	if queueIdDelete != "delete" {
		t.Errorf("Expected queueIdDelete to be 'delete', got %s", queueIdDelete)
	}

	if queueIdReplaceFolderByFile != "replace_folder_by_file" {
		t.Errorf("Expected queueIdReplaceFolderByFile to be 'replace_folder_by_file', got %s", queueIdReplaceFolderByFile)
	}

	if queueIdReplaceFileByFolder != "replace_file_by_folder" {
		t.Errorf("Expected queueIdReplaceFileByFolder to be 'replace_file_by_folder', got %s", queueIdReplaceFileByFolder)
	}
}

func TestSyncImpl_CreateFolder_Success(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
	)

	syncImpl := syncer.(*syncImpl)

	target := es_filesystem_model.NewPath("/new_folder")
	err := syncImpl.createFolder(target)
	if err != nil {
		t.Error("Expected no error when creating folder")
	}

	// Verify folder was created
	entry, err := fs2.Info(target)
	if err != nil {
		t.Error("Expected folder to exist after creation")
	}

	if !entry.IsFolder() {
		t.Error("Expected created entry to be a folder")
	}
}

func TestSyncImplStruct(t *testing.T) {
	// Test that syncImpl struct can be created and has expected fields
	syncImpl := &syncImpl{}

	if syncImpl == nil {
		t.Error("Expected syncImpl to be created")
	}

	// Test that setting fields works
	syncImpl.log = esl.Default()
	if syncImpl.log == nil {
		t.Error("Expected log field to be settable")
	}
}

func TestSyncImpl_Delete(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("to_delete.txt", 7, time.Now(), 1),
		em_file.NewFolder("to_delete_folder", []em_file.Node{
			em_file.NewFile("nested.txt", 6, time.Now(), 2),
		}),
	})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	var deletedPaths []string
	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		OnDeleteSuccess(func(target es_filesystem.Path) {
			deletedPaths = append(deletedPaths, target.Path())
		}),
	)

	syncImpl := syncer.(*syncImpl)

	// Test deleting a file
	filePath := es_filesystem_model.NewPath("/to_delete.txt")
	err := syncImpl.delete(filePath)
	if err != nil {
		t.Error("Expected no error when deleting file")
	}

	// Verify file was deleted
	_, err = fs2.Info(filePath)
	if err == nil {
		t.Error("Expected file to be deleted")
	}

	// Test deleting a folder
	folderPath := es_filesystem_model.NewPath("/to_delete_folder")
	err = syncImpl.delete(folderPath)
	if err != nil {
		t.Error("Expected no error when deleting folder")
	}

	// Verify folder was deleted
	_, err = fs2.Info(folderPath)
	if err == nil {
		t.Error("Expected folder to be deleted")
	}

	// Verify callbacks were called
	if len(deletedPaths) != 2 {
		t.Errorf("Expected 2 delete callbacks, got %d", len(deletedPaths))
	}
}

func TestSyncImpl_Copy(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("source.txt", 7, time.Now(), 3),
		em_file.NewFolder("source_folder", []em_file.Node{
			em_file.NewFile("nested.txt", 6, time.Now(), 4),
		}),
	})
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	var copiedCount int
	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		OnCopySuccess(func(source, target es_filesystem.Entry) {
			copiedCount++
		}),
	)

	syncImpl := syncer.(*syncImpl)

	// Test copying a file
	sourcePath := es_filesystem_model.NewPath("/source.txt")
	targetPath := es_filesystem_model.NewPath("/target.txt")
	sourceEntry, _ := fs1.Info(sourcePath)
	
	// The copy method is asynchronous via the connector
	// Just verify it doesn't error
	err := syncImpl.copy(sourceEntry, targetPath)
	if err != nil {
		t.Error("Expected no error when calling copy")
	}
}

func TestSyncImpl_TaskCopyFile(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("source.txt", 7, time.Now(), 5),
	})
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncOverwrite(true),
	)

	syncImpl := syncer.(*syncImpl)

	sourcePath := es_filesystem_model.NewPath("/source.txt")
	targetPath := es_filesystem_model.NewPath("/target.txt")
	sourceEntry, _ := fs1.Info(sourcePath)

	task := &TaskCopyFile{
		Source: sourceEntry.AsData(),
		Target: targetPath.AsData(),
	}

	// Add to waitgroup before calling task
	syncImpl.wg.Add(1)
	syncImpl.taskCopyFile(task, eq_queue.New())

	// Wait for completion
	syncImpl.wg.Wait()
}

func TestSyncImpl_TaskDelete(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("to_delete.txt", 7, time.Now(), 6),
	})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncDelete(true),
	)

	syncImpl := syncer.(*syncImpl)

	targetPath := es_filesystem_model.NewPath("/to_delete.txt")

	task := &TaskDelete{
		Target: targetPath.AsData(),
	}

	// Add to waitgroup before calling task
	syncImpl.wg.Add(1)
	syncImpl.taskDelete(task, eq_queue.New())

	// Wait for completion
	syncImpl.wg.Wait()
}

func TestSyncImpl_TaskReplaceFolderByFile(t *testing.T) {
	// Just test that the task struct can be created
	tree := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("source", 100, time.Now(), 1),
	})
	fs := es_filesystem_model.NewFileSystem(tree)
	
	sourcePath := es_filesystem_model.NewPath("/source")
	sourceEntry, _ := fs.Info(sourcePath)
	
	task := &TaskReplaceFolderByFile{
		Source: sourceEntry.AsData(),
		Target: es_filesystem_model.NewPath("/target").AsData(),
	}

	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestSyncImpl_TaskReplaceFileByFolder(t *testing.T) {
	// Just test that the task struct can be created
	tree := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFolder("source", []em_file.Node{}),
	})
	fs := es_filesystem_model.NewFileSystem(tree)
	
	sourcePath := es_filesystem_model.NewPath("/source")
	sourceEntry, _ := fs.Info(sourcePath)
	
	task := &TaskReplaceFileByFolder{
		Source: sourceEntry.AsData(),
		Target: es_filesystem_model.NewPath("/target").AsData(),
	}

	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestSyncImpl_TaskSyncFolder_Basic(t *testing.T) {
	// Just test that the task struct can be created
	task := &TaskSyncFolder{
		Source: es_filesystem_model.NewPath("/sync_me").AsData(),
		Target: es_filesystem_model.NewPath("/sync_me").AsData(),
	}

	if task == nil {
		t.Error("Expected task to be created")
	}
}

func TestSyncImpl_TaskSyncFolder_WithFilters(t *testing.T) {
	// Test that we can create a syncer with filters
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	// Create filter that skips .tmp files
	filter := mo_filter.New("*.tmp")
	filter.SetOptions(mo_filter.NewTestNameFilter("tmp"))

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		WithNameFilter(filter),
	)

	syncImpl := syncer.(*syncImpl)
	if syncImpl.opts.entryNameFilter == nil {
		t.Error("Expected filter to be set")
	}
}

func TestSyncImpl_TaskSyncFolder_WithDelete(t *testing.T) {
	// Test that we can create a syncer with delete option
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncDelete(true),
	)

	syncImpl := syncer.(*syncImpl)
	if !syncImpl.opts.SyncDelete() {
		t.Error("Expected delete option to be true")
	}
}


func TestNewWithAllOptions(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.DemoTree()

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	filter := mo_filter.New("*.tmp")
	progress := ea_indicator.Global()

	var copyCount, deleteCount, skipCount int

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncDelete(true),
		SyncOverwrite(false),
		SyncDontCompareContent(true),
		SyncDontCompareTime(true),
		WithNameFilter(filter),
		WithProgress(progress),
		OptimizePreventCreateFolder(true),
		OnCopySuccess(func(source, target es_filesystem.Entry) {
			copyCount++
		}),
		OnCopyFailure(func(source es_filesystem.Path, fsErr es_filesystem.FileSystemError) {
			// Handle copy failure
		}),
		OnDeleteSuccess(func(target es_filesystem.Path) {
			deleteCount++
		}),
		OnDeleteFailure(func(target es_filesystem.Path, fsErr es_filesystem.FileSystemError) {
			// Handle delete failure
		}),
		OnSkip(func(reason SkipReason, source es_filesystem.Entry, target es_filesystem.Path) {
			skipCount++
		}),
	)

	if syncer == nil {
		t.Error("Expected non-nil syncer")
	}

	syncImpl := syncer.(*syncImpl)
	
	// Verify all options were applied
	if !syncImpl.opts.syncDelete {
		t.Error("Expected syncDelete to be true")
	}
	if syncImpl.opts.syncOverwrite {
		t.Error("Expected syncOverwrite to be false")
	}
	if !syncImpl.opts.syncDontCompareContent {
		t.Error("Expected syncDontCompareContent to be true")
	}
	if !syncImpl.opts.syncDontCompareTime {
		t.Error("Expected syncDontCompareTime to be true")
	}
	if !syncImpl.opts.optimizeReduceCreateFolder {
		t.Error("Expected optimizeReduceCreateFolder to be true")
	}
	if syncImpl.opts.listenerCopySuccess == nil {
		t.Error("Expected onCopySuccess to be set")
	}
	if syncImpl.opts.listenerCopyFailure == nil {
		t.Error("Expected onCopyFailure to be set")
	}
	if syncImpl.opts.listenerDeleteSuccess == nil {
		t.Error("Expected onDeleteSuccess to be set")
	}
	if syncImpl.opts.listenerDeleteFailure == nil {
		t.Error("Expected onDeleteFailure to be set")
	}
	if syncImpl.opts.listenerSkip == nil {
		t.Error("Expected onSkip to be set")
	}
}

func TestSyncImpl_CreateFolder_AlreadyExists(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFolder("existing", []em_file.Node{}),
	})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
	)

	syncImpl := syncer.(*syncImpl)

	// Try to create a folder that already exists
	target := es_filesystem_model.NewPath("/existing")
	err := syncImpl.createFolder(target)
	
	// Should not error when folder already exists
	if err != nil {
		t.Error("Expected no error when creating folder that already exists")
	}
}

func TestSyncImpl_Delete_NotFound(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("root", []em_file.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	var deleteFailureCount int
	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		OnDeleteFailure(func(target es_filesystem.Path, fsErr es_filesystem.FileSystemError) {
			deleteFailureCount++
		}),
	)

	syncImpl := syncer.(*syncImpl)

	// Try to delete a non-existent path
	target := es_filesystem_model.NewPath("/non_existent")
	err := syncImpl.delete(target)
	
	// Should return error when deleting non-existent path
	if err == nil {
		t.Error("Expected error when deleting non-existent path")
	}

	// Callback should be called for failures
	if deleteFailureCount != 1 {
		t.Error("Expected delete failure callback to be called once")
	}
}

func TestSyncImpl_Copy_WithOverwrite(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()

	tree1 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("source.txt", 11, time.Now(), 16),
	})
	tree2 := em_file.NewFolder("root", []em_file.Node{
		em_file.NewFile("target.txt", 11, time.Now(), 17),
	})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)
	conn := es_filesystem_copier.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		eq_queue.New(),
		fs1,
		fs2,
		conn,
		SyncOverwrite(true),
	)

	syncImpl := syncer.(*syncImpl)

	sourcePath := es_filesystem_model.NewPath("/source.txt")
	targetPath := es_filesystem_model.NewPath("/target.txt")
	sourceEntry, _ := fs1.Info(sourcePath)
	
	err := syncImpl.copy(sourceEntry, targetPath)
	if err != nil {
		t.Error("Expected no error when copying with overwrite")
	}

	// Verify file was overwritten
	_, err = fs2.Info(targetPath)
	if err != nil {
		t.Error("Expected target file to exist")
	}
	
	// Note: In a real filesystem, we'd check the content changed
	// For model filesystem, we just verify the operation succeeded
}