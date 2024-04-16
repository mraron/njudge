package memory

import (
	"fmt"
	"strconv"
)

// Amount represents an amount of bytes of memory.
type Amount int64

const (
	Byte Amount = 1
	KB          = 1000 * Byte
	KiB         = 1024 * Byte
	MB   Amount = 1000 * KB
	MiB         = 1024 * KiB
	GB   Amount = 1000 * MB
	GiB         = 1024 * MiB
)

func (x *Amount) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%dB\"", *x)), nil
}

func (x *Amount) UnmarshalJSON(bs []byte) error {
	if len(bs) == 0 {
		return nil
	}

	if bs[0] != '"' {
		tmp, err := strconv.Atoi(string(bs))
		if err != nil {
			return err
		}
		*x = Amount(tmp) * KiB
		return nil
	}
	tmp, err := strconv.Atoi(string(bs[1 : len(bs)-2]))
	if err != nil {
		return err
	}
	*x = Amount(tmp)
	return nil
}
