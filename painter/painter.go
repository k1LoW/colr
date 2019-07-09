package painter

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/labstack/gommon/color"
)

type Painter struct {
	ctx      context.Context
	replaces []string
	regexps  []*regexp.Regexp
}

type colorFunc func(interface{}, ...string) string

var colorFuncs = []colorFunc{
	color.Red,
	color.Yellow,
	color.Magenta,
	color.Green,
	color.Cyan,
	color.Blue,
}

// NewPainter ...
func NewPainter(ctx context.Context, strs []string) *Painter {
	replaces := []string{}
	regexps := []*regexp.Regexp{}
	for _, s := range strs {
		if strings.IndexAny(s, ".*[(^-?") < 0 {
			replaces = append(replaces, s)
		} else {
			regexps = append(regexps, regexp.MustCompile("("+s+")"))
		}
	}
	return &Painter{
		ctx:      ctx,
		replaces: replaces,
		regexps:  regexps,
	}
}

func (p *Painter) AddColor(inn io.Reader) <-chan string {
	in := bufio.NewReader(inn)
	out := make(chan string)
	fLen := len(colorFuncs)
	rLen := len(p.replaces)

	go func() {
	L:
		for {
			s, err := in.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			for i, r := range p.replaces {
				s = strings.ReplaceAll(s, r, colorFuncs[i%fLen](r, color.B))
			}
			for i, re := range p.regexps {
				s = re.ReplaceAllString(s, colorFuncs[(i+rLen)%fLen]("$1", color.B))
			}
			select {
			case <-p.ctx.Done():
				break L
			default:
				out <- s
			}
		}
	}()

	return out
}
