package traverse

type TraverseFile interface {
	Prepare() error
	Close() error
	Scan() error
}
