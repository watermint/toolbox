package app_feature

import (
	"testing"
)

type SampleOptIn struct {
	OptInStatus
}

func TestOptInFrom(t *testing.T) {
	v := map[string]interface{}{
		"timestamp": "2020-04-08T17:20:23+09:00",
		"user":      "john",
		"status":    true,
	}
	soi := &SampleOptIn{}
	if err := OptInFrom(v, soi); err != nil {
		t.Error(err)
	}
	if soi.User != "john" {
		t.Error(soi.User)
	}
	if !soi.Status {
		t.Error(soi.Status)
	}
	if soi.Timestamp != "2020-04-08T17:20:23+09:00" {
		t.Error(soi.Timestamp)
	}
}

func TestOptInStatus_OptInCommit(t *testing.T) {
	soi := &SampleOptIn{}
	soi.OptInCommit(true)
	if soi.OptInUser() == "" {
		t.Error(soi.User)
	}
	if !soi.OptInIsEnabled() {
		t.Error(soi.Status)
	}
	if soi.OptInTimestamp() == "" {
		t.Error(soi.Timestamp)
	}
}

func TestOptInStatus_OptInName(t *testing.T) {
	soi := &SampleOptIn{}
	name := OptInName(soi)
	if name != "infra.control.app_feature.sample_opt_in" {
		t.Error(name)
	}
}

func TestOptInStatus_OptInMessages(t *testing.T) {
	soi := &SampleOptIn{}
	ma := OptInAgreement(soi)
	if ma.Key() != "infra.control.app_feature.sample_opt_in.agreement" {
		t.Error(ma.Key())
	}
	md := OptInDisclaimer(soi)
	if md.Key() != "infra.control.app_feature.sample_opt_in.disclaimer" {
		t.Error(md.Key())
	}
	mc := OptInDescription(soi)
	if mc.Key() != "infra.control.app_feature.sample_opt_in.desc" {
		t.Error(mc.Key())
	}
}

func TestOptInStatus_OptInIsEnabled(t *testing.T) {
	soi := &SampleOptIn{}
	if soi.OptInIsEnabled() {
		t.Error(soi.OptInIsEnabled())
	}
	if soi.OptInTimestamp() != "" {
		t.Error(soi.OptInTimestamp())
	}
	if soi.OptInUser() != "" {
		t.Error(soi.OptInUser())
	}
}
