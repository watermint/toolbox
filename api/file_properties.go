package api

import "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_properties"

type ApiFileProperties struct {
	Context *ApiContext
}

func (a *ApiFileProperties) Compat() file_properties.Client {
	return file_properties.New(a.Context.compatConfig())
}

func (a *ApiFileProperties) PropertiesAdd(arg *file_properties.AddPropertiesArg) (err error) {
	return a.Compat().PropertiesAdd(arg)
}
func (a *ApiFileProperties) PropertiesOverwrite(arg *file_properties.OverwritePropertyGroupArg) (err error) {
	return a.Compat().PropertiesOverwrite(arg)
}
func (a *ApiFileProperties) PropertiesRemove(arg *file_properties.RemovePropertiesArg) (err error) {
	return a.Compat().PropertiesRemove(arg)
}
func (a *ApiFileProperties) PropertiesSearch(arg *file_properties.PropertiesSearchArg) (res *file_properties.PropertiesSearchResult, err error) {
	return a.Compat().PropertiesSearch(arg)
}
func (a *ApiFileProperties) PropertiesSearchContinue(arg *file_properties.PropertiesSearchContinueArg) (res *file_properties.PropertiesSearchResult, err error) {
	return a.Compat().PropertiesSearchContinue(arg)
}
func (a *ApiFileProperties) PropertiesUpdate(arg *file_properties.UpdatePropertiesArg) (err error) {
	return a.Compat().PropertiesUpdate(arg)
}
func (a *ApiFileProperties) TemplatesAddForTeam(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().TemplatesAddForTeam(arg)
}
func (a *ApiFileProperties) TemplatesAddForUser(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().TemplatesAddForUser(arg)
}
func (a *ApiFileProperties) TemplatesGetForTeam(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().TemplatesGetForTeam(arg)
}
func (a *ApiFileProperties) TemplatesGetForUser(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().TemplatesGetForUser(arg)
}
func (a *ApiFileProperties) TemplatesListForTeam() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().TemplatesListForTeam()
}
func (a *ApiFileProperties) TemplatesListForUser() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().TemplatesListForUser()
}
func (a *ApiFileProperties) TemplatesRemoveForTeam(arg *file_properties.RemoveTemplateArg) (err error) {
	return a.Compat().TemplatesRemoveForTeam(arg)
}
func (a *ApiFileProperties) TemplatesRemoveForUser(arg *file_properties.RemoveTemplateArg) (err error) {
	return a.Compat().TemplatesRemoveForUser(arg)
}
func (a *ApiFileProperties) TemplatesUpdateForTeam(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().TemplatesUpdateForTeam(arg)
}
func (a *ApiFileProperties) TemplatesUpdateForUser(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().TemplatesUpdateForUser(arg)
}
