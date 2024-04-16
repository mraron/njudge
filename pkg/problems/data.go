package problems

import (
	"bytes"
	"io"
)

type Data interface {
	Value() ([]byte, error)
	ValueReader() (io.ReadCloser, error)
	Type() string
	IsText() bool
	IsHTML() bool
	IsPDF() bool
	String() string
}

type NamedData interface {
	Data
	Name() string
}

type LocalizedData interface {
	Data
	Locale() string
}

type BytesData struct {
	Loc string //Locale
	Val []byte //Value
	Typ string //Type (for example "text", "application/pdf", "text/html")
	Nam string //Name
}

func (s BytesData) Name() string {
	return s.Nam
}

func (s BytesData) Value() ([]byte, error) {
	return s.Val, nil
}

func (s BytesData) ValueReader() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewBuffer(s.Val)), nil
}

func (s BytesData) Type() string {
	return s.Typ
}

var (
	DataTypeText = "text"
	DataTypeHTML = "text/html"
	DataTypePDF  = "application/pdf"
)

func (s BytesData) IsText() bool {
	return s.Typ == DataTypeText
}

func (s BytesData) IsHTML() bool {
	return s.Typ == DataTypeHTML
}

func (s BytesData) IsPDF() bool {
	return s.Typ == DataTypePDF
}

func (s BytesData) Locale() string {
	return s.Loc
}

func (s BytesData) String() string {
	return string(s.Val)
}

type (
	Contents    []LocalizedData
	Attachments []NamedData
)

func (cs Contents) Locales() []string {
	lst := make(map[string]bool)
	for _, val := range cs {
		lst[val.Locale()] = true
	}

	ans := make([]string, len(lst))

	ind := 0
	for key := range lst {
		ans[ind] = key
		ind++
	}

	return ans
}

func (cs Contents) FilterByType(mime string) Contents {
	res := make(Contents, 0)
	for _, val := range cs {
		if mime == val.Type() {
			res = append(res, val)
		}
	}

	return res
}

func (cs Contents) FilterByLocale(locale string) Contents {
	res := make(Contents, 0)
	for _, val := range cs {
		if locale == val.Locale() {
			res = append(res, val)
		}
	}

	return res
}
