package api_util

import (
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"strings"
	"testing"
)

var (
	errorSummarySamples = []string{
		"access_denied/...",
		"account_id_not_found/...",
		"admin_not_active/...",
		"already_mounted/...",
		"app_id_mismatch/...",
		"app_lacks_access/...",
		"app_not_found/...",
		"bad_cursor/...",
		"bad_member/invalid_dropbox_id/...",
		"banned_member/...",
		"cannot_keep_account_and_delete_data/...",
		"cannot_keep_account_and_transfer/...",
		"cannot_keep_invited_user_account/...",
		"cannot_set_permissions/...",
		"cant_copy_shared_folder/...",
		"cant_move_folder_into_itself/...",
		"cant_move_shared_folder/...",
		"cant_nest_shared_folder/...",
		"cant_share_outside_team/...",
		"cant_transfer_ownership/...",
		"closed/...",
		"conflicting_property_names/...",
		"content_malformed/...",
		"conversion_error/...",
		"device_session_not_found/...",
		"directory_restricted_off/...",
		"disabled_for_team/...",
		"disallowed_shared_link_policy/...",
		"doc_archived/...",
		"doc_deleted/...",
		"doc_length_exceeded/...",
		"doc_not_found/...",
		"does_not_fit_template/...",
		"download_failed/...",
		"duplicate_user/...",
		"duplicated_or_nested_paths/...",
		"email_address_too_long_to_be_disabled/...",
		"email_not_verified/...",
		"email_reserved_for_other_user/...",
		"email_unverified/...",
		"empty_features_list/...",
		"expired_access_token/...",
		"external_id_already_in_use/...",
		"external_id_and_new_external_id_unsafe/...",
		"external_id_used_by_other_user/...",
		"folder_name_already_used/...",
		"folder_name_reserved/...",
		"folder_not_found/...",
		"folder_owner/...",
		"group_access/...",
		"group_already_deleted/...",
		"group_name_already_used/...",
		"group_name_invalid/...",
		"group_not_found/...",
		"group_not_in_team/...",
		"group_not_on_team/...",
		"image_size_exceeded/...",
		"in_progress/...",
		"inside_shared_folder/...",
		"insufficient_permissions/...",
		"insufficient_plan/...",
		"insufficient_quota/...",
		"internal_error/...",
		"invalid_access_token/...",
		"invalid_arg/...",
		"invalid_async_job_id/...",
		"invalid_comment/...",
		"invalid_copy_reference/...",
		"invalid_cursor/...",
		"invalid_dropbox_id/...",
		"invalid_folder_name/...",
		"invalid_id/...",
		"invalid_location/...",
		"invalid_member/...",
		"invalid_oauth1_token_info/...",
		"invalid_revision/...",
		"invalid_select_admin/...",
		"invalid_select_user/...",
		"invalid_time_range/...",
		"invalid_url/...",
		"last_admin/...",
		"list_error/...",
		"mapping_not_found/...",
		"member_not_found/...",
		"member_not_in_group/...",
		"mounted/...",
		"new_owner_email_unverified/...",
		"new_owner_not_a_member/...",
		"new_owner_unmounted/...",
		"no_account/...",
		"no_explicit_access/...",
		"no_new_data_specified/...",
		"no_permission/...",
		"not_a_folder/...",
		"not_a_member/...",
		"not_closed/...",
		"not_found/...",
		"not_mountable/...",
		"not_on_team/...",
		"not_unmountable/...",
		"other/...",
		"param_cannot_be_empty/...",
		"persistent_id_disabled/...",
		"persistent_id_used_by_other_user/...",
		"properties_error/does_not_fit_template/...",
		"property_field_too_large/...",
		"property_group_already_exists/...",
		"rate_limit/...",
		"recipient_not_verified/...",
		"remove_last_admin/...",
		"removed_and_transfer_admin_should_differ/...",
		"removed_and_transfer_dest_should_differ/...",
		"reset/...",
		"restricted_content/...",
		"revision_mismatch/...",
		"set_profile_disallowed/...",
		"shared_link_access_denied/...",
		"shared_link_is_directory/...",
		"shared_link_malformed/...",
		"shared_link_not_found/...",
		"some_users_are_excluded/...",
		"suspend_inactive_user/...",
		"suspend_last_admin/...",
		"system_managed_group_disallowed/...",
		"team_folder/...",
		"team_license_limit/...",
		"team_policy_disallows_member_policy/...",
		"template_attribute_too_large/...",
		"too_large/...",
		"too_many_files/...",
		"too_many_invitees/...",
		"too_many_properties/...",
		"too_many_shared_folder_targets/...",
		"too_many_templates/...",
		"too_many_users/...",
		"too_many_write_operations/...",
		"transfer_admin_is_not_admin/...",
		"transfer_admin_user_not_found/...",
		"transfer_admin_user_not_in_team/...",
		"transfer_dest_user_not_found/...",
		"transfer_dest_user_not_in_team/...",
		"unmounted/...",
		"unspecified_transfer_admin_id/...",
		"unsupported_content/...",
		"unsupported_extension/...",
		"unsupported_file/...",
		"unsupported_folder/...",
		"unsupported_image/...",
		"unsupported_link_type/...",
		"unsuspend_non_suspended_member/...",
		"user_cannot_be_manager_of_company_managed_group/...",
		"user_data_already_transferred/...",
		"user_data_cannot_be_transferred/...",
		"user_data_is_being_transferred/...",
		"user_error/email_unverified/...",
		"user_must_be_active_to_be_owner/...",
		"user_not_found/...",
		"user_not_in_team/...",
		"user_not_removed/...",
		"user_suspended/...",
		"user_unrecoverable/...",
		"users_not_in_team/...",
		"validation_error/...",
		"path/not_found/.",
	}
)

func TestErrorSummary(t *testing.T) {
	for _, e := range errorSummarySamples {
		ae := api_rpc.ApiError{
			ErrorSummary: e,
		}
		re := ErrorSummary(ae)

		if e == re {
			t.Error("invalid")
		}
		if strings.HasSuffix(re, ".") {
			t.Error("has suffix `.`", re)
		}
		if strings.HasSuffix(re, "/") {
			t.Error("has suffix `/`", re)
		}
	}
}