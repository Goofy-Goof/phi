package auth

import (
	ory "github.com/ory/kratos-client-go"
)

func (c *Client) Login(username string, password string) (Session, error) {
	flow, _, err := c.OryClient.V0alpha2Api.InitializeSelfServiceLoginFlowWithoutBrowser(c.ctx).Execute()
	if err != nil {
		return Session{}, err
	}

	result, _, err := c.OryClient.V0alpha2Api.SubmitSelfServiceLoginFlow(c.ctx).Flow(flow.Id).SubmitSelfServiceLoginFlowBody(
		ory.SubmitSelfServiceLoginFlowWithPasswordMethodBodyAsSubmitSelfServiceLoginFlowBody(&ory.SubmitSelfServiceLoginFlowWithPasswordMethodBody{
			Method:             "password",
			Password:           password,
			PasswordIdentifier: username,
		}),
	).Execute()
	if err != nil {
		return Session{}, err
	}

	return Session{
		Session: result.Session,
		Token:   *result.SessionToken,
	}, nil
}
