package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient       *http.Client
	httpHeaders      map[string]string
	address          string
	apiKey           string
	timeout          time.Duration
	numRetries       int
	retryStatusCodes []string
}

func NewClient(
	address string,
	apiKey string,
	timeout time.Duration,
	numRetries int,
	retryStatusCodes []string) (*Client, error) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpClient := &http.Client{Transport: tr}

	return &Client{
		httpClient: httpClient,
		address:    address,
		apiKey:     apiKey,
		timeout:    10 * time.Second,
	}, nil
}

func (c *Client) SetTimeout(t time.Duration) {
	c.timeout = t
}

func (c *Client) request(method, requestPath string, body io.Reader, responseStruct interface{}) error {

	var resp *http.Response
	urlPath, err := url.JoinPath(c.address, requestPath)
	req, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return err
	}

	if c.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	if c.httpHeaders != nil {
		for k, v := range c.httpHeaders {
			req.Header.Add(k, v)
		}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	//fmt.Println(req)

	resp, err = c.httpClient.Do(req.WithContext(ctx))
	if err != nil || !(resp.StatusCode > 199 && resp.StatusCode < 300) {
		//fmt.Println("error happend", err, resp.Status)
		return err

		// TODO: handle non-200 status codes
	}

	//  POST-Request: body <-- bytes.NewBuffer(json_data) <--- json.Marshal(map[string]string)

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(resBody, responseStruct)
	if err != nil {
		return err
	}
	return nil

}
