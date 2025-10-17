package form

import (
	"bytes"
)

type (
	V  struct{ Bgn, End int }
	Kv struct {
		K []byte
		V
	}
)

const (
	aPre = 'A' - 1
	zSu  = 'Z' + 1
	kSep = '='
	vSep = '&'
)

func Parse(x []byte, kvs ...*Kv) {
	xLen := len(x)
	kvsLen := len(kvs)
	i := 0
	for i < xLen {
		j := i
		for i < xLen && x[i] != kSep {
			if x[i] > aPre && x[i] < zSu {
				x[i] |= 0b100000
			}
			i++
		}
		k := i
		for i < xLen && x[i] != vSep {
			i++
		}
		for l := 0; l < kvsLen; l++ {
			//fmt.Println(kvs[l].K, x[j:k])
			if bytes.Equal(
				kvs[l].K, x[j:k],
			) {
				kvs[l].V = V{k + 1, i}
				//copy(kvs[l:], kvs[l+1:])
				kvs[l] = kvs[kvsLen-1]
				kvsLen--
				if kvsLen == 0 {
					return
				}
				break
			}
		}
		i++
	}
}
