package judge

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type submission struct {
	Submission

	Stream      bool   `json:"stream"`
	CallbackUrl string `json:"callback_url"`

	c    Callbacker
	done chan bool
}

type Submission struct {
	Id       string `json:"id"`
	Problem  string `json:"problem"`
	Language string `json:"language"`
	Source   []byte `json:"source"`
}

type Client interface {
	SubmitCallback(context.Context, Submission, string) error
	SubmitStream(context.Context, Submission, chan Status) error
	Status(ctx context.Context) (ServerStatus, error)
}

type defaultClient struct {
	client *http.Client

	url   string
	token string
}

func NewClient(url, token string) Client {
	return &defaultClient{url: url, token: token, client: &http.Client{}}
}

func (dc defaultClient) submit(ctx context.Context, sub submission) (*http.Response, error) {
	dst := dc.url + "/judge"

	buf := bytes.Buffer{}

	enc := json.NewEncoder(&buf)
	err := enc.Encode(sub)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", dst, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if dc.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", dc.token))
	}

	resp, err := dc.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (dc defaultClient) SubmitCallback(ctx context.Context, sub Submission, callback string) error {
	resp, err := dc.submit(ctx, submission{Submission: sub, Stream: false, CallbackUrl: callback})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(data) != "queued" {
		return errors.New(string(data))
	}

	return nil
}

func (dc defaultClient) SubmitStream(ctx context.Context, sub Submission, res chan Status) error {
	var err error

	resp, err := dc.submit(ctx, submission{Submission: sub, Stream: true, CallbackUrl: ""})
	if err != nil {
		return err
	}

	done := make(chan bool, 1)

	go func() {
		s := bufio.NewScanner(resp.Body)

		for s.Scan() {
			status := Status{}
			if err = json.Unmarshal([]byte(s.Text()), &status); err != nil {
				return
			}

			res <- status
		}

		err = resp.Body.Close()
		done <- true
	}()

	select {
	case <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (dc defaultClient) Status(ctx context.Context) (ServerStatus, error) {
	dst := dc.url + "/status"

	req, err := http.NewRequestWithContext(ctx, "GET", dst, nil)
	if err != nil {
		return ServerStatus{}, err
	}

	if dc.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", dc.token))
	}

	resp, err := dc.client.Do(req)
	if err != nil {
		return ServerStatus{}, err
	}

	ans := ServerStatus{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ans)
	if err != nil {
		return ServerStatus{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ServerStatus{}, errors.New("judger returned: " + resp.Status)
	}

	return ans, nil
}
