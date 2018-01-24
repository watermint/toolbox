package api

import "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"

type ApiUsers struct {
	Context *ApiContext
}

func (a *ApiUsers) Compat() users.Client {
	return users.New(a.Context.compatConfig())
}

func (a *ApiUsers) GetAccount(arg *users.GetAccountArg) (res *users.BasicAccount, err error) {
	return a.Compat().GetAccount(arg)
}
func (a *ApiUsers) GetAccountBatch(arg *users.GetAccountBatchArg) (res []*users.BasicAccount, err error) {
	return a.Compat().GetAccountBatch(arg)
}
func (a *ApiUsers) GetCurrentAccount() (res *users.FullAccount, err error) {
	return a.Compat().GetCurrentAccount()
}
func (a *ApiUsers) GetSpaceUsage() (res *users.SpaceUsage, err error) {
	return a.Compat().GetSpaceUsage()
}
