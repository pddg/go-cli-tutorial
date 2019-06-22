package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_printVersion(t *testing.T) {
	out := new(bytes.Buffer)
	printVersion(out)
	want := fmt.Sprintf("dailyrepo %s\n", version)
	got := out.String()
	if got != want {
		t.Fatalf("printVersion() returned %#v, but want %#v\n", got, want)
	}
}
