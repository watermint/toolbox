package dc_license

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	confidenceThreshold = 0.9
)

var (
	gitRemotes = []string{
		"origin",
		"upstream",
	}

	//supportedLicenseTypes = map[licenses.Type]bool{
	//	licenses.Unknown:      false,
	//	licenses.Restricted:   false,
	//	licenses.Reciprocal:   false,
	//	licenses.Notice:       true,
	//	licenses.Permissive:   true,
	//	licenses.Unencumbered: true,
	//	licenses.Forbidden:    false,
	//}

	ignoreLibraryList = map[string]bool{
		"github.com/watermint/toolbox":          true, // myself
		"github.com/vbauerster/mpb/v5/cwriter":  true, // same as github.com/vbauerster/mpb
		"github.com/vbauerster/mpb/v5/decor":    true, // same as github.com/vbauerster/mpb
		"github.com/vbauerster/mpb/v5/internal": true, // same as github.com/vbauerster/mpb
		"github.com/golang/freetype/truetype":   true, // dual license GPL/FTL (BSD like)
		"github.com/golang/freetype/raster":     true, // dual license GPL/FTL (BSD like)
	}

	approvedLibraryList = map[string]LicenseInfo{
		"github.com/vbauerster/mpb/v5": {
			Package:     "github.com/vbauerster/mpb",
			Url:         "https://github.com/vbauerster/mpb",
			LicenseType: "The Unlicense",
			LicenseBody: "This is free and unencumbered software released into the public domain.\n\nAnyone is free to copy, modify, publish, use, compile, sell, or\ndistribute this software, either in source code form or as a compiled\nbinary, for any purpose, commercial or non-commercial, and by any\nmeans.\n\nIn jurisdictions that recognize copyright laws, the author or authors\nof this software dedicate any and all copyright interest in the\nsoftware to the public domain. We make this dedication for the benefit\nof the public at large and to the detriment of our heirs and\nsuccessors. We intend this dedication to be an overt act of\nrelinquishment in perpetuity of all present and future rights to this\nsoftware under copyright law.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND,\nEXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF\nMERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.\nIN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR\nOTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,\nARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR\nOTHER DEALINGS IN THE SOFTWARE.\n\nFor more information, please refer to <http://unlicense.org/>",
		},
		"github.com/golang/freetype": {
			Package:     "github.com/golang/freetype",
			Url:         "https://github.com/golang/freetype/blob/master/LICENSE",
			LicenseType: "The FreeType License",
			LicenseBody: "                    The FreeType Project LICENSE\n                    ----------------------------\n\n                            2006-Jan-27\n\n                    Copyright 1996-2002, 2006 by\n          David Turner, Robert Wilhelm, and Werner Lemberg\n\n\n\nIntroduction\n============\n\n  The FreeType  Project is distributed in  several archive packages;\n  some of them may contain, in addition to the FreeType font engine,\n  various tools and  contributions which rely on, or  relate to, the\n  FreeType Project.\n\n  This  license applies  to all  files found  in such  packages, and\n  which do not  fall under their own explicit  license.  The license\n  affects  thus  the  FreeType   font  engine,  the  test  programs,\n  documentation and makefiles, at the very least.\n\n  This  license   was  inspired  by  the  BSD,   Artistic,  and  IJG\n  (Independent JPEG  Group) licenses, which  all encourage inclusion\n  and  use of  free  software in  commercial  and freeware  products\n  alike.  As a consequence, its main points are that:\n\n    o We don't promise that this software works. However, we will be\n      interested in any kind of bug reports. (`as is' distribution)\n\n    o You can  use this software for whatever you  want, in parts or\n      full form, without having to pay us. (`royalty-free' usage)\n\n    o You may not pretend that  you wrote this software.  If you use\n      it, or  only parts of it,  in a program,  you must acknowledge\n      somewhere  in  your  documentation  that  you  have  used  the\n      FreeType code. (`credits')\n\n  We  specifically  permit  and  encourage  the  inclusion  of  this\n  software, with  or without modifications,  in commercial products.\n  We  disclaim  all warranties  covering  The  FreeType Project  and\n  assume no liability related to The FreeType Project.\n\n\n  Finally,  many  people  asked  us  for  a  preferred  form  for  a\n  credit/disclaimer to use in compliance with this license.  We thus\n  encourage you to use the following text:\n\n   \"\"\"  \n    Portions of this software are copyright Š <year> The FreeType\n    Project (www.freetype.org).  All rights reserved.\n   \"\"\"\n\n  Please replace <year> with the value from the FreeType version you\n  actually use.\n\n\nLegal Terms\n===========\n\n0. Definitions\n--------------\n\n  Throughout this license,  the terms `package', `FreeType Project',\n  and  `FreeType  archive' refer  to  the  set  of files  originally\n  distributed  by the  authors  (David Turner,  Robert Wilhelm,  and\n  Werner Lemberg) as the `FreeType Project', be they named as alpha,\n  beta or final release.\n\n  `You' refers to  the licensee, or person using  the project, where\n  `using' is a generic term including compiling the project's source\n  code as  well as linking it  to form a  `program' or `executable'.\n  This  program is  referred to  as  `a program  using the  FreeType\n  engine'.\n\n  This  license applies  to all  files distributed  in  the original\n  FreeType  Project,   including  all  source   code,  binaries  and\n  documentation,  unless  otherwise  stated   in  the  file  in  its\n  original, unmodified form as  distributed in the original archive.\n  If you are  unsure whether or not a particular  file is covered by\n  this license, you must contact us to verify this.\n\n  The FreeType  Project is copyright (C) 1996-2000  by David Turner,\n  Robert Wilhelm, and Werner Lemberg.  All rights reserved except as\n  specified below.\n\n1. No Warranty\n--------------\n\n  THE FREETYPE PROJECT  IS PROVIDED `AS IS' WITHOUT  WARRANTY OF ANY\n  KIND, EITHER  EXPRESS OR IMPLIED,  INCLUDING, BUT NOT  LIMITED TO,\n  WARRANTIES  OF  MERCHANTABILITY   AND  FITNESS  FOR  A  PARTICULAR\n  PURPOSE.  IN NO EVENT WILL ANY OF THE AUTHORS OR COPYRIGHT HOLDERS\n  BE LIABLE  FOR ANY DAMAGES CAUSED  BY THE USE OR  THE INABILITY TO\n  USE, OF THE FREETYPE PROJECT.\n\n2. Redistribution\n-----------------\n\n  This  license  grants  a  worldwide, royalty-free,  perpetual  and\n  irrevocable right  and license to use,  execute, perform, compile,\n  display,  copy,   create  derivative  works   of,  distribute  and\n  sublicense the  FreeType Project (in  both source and  object code\n  forms)  and  derivative works  thereof  for  any  purpose; and  to\n  authorize others  to exercise  some or all  of the  rights granted\n  herein, subject to the following conditions:\n\n    o Redistribution of  source code  must retain this  license file\n      (`FTL.TXT') unaltered; any  additions, deletions or changes to\n      the original  files must be clearly  indicated in accompanying\n      documentation.   The  copyright   notices  of  the  unaltered,\n      original  files must  be  preserved in  all  copies of  source\n      files.\n\n    o Redistribution in binary form must provide a  disclaimer  that\n      states  that  the software is based in part of the work of the\n      FreeType Team,  in  the  distribution  documentation.  We also\n      encourage you to put an URL to the FreeType web page  in  your\n      documentation, though this isn't mandatory.\n\n  These conditions  apply to any  software derived from or  based on\n  the FreeType Project,  not just the unmodified files.   If you use\n  our work, you  must acknowledge us.  However, no  fee need be paid\n  to us.\n\n3. Advertising\n--------------\n\n  Neither the  FreeType authors and  contributors nor you  shall use\n  the name of the  other for commercial, advertising, or promotional\n  purposes without specific prior written permission.\n\n  We suggest,  but do not require, that  you use one or  more of the\n  following phrases to refer  to this software in your documentation\n  or advertising  materials: `FreeType Project',  `FreeType Engine',\n  `FreeType library', or `FreeType Distribution'.\n\n  As  you have  not signed  this license,  you are  not  required to\n  accept  it.   However,  as  the FreeType  Project  is  copyrighted\n  material, only  this license, or  another one contracted  with the\n  authors, grants you  the right to use, distribute,  and modify it.\n  Therefore,  by  using,  distributing,  or modifying  the  FreeType\n  Project, you indicate that you understand and accept all the terms\n  of this license.\n\n4. Contacts\n-----------\n\n  There are two mailing lists related to FreeType:\n\n    o freetype@nongnu.org\n\n      Discusses general use and applications of FreeType, as well as\n      future and  wanted additions to the  library and distribution.\n      If  you are looking  for support,  start in  this list  if you\n      haven't found anything to help you in the documentation.\n\n    o freetype-devel@nongnu.org\n\n      Discusses bugs,  as well  as engine internals,  design issues,\n      specific licenses, porting, etc.\n\n  Our home page can be found at\n\n    http://www.freetype.org\n\n\n--- end of FTL.TXT ---",
		},
	}
)

