package qt_recipe

func HumanInteractionRequired() error {
	return &ErrorHumanInteractionRequired{}
}

type ErrorHumanInteractionRequired struct {
}

func (z *ErrorHumanInteractionRequired) Error() string {
	return "human interaction require"
}

func NoTestRequired() error {
	return &ErrorNoTestRequired{}
}

type ErrorNoTestRequired struct {
}

func (z *ErrorNoTestRequired) Error() string {
	return "no test required"
}

func ScenarioTest() error {
	return &ErrorScenarioTest{}
}

type ErrorScenarioTest struct {
}

func (z *ErrorScenarioTest) Error() string {
	return "scenario test"
}

func ImplementMe() error {
	return &ErrorImplementMe{}
}

type ErrorImplementMe struct {
}

func (z *ErrorImplementMe) Error() string {
	return "implement me"
}

func NotEnoughResource() error {
	return &ErrorNotEnoughResource{}
}

type ErrorNotEnoughResource struct {
}

func (z *ErrorNotEnoughResource) Error() string {
	return "not enough resource"
}
