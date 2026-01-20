package session

import (
	"io"
)

type Session struct {
	host string
	port string
	in   io.Writer
	out  io.Reader
}

func NewSession(host string, port string,
	in io.Writer, out io.Reader,
) Session {
	return Session{
		host: host,
		port: port,
		in:   in,
		out:  out,
	}
}

func (s *Session) GetWriter() io.Writer {
	return s.in
}

func (s *Session) GetReader() io.Reader {
	return s.out
}
