package app_validate

import "github.com/watermint/toolbox/atbx/app_msg"

func AssertEmailFormat(email string) error {
	return nil
}

type InvalidRowError error

func InvalidRow(key string, p ...app_msg.Param) InvalidRowError {
	return nil
}

type NoDataRowError error

func NoDataRow() NoDataRowError {
	return nil
}
