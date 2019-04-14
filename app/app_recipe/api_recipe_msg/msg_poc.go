package api_recipe_msg

type PlaceHolder func() *placeHolder
type placeHolder struct {
	key   string
	value interface{}
}

func P(key string, value interface{}) PlaceHolder {
	return func() *placeHolder {
		ph := &placeHolder{}
		ph.key = key
		ph.value = value
		return ph
	}
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

type K func(key string, placeHolders ...PlaceHolder) Message
