package sv_file

import "github.com/watermint/toolbox/domain/model/mo_file"

type Search interface {
	// `files/search`
	// options: max_results
	ByName(query string) (entries []mo_file.Entry, err error)

	// `files/search`
	// options: max_results
	ByContent(query string) (entries []mo_file.Entry, err error)

	// `files/search`
	// options: max_results
	Deleted(query string) (entries []mo_file.Entry, err error)
}
