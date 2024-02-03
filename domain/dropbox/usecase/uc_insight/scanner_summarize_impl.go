package uc_insight

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

func (z tsImpl) defineSummarizeQueues(s eq_sequence.Stage) {
	s.Define(teamSummarizeFolderImmediate, z.summarizeFolderImmediateCount)
	s.Define(teamSummarizeFolderPath, z.summarizeFolderPaths)
	s.Define(teamSummarizeFolderRecursive, z.summarizeFolderRecursive)
	s.Define(teamSummarizeNamespace, z.summarizeNamespace)
	s.Define(teamSummarizeEntry, z.summarizeEntry)
	s.Define(teamSummarizeTeamFolder, z.summarizeTeamFolder, s)
	s.Define(teamSummarizeTeamFolderEntry, z.summarizeTeamFolderEntry, s)
}

// Stage1: summarize namespaces
func (z tsImpl) summarizeStage1() error {
	l := z.ctl.Log()
	var lastErr error

	namespaceModel := &Namespace{}
	namespaceRows, err := z.adb.Model(namespaceModel).Distinct("namespace_id").Select("namespace_id").Rows()
	if err != nil {
		l.Debug("Unable to get namespace rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = namespaceRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qNamespace := s.Get(teamSummarizeNamespace)
		for namespaceRows.Next() {
			namespaceModel = &Namespace{}
			if err := z.adb.ScanRows(namespaceRows, namespaceModel); err != nil {
				l.Debug("Unable to scan namespace row", esl.Error(err))
				lastErr = err
				return
			}
			qNamespace.Enqueue(namespaceModel.NamespaceId)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))
	_ = namespaceRows.Close()

	if lastErr != nil {
		l.Debug("Failure on processing namespace rows", esl.Error(lastErr))
		return lastErr
	}

	return nil
}

// Stage2: summarize folders
func (z tsImpl) summarizeStage2() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
	if err != nil {
		l.Debug("Unable to get folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = folderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qFolderPath := s.Get(teamSummarizeFolderPath)
		qFolderImmediate := s.Get(teamSummarizeFolderImmediate)
		for folderRows.Next() {
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qFolderPath.Enqueue(folderEntry.FileId)
			qFolderImmediate.Enqueue(folderEntry.FileId)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}

	return nil
}

// Stage3: summarize files
func (z tsImpl) summarizeStage3() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
	if err != nil {
		l.Debug("Unable to get folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = folderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qFolderRecursive := s.Get(teamSummarizeFolderRecursive)
		for folderRows.Next() {
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qFolderRecursive.Enqueue(folderEntry.FileId)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}
	return nil
}

// Stage4: summarize entries
func (z tsImpl) summarizeStage4() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
	if err != nil {
		l.Debug("Unable to get folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = folderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qEntry := s.Get(teamSummarizeEntry)
		for folderRows.Next() {
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qEntry.Enqueue(folderEntry.FileId)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}
	return nil
}

// Stage5: summarize team folder child entries
func (z tsImpl) summarizeStage5() error {
	l := z.ctl.Log()
	var lastErr error

	teamFolder := &TeamFolder{}
	teamFolderRows, err := z.adb.Model(teamFolder).Rows()
	if err != nil {
		l.Debug("Unable to get team folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = teamFolderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qEntry := s.Get(teamSummarizeTeamFolder)
		for teamFolderRows.Next() {
			teamFolder := &TeamFolder{}
			if err := z.adb.ScanRows(teamFolderRows, teamFolder); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qEntry.Enqueue(teamFolder)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}
	return nil
}

func (z tsImpl) Summarize() error {
	l := z.ctl.Log()

	summaryTables := []interface{}{
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
		&SummaryTeamFolderEntry{},
	}
	for _, st := range summaryTables {
		z.adb.Delete(st)
	}

	if err := z.summarizeStage1(); err != nil {
		l.Debug("Stage 1 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage2(); err != nil {
		l.Debug("Stage 2 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage3(); err != nil {
		l.Debug("Stage 3 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage4(); err != nil {
		l.Debug("Stage 4 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage5(); err != nil {
		l.Debug("Stage 5 failed", esl.Error(err))
		return err
	}

	db, err := z.adb.DB()
	if err != nil {
		l.Debug("Unable to get DB", esl.Error(err))
		return err
	}
	_ = db.Close()

	return nil
}
