package dbx_async_impl

const (
	// files/create_folder_batch

	createFolderBatchForceAsyncFalseRes = `{
  ".tag": "complete",
  "entries": [
    {
      ".tag": "success",
      "metadata": {
        "name": "create_folder_batch",
        "path_lower": "/toolbox-testsuite/create_folder_batch",
        "path_display": "/toolbox-testsuite/create_folder_batch",
        "id": "id:xxxxxxxxxxxxxxxxxxxxxx"
      }
    }
  ]
}`
	createFolderBatchForceAsyncTrueRes = `{
  ".tag": "async_job_id",
  "async_job_id": "dbjid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxx"
}`
	createFolderBatchCheckInvalidRes = `{
  "error_summary": "invalid_async_job_id/...",
  "error": {
    ".tag": "invalid_async_job_id"
  }
}`
	createFolderBatchCheckInProgressRes = `{
  ".tag": "in_progress"
}`
	createFolderBatchCheckSuccessRes = `{
  ".tag": "complete",
  "entries": [
    {
      ".tag": "success",
      "metadata": {
        "name": "test (1)",
        "path_lower": "/toolbox-testsuite/create_folder_batch/test (1)",
        "path_display": "/toolbox-testsuite/create_folder_batch/test (1)",
        "id": "id:xxxxxxxxxxxxxxxxxxxxxx"
      }
    }
  ]
}`

	// sharing/share_folder

	shareFolderForceAsyncTrue = `{
  ".tag": "async_job_id",
  "async_job_id": "dbjid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-xxx-xx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-x-x"
}`
	shareFolderCheckShareJobStatusInProgressRes = `{
  ".tag": "in_progress"
}`
	shareFoldercheckShareJobStatusCompletedRes = `{
  ".tag": "complete",
  "access_type": {
    ".tag": "owner"
  },
  "is_inside_team_folder": false,
  "is_team_folder": false,
  "owner_team": {
    "id": "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "name": "xxxxxxxxx xxx"
  },
  "path_lower": "/toolbox-testsuite/share_folder/test",
  "name": "test",
  "policy": {
    "member_policy": {
      ".tag": "anyone"
    },
    "resolved_member_policy": {
      ".tag": "anyone"
    },
    "acl_update_policy": {
      ".tag": "editors"
    },
    "shared_link_policy": {
      ".tag": "anyone"
    },
    "viewer_info_policy": {
      ".tag": "enabled"
    }
  },
  "preview_url": "https://www.dropbox.com/scl/fo/xxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxx?dl=0",
  "shared_folder_id": "xxxxxxxxxx",
  "time_invited": "2019-03-11T23:03:47Z",
  "access_inheritance": {
    ".tag": "inherit"
  }
}`

	// sharing/unshare_folder

	unshareFolderRes = `{
  ".tag": "async_job_id",
  "async_job_id": "dbjid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-xxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-xxxx"
}`
	unshareFolderCheckJobStatusInProgressRes = `{
  ".tag": "in_progress"
}`
	unshareFolderCheckJobStatusErrorRes = `{
  "error_summary": "internal_error/",
  "error": {
    ".tag": "internal_error"
  }
}`
)

var (
	seqCreateFolderBatchImmediateResponse = []string{
		createFolderBatchForceAsyncFalseRes,
	}
	seqCreateFolderBatchSuccess = []string{
		createFolderBatchForceAsyncTrueRes,
		createFolderBatchCheckSuccessRes,
	}
	seqCreateFolderBatchInProgressSuccess = []string{
		createFolderBatchForceAsyncTrueRes,
		createFolderBatchCheckInProgressRes,
		createFolderBatchCheckSuccessRes,
	}
	seqCreateFolderBatchInvalid = []string{
		createFolderBatchForceAsyncTrueRes,
		createFolderBatchCheckInvalidRes,
	}
	seqCreateFolderBatchInProgressInvalid = []string{
		createFolderBatchForceAsyncTrueRes,
		createFolderBatchCheckInProgressRes,
		createFolderBatchCheckInvalidRes,
	}
	seqShareFolderInProgressCompleted = []string{
		shareFolderForceAsyncTrue,
		shareFolderCheckShareJobStatusInProgressRes,
		shareFoldercheckShareJobStatusCompletedRes,
	}
	seqShareFolderCompleted = []string{
		shareFolderForceAsyncTrue,
		shareFoldercheckShareJobStatusCompletedRes,
	}
	seqUnshareFolderInProgressError = []string{
		unshareFolderRes,
		unshareFolderCheckJobStatusInProgressRes,
		unshareFolderCheckJobStatusErrorRes,
	}
	seqUnshareFolderError = []string{
		unshareFolderRes,
		unshareFolderCheckJobStatusErrorRes,
	}
)
