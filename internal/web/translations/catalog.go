// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"hu_HU": &dictionary{index: hu_HUIndex, data: hu_HUData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{}

var en_USIndex = []uint32{ // 1 elements
	0x00000000,
} // Size: 28 bytes

const en_USData string = ""

var hu_HUIndex = []uint32{ // 1 elements
	0x00000000,
} // Size: 28 bytes

const hu_HUData string = ""

// Total table size 56 bytes (0KiB); checksum: 32CC3CAD
