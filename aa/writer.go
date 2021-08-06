package aa

import "io"

type AAWriter struct {
	w io.Writer
}

func NewAAWriter(w io.Writer) AAWriter {
	return AAWriter{
		w: w,
	}
}
