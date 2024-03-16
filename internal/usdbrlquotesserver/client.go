package usdbrlquotesserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	APPLICATION_JSON_UTF8 = "application/json; charset=utf-8"
	BASE_URL              = "http://localhost:8080/cotacao"
)

type Client struct {
	BaseURL *url.URL
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient() *Client {
	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		BaseURL: baseUrl,
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", APPLICATION_JSON_UTF8)
	req.Header.Set("Accept", APPLICATION_JSON_UTF8)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("erro desconhecido, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetQuote(ctx *context.Context) (float64, error) {

	req, err := http.NewRequestWithContext(*ctx, "GET", c.BaseURL.String(), nil)
	if err != nil {
		return 0, err
	}

	var res float64
	if err := c.sendRequest(req, &res); err != nil {
		return 0, err
	}

	return res, nil
}
