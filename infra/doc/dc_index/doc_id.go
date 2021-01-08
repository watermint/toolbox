package dc_index

import (
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
)

const (
	DocRootReadme DocId = iota
	DocRootLicense
	DocRootBuild
	DocRootContributing
	DocRootCodeOfConduct
	DocRootSecurityAndPrivacy

	// -- Documents under `doc/generated(_\w{2})?`

	// Changes from the previous release
	DocManualChanges
	// Individual manual for each commands
	DocManualCommand
	// Supplemental:
	DocSupplementalPathVariables
	DocSupplementalExperimentalFeature
	DocSupplementalTroubleshooting
	DocSupplementalDropboxBusiness
)

const (
	GeneratedDocPath = "doc/generated"
)

type DocId int

var (
	RootDocs = []DocId{
		DocRootReadme,
		DocRootLicense,
		DocRootBuild,
		DocRootContributing,
		DocRootCodeOfConduct,
		DocRootSecurityAndPrivacy,
	}

	GeneratedDocs = []DocId{
		DocManualChanges,
		DocManualCommand,
		DocSupplementalPathVariables,
		DocSupplementalExperimentalFeature,
		DocSupplementalTroubleshooting,
		DocSupplementalDropboxBusiness,
	}
)

func GeneratedPath(l lang.Lang, name string) string {
	return GeneratedDocPath + l.Suffix() + "/" + name
}

func SupplementalDocPath(l lang.Lang, name string) string {
	return GeneratedPath(l, "supplemental_"+name)
}

type NameOpts struct {
	CommandName string `json:"command_name"`
}

func (z NameOpts) Apply(opts []NameOpt) NameOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type NameOpt func(opts NameOpts) NameOpts

// Document path and name (without extension)
func DocName(id DocId, lg lang.Lang, opts ...NameOpt) string {
	nameOpts := NameOpts{}.Apply(opts)

	switch id {
	case DocRootReadme:
		return "README" + lg.Suffix()
	case DocRootLicense:
		return "LICENSE" + lg.Suffix()
	case DocRootBuild:
		return "BUILD" + lg.Suffix()
	case DocRootContributing:
		return "CONTRIBUTING" + lg.Suffix()
	case DocRootCodeOfConduct:
		return "CODE_OF_CONDUCT" + lg.Suffix()
	case DocRootSecurityAndPrivacy:
		return "SECURITY_AND_PRIVACY" + lg.Suffix()
	case DocManualChanges:
		return GeneratedPath(lg, "changes")
	case DocManualCommand:
		return GeneratedPath(lg, nameOpts.CommandName)
	case DocSupplementalPathVariables:
		return SupplementalDocPath(lg, "path_variables")
	case DocSupplementalExperimentalFeature:
		return SupplementalDocPath(lg, "experimental_features")
	case DocSupplementalTroubleshooting:
		return SupplementalDocPath(lg, "troubleshooting")
	case DocSupplementalDropboxBusiness:
		return SupplementalDocPath(lg, "dropbox_business")
	}
	esl.Default().Warn("Invalid document id", esl.Int("documentId", int(id)))
	panic("invalid document id")
}