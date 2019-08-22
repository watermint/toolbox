package app_workspace

type User interface {
	// Secrets path
	Secrets() string
}

type MultiUser interface {
	User

	// User home path
	UserHome() string
}
