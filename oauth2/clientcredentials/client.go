package clientcredentials

import (
	"encoding/json"

	_oauth2 "github.com/codeslala/gotil/oauth2"
	"github.com/codeslala/gotil/postutil"
	"github.com/codeslala/gotil/util/m"
	"github.com/codeslala/gotil/util/must"
	"golang.org/x/oauth2"

	"net/http"
	"time"
)

// This file is derived from "golang.org/x/oauth2/clientcredentials/clientcredentials.go",
// which is a lite version without context control and other optional params.

type Config struct {
	ClientID     string
	ClientSecret string
}

func (c *Config) Client() *http.Client {
	return &http.Client{
		Transport: &Transport{
			Source: oauth2.ReuseTokenSource(nil, c.TokenSource()),
		},
	}
}

func (c *Config) TokenSource() oauth2.TokenSource {
	return &tokenSource{
		conf: c,
	}
}

type tokenSource struct {
	conf *Config
}

func (c *tokenSource) Token() (*oauth2.Token, error) {
	cli := postutil.Client()
	body, err := cli.PostWithUrlencoded(m.M{
		"client_id":     c.conf.ClientID,
		"client_secret": c.conf.ClientSecret,
		"grant_type":    "client_credentials",
	}, _oauth2.TokenURL()).ResponseBody()
	if err != nil {
		return nil, err
	}
	resM := make(map[string]interface{})
	must.Must(json.Unmarshal(body, &resM))
	return &oauth2.Token{
		AccessToken: resM["access_token"].(string),
		TokenType:   resM["token_type"].(string),
		Expiry:      time.Now().Add(time.Duration(resM["expires_in"].(float64)) * time.Second),
	}, nil
}
