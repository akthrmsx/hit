package main

import (
	"strings"
	"testing"
)

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}

func testRun(args ...string) (*testEnv, error) {
	var te testEnv
	e := env{
		stdout: &te.stdout,
		stderr: &te.stderr,
		args:   append([]string{"hit"}, args...),
		dryrun: true,
	}
	err := run(&e)
	return &te, err
}

func TestRun(t *testing.T) {
	t.Parallel()

	te, err := testRun("http://test")

	if err != nil {
		t.Fatalf("\nwant: nil\ngot: %#v\n", err)
	}

	if n := te.stdout.Len(); n == 0 {
		t.Errorf("\nwant: >0\ngot: 0\n")
	}

	if n := te.stderr.Len(); n != 0 {
		t.Errorf("\nwant: 0\ngot: %d\n", te.stdout.Len())
	}
}

func TestRunError(t *testing.T) {
	t.Parallel()

	te, err := testRun("test")

	if err == nil {
		t.Fatalf("\nwant: error\ngot: nil\n")
	}

	if n := te.stderr.Len(); n == 0 {
		t.Errorf("\nwant: >0\ngot: 0\n")
	}
}
