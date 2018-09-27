package goodreads

import (
	"net/http"
)

// GetCurrentUserID returns id for oauth user
func (c *Client) GetCurrentUserID() (*User, error) {
	data, err := c.doRequest(http.MethodGet, CurrentUserEndpoint, nil)
	if err != nil {
		return nil, err
	}

	return &data.User, nil
}
