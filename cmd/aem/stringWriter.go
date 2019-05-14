package main

type stringWriter struct {
	contentBytes []byte
}

func (w *stringWriter) Write(p []byte) (n int, err error) {
	w.contentBytes = append(w.contentBytes, p...)
	return len(p), nil
}

func (w *stringWriter) String() string {
	return string(w.contentBytes)
}
