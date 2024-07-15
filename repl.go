package rush

import (
	"context"
	"fmt"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

func REPL(cfg Config) error {
	r, err := interp.New(interp.StdIO(cfg.StdIn, cfg.StdOut, cfg.StdErr))
	if err != nil {
		return err
	}
	parser := syntax.NewParser()
	fmt.Fprintf(cfg.StdOut, "$ ")
	var runErr error
	fn := func(stmts []*syntax.Stmt) bool {
		if parser.Incomplete() {
			fmt.Fprintf(cfg.StdOut, "> ")
			return true
		}
		ctx := context.Background()
		for _, stmt := range stmts {
			runErr = r.Run(ctx, stmt)
			if r.Exited() {
				return false
			}
		}
		fmt.Fprintf(cfg.StdOut, "$ ")
		return true
	}
	if err := parser.Interactive(cfg.StdIn, fn); err != nil {
		return err
	}
	return runErr
}
