package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	domaindeeplapideepl_conn_impl "github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl"
	domaindropboxapidbx_conn_impl "github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	domaindropboxapidbx_error "github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	domaindropboxapidbx_list_impl "github.com/watermint/toolbox/domain/dropbox/api/dbx_list_impl"
	domaindropboxfilesystemdbx_fs "github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	domaindropboxmodelmo_file_filter "github.com/watermint/toolbox/domain/dropbox/model/mo_file_filter"
	domaindropboxmodelmo_sharedfolder_member "github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	domaindropboxusecaseuc_compare_local "github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_local"
	domaindropboxusecaseuc_compare_paths "github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_paths"
	domaindropboxusecaseuc_file_merge "github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_merge"
	domaindropboxusecaseuc_file_relocation "github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_relocation"
	domaindropboxsignapihs_conn_impl "github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl"
	domainfigmaservicesv_file "github.com/watermint/toolbox/domain/figma/service/sv_file"
	domainfigmaservicesv_project "github.com/watermint/toolbox/domain/figma/service/sv_project"
	domaingooglemailservicesv_label "github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	domaingooglemailservicesv_message "github.com/watermint/toolbox/domain/google/mail/service/sv_message"
	essentialsapiapi_auth_basic "github.com/watermint/toolbox/essentials/api/api_auth_basic"
	essentialsapiapi_auth_key "github.com/watermint/toolbox/essentials/api/api_auth_key"
	essentialsapiapi_auth_oauth "github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	essentialsapiapi_callback "github.com/watermint/toolbox/essentials/api/api_callback"
	essentialsapiapi_doc "github.com/watermint/toolbox/essentials/api/api_doc"
	essentialslogesl_rotate "github.com/watermint/toolbox/essentials/log/esl_rotate"
	essentialsmodelmo_filter "github.com/watermint/toolbox/essentials/model/mo_filter"
	essentialsnetworknw_diag "github.com/watermint/toolbox/essentials/network/nw_diag"
	infracontrolapp_error "github.com/watermint/toolbox/infra/control/app_error"
	infracontrolapp_job_impl "github.com/watermint/toolbox/infra/control/app_job_impl"
	infradatada_griddata "github.com/watermint/toolbox/infra/data/da_griddata"
	infradatada_json "github.com/watermint/toolbox/infra/data/da_json"
	infradocdc_command "github.com/watermint/toolbox/infra/doc/dc_command"
	infradocdc_contributor "github.com/watermint/toolbox/infra/doc/dc_contributor"
	infradocdc_license "github.com/watermint/toolbox/infra/doc/dc_license"
	infradocdc_options "github.com/watermint/toolbox/infra/doc/dc_options"
	infradocdc_supplemental "github.com/watermint/toolbox/infra/doc/dc_supplemental"
	infrafeedfd_file_impl "github.com/watermint/toolbox/infra/feed/fd_file_impl"
	infrareciperc_exec "github.com/watermint/toolbox/infra/recipe/rc_exec"
	infrareciperc_group "github.com/watermint/toolbox/infra/recipe/rc_group"
	infrareciperc_group_impl "github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	infrareciperc_spec "github.com/watermint/toolbox/infra/recipe/rc_spec"
	infrareciperc_value "github.com/watermint/toolbox/infra/recipe/rc_value"
	infrareportrp_model_impl "github.com/watermint/toolbox/infra/report/rp_model_impl"
	infrareportrp_writer_impl "github.com/watermint/toolbox/infra/report/rp_writer_impl"
	infrauiapp_ui "github.com/watermint/toolbox/infra/ui/app_ui"
	ingredientig_file "github.com/watermint/toolbox/ingredient/ig_file"
	recipedevdiag "github.com/watermint/toolbox/recipe/dev/diag"
	recipefileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	recipegroupmember "github.com/watermint/toolbox/recipe/group/member"
	recipegroupmemberbatch "github.com/watermint/toolbox/recipe/group/member/batch"
	recipemember "github.com/watermint/toolbox/recipe/member"
	recipememberupdate "github.com/watermint/toolbox/recipe/member/update"
	recipeteamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	recipeutiltidymove "github.com/watermint/toolbox/recipe/util/tidy/move"
)

