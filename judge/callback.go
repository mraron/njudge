package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/mraron/njudge/utils/problems"
	"io"
	"net/http"
	"time"
)

type Callbacker interface {
	Callback(string, problems.Status, bool) error
}

//@TODO Create cached callback

type CombineCallback struct {
	lst []Callbacker
}

func NewCombineCallback(lst ...Callbacker) CombineCallback {
	return CombineCallback{lst}
}

func (c CombineCallback) Callback(test string, status problems.Status, done bool) error {
	var err error
	for _, cb := range c.lst {
		multierror.Append(err, cb.Callback(test, status, done))
	}

	return err
}

func (c *CombineCallback) Append(callbacker Callbacker) {
	c.lst = append(c.lst, callbacker)
}

type WriterCallback struct {
	enc *json.Encoder
	err error
	afterFunc func()
}

func NewWriterCallback(w io.Writer, afterFunc func()) *WriterCallback {
	return &WriterCallback{enc: json.NewEncoder(w), err: nil, afterFunc: afterFunc}
}

func (wc* WriterCallback) Callback(test string, status problems.Status, done bool) error {
	if wc.err != nil {
		return wc.err
	}

	wc.err = wc.enc.Encode(Status{test, status, done, time.Now()})
	if wc.err == nil {
		wc.afterFunc()
	}

	return wc.err
}

func (wc *WriterCallback) Error() error {
	return wc.err
}

type HTTPCallback struct {
	url string
}

func NewHTTPCallback(url string) HTTPCallback {
	return HTTPCallback{url}
}

func (h HTTPCallback) Callback(test string, status problems.Status, done bool) error {
	raw := Status{test, status, done, time.Now()}

	buf := &bytes.Buffer{}

	data := json.NewEncoder(buf)
	err := data.Encode(raw)
	if err != nil {
		return err
	}

	resp, err := http.Post(h.url, "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Callback error: %d %s", resp.Status, resp.Body)
	}

	return nil
}
