package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SubmissionStatus struct {
	Response
	Time   time.Time
}

type Callbacker interface {
	Callback(Response) error
}

type ChanCallback struct {
	c chan Response
}

func NewChanCallback(c chan Response) *ChanCallback {
	return &ChanCallback{c}
}

func (cc *ChanCallback) Callback(r Response) error {
	cc.c <- r
	if (r.Done) {
		close(cc.c)
	}

	return nil
}

type WriterCallback struct {
	enc       *json.Encoder
	err       error
	afterFunc func()
}

func NewWriterCallback(w io.Writer, afterFunc func()) *WriterCallback {
	return &WriterCallback{enc: json.NewEncoder(w), err: nil, afterFunc: afterFunc}
}

func (wc *WriterCallback) Callback(r Response) error {
	if wc.err != nil {
		return wc.err
	}

	wc.err = wc.enc.Encode(SubmissionStatus{Response: r, Time: time.Now()})
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

func (h HTTPCallback) Callback(r Response) error {
	raw := SubmissionStatus{Response: r, Time: time.Now()}

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
		return fmt.Errorf("Callback error: %s %s", resp.Status, resp.Body)
	}

	return nil
}
