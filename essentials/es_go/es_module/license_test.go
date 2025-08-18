package es_module

import (
	"os"
	"testing"
)

func testFile(t *testing.T, filename string, s func(body string)) {
	body, err := os.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	s(string(body))
}

func testSingleLicenseFile(t *testing.T, filename string, s func(license License)) {
	testFile(t, filename, func(body string) {
		licenses := ScanLicenseBody(body)
		if n := len(licenses); n != 1 {
			if n < 1 {
				t.Error("no license")
			} else {
				licTypes := make([]string, len(licenses))
				for i, l := range licenses {
					licTypes[i] = string(l.Type())
				}
				t.Error(n, licTypes)
			}
		}
		s(licenses[0])
	})
}

func TestLicenseApache(t *testing.T) {
	testSingleLicenseFile(t, "license_apache.dat", func(license License) {
		if license.Type() != LicenseTypeApache20 {
			t.Error(license.Type())
		}
	})
}

func TestLicenseAGPL(t *testing.T) {
	testSingleLicenseFile(t, "license_agpl.dat", func(license License) {
		if license.Type() != LicenseTypeAGPL {
			t.Error(license.Type())
		}
	})
}

func TestLicenseBSD2(t *testing.T) {
	testSingleLicenseFile(t, "license_bsd2.dat", func(license License) {
		if license.Type() != LicenseTypeBSD2Clause {
			t.Error(license.Type())
		}
	})
}

func TestLicenseBSD3(t *testing.T) {
	testSingleLicenseFile(t, "license_bsd3.dat", func(license License) {
		if license.Type() != LicenseTypeBSD3Clause {
			t.Error(license.Type())
		}
	})
}

func TestLicenseCDDL(t *testing.T) {
	testSingleLicenseFile(t, "license_cddl.dat", func(license License) {
		if license.Type() != LicenseTypeCDDL {
			t.Error(license.Type())
		}
	})
}

func TestLicenseEclipse(t *testing.T) {
	testSingleLicenseFile(t, "license_eclipse.dat", func(license License) {
		if license.Type() != LicenseTypeEclipse {
			t.Error(license.Type())
		}
	})
}

func TestLicenseGPL(t *testing.T) {
	testSingleLicenseFile(t, "license_gpl2.dat", func(license License) {
		if license.Type() != LicenseTypeGPL {
			t.Error(license.Type())
		}
	})
	testSingleLicenseFile(t, "license_gpl3.dat", func(license License) {
		if license.Type() != LicenseTypeGPL {
			t.Error(license.Type())
		}
	})
}

func TestLicenseLGPL(t *testing.T) {
	testSingleLicenseFile(t, "license_lgpl.dat", func(license License) {
		if license.Type() != LicenseTypeLGPL {
			t.Error(license.Type())
		}
	})
}

func TestLicenseMIT(t *testing.T) {
	testSingleLicenseFile(t, "license_mit.dat", func(license License) {
		if license.Type() != LicenseTypeMIT {
			t.Error(license.Type())
		}
	})
}

func TestLicenseMPL(t *testing.T) {
	testSingleLicenseFile(t, "license_mpl.dat", func(license License) {
		if license.Type() != LicenseTypeMPL {
			t.Error(license.Type())
		}
	})
}

func TestLicenseUnlicense(t *testing.T) {
	testSingleLicenseFile(t, "license_unlicense.dat", func(license License) {
		if license.Type() != LicenseTypeUnlicense {
			t.Error(license.Type())
		}
	})
}