func Detect(c app_control.Control) (inventory []LicenseInfo, err error) {
	return make([]LicenseInfo, 0), nil
	//l := c.Log()
	//cf, err := licenses.NewClassifier(confidenceThreshold)
	//if err != nil {
	//	l.Debug("Unable to create the classifier", esl.Error(err))
	//	return nil, err
	//}
	//
	//libs, err := licenses.Libraries(context.Background(), cf, []string{})
	//if err != nil {
	//	l.Debug("Unable to load libraries", esl.Error(err))
	//	return nil, err
	//}
	//
	//inventory = make([]LicenseInfo, 0)
	//
	//var lastErr error
	//for _, lib := range libs {
	//	ll := l.With(esl.String("library", lib.Name()), esl.String("path", lib.LicensePath))
	//	licName, licType, err := cf.Identify(lib.LicensePath)
	//	if err != nil {
	//		ll.Debug("Unable to determine the license type", esl.Error(err))
	//		return nil, err
	//	}
	//	ll = ll.With(esl.String("licenseName", licName))
	//
	//	if ignore, ok := ignoreLibraryList[lib.Name()]; ok && ignore {
	//		ll.Debug("Ignore the library")
	//		continue
	//	}
	//
	//	if li, ok := approvedLibraryList[lib.Name()]; ok {
	//		ll.Debug("Approved library found", esl.Any("library", li))
	//		inventory = append(inventory, li)
	//		continue
	//	}
	//
	//	//if supported, ok := supportedLicenseTypes[licType]; !ok {
	//	//	ll.Warn("Unknown license type", esl.String("licenseType", licType.String()))
	//	//	lastErr = errors.New("unknown license type")
	//	//} else if !supported {
	//	//	ll.Warn("Unsupported license type library found", esl.String("licenseType", licType.String()))
	//	//	lastErr = errors.New("unsupported license type")
	//	//}
	//
	//	var licenseUrl string
	//	if lib.LicensePath != "" {
	//		repo, err := licenses.FindGitRepo(lib.LicensePath)
	//		if err != nil {
	//			ll.Debug("Unable to find the git repository", esl.Error(err))
	//			derivedUrl, err := lib.FileURL(context.Background(), lib.LicensePath)
	//			if err != nil {
	//				ll.Debug("Unable to determine the library url", esl.Error(err))
	//			} else {
	//				licenseUrl = derivedUrl
	//			}
	//		} else {
	//			for _, remote := range gitRemotes {
	//				url, err := repo.FileURL(lib.LicensePath, remote)
	//				if err != nil {
	//					ll.Debug("Unable to determine the library url", esl.Error(err))
	//				} else {
	//					licenseUrl = url.String()
	//				}
	//			}
	//		}
	//	}
	//
	//	body, err := ioutil.ReadFile(lib.LicensePath)
	//	if err != nil {
	//		ll.Warn("Unable to read the license file", esl.Error(err))
	//		lastErr = err
	//	}
	//
	//	inventory = append(inventory, LicenseInfo{
	//		Package:     lib.Name(),
	//		Url:         licenseUrl,
	//		LicenseType: licName,
	//		LicenseBody: string(body),
	//	})
	//}
	//
	//return inventory, lastErr
}
