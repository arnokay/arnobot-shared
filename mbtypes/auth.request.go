package mbtypes

import "arnobot-shared/data"

type (
	AuthSessionTokenRequest          = Request[string]
	AuthSessionTokenValidateResponse = Response[bool]
	AuthSessionTokenExchangeResponse = Response[*data.User]
	AuthProviderGetRequest           = Request[data.AuthProviderGet]
	AuthProviderGetResponse          = Response[*data.AuthProvider]
)
