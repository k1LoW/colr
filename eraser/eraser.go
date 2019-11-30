package eraser

import (
	"bufio"
	"context"
	"io"

	"github.com/mattn/go-colorable"
)

type Eraser struct {
	out io.Writer
}

func NewEraser(out io.Writer) *Eraser {
	return &Eraser{
		out: out,
	}
}

func (e *Eraser) Handle(ctx context.Context, inn io.Reader) error {
	in := bufio.NewReader(inn)
	out := colorable.NewNonColorable(e.out)

	for {
		s, err := in.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			break
		default:
			_, err = out.Write(s)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
