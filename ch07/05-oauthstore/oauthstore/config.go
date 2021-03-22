package oauthstore

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// Config wraps the default oauth2.Config and adds our storage
type Config struct {
	*oauth2.Config
	Storage // token storage interface. filestorage
}

// Exchange stores a token after retrieval
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// TokenSource can be passed a token which
// is stored, or when a new one is retrieved,
// that's stored
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, t)
}

// Client is attached to our TokenSource
func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}

// GetToken retrieves a github oauth2 token
func GetToken(ctx context.Context, conf Config) (*oauth2.Token, error) {
	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Print("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
