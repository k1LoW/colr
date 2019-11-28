package eraser

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/acarl005/stripansi"
)

type Eraser struct {
	out chan string
}

func NewEraser() *Eraser {
	return &Eraser{
		out: make(chan string),
	}
}

func (e *Eraser) Erase(s string) string {
	return stripansi.Strip(s)
}

func (e *Eraser) Handle(ctx context.Context, inn io.Reader) <-chan string {
	in := bufio.NewReader(inn)

	go func() {
		defer close(e.out)
	L:
		for {
			s, err := in.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			select {
			case <-ctx.Done():
				break L
			default:
				e.out <- e.Erase(s)
			}
		}
	}()

	return e.out
}
