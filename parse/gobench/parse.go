package gobench

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/benchmark/parse"
)

var packageRe = regexp.MustCompile(`^(ok|FAIL)\s+(.+?)\s+(.+)$`)

// Parse a bunch of gobench statistics out of test output
func Parse(in io.Reader) (map[string][]*parse.Benchmark, error) {
	var (
		benchmarks []*parse.Benchmark
		out        = make(map[string][]*parse.Benchmark)
	)

	lines := bufio.NewScanner(in)
	for lines.Scan() {
		if strings.HasPrefix(lines.Text(), "Benchmark") {
			// try to parse a benchmark
			bm, err := parse.ParseLine(lines.Text())
			if err != nil {
				fmt.Println(err)
				continue
			}
			benchmarks = append(benchmarks, bm)
		} else if matches := packageRe.FindSubmatch(lines.Bytes()); len(matches) != 0 && benchmarks != nil {
			// that sounds like a package name following the benchmarks. We need
			// that too!
			out[string(matches[2])] = benchmarks
			benchmarks = nil
		}
	}
	if err := lines.Err(); err != nil {
		return nil, errors.Wrap(err, "scanning failed")
	}

	if benchmarks != nil {
		out["unknown-package"] = benchmarks
	}

	return out, nil
}
