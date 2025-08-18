package dc_supplemental

import (
	"testing"

	"github.com/watermint/toolbox/infra/doc/dc_index"
)

// Test basic document functionality
func TestDocuments(t *testing.T) {
	// Test PathVariable
	pv := &PathVariable{}
	pv.DocId()
	pv.DocDesc()
	pv.Sections()

	// Test ExperimentalFeature
	ef := &ExperimentalFeature{}
	ef.DocId()
	ef.DocDesc()
	ef.Sections()

	// Test Troubleshooting
	ts := &Troubleshooting{}
	ts.DocId()
	ts.DocDesc()
	ts.Sections()

	// Test ReportingOptions
	ro := &ReportingOptions{}
	ro.DocId()
	ro.DocDesc()
	ro.Sections()

	// Test AuthenticationGuide
	ag := &AuthenticationGuide{}
	ag.DocId()
	ag.DocDesc()
	ag.Sections()

	// Test ErrorHandlingGuide
	eg := &ErrorHandlingGuide{}
	eg.DocId()
	eg.DocDesc()
	eg.Sections()

	// Test BestPracticesGuide
	bg := &BestPracticesGuide{}
	bg.DocId()
	bg.DocDesc()
	bg.Sections()

	// Test ReportingGuide
	rg := &ReportingGuide{}
	rg.DocId()
	rg.DocDesc()
	rg.Sections()
}

// Test factory methods
func TestFactories(t *testing.T) {
	// Test NewDocSpecChange
	NewDocSpecChange()

	// Test NewDropboxBusiness
	NewDropboxBusiness(dc_index.MediaRepository)
	NewDropboxBusiness(dc_index.MediaWeb)
	NewDropboxBusiness(dc_index.MediaKnowledge)
}

// Test Docs function
func TestDocsFunc(t *testing.T) {
	Docs(dc_index.MediaRepository)
	Docs(dc_index.MediaWeb)
	Docs(dc_index.MediaKnowledge)
}

// Test section definitions
func TestSections(t *testing.T) {
	// Test PathVariableDefinitions
	pvd := &PathVariableDefinitions{}
	pvd.Title()

	// Test ExperimentalFeatureDefinitions
	efd := &ExperimentalFeatureDefinitions{}
	efd.Title()

	// Test some auth sections
	aos := &AuthOverviewSection{}
	aos.Title()

	das := &DropboxAuthSection{}
	das.Title()

	tms := &TokenManagementSection{}
	tms.Title()

	ats := &AuthTroubleshootingSection{}
	ats.Title()

	sts := &SecurityTipsSection{}
	sts.Title()

	// Test some error sections
	ces := &CommonErrorsSection{}
	ces.Title()

	nes := &NetworkErrorsSection{}
	nes.Title()

	aes := &AuthenticationErrorsSection{}
	aes.Title()

	fes := &FileSystemErrorsSection{}
	fes.Title()

	rle := &RateLimitErrorsSection{}
	rle.Title()

	apis := &APIErrorsSection{}
	apis.Title()

	dts := &DebugTechniquesSection{}
	dts.Title()
}
