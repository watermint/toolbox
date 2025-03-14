package dbx_filesystem

type BaseNamespaceType string

const (
	BaseNamespaceRoot BaseNamespaceType = "root"
	BaseNamespaceHome BaseNamespaceType = "home"
)

var (
	BaseNamespaceTypes   = []BaseNamespaceType{BaseNamespaceRoot, BaseNamespaceHome}
	BaseNamespaceDefault = BaseNamespaceRoot

	BaseNamespaceTypesInString = []string{
		string(BaseNamespaceRoot),
		string(BaseNamespaceHome),
	}
	BaseNamespaceDefaultInString = string(BaseNamespaceRoot)
)

type RootNamespaceResolver interface {
	ResolveIndividual() (namespaceId string, err error)
	ResolveTeamMember(teamMemberId string) (namespaceId string, err error)
}

func AsNamespaceType(s string) BaseNamespaceType {
	for _, t := range BaseNamespaceTypes {
		if string(t) == s {
			return t
		}
	}
	return BaseNamespaceDefault
}
