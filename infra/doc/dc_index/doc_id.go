package dc_index

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_lang"
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
	DocWebKnowledge

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
	DocSupplementalSpecChange
	DoCSupplementalReportingOptions

	// Contributor documents
	DocContributorRecipeValues
)

const (
	MediaRepository MediaType = iota
	MediaWeb
	MediaKnowledge
)

type MediaType int

const (
	WebCategoryHome WebCategory = iota
	WebCategoryCommand
	WebCategoryGuide
	WebCategoryKnowledge
	WebCategoryContributor
)

type WebCategory int

const (
	GeneratedDocPath = "docs/generated"
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
		DocSupplementalSpecChange,
		DocContributorRecipeValues,
	}

	WebDocs = []DocId{
		DocRootBuild,
		DocRootContributing,
		DocRootCodeOfConduct,
		DocRootSecurityAndPrivacy,
		DocWebHome,
		DocWebLicense,
		DocWebKnowledge,
		DocManualCommand,
		DocSupplementalPathVariables,
		DocSupplementalExperimentalFeature,
		DocSupplementalTroubleshooting,
		DocSupplementalDropboxBusiness,
		DocSupplementalSpecChange,
		DocContributorRecipeValues,
	}

	AllMedia = []MediaType{
		MediaRepository,
		MediaWeb,
	}
)

func GeneratedPath(l es_lang.Lang, name string) string {
	return GeneratedDocPath + l.Suffix() + "/" + name
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

func WebDocPath(refPath bool, cat WebCategory, name string, lg es_lang.Lang) string {
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
	case WebCategoryKnowledge:
		return basePath + pathLang + "knowledge/" + name + suffix
	case WebCategoryContributor:
		return basePath + pathLang + "contributor/" + name + suffix
	}

	esl.Default().Warn("Invalid web category id", esl.Int("category", int(cat)))
	panic("invalid category")
}

// DocName Document path and name (without extension)
func DocName(media MediaType, id DocId, lg es_lang.Lang, opts ...NameOpt) string {
	nameOpts := NameOpts{}.Apply(opts)
	l := esl.Default()

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
			return WebDocPath(nameOpts.RefPath, WebCategoryCommand, nameOpts.CommandName, lg)
		case DocSupplementalPathVariables:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "path-variables", lg)
		case DocSupplementalExperimentalFeature:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "experimental-features", lg)
		case DocSupplementalTroubleshooting:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "troubleshooting", lg)
		case DocSupplementalDropboxBusiness:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "dropbox-business", lg)
		case DocSupplementalSpecChange:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "spec-change", lg)
		case DocContributorRecipeValues:
			return WebDocPath(nameOpts.RefPath, WebCategoryContributor, "recipe_values", lg)
		default:
			l.Warn("Invalid document id", esl.Int("mediaType", int(media)), esl.Int("documentId", int(id)))
			panic(fmt.Sprintf("Invalid document id: %d", id))
		}

	case MediaWeb, MediaKnowledge:
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
		case DocWebKnowledge:
			return WebDocPath(nameOpts.RefPath, WebCategoryKnowledge, "knowledge", lg)
		case DocSupplementalPathVariables:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "path-variables", lg)
		case DocSupplementalExperimentalFeature:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "experimental-features", lg)
		case DocSupplementalTroubleshooting:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "troubleshooting", lg)
		case DocSupplementalDropboxBusiness:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "dropbox-business", lg)
		case DoCSupplementalReportingOptions:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "reporting-options", lg)
		case DocSupplementalSpecChange:
			return WebDocPath(nameOpts.RefPath, WebCategoryGuide, "spec-change", lg)
		case DocContributorRecipeValues:
			return WebDocPath(nameOpts.RefPath, WebCategoryContributor, "recipe_values", lg)
		default:
			l.Warn("Invalid document id", esl.Int("mediaType", int(media)), esl.Int("documentId", int(id)))
			panic(fmt.Sprintf("Invalid document id: %d", id))
		}
	}

	esl.Default().Warn("Invalid document id", esl.Int("mediaType", int(media)), esl.Int("documentId", int(id)))
	panic("invalid document id")
}
