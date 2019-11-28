package eraser

import "testing"

func TestErase(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"Hello", "Hello"},
		{"\033[1;34mHello\033[0m", "Hello"},
		{"\x1b[31;1mHe\x1b[0ml\x1b[36;1mlo\x1b[0m", "Hello"},
	}

	e := NewEraser()
	for _, tt := range tests {
		got := e.Erase(tt.in)
		if tt.want != got {
			t.Errorf("Erase(): got = %v ,want = %v", got, tt.want)
		}
	}
}
