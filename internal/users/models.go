package users

type NewUserModel struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	PictureURL string `json:"picture_url"`
	ProviderID uint   `json:"provider_id"`
}
