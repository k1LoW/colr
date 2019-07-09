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
	ctx        context.Context
	replaces   []string
	regexps    []*regexp.Regexp
	colorFuncs []colorFunc
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
	c := color.New()
	c.Enable()
	return &Painter{
		ctx:      ctx,
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

func (p *Painter) AddColor(inn io.Reader) <-chan string {
	in := bufio.NewReader(inn)
	out := make(chan string)
	fLen := len(p.colorFuncs)
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
				s = strings.ReplaceAll(s, r, p.colorFuncs[i%fLen](r, color.B))
			}
			for i, re := range p.regexps {
				s = re.ReplaceAllString(s, p.colorFuncs[(i+rLen)%fLen]("$1", color.B))
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
