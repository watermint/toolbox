package api

import (
	"encoding/json"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

func parseErrorFilesSearch(body []byte) error {
	var apiErr files.SearchAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}

func parseErrorListFolder(body []byte) error {
	var apiErr files.ListFolderAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}
