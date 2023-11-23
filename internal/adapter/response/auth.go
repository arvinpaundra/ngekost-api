package response

type (
	Login struct {
		UserId       string  `json:"user_id"`
		Role         string  `json:"role"`
		IssuedAt     string  `json:"issued_at"`
		ExpiredAt    string  `json:"expired_at"`
		AccessToken  string  `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}
)
