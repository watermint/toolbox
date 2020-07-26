package queue

type SessionId int

type Factory interface {
	Create() Session
	Resume(id SessionId) (s Session, err error)
}

type Session interface {
}

type Handler func(session Session)
