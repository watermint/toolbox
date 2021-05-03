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

	DocWebHome
	DocWebLicense
	DocWebCommandTableOfContent

	// -- Documents under `doc/generated(_\w{2})?`

	// Changes from the previous release
	DocManualChanges
	// Individual manual for each commands
	DocManualCommand

	// Supplemental
	DocSupplementalPathVariables
	DocSupplementalExperimentalFeature
	DocSupplementalTroubleshooting
	DocSupplementalDropboxBusiness
)

const (
	MediaRepository MediaType = iota
	MediaWeb
)

type MediaType int

const (
	WebCategoryHome WebCategory = iota
	WebCategoryCommand
	WebCategoryGuide
)

type WebCategory int

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

	WebDocs = []DocId{
		DocRootBuild,
		DocRootContributing,
		DocRootCodeOfConduct,
		DocRootSecurityAndPrivacy,
		DocWebHome,
		DocWebLicense,
		DocManualCommand,
		DocSupplementalPathVariables,
		DocSupplementalExperimentalFeature,
		DocSupplementalTroubleshooting,
		DocSupplementalDropboxBusiness,
	}

	AllMedia = []MediaType{
		MediaRepository,
		MediaWeb,
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
	RefPath     bool   `json:"ref_path"`
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

func CommandName(name string) NameOpt {
	return func(opts NameOpts) NameOpts {
		opts.CommandName = name
		return opts
	}
}
func RefPath(enable bool) NameOpt {
	return func(opts NameOpts) NameOpts {
		opts.RefPath = enable
		return opts
	}
}

const (
	WebDocPathRoot = "docs/"
)

func WebDocPath(refPath bool, cat WebCategory, name string, lg lang.Lang) string {
	basePath := WebDocPathRoot
	suffix := ".md"
	if name == "" {
		suffix = ""
	}
	if refPath {
		basePath = "{{ site.baseurl }}/"
		if name != "" {
			suffix = ".html"
		}
	}
	pathLang := ""
	if !lg.IsDefault() {
		pathLang = lg.String() + "/"
	}
	switch cat {
	case WebCategoryHome:
		return basePath + pathLang + name + suffix
	case WebCategoryCommand:
		return basePath + pathLang + "commands/" + name + suffix
	case WebCategoryGuide:
		return basePath + pathLang + "guides/" + name + suffix
	}

	esl.Default().Warn("Invalid web category id", esl.Int("category", int(cat)))
	panic("invalid category")
}

// DocName Document path and name (without extension)
func DocName(media MediaType, id DocId, lg lang.Lang, opts ...NameOpt) string {
	nameOpts := NameOpts{}.Apply(opts)

	switch media {
	case MediaRepository:
		switch id {
		case DocRootReadme:
			return "README" + lg.Suffix() + ".md"
		case DocRootLicense:
			return "LICENSE" + lg.Suffix() + ".md"
		case DocRootBuild:
			return "BUILD" + lg.Suffix() + ".md"
		case DocRootContributing:
			return "CONTRIBUTING" + lg.Suffix() + ".md"
		case DocRootCodeOfConduct:
			return "CODE_OF_CONDUCT" + lg.Suffix() + ".md"
		case DocRootSecurityAndPrivacy:
			return "SECURITY_AND_PRIVACY" + lg.Suffix() + ".md"
		case DocManualChanges:
			return GeneratedPath(lg, "changes") + ".md"
		case DocManualCommand:
			if nameOpts.CommandName != "" {
				return GeneratedPath(lg, nameOpts.CommandName) + ".md"
			} else {
				return GeneratedPath(lg, "")
			}
		case DocSupplementalPathVariables:
			return SupplementalDocPath(lg, "path_variables") + ".md"
		case DocSupplementalExperimentalFeature:
			return SupplementalDocPath(lg, "experimental_features") + ".md"
		case DocSupplementalTroubleshooting:
			return SupplementalDocPath(lg, "troubleshooting") + ".md"
		case DocSupplementalDropboxBusiness:
			return SupplementalDocPath(lg, "dropbox_business") + ".md"
		}

	case MediaWeb:
		switch id {
		case DocRootBuild:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "build", lg)
		case DocRootContributing:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "contributing", lg)
		case DocRootCodeOfConduct:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "code_of_conduct", lg)
		case DocRootSecurityAndPrivacy:
			return WebDocPath(nameOpts.RefPath, WebCategoryHome, "security_and_privacy", lg)
		case DocWebHome:
			return WebDocPath(nameOpts.RefPath, WebCategoryHome, "home", lg)
		case DocWebLicense:
			return WebDocPath(nameOpts.RefPath, WebCategoryHome, "license", lg)
		case DocWebCommandTableOfContent:
			return WebDocPath(nameOpts.RefPath, WebCategoryCommand, "toc", lg)
		case DocManualCommand:
			return WebDocPath(nameOpts.RefPath, WebCategoryCommand, nameOpts.CommandName, lg)
		case DocSupplementalPathVariables:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "path-variables", lg)
		case DocSupplementalExperimentalFeature:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "experimental-features", lg)
		case DocSupplementalTroubleshooting:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "troubleshooting", lg)
		case DocSupplementalDropboxBusiness:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "dropbox-business", lg)
		}
	}

	esl.Default().Warn("Invalid document id", esl.Int("mediaType", int(media)), esl.Int("documentId", int(id)))
	panic("invalid document id")
}
