package form

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x := []byte("BC=123&AB=234")
	ab := Kv{K: []byte("ab")}
	bc := Kv{K: []byte("bc")}
	Parse(x, &ab, &bc)
	fmt.Printf(
		"%s: %s, %s: %s\n",
		ab.K, x[ab.Bgn:ab.End],
		bc.K, x[bc.Bgn:bc.End],
	)
}
