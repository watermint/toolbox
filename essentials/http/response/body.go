package response

type Body interface {
	// Length of the read content in bytes
	ContentLength() int64

	// Body bytes. Returns nil if the body written in to the file.
	Body() []byte

	// Body file. Return empty string if the body loaded on the memory.
	File() string

	// True when the body written into the file.
	IsFile() bool
}

func newMemoryBody(content []byte) Body {
	return &bodyMemoryImpl{
		content: content,
	}
}

func newFileBody(path string, contentLength int64) Body {
	return &bodyFileImpl{
		path:          path,
		contentLength: contentLength,
	}
}

type bodyFileImpl struct {
	path          string
	contentLength int64
}

func (z bodyFileImpl) ContentLength() int64 {
	return z.contentLength
}

func (z bodyFileImpl) Body() []byte {
	return nil
}

func (z bodyFileImpl) File() string {
	return z.path
}

func (z bodyFileImpl) IsFile() bool {
	return true
}

type bodyMemoryImpl struct {
	content []byte
}

func (z bodyMemoryImpl) ContentLength() int64 {
	return int64(len(z.content))
}

func (z bodyMemoryImpl) Body() []byte {
	return z.content
}

func (z bodyMemoryImpl) File() string {
	return ""
}

func (z bodyMemoryImpl) IsFile() bool {
	return false
}
