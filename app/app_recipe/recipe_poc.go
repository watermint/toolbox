package app_recipe

type PlaceHolder func(ph *placeHolder) *placeHolder
type placeHolder struct {
	key   string
	value interface{}
}
type MessageContainer interface {
	Message(key string)
}
type Message interface {
	Key() string
	Text(placeHolders ...PlaceHolder) string
}

func M(key string, placeHolders ...PlaceHolder) Message {
	panic("implement me")
}

type ValueObjectValidator struct {
}

func (z *ValueObjectValidator) Error(msg string) {
}

type ValueObject interface {
	Validate(t *ValueObjectValidator)
}
type Recipe struct {
	Name  string
	Usage string
	Value ValueObject
	Exec  Cook
}

type Cook interface {
	Exec(rc RecipeContext) error
}

type RecipeContext interface {
	Value() ValueObject
}

type ApiRecipeContext struct {
}

func WithBusinessFile(exec func(rc *ApiRecipeContext) error) Cook {
	panic("implement me")
}
