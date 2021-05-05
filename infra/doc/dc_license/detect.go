package dc_license

import (
	"context"
	"errors"
	"github.com/google/go-licenses/licenses"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"io/ioutil"
)

const (
	confidenceThreshold = 0.9
)

var (
	gitRemotes = []string{
		"origin",
		"upstream",
	}

	supportedLicenseTypes = map[licenses.Type]bool{
		licenses.Unknown:      false,
		licenses.Restricted:   false,
		licenses.Reciprocal:   false,
		licenses.Notice:       true,
		licenses.Permissive:   true,
		licenses.Unencumbered: true,
		licenses.Forbidden:    false,
	}

	ignoreLibraryList = map[string]bool{
		"github.com/watermint/toolbox":          true, // myself
		"github.com/vbauerster/mpb/v5/cwriter":  true, // same as github.com/vbauerster/mpb
		"github.com/vbauerster/mpb/v5/decor":    true, // same as github.com/vbauerster/mpb
		"github.com/vbauerster/mpb/v5/internal": true, // same as github.com/vbauerster/mpb
	}

	approvedLibraryList = map[string]LicenseInfo{
		"github.com/vbauerster/mpb/v5": {
			Package:     "github.com/vbauerster/mpb",
			Url:         "https://github.com/vbauerster/mpb",
			LicenseType: "The Unlicense",
			LicenseBody: "This is free and unencumbered software released into the public domain.\n\nAnyone is free to copy, modify, publish, use, compile, sell, or\ndistribute this software, either in source code form or as a compiled\nbinary, for any purpose, commercial or non-commercial, and by any\nmeans.\n\nIn jurisdictions that recognize copyright laws, the author or authors\nof this software dedicate any and all copyright interest in the\nsoftware to the public domain. We make this dedication for the benefit\nof the public at large and to the detriment of our heirs and\nsuccessors. We intend this dedication to be an overt act of\nrelinquishment in perpetuity of all present and future rights to this\nsoftware under copyright law.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND,\nEXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF\nMERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.\nIN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR\nOTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,\nARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR\nOTHER DEALINGS IN THE SOFTWARE.\n\nFor more information, please refer to <http://unlicense.org/>",
		},
	}
)

func Detect(c app_control.Control) (inventory []LicenseInfo, err error) {
	l := c.Log()
	cf, err := licenses.NewClassifier(confidenceThreshold)
	if err != nil {
		l.Debug("Unable to create the classifier", esl.Error(err))
		return nil, err
	}

	libs, err := licenses.Libraries(context.Background(), cf)
	if err != nil {
		l.Debug("Unable to load libraries", esl.Error(err))
		return nil, err
	}

	inventory = make([]LicenseInfo, 0)

	var lastErr error
	for _, lib := range libs {
		ll := l.With(esl.String("library", lib.Name()), esl.String("path", lib.LicensePath))
		licName, licType, err := cf.Identify(lib.LicensePath)
		if err != nil {
			ll.Debug("Unable to determine the license type", esl.Error(err))
			return nil, err
		}
		ll = ll.With(esl.String("licenseName", licName))

		if ignore, ok := ignoreLibraryList[lib.Name()]; ok && ignore {
			ll.Debug("Ignore the library")
			continue
		}

		if li, ok := approvedLibraryList[lib.Name()]; ok {
			ll.Debug("Approved library found", esl.Any("library", li))
			inventory = append(inventory, li)
			continue
		}

		if supported, ok := supportedLicenseTypes[licType]; !ok {
			ll.Warn("Unknown license type", esl.String("licenseType", licType.String()))
			lastErr = errors.New("unknown license type")
		} else if !supported {
			ll.Warn("Unsupported license type library found", esl.String("licenseType", licType.String()))
			lastErr = errors.New("unsupported license type")
		}

		var licenseUrl string
		if lib.LicensePath != "" {
			repo, err := licenses.FindGitRepo(lib.LicensePath)
			if err != nil {
				ll.Debug("Unable to find the git repository", esl.Error(err))
				derivedUrl, err := lib.FileURL(lib.LicensePath)
				if err != nil {
					ll.Debug("Unable to determine the library url", esl.Error(err))
				} else {
					licenseUrl = derivedUrl.String()
				}
			} else {
				for _, remote := range gitRemotes {
					url, err := repo.FileURL(lib.LicensePath, remote)
					if err != nil {
						ll.Debug("Unable to determine the library url", esl.Error(err))
					} else {
						licenseUrl = url.String()
					}
				}
			}
		}

		body, err := ioutil.ReadFile(lib.LicensePath)
		if err != nil {
			ll.Warn("Unable to read the license file", esl.Error(err))
			lastErr = err
		}

		inventory = append(inventory, LicenseInfo{
			Package:     lib.Name(),
			Url:         licenseUrl,
			LicenseType: licName,
			LicenseBody: string(body),
		})
	}

	return inventory, lastErr
}
