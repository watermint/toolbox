package app_template

type Template interface {
	// Define template with given name and resource names.
	Define(name string, resNames ...string) error

	// Render template with given name and data.
	Render(name string, d ...D) string
}

type D map[string]interface{}
