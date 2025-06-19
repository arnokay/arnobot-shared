package apptype

import "github.com/arnokay/arnobot-shared/data"

type (
	AuthSessionTokenRequest          = Request[string]
	AuthSessionTokenValidateResponse = Response[bool]
	AuthSessionTokenExchangeResponse = Response[*data.User]

	AuthProviderGetRequest  = Request[data.AuthProviderGet]
	AuthProviderGetResponse = Response[*data.AuthProvider]

	AuthProviderUpdateTokensRequest  = Request[AuthProviderUpdateTokensPayload]
	AuthProviderUpdateTokensResponse = Response[bool]
)

type AuthProviderUpdateTokensPayload struct {
	ID   int32
	Data data.AuthProviderUpdateTokens
}
