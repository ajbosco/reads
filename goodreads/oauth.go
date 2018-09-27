package goodreads

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mrjones/oauth"
)

// GetAccessToken returns an oauth access token
func GetAccessToken(developerKey string, developerSecret string) (*oauth.AccessToken, error) {
	c := oauth.NewConsumer(developerKey, developerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   RequestTokenURI,
			AuthorizeTokenUrl: AuthorizeTokenURI,
			AccessTokenUrl:    AccessTokenURI,
		})

	requestToken, url, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		return nil, err
	}
	fmt.Println("You need to grant Goodreads CLI access")
	fmt.Println("Go to and grant access: " + url)
	fmt.Println("Press enter after authorizing")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	accessToken, err := c.AuthorizeToken(requestToken, "")
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
