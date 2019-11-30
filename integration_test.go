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
		got  []byte
		want string
		err  error
	)
	if runtime.GOOS == "windows" {
		got, err = exec.Command("cmd", "/c", `echo Hello| .\colr.exe He lo xx | .\colr.exe --erase`).Output()
		want = "Hello\r\n"
	} else {
		got, err = exec.Command("bash", "-c", `echo Hello| ./colr He lo xx | ./colr --erase`).Output()
		want = "Hello\n"
	}
	if err != nil {
		t.Fatal(err)
	}

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
