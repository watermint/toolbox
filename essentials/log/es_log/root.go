package es_log

type RootLogger interface {
	// Set log level for the console.
	SetLevelConsole(level Level)

	// Add logger to root log audience.
	AddSubscriber(l Logger)

	// Remove logger from root log audience.
	RemoveSubscriber(l Logger)

	// Get current logger.
	Current() Logger
}
