package painter

import (
	"testing"
)

func TestPaint(t *testing.T) {
	var tests = []struct {
		strs []string
		in   string
		want string
	}{
		{[]string{}, "Hello", "Hello"},
		{[]string{"Hello"}, "Hello", "\x1b[31;1mHello\x1b[0m"},
		{[]string{"He"}, "Hello", "\x1b[31;1mHe\x1b[0mllo"},
		{[]string{"l"}, "Hello", "He\x1b[31;1ml\x1b[0m\x1b[31;1ml\x1b[0mo"},
	}

	for _, tt := range tests {
		p := NewPainter(tt.strs)
		got := p.Paint(tt.in)
		if tt.want != got {
			t.Errorf("Paint(): got = %#v ,want = %#v", got, tt.want)
		}
	}
}
