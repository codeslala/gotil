package clientcredentials

import (
	"context"
	"encoding/json"

	"github.com/codeslala/gotil/oauth2"
	"github.com/codeslala/gotil/postutil"
	"github.com/codeslala/gotil/util/m"
	"github.com/codeslala/gotil/util/must"
)

// Authentication implements PerRPCCredentials interface.
type Authentication struct {
	ClientID     string
	ClientSecret string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	cli := postutil.Client()
	body, err := cli.PostWithUrlencoded(m.M{
		"client_id":     a.ClientID,
		"client_secret": a.ClientSecret,
		"grant_type":    "client_credentials",
	}, oauth2.TokenURL()).ResponseBody()
	if err != nil {
		return nil, err
	}
	resM := make(map[string]interface{})
	must.Must(json.Unmarshal(body, &resM))

	return map[string]string{
		"authorization": authorization(resM["access_token"].(string)),
	}, nil

}

// RequireTransportSecurity indicates whether the credentials requires transport security,
// TODO:
// it's always return false because we do not implement any transport security protocols.
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
