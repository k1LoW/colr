// +build integration

package main

import (
	"os/exec"
	"runtime"
	"testing"
)

func init() {
	err := exec.Command("go", "build").Run()
	if err != nil {
		panic(err)
	}
}

func TestPaintAndErase(t *testing.T) {
	var (
		got []byte
		err error
	)
	if runtime.GOOS == "windows" {
		got, err = exec.Command("cmd", "/c", `echo Hello | .\colr.exe He | .\colr.exe --erase`).Output()
	} else {
		got, err = exec.Command("bash", "-c", `echo Hello | ./colr He | ./colr --erase`).Output()
	}
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello\n"
	if want != string(got) {
		t.Errorf("got = %#v ,want = %#v", string(got), want)
	}
}

type testcase struct {
	name            string
	cmd             []string
	want            string
	processFinished bool
}
