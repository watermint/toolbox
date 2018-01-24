package api

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/paper"
	"io"
)

type ApiPaper struct {
	Context *ApiContext
}

func (a *ApiPaper) Compat() paper.Client {
	return paper.New(a.Context.compatConfig())
}

func (a *ApiPaper) DocsArchive(arg *paper.RefPaperDoc) (err error) {
	return a.Compat().DocsArchive(arg)
}
func (a *ApiPaper) DocsCreate(arg *paper.PaperDocCreateArgs, content io.Reader) (res *paper.PaperDocCreateUpdateResult, err error) {
	return a.Compat().DocsCreate(arg, content)
}
func (a *ApiPaper) DocsDownload(arg *paper.PaperDocExport) (res *paper.PaperDocExportResult, content io.ReadCloser, err error) {
	return a.Compat().DocsDownload(arg)
}
func (a *ApiPaper) DocsFolderUsersList(arg *paper.ListUsersOnFolderArgs) (res *paper.ListUsersOnFolderResponse, err error) {
	return a.Compat().DocsFolderUsersList(arg)
}
func (a *ApiPaper) DocsFolderUsersListContinue(arg *paper.ListUsersOnFolderContinueArgs) (res *paper.ListUsersOnFolderResponse, err error) {
	return a.Compat().DocsFolderUsersListContinue(arg)
}
func (a *ApiPaper) DocsGetFolderInfo(arg *paper.RefPaperDoc) (res *paper.FoldersContainingPaperDoc, err error) {
	return a.Compat().DocsGetFolderInfo(arg)
}
func (a *ApiPaper) DocsList(arg *paper.ListPaperDocsArgs) (res *paper.ListPaperDocsResponse, err error) {
	return a.Compat().DocsList(arg)
}
func (a *ApiPaper) DocsListContinue(arg *paper.ListPaperDocsContinueArgs) (res *paper.ListPaperDocsResponse, err error) {
	return a.Compat().DocsListContinue(arg)
}
func (a *ApiPaper) DocsPermanentlyDelete(arg *paper.RefPaperDoc) (err error) {
	return a.Compat().DocsPermanentlyDelete(arg)
}
func (a *ApiPaper) DocsSharingPolicyGet(arg *paper.RefPaperDoc) (res *paper.SharingPolicy, err error) {
	return a.Compat().DocsSharingPolicyGet(arg)
}
func (a *ApiPaper) DocsSharingPolicySet(arg *paper.PaperDocSharingPolicy) (err error) {
	return a.Compat().DocsSharingPolicySet(arg)
}
func (a *ApiPaper) DocsUpdate(arg *paper.PaperDocUpdateArgs, content io.Reader) (res *paper.PaperDocCreateUpdateResult, err error) {
	return a.Compat().DocsUpdate(arg, content)
}
func (a *ApiPaper) DocsUsersAdd(arg *paper.AddPaperDocUser) (res []*paper.AddPaperDocUserMemberResult, err error) {
	return a.Compat().DocsUsersAdd(arg)
}
func (a *ApiPaper) DocsUsersList(arg *paper.ListUsersOnPaperDocArgs) (res *paper.ListUsersOnPaperDocResponse, err error) {
	return a.Compat().DocsUsersList(arg)
}
func (a *ApiPaper) DocsUsersListContinue(arg *paper.ListUsersOnPaperDocContinueArgs) (res *paper.ListUsersOnPaperDocResponse, err error) {
	return a.Compat().DocsUsersListContinue(arg)
}
func (a *ApiPaper) DocsUsersRemove(arg *paper.RemovePaperDocUser) (err error) {
	return a.Compat().DocsUsersRemove(arg)
}
