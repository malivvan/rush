package rush

import "io"

type Config struct {
	Args   []string
	StdIn  io.Reader
	StdOut io.Writer
	StdErr io.Writer
}
