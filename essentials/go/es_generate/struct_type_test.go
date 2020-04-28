package es_generate

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestUniqSortedPackages(t *testing.T) {
	sts := []*StructType{
		{
			"fruit",
			"apple",
		},
		{
			"fruit",
			"orange",
		},
		{
			"veggies",
			"carrot",
		},
		{
			"veggies",
			"onion",
		},
	}

	usp := UniqSortedPackages(sts)
	d := cmp.Diff(usp, []string{"fruit", "veggies"})
	if d != "" {
		t.Error(d)
	}
}

func TestSortedStructTypes(t *testing.T) {
	sts := []*StructType{
		{
			"fruit",
			"apple",
		},
		{
			"veggies",
			"carrot",
		},
		{
			"fruit",
			"orange",
		},
		{
			"veggies",
			"onion",
		},
	}

	manuallySorted := []*StructType{
		{
			"fruit",
			"apple",
		},
		{
			"fruit",
			"orange",
		},
		{
			"veggies",
			"carrot",
		},
		{
			"veggies",
			"onion",
		},
	}

	sorted := SortedStructTypes(sts)
	d := cmp.Diff(sorted, manuallySorted)
	if d != "" {
		t.Error(d)
	}
}
