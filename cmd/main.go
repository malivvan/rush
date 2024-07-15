package main

import (
	"github.com/malivvan/rush"
	"os"
)

func main() {
	cfg := rush.Config{
		StdIn:  os.Stdin,
		StdOut: os.Stdout,
		StdErr: os.Stderr,
		Args:   os.Args[1:],
	}
	rush.RunString(cfg, "test", "echo hello")
	rush.REPL(cfg)
}
