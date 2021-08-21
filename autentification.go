package tool_library

type AuthConfiguration struct {
	AccessTokenLifeTime  int64 `json:"accessTokenLifeTime"`
	RefreshTokenLifeTime int64 `json:"refreshTokenLifeTime"`
}
