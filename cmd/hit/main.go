package main

import (
	"fmt"
	"io"
	"os"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

const (
	defaultRequests    = 100
	defaultConcurrency = 1
)

type env struct {
	stdout io.Writer
	stderr io.Writer
	args   []string
	dryrun bool
}

func main() {
	e := env{
		stdout: os.Stdout,
		stderr: os.Stderr,
		args:   os.Args,
	}

	if err := run(&e); err != nil {
		os.Exit(1)
	}
}

func run(e *env) error {
	c := config{
		requests:    defaultRequests,
		concurrency: defaultConcurrency,
	}

	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
		return err
	}

	fmt.Fprintf(e.stdout, "%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.requests, c.url, c.concurrency)

	if e.dryrun {
		return nil
	}

	if err := runHit(&c, e.stdout); err != nil {
		fmt.Fprintf(e.stderr, "\nerror occurred: %v\n", err)
		return err
	}

	return nil
}

func runHit(_ *config, _ io.Writer) error {
	return nil
}
