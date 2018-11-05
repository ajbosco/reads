package goodreads

import (
	"fmt"
	"net/http"
	"strconv"
)

// ListShelves for current user.
func (c *Client) ListShelves(userID string) ([]UserShelf, error) {
	args := make(map[string]string)
	args["user_id"] = userID
	data, err := c.doRequest(http.MethodGet, ShelfListEndpoint, args)
	if err != nil {
		return nil, err
	}

	return data.Shelves.UserShelves, nil
}

// ListShelfBooks returns books by from a shelf.
func (c *Client) ListShelfBooks(shelf string, userID string) ([]Book, error) {
	var books []Book
	args := make(map[string]string)
	args["shelf"] = shelf
	args["v"] = "2"
	args["per_page"] = "200"

	for page := 1; ; page++ {
		args["page"] = strconv.Itoa(page)
		listShelfEndpoint := fmt.Sprintf("%s/%s.xml", ListShelfEndpoint, userID)
		data, err := c.doRequest(http.MethodGet, listShelfEndpoint, args)
		if err != nil {
			return nil, err
		}
		books = append(books, data.Reviews.Books...)
		if data.Reviews.AttrEnd == data.Reviews.AttrTotal {
			break
		}
	}

	return books, nil
}

// AddToShelf adds a book to shelf.
func (c *Client) AddToShelf(shelf string, bookID string) error {
	args := make(map[string]string)
	args["name"] = shelf
	args["book_id"] = bookID
	_, err := c.doRequest(http.MethodPost, AddToShelfEndpoint, args)
	if err != nil {
		return err
	}

	return nil
}
