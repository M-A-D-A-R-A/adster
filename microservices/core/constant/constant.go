package constant

import "os"

var TableFiles string = "ad_data"

type RequestURL struct {
	ThirdPartyBaseURL string
	Endpoint string
}

type Token struct {
	Token string
}

func ForcastURL() string {

	requestParams := RequestURL{
		ThirdPartyBaseURL: os.Getenv("THIRD_PARTY_BASE_URL"),
		Endpoint: "/forecast",
	}
	requestURL := requestParams.ThirdPartyBaseURL+requestParams.Endpoint

	return requestURL
}

func GetToken() string {
	jwtToken := Token{
		Token: os.Getenv("BEARER_SYSTEM_JWT"),
	}
	return jwtToken.Token
}

