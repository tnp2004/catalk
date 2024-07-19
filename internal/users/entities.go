package users

import "time"

type UserEntity struct {
	ID         string    `db:"id"`
	Email      string    `db:"email"`
	Username   string    `db:"username"`
	PictureURL string    `db:"picture_url"`
	ProviderID uint      `db:"provider_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
