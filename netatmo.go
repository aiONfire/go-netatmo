// package netatmo
package netatmo

import (
	"net/http"

	"golang.org/x/oauth2"
)

const (

	// BaseURL is netatmo api url
	BaseURL = "https://api.netatmo.com/"

	// AuthURL is netatmo auth url
	AuthURL = BaseURL + "oauth2/authorize"

	// TokenURL is netatmo token url
	TokenURL = BaseURL + "oauth2/token"

	// deviceURL is netatmo device url
	DeviceURL = BaseURL + "/api/getstationsdata"
)

// https://github.com/exzz/netatmo-api-go/blob/master/weather.go
// Config is used to specify credential to Netatmo API
// ClientID : Client ID from netatmo app registration at http://dev.netatmo.com/dev/listapps
// ClientSecret : Client app secret
// Username : Your netatmo account username
// Password : Your netatmo account password
type Config struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// GrandType is used to specify the grant type to use for
// OAuth 2 authorization process. Depending on the selected
// type, different properties have to be provided.
// This standard method is the "Authorization code" grant
// type. This type is for personal use only (the account
// with which the API application have been created).
// As second method the "Client credentials" grant type can
// be used in which the username and the password is sent
// along with the request.
// In both cases, you will retrieve one access token and one
// refresh token. The access token represent a key for each
// user to access the data and it should be sent along with
// the request you will make to Netatmo backend.
//
// To know how you should send the access token, check out
// the "Using a token" part. As all access token expires
// after a certain duration, you need to refresh it using
// the refresh token. This process is described in the
// "Refreshing a token" part of the documentation.
/*
type GrantType string

const (
	AuthorizationCode GrantType = "AuthorizationCode"
	ClientCredentials           = "ClientCredentials"
)*/

// AuthorizationCode represents the data which is necessary
// for using the "Authorization code" grant type.
// This standard method. This method is for personal use only
// (the account with which the API application have been
// created). As alternative "Client credentials" grant type
// can be used.
type AuthorizationCode struct {
	GrantType   string   `json:"grant_type"`
	ClientID    string   `json:"clint_id"`
	RedirectURI string   `json:"redirect_uri"`
	Scope       []string `json:"scope"`
	State       string   `json:"state"`
}

func NewAuthorizationCode(clientID string, redirectURI string, scope []string) (authorisationCode *AuthorizationCode, err error) {
	authorisationCode = new(AuthorizationCode)
	authorisationCode.GrantType = "AuthorizationCode"
	authorisationCode.ClientID = clientID
	authorisationCode.RedirectURI = redirectURI
	authorisationCode.Scope = scope

	// do some testing on given values and report
	// an error if something is wrong or missing
	return authorisationCode, nil
}

type ClientCredentials struct {
	GrantType string `json:"grant_type"`

	ClientID     string `json:"clint_id"`
	ClientSecret string `json:"client_secret"`

	Scope []string `json:"scope"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func NewClientCredentials(clientID string, clientSecret string, scope []string, username string, password string) (clientCredentials *ClientCredentials, err error) {
	clientCredentials = new(ClientCredentials)
	clientCredentials.GrantType = "ClientCredentials"
	clientCredentials.ClientID = clientID
	clientCredentials.ClientSecret = clientSecret
	clientCredentials.Scope = scope
	clientCredentials.Username = username
	clientCredentials.Password = password
	return clientCredentials, nil
}

type RefreshToken struct {
}

type Token struct {
	AccessToken  string
	ExpiresIn    string
	RefreshToken RefreshToken
}

// Client is a nuse to make request to Netatmo API
type Client struct {
	oauth        *oauth2.Config
	httpClient   *http.Client
	httpResponse *http.Response
	Dc           *DeviceCollection
}

func NewClient(grantType interface{}) (client *Client, err error) {

	//gT := reflect.TypeOf(grantType)

	switch grantType.(type) {
	case AuthorizationCode:

		config := new(oauth2.Config)
		config.ClientID = grantType.(AuthorizationCode).ClientID
		config.Scopes = grantType.(AuthorizationCode).Scope
		config.RedirectURL = grantType.(AuthorizationCode).RedirectURI
		config.Endpoint = oauth2.Endpoint{
			AuthURL:  AuthURL,  //"https://api.netatmo.com/oauth2/authorize", // netatmo.BaseURL, // "http://localhost:9096/authorize",
			TokenURL: TokenURL, //"https://api.netatmo.com/oauth2/token",     //netatmo.AuthURL, // "http://localhost:9096/token",
		}

		return nil, nil
	case ClientCredentials:
		return nil, nil
	default:
		return nil, nil
	}
}

func Echo(echo string) string {
	return echo
}
