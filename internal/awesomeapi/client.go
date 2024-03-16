package awesomeapi

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
	BASE_URL              = "https://economia.awesomeapi.com.br/json/last"
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

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetQuote(ctx *context.Context, pair string) (Quotes, error) {

	url := c.BaseURL.JoinPath(pair).String()

	req, err := http.NewRequestWithContext(*ctx, "GET", url, nil)
	if err != nil {
		return Quotes{}, err
	}

	var res Quotes
	if err := c.sendRequest(req, &res); err != nil {
		return Quotes{}, err
	}

	return res, nil
}

func (c *Client) GetUSDBRLQuote(ctx *context.Context) (Quote, error) {

	quotes, err := c.GetQuote(ctx, "USD-BRL")
	if err != nil {
		return Quote{}, err
	}

	return quotes.USDBRL, nil
}
