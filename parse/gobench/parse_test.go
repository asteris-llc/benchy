package gobench_test

import (
	"bytes"
	"testing"

	"github.com/asteris-llc/benchy/parse/gobench"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("basic", func(t *testing.T) {
		buf := bytes.NewBuffer(testBasicInput)

		out, err := gobench.Parse(buf)
		assert.NoError(t, err)
		assert.Contains(t, out, "github.com/asteris-llc/converge/graph")
	})

	t.Run("no-benchmarks", func(t *testing.T) {
		buf := bytes.NewBuffer(testNoBenchmarks)

		out, err := gobench.Parse(buf)
		assert.NoError(t, err)
		assert.NotContains(t, out, "github.com/asteris-llc/converge/helpers/fakerenderer")
	})

	t.Run("no-package", func(t *testing.T) {
		buf := bytes.NewBuffer(testNoPackage)

		out, err := gobench.Parse(buf)
		assert.NoError(t, err)
		if assert.Contains(t, out, "unknown-package") {
			assert.Len(t, out["unknown-package"], 1)
		}
	})
}

var (
	testBasicInput = []byte(`BenchmarkAddThenGet-4     	 1000000	      2896 ns/op	     407 B/op	       4 allocs/op
BenchmarkCopyParallel-4   	  500000	      2688 ns/op	    2944 B/op	      68 allocs/op
ok  	github.com/asteris-llc/converge/graph	4.352s`)
	testNoBenchmarks = []byte(`ok  	github.com/asteris-llc/converge/helpers/fakerenderer	0.032s`)
	testNoPackage    = []byte(`BenchmarkCopyParallel-4   	  500000	      2688 ns/op	    2944 B/op	      68 allocs/op`)
)
