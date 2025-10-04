package main

import (
	"io"
	"testing"
)

type test struct {
	name string
	args []string
	want config
}

func TestParseArgs(t *testing.T) {
	t.Parallel()

	tests := []test{
		{
			name: "all_flags",
			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
			want: config{url: "http://test", requests: 10, concurrency: 5, rps: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config

			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
				t.Fatalf("\nwant: nil\ngot: %#v\n", err)
			}

			if got != tt.want {
				t.Errorf("\nwant: %#v\ngot: %#v\n", tt.want, got)
			}
		})
	}
}

func TestParseArgsError(t *testing.T) {
	t.Parallel()

	tests := []test{
		{
			name: "n_syntax",
			args: []string{"-n=one", "http://test"},
		},
		{
			name: "n_zero",
			args: []string{"-n=0", "http://test"},
		},
		{
			name: "n_negative",
			args: []string{"-n=-1", "http://test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := parseArgs(&config{}, tt.args, io.Discard); err == nil {
				t.Fatalf("\nwant: error\ngot: nil\n")
			}
		})
	}
}
