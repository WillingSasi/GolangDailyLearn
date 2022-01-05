package main

import "io"

type Socket struct {
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

func (socket *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (socket *Socket) Close() error {
	return nil
}

func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

func usingCloser(closer io.Closer) {
	closer.Close()
}

func main() {
	s := new(Socket)

	usingWriter(s)

	usingCloser(s)
}
