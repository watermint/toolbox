package rc_doc

type Value struct {
	Name     string      `json:"name"`
	Desc     string      `json:"desc"`
	Default  string      `json:"default"`
	TypeName string      `json:"type_name"`
	TypeAttr interface{} `json:"type_attr"`
}
