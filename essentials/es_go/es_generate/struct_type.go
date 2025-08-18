package es_generate

import (
	"slices"
)

type StructType struct {
	// Relative package path from the scan root.
	Package string `json:"package"`

	// Name of the struct type.
	Name string `json:"name"`
}

func UniqSortedPackages(sts []*StructType) []string {
	// Collect unique packages using a temporary map
	pkgMap := make(map[string]bool)
	for _, st := range sts {
		pkgMap[st.Package] = true
	}

	// Get keys from the map
	pkgSorted := make([]string, 0, len(pkgMap))
	for pkg := range pkgMap {
		pkgSorted = append(pkgSorted, pkg)
	}

	// Sort using slices.Sort
	slices.Sort(pkgSorted)
	return pkgSorted
}

func SortedStructTypes(sts []*StructType) (sorted []*StructType) {
	stsMap := make(map[string]*StructType)
	for _, st := range sts {
		stsMap[st.Package+"/"+st.Name] = st
	}

	// Pre-allocate capacity for memory efficiency
	stsKey := make([]string, 0, len(stsMap))
	for key := range stsMap {
		stsKey = append(stsKey, key)
	}

	// Sort using slices.Sort
	slices.Sort(stsKey)

	sorted = make([]*StructType, len(stsKey))
	for i, key := range stsKey {
		sorted[i] = stsMap[key]
	}
	return sorted
}
