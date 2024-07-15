package rush

import (
	"context"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
	"os"
	"strings"
)

func RunString(cfg Config, name string, code string) error {
	r, err := interp.New(interp.StdIO(cfg.StdIn, cfg.StdOut, cfg.StdErr))
	if err != nil {
		return err
	}
	prog, err := syntax.NewParser().Parse(strings.NewReader(code), name)
	if err != nil {
		return err
	}
	r.Reset()
	ctx := context.Background()
	return r.Run(ctx, prog)
}

func RunPath(cfg Config, path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return RunString(cfg, path, string(f))
}
