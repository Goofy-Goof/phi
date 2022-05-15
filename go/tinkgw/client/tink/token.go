package tink

import (
	"encoding/json"
	"fmt"
	"github.com/clstb/phi/go/tinkgw/config"
	"net/http"
	"net/url"
	"os"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func (c *Client) GetToken(
	code,
	refreshToken,
	clientId,
	clientSecret,
	grantType,
	scope string,
) (token Token, err error) {
	form := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"grant_type":    {grantType},
		"scope":         {scope},
	}
	switch grantType {
	case "refresh_token":
		form.Add("refresh_token", refreshToken)
	case "authorization_code":
		form.Add("code", code)
	}

	res, err := c.httpClient.httpClient.PostForm(c.url+"/api/v1/oauth/token", form)
	if err != nil {
		return token, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(res.Body).Decode(&token)
	default:
		err = fmt.Errorf("unimplemented status: %d", res.StatusCode)
	}

	return
}

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IdHint       string `json:"id_hint"`
}

func (c *Client) ExchangeAccessCodeForToken(code string) (string, error) {
	data := url.Values{
		"code":          {code},
		"client_id":     {os.Getenv("TINK_CLIENT_ID")},
		"client_secret": {os.Getenv("TINK_CLIENT_SECRET")},
		"grant_type":    {config.AuthorizationCodeGrantType},
	}

	res, err := c.httpClient.PostForm(config.TinkTokenUri, data)
	if err != nil {
		return "", err
	}

	var respBody TokenResponse
	err = json.Unmarshal(res, &respBody)
	return respBody.AccessToken, nil
}
