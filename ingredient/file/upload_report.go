package file

import "time"

type UploadRow struct {
	File string `json:"file"`
	Size int64  `json:"size"`
}

type UploadSummary struct {
	UploadStart      time.Time `json:"upload_start"`
	UploadEnd        time.Time `json:"upload_end"`
	NumBytes         int64     `json:"num_bytes"`
	NumFilesError    int64     `json:"num_files_error"`
	NumFilesUpload   int64     `json:"num_files_upload"`
	NumFilesSkip     int64     `json:"num_files_skip"`
	NumFolderCreated int64     `json:"num_folder_created"`
	NumDeleted       int64     `json:"num_delete"`
	NumApiCall       int64     `json:"num_api_call"`
}
