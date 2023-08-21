package api

type OauthSpec struct {
	Code    string `json:"access_code"`
	Account int    `json:"account_type"`
}
