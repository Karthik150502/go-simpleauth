package types

type SignUpResponse struct {
	UserId string `json:"userId"`
}

type SignInResponse struct {
	UserId string `json:"userId"`
}

type UserSignInResponse struct {
	StatusCode   int    `json:"statusCode"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type MessageResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type ApiResponse[T SignUpResponse | SignInResponse] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Payload    *T     `json:"payload,omitempty"`
}
