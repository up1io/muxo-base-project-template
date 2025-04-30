package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(1)
}
