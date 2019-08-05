package app

import (
	"context"
	"github.com/beefsack/go-rate"
	"github.com/jzelinskie/geddit"
	"github.com/ropenttd/tsubasa/generics/pkg/environment"
	"golang.org/x/oauth2"
	"net/http"
)

type transport struct {
	http.RoundTripper
	useragent string
}

// Any request headers can be modified here.
func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", t.useragent)
	return t.RoundTripper.RoundTrip(req)
}

type ExtendedOAuthSession struct {
	geddit.OAuthSession
	ctx      context.Context
	throttle *rate.RateLimiter
}

func (o *ExtendedOAuthSession) GetCtx() *context.Context {
	return &o.ctx
}

func (o *ExtendedOAuthSession) CodeAuthWithToken(code string) (token *oauth2.Token, err error) {
	t, err := o.OAuthConfig.Exchange(o.ctx, code)
	if err != nil {
		return nil, err
	}
	o.Client = o.OAuthConfig.Client(o.ctx, t)
	return t, nil
}

func NewRedditOAuthSession() (session *ExtendedOAuthSession, err error) {
	o := &ExtendedOAuthSession{}

	// Set OAuth config
	o.OAuthConfig = &oauth2.Config{
		ClientID:     environment.GetEnv("reddit_oauth_id", "TEST"),
		ClientSecret: environment.GetEnv("reddit_oauth_secret", "TEST"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.reddit.com/api/v1/authorize",
			TokenURL: "https://www.reddit.com/api/v1/access_token",
		},
		RedirectURL: "http://localhost:8000/api/user/reddit/auth/callback",
	}
	// Inject our custom HTTP client so that a user-defined UA can
	// be passed during any authentication requests.
	c := &http.Client{}
	c.Transport = &transport{http.DefaultTransport, "tsubasa beta/1-pre"}
	o.ctx = context.WithValue(context.Background(), oauth2.HTTPClient, c)
	return o, nil
}

func GetRedditOAuthURL() (url *string, err error) {
	// Redirect user to consent page to ask for permission
	// for the scopes specified.
	s, err := NewRedditOAuthSession()
	if err != nil {
		return nil, err
	}
	url_s := s.AuthCodeURL("state", []string{"identity", "mysubreddits"})
	return &url_s, nil
}
