package api

import "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_requests"

type ApiFileRequests struct {
	Context *ApiContext
}

func (a *ApiFileRequests) Compat() file_requests.Client {
	return file_requests.New(a.Context.compatConfig())
}

func (a *ApiFileRequests) Create(arg *file_requests.CreateFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Create(arg)
}
func (a *ApiFileRequests) Get(arg *file_requests.GetFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Get(arg)
}
func (a *ApiFileRequests) List() (res *file_requests.ListFileRequestsResult, err error) {
	return a.Compat().List()
}
func (a *ApiFileRequests) Update(arg *file_requests.UpdateFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Update(arg)
}
