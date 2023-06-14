package sandbox

import "io"

type PrefixBuffer struct {
	prefix    []byte
	prefixLen int

	io.Writer
}

func NewPrefixBuffer(w io.Writer, len int) *PrefixBuffer {
	return &PrefixBuffer{make([]byte, 0, len), len, w}
}

func (pb *PrefixBuffer) Write(bs []byte) (int, error) {
	for j := 0; len(pb.prefix) < pb.prefixLen && j < len(bs); j++ {
		pb.prefix = append(pb.prefix, bs[j])
	}

	if pb.Writer != nil {
		return pb.Writer.Write(bs)
	}

	return len(bs), nil
}

func (pb *PrefixBuffer) Prefix() []byte {
	return pb.prefix
}
