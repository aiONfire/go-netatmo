//
package go-netatmo

import (
	"strings"
)

const (

	// baseURL is netatmo api url
	baseURL = "https://api.netatmo.net/"

	// authURL is netatmo auth url
	authURL = baseURL + "oauth2/token"

	// deviceURL is netatmo device url
	deviceURL = baseURL + "/api/getstationsdata"
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
type GrantType string

const(
    AuthorizationCode GrantType = "AuthorizationCode"
    ClientCredentials = "ClientCredentials"
)

// AuthorizationCode represents the data which is necessary
// for using the "Authorization code" grant type. 
// This standard method. This method is for personal use only 
// (the account with which the API application have been 
// created). As alternative "Client credentials" grant type 
// can be used.
type AuthorizationCode {

	ClientID     string	`json:"clint_id"`
	RedirectURI	string	`json:"redirect_uri"`
	Scope string	`json:"scope"`
	State string	`json:"state"`


}

type ClientCredentials {

	GrantType GrandType `json:"grant_type"`

	ClientID     string	`json:"clint_id"`
	ClientSecret string	`json:"client_secret"`
	
	Username     string	`json:"username"`
	Password     string	`json:"password"`
	Scope []string 	`json:"scope"`
}

type RefreshToken {

}

type Token {
	AccessToken
	ExpiresIn
	RefreshToken
}