package dc_license

type Licenses struct {
	Project    LicenseInfo   `json:"project"`
	ThirdParty []LicenseInfo `json:"third_party"`
}

type LicenseInfo struct {
	Package     string `json:"package"`
	Version     string `json:"version"`
	Url         string `json:"url"`
	LicenseType string `json:"license_type"`
	LicenseBody string `json:"license_body"`
}
