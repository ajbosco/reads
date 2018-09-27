package goodreads

import (
	"net/http"
)

// SearchBooks returns books by author, title, id.
func (c *Client) SearchBooks(query string) ([]Work, error) {
	args := make(map[string]string)
	args["q"] = query
	data, err := c.doRequest(http.MethodGet, SearchEndpoint, args)
	if err != nil {
		return nil, err
	}

	return data.Search.Works, nil
}
