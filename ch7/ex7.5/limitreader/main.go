package limitreader

import "io"

type limitReader struct {
	reader io.Reader
	limit  int
	bytes  int
}

func (lr *limitReader) Read(p []byte) (bytes int, err error) {
	// Check for exhaustion before, because the reader may have been read from before.
	if lr.isExhausted() {
		return 0, io.EOF
	}

	bytes, err = lr.reader.Read(p[:lr.limit]) // Read bytes from p up to the limit.
	lr.bytes += bytes                         // Keep track of how many bytes we've read.

	// Check for exhaustion after because the reader may have been read for the first time.
	if lr.isExhausted() {
		err = io.EOF
	}

	return
}

// If the limit has been exhausted we need to return an EOF error.
func (lr *limitReader) isExhausted() bool {
	return lr.bytes >= lr.limit
}

func LimitReader(reader io.Reader, limit int) io.Reader {
	return &limitReader{reader: reader, limit: limit}
}
