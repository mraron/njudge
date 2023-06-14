package sandbox_test

import (
	"bytes"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixBuffer(t *testing.T) {
	tests := []struct {
		name    string
		len     int
		written [][]byte
		want    []byte
	}{
		{"empty", 0, [][]byte{[]byte("abc")}, []byte("")},
		{"len=1", 1, [][]byte{[]byte("abc")}, []byte("a")},
		{"len=2", 2, [][]byte{[]byte("abc")}, []byte("ab")},
		{"len=3", 3, [][]byte{[]byte("abc")}, []byte("abc")},
		{"len=40 too long", 40, [][]byte{[]byte("abc")}, []byte("abc")},
		{"parts", 40, [][]byte{[]byte("a"), []byte("b"), []byte("d")}, []byte("abd")},
		{"parts2", 2, [][]byte{[]byte("a"), []byte("q"), []byte("d")}, []byte("aq")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := &bytes.Buffer{}

			pb := sandbox.NewPrefixBuffer(buf, test.len)
			for _, toWrite := range test.written {
				n, err := pb.Write(toWrite)
				assert.Equal(t, n, len(toWrite))
				assert.Nil(t, err)
			}

			assert.Equal(t, pb.Prefix(), test.want)
		})
	}
}
