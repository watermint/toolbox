package ut_source

import "sort"

type StructType struct {
	// Relative package path from the scan root.
	Package string `json:"package"`

	// Name of the struct type.
	Name string `json:"name"`
}

func UniqSortedPackages(sts []*StructType) []string {
	pkgMap := make(map[string]bool)
	for _, st := range sts {
		pkgMap[st.Package] = true
	}
	pkgSorted := make([]string, 0)
	for pkg := range pkgMap {
		pkgSorted = append(pkgSorted, pkg)
	}
	sort.Strings(pkgSorted)
	return pkgSorted
}

func SortedStructTypes(sts []*StructType) (sorted []*StructType) {
	stsMap := make(map[string]*StructType)
	for _, st := range sts {
		stsMap[st.Package+"/"+st.Name] = st
	}
	stsKey := make([]string, 0)
	for key := range stsMap {
		stsKey = append(stsKey, key)
	}
	sort.Strings(stsKey)
	sorted = make([]*StructType, len(stsKey))
	for i, key := range stsKey {
		sorted[i] = stsMap[key]
	}
	return sorted
}
