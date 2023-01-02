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

type Submission struct {
	Id       string `json:"id"`
	Problem  string `json:"problem"`
	Language string `json:"language"`
	Source   []byte `json:"source"`

	Stream      bool   `json:"stream"`
	CallbackUrl string `json:"callback_url"`

	c    Callbacker
	done chan bool
}

type Client struct {
	client *http.Client

	url   string
	token string
}

type ClientOption func(*Client)

func NewClient(url string, opts ...ClientOption) *Client {
	client := &Client{url: url, client: &http.Client{}}
	for i := range opts {
		opts[i](client)
	}

	return client
}

func (dc Client) submit(ctx context.Context, sub Submission) (*http.Response, error) {
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

func (dc Client) SubmitCallback(ctx context.Context, sub Submission, callback string) error {
	sub.Stream = false
	sub.CallbackUrl = callback
	resp, err := dc.submit(ctx, sub)
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

func (dc Client) SubmitStream(ctx context.Context, sub Submission, res chan SubmissionStatus) error {
	var err error

	sub.Stream = true
	sub.CallbackUrl = ""
	resp, err := dc.submit(ctx, sub)
	if err != nil {
		return err
	}

	done := make(chan bool, 1)

	go func() {
		s := bufio.NewScanner(resp.Body)

		for s.Scan() {
			status := SubmissionStatus{}
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

func (dc Client) Status(ctx context.Context) (Status, error) {
	dst := dc.url + "/status"

	req, err := http.NewRequestWithContext(ctx, "GET", dst, nil)
	if err != nil {
		return Status{}, err
	}

	if dc.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", dc.token))
	}

	resp, err := dc.client.Do(req)
	if err != nil {
		return Status{}, err
	}

	ans := Status{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ans)
	if err != nil {
		return Status{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Status{}, errors.New("judger returned: " + resp.Status)
	}

	return ans, nil
}
