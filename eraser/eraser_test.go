package eraser

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

func TestErase(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"Hello", "Hello"},
		{"\033[1;34mHello\033[0m", "Hello"},
		{"\x1b[31;1mHe\x1b[0ml\x1b[36;1mlo\x1b[0m", "Hello"},
	}

	for _, tt := range tests {
		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		ctx := context.Background()
		e := NewEraser(stdout)
		stdin.WriteString(fmt.Sprintf("%s\n", tt.in))
		err := e.Handle(ctx, stdin)
		if err != nil {
			t.Fatal(err)
		}
		got, err := stdout.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		want := fmt.Sprintf("%s\n", tt.want)
		if want != got {
			t.Errorf("Paint(): got = %#v ,want = %#v", got, want)
		}
	}
}
