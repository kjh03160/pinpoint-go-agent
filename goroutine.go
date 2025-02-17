package pinpoint

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
)

type goroutine struct {
	id     uint64
	header string
	state  string
	trace  string

	span activeSpanInfo

	frozen bool
	buf    *bytes.Buffer
}

func (g *goroutine) addLine(line string) {
	if g.frozen {
		return
	}

	g.buf.WriteString(line)
	g.buf.WriteString("\n")
}

func (g *goroutine) freeze() {
	if g.frozen {
		return
	}

	g.frozen = true
	g.trace = g.buf.String()
	g.buf = nil
}

func newGoroutine(line string) *goroutine {
	idx := strings.Index(line, "[")
	parts := strings.Split(line[idx+1:len(line)-2], ",")
	state := strings.TrimSpace(parts[0])

	if id, err := strconv.Atoi(strings.TrimSpace(line[9:idx])); err == nil {
		return &goroutine{
			id:     uint64(id),
			header: line[0 : idx-1],
			state:  state,
			buf:    &bytes.Buffer{},
		}
	} else {
		Log("cmd").Errorf("convert goroutine id: %v", err)
	}

	return nil
}

type goroutineDump struct {
	goroutines []*goroutine
}

func (gd *goroutineDump) add(g *goroutine) {
	gd.goroutines = append(gd.goroutines, g)
}

func (gd *goroutineDump) search(s string) *goroutine {
	for _, g := range gd.goroutines {
		if g.header == s {
			return g
		}
	}

	return nil
}

func newGoroutineDump() *goroutineDump {
	return &goroutineDump{
		goroutines: []*goroutine{},
	}
}

func dumpGoroutine() (dump *goroutineDump) {
	var b bytes.Buffer
	buf := bufio.NewWriter(&b)

	defer func() {
		if e := recover(); e != nil {
			Log("cmd").Errorf("profile goroutine: %v", e)
			dump = nil
		}
	}()

	if p := pprof.Lookup("goroutine"); p != nil {
		if err := p.WriteTo(buf, 2); err != nil {
			Log("cmd").Errorf("profile goroutine: %v", err)
			return nil
		}
	}

	dump = parseProfile(&b)
	return
}

func parseProfile(r io.Reader) *goroutineDump {
	dump := newGoroutineDump()
	var g *goroutine

	g = nil
	scanner := bufio.NewScanner(r)
	startLinePattern := regexp.MustCompile(`^goroutine\s+(\d+)\s+\[(.*)\]:$`)

	for scanner.Scan() {
		line := scanner.Text()
		if startLinePattern.MatchString(line) {
			g = newGoroutine(line)
			if g == nil {
				return nil
			}

			if v, ok := realTimeActiveSpan.Load(g.id); ok {
				g.span = v.(activeSpanInfo)
				dump.add(g)
			}
		} else if line == "" {
			// End of a goroutine section.
			if g != nil {
				g.freeze()
			}
			g = nil
		} else if g != nil {
			g.addLine(line)
		}
	}

	if err := scanner.Err(); err != nil {
		Log("cmd").Errorf("scan goroutine profile: %v", err)
		return nil
	}

	if g != nil {
		g.freeze()
	}

	return dump
}

// Sourced: https://github.com/golang/net/blob/master/http2/gotrack.go

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	bp := littleBuf.Get().(*[]byte)
	defer littleBuf.Put(bp)
	b := *bp
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := parseUintBytes(b, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	fmt.Sprintf("curGoroutineID2: %v", n)
	return n
}

var littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}

// parseUintBytes is like strconv.ParseUint, but using a []byte.
func parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
	var cutoff, maxVal uint64

	if bitSize == 0 {
		bitSize = int(strconv.IntSize)
	}

	s0 := s
	switch {
	case len(s) < 1:
		err = strconv.ErrSyntax
		goto Error

	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal, hex prefix.
		switch {
		case s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X'):
			base = 16
			s = s[2:]
			if len(s) < 1 {
				err = strconv.ErrSyntax
				goto Error
			}
		case s[0] == '0':
			base = 8
		default:
			base = 10
		}

	default:
		err = errors.New("invalid base " + strconv.Itoa(base))
		goto Error
	}

	n = 0
	cutoff = cutoff64(base)
	maxVal = 1<<uint(bitSize) - 1

	for i := 0; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}
		if int(v) >= base {
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}

		if n >= cutoff {
			// n*base overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n *= uint64(base)

		n1 := n + uint64(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n = n1
	}

	return n, nil

Error:
	return n, &strconv.NumError{Func: "ParseUint", Num: string(s0), Err: err}
}

// Return the first number n such that n*base >= 1<<64.
func cutoff64(base int) uint64 {
	if base < 2 {
		return 0
	}
	return (1<<64-1)/uint64(base) + 1
}
