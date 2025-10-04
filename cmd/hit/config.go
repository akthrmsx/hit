package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type config struct {
	url         string
	requests    int
	concurrency int
	rps         int
}

func parseArgs(c *config, args []string, stderr io.Writer) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.Var(asPositiveInt(&c.requests), "n", "Number of requests")
	fs.Var(asPositiveInt(&c.concurrency), "c", "Concurrency level")
	fs.Var(asPositiveInt(&c.rps), "rps", "Requests per second")

	if err := fs.Parse(args); err != nil {
		return err
	}

	c.url = fs.Arg(0)

	if err := validateArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

	return nil
}

func validateArgs(c *config) error {
	u, err := url.Parse(c.url)

	if err != nil {
		return fmt.Errorf("invalid value %q for url: %w", c.url, err)
	}

	if c.url == "" || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid value %q for url: requires a valid url", c.url)
	}

	if c.requests < c.concurrency {
		return fmt.Errorf("invalid value %d for flag -n: should be greater than flag -c: %d", c.requests, c.concurrency)
	}

	return nil
}

type positiveInt int

func asPositiveInt(p *int) *positiveInt {
	return (*positiveInt)(p)
}

func (pi *positiveInt) String() string {
	return strconv.Itoa(int(*pi))
}

func (pi *positiveInt) Set(s string) error {
	n, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	if n <= 0 {
		return fmt.Errorf("should be greater than zero")
	}

	*pi = positiveInt(n)
	return nil
}
