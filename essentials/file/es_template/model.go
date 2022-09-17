package es_template

type Folder struct {
	Name    string   `json:"name"`
	Files   []File   `json:"files,omitempty"`
	Folders []Folder `json:"folders,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type File struct {
	Name    string   `json:"name"`
	Content string   `json:"content,omitempty"`
	Source  string   `json:"source,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type Root struct {
	Files   []File   `json:"files,omitempty"`
	Folders []Folder `json:"folders,omitempty"`
}
