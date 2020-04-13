package gh_context

type ApiError struct {
	Message          string `path:"message" json:"message"`
	DocumentationUrl string `path:"documentation_url" json:"documentation_url"`
}

func (z ApiError) Error() string {
	return z.Message
}