func AutoDetectedMessageObjects() []interface{} {
	return []interface{}{
		&domaindeeplapideepl_conn_impl.MsgDeeplApi{},
		&domaindropboxapidbx_conn_impl.MsgConnect{},
		&domaindropboxapidbx_error.MsgHandler{},
		&domaindropboxapidbx_list_impl.MsgList{},
		&domaindropboxfilesystemdbx_fs.MsgFileSystemCached{},
		&domaindropboxmodelmo_file_filter.MsgFileFilterOpt{},
		&domaindropboxmodelmo_sharedfolder_member.MsgExternalOpt{},
		&domaindropboxmodelmo_sharedfolder_member.MsgInternalOpt{},
		&domaindropboxusecaseuc_compare_local.MsgCompare{},
		&domaindropboxusecaseuc_compare_paths.MsgCompare{},
		&domaindropboxusecaseuc_file_merge.MsgMerge{},
		&domaindropboxusecaseuc_file_relocation.MsgRelocation{},
		&domaindropboxsignapihs_conn_impl.MsgDropboxSign{},
		&domainfigmaservicesv_file.MsgFile{},
		&domainfigmaservicesv_project.MsgProject{},
		&domaingooglemailservicesv_label.MsgFindLabel{},
		&domaingooglemailservicesv_message.MsgProgress{},
		&essentialsapiapi_auth_basic.MsgConsole{},
		&essentialsapiapi_auth_key.MsgConsole{},
		&essentialsapiapi_auth_oauth.MsgApiAuth{},
		&essentialsapiapi_callback.MsgCallback{},
		&essentialsapiapi_doc.MsgApiDoc{},
		&essentialslogesl_rotate.MsgOut{},
		&essentialslogesl_rotate.MsgPurge{},
		&essentialslogesl_rotate.MsgRotate{},
		&essentialsmodelmo_filter.MsgFilter{},
		&essentialsnetworknw_diag.MsgNetwork{},
		&infracontrolapp_error.MsgErrorReport{},
		&infracontrolapp_job_impl.MsgLauncher{},
		&infradatada_griddata.MsgGridDataInput{},
		&infradatada_json.MsgJsonInput{},
		&infradocdc_command.MsgHeader{},
		&infradocdc_contributor.MsgDeveloper{},
		&infradocdc_license.MsgLicense{},
		&infradocdc_options.MsgDoc{},
		&infradocdc_supplemental.MsgDropboxBusiness{},
		&infradocdc_supplemental.MsgExperimentalFeature{},
		&infradocdc_supplemental.MsgPathVariable{},
		&infradocdc_supplemental.MsgTroubleshooting{},
		&infrafeedfd_file_impl.MsgRowFeed{},
		&infrareciperc_exec.MsgPanic{},
		&infrareciperc_group.MsgHeader{},
		&infrareciperc_group_impl.MsgGroup{},
		&infrareciperc_spec.MsgSelfContained{},
		&infrareciperc_value.MsgRepository{},
		&infrareciperc_value.MsgValFdFileRowFeed{},
		&infrareportrp_model_impl.MsgColumnSpec{},
		&infrareportrp_model_impl.MsgTransactionReport{},
		&infrareportrp_writer_impl.MsgSortedWriter{},
		&infrareportrp_writer_impl.MsgUIWriter{},
		&infrareportrp_writer_impl.MsgXlsxWriter{},
		&infrauiapp_ui.MsgConsole{},
		&infrauiapp_ui.MsgProgress{},
		&ingredientig_file.MsgUpload{},
		&recipedevdiag.MsgLoader{},
		&recipefileimportbatch.MsgUrl{},
		&recipegroupmember.MsgList{},
		&recipegroupmemberbatch.MsgOperation{},
		&recipemember.MsgInvite{},
		&recipememberupdate.MsgEmail{},
		&recipeteamsharedlink.MsgList{},
		&recipeutiltidymove.MsgLocal{},
	}
}
