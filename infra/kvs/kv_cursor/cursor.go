package kv_cursor

type Cursor interface {
	First() (key string, value []byte, exist bool)
	Last() (key string, value []byte, exist bool)
	Next() (key string, value []byte, exist bool)
	Prev() (key string, value []byte, exist bool)
	Seek(seek string) (key string, value []byte, exist bool)
	Delete() error
}
