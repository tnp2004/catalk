package users

type NewUserModel struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	PictureURL string `json:"picture_url"`
	ProviderID uint8  `json:"provider_id"`
}

type NewGoogleUserModel struct {
	Email      string `json:"email"`
	Username   string `json:"given_name"`
	PictureURL string `json:"picture"`
	ProviderID uint8  `json:"provider_id"`
}

type Providers struct {
	Google  uint8
	Discord uint8
}

var Provider = &Providers{
	Google:  1,
	Discord: 2,
}
