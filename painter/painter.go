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

type colorFunc func(interface{}, ...string) string

type Painter struct {
	out        chan string
	replaces   []string
	regexps    []*regexp.Regexp
	colorFuncs []colorFunc
}

// NewPainter ...
func NewPainter(strs []string) *Painter {
	replaces := []string{}
	regexps := []*regexp.Regexp{}
	for _, s := range strs {
		if !strings.ContainsAny(s, ".*[(^-?") {
			replaces = append(replaces, s)
		} else {
			regexps = append(regexps, regexp.MustCompile("("+s+")"))
		}
	}
	c := color.New()
	c.Enable()
	return &Painter{
		out:      make(chan string),
		replaces: replaces,
		regexps:  regexps,
		colorFuncs: []colorFunc{
			func(msg interface{}, styles ...string) string {
				return c.Red(msg, styles...)
			},
			func(msg interface{}, styles ...string) string {
				return c.Cyan(msg, styles...)
			},
			func(msg interface{}, styles ...string) string {
				return c.Yellow(msg, styles...)
			},
			func(msg interface{}, styles ...string) string {
				return c.Magenta(msg, styles...)
			},
			func(msg interface{}, styles ...string) string {
				return c.Green(msg, styles...)
			},
			func(msg interface{}, styles ...string) string {
				return c.Blue(msg, styles...)
			},
		},
	}
}

func (p *Painter) Paint(s string) string {
	fLen := len(p.colorFuncs)
	rLen := len(p.replaces)
	for i, r := range p.replaces {
		s = strings.ReplaceAll(s, r, p.colorFuncs[i%fLen](r, color.B))
	}
	for i, re := range p.regexps {
		s = re.ReplaceAllString(s, p.colorFuncs[(i+rLen)%fLen]("$1", color.B))
	}
	return s
}

func (p *Painter) Handle(ctx context.Context, inn io.Reader) <-chan string {
	in := bufio.NewReader(inn)
	go func() {
		defer close(p.out)
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
				p.out <- p.Paint(s)
			}
		}
	}()

	return p.out
}
