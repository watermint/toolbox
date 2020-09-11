package file

import "time"

type SyncRow struct {
	File string `json:"file"`
	Size int64  `json:"size"`
}

type Summary struct {
	Start               time.Time `json:"start"`
	End                 time.Time `json:"end"`
	NumBytes            int64     `json:"num_bytes"`
	NumFilesError       int64     `json:"num_files_error"`
	NumFilesTransferred int64     `json:"num_files_transferred"`
	NumFilesSkip        int64     `json:"num_files_skip"`
	NumFolderCreated    int64     `json:"num_folder_created"`
	NumDeleted          int64     `json:"num_delete"`
	NumApiCall          int64     `json:"num_api_call"`
}
