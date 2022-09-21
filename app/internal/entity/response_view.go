package entity

type ResponseView struct {
	IsSuccess bool        `json:"is_success"`
	Data      interface{} `json:"data"`
}

type TokensResponseView struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
