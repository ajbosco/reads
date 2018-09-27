package goodreads

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mrjones/oauth"
	"github.com/pkg/errors"
)

// Client contains information for authenticating with Goodreads API
type Client struct {
	baseURI      string
	developerKey string
	client       *http.Client
}

// NewClient creates a new Gooodreads API client
func NewClient(config *Config) (*Client, error) {
	user := oauth.AccessToken{
		Secret: config.AccessSecret,
		Token:  config.AccessToken,
	}
	consumer := oauth.NewConsumer(config.DeveloperKey, config.DeveloperSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   RequestTokenURI,
			AuthorizeTokenUrl: AuthorizeTokenURI,
			AccessTokenUrl:    AccessTokenURI,
		})
	client, err := consumer.MakeHttpClient(&user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create client")
	}
	return &Client{
		baseURI:      APIUri,
		developerKey: config.DeveloperKey,
		client:       client,
	}, nil
}

func (c *Client) doRequest(method, endpoint string, args Arguments) (*Response, error) {
	client := c.client

	params := args.ToURLValues()
	params.Set("key", c.developerKey)

	// Create the request.
	uri := fmt.Sprintf("%s/%s", c.baseURI, strings.Trim(endpoint, "/"))
	uriWithParams := fmt.Sprintf("%s?%s", uri, params.Encode())
	req, err := http.NewRequest(method, uriWithParams, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("creating %s request to %s failed", method, uri))
	}

	// Set the proper headers.
	req.Header.Add("Accept", "application/xml")

	// Do the request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("performing %s request to %s failed", method, uri))
	}
	defer resp.Body.Close()

	// Check that the response status code was OK.
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("invalid developer key")
	case http.StatusForbidden:
		return nil, fmt.Errorf("unauthorized access to endpoint")
	case http.StatusNotFound:
		return nil, fmt.Errorf("the requested uri does not exsit")
	default:
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	// Decode the response into a Goodreads response object.
	var r Response
	if err := decodeResponse(resp, &r); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		body, _ := ioutil.ReadAll(resp.Body)

		return nil, errors.Wrap(err, fmt.Sprintf("decoding response from %s request to %s failed: body -> %s\n", method, uri, string(body)))
	}

	// Return errors on the API errors.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &r, errors.Wrap(err, "API error")
	}

	return &r, nil
}

func decodeResponse(resp *http.Response, v interface{}) error {
	// Copy buffer so we have a backup.
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return err
	}

	return xml.Unmarshal(buf.Bytes(), v)
}
