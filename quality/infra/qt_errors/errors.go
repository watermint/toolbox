package qt_errors

import "errors"

var (
	// Marker error: Skip end to end test
	ErrorSkipEndToEndTest = errors.New("skip end to end test")

	// Marker error: The test requires human interaction. Will not do automated test
	ErrorHumanInteractionRequired = errors.New("human interaction required")

	// Marker error: The test requires no test
	ErrorNoTestRequired = errors.New("no test required")

	// Marker error: The test will be done separately as a scenario test.
	ErrorScenarioTest = errors.New("scenario test")

	// Marker error: The test is not yet implemented
	ErrorImplementMe = errors.New("implement me")

	// Marker error: The test requires some resource, but the resource is not available.
	ErrorNotEnoughResource = errors.New("not enough resource")

	// Unsupported UI
	ErrorUnsupportedUI = errors.New("unsupported UI for this auth scope")

	// Marker error: Mock
	ErrorMock = errors.New("mock error")
)
