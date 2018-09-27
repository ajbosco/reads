package goodreads

import "errors"

const (
	// APIUri holds the Goodreads API uri.
	APIUri = "https://www.goodreads.com"
	// SearchEndpoint is the API endpoint for searching for books.
	SearchEndpoint = "/search/index.xml"
	// ShelfListEndpoint is the API endpoint for listing shelves
	ShelfListEndpoint = "/shelf/list.xml"
	// CurrentUserEndpoint is the API endpoint to get the current user
	CurrentUserEndpoint = "/api/auth_user"
	// ListShelfEndpoint is the API endpoint to list books on a shelf
	ListShelfEndpoint = "/review/list"
	// AddToShelfEndpoint is the API endpoint to add a book to a shelf
	AddToShelfEndpoint = "/shelf/add_to_shelf.xml"
	// RequestTokenURI holds the Goodreads request token uri
	RequestTokenURI = "https://www.goodreads.com/oauth/request_token"
	// AuthorizeTokenURI holds the Goodreads authorize token uri
	AuthorizeTokenURI = "https://www.goodreads.com/oauth/authorize"
	// AccessTokenURI holds the Goodreads access token uri
	AccessTokenURI = "https://www.goodreads.com/oauth/access_token"
)

var (
	// ErrorEmptyResult defines the error when the result is empty.
	ErrorEmptyResult = errors.New("Empty result")
)
