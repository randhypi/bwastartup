package user

import "time"

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	Email          string `json:"email"`
	PasswordHash   string
	AvatarFileName string `json:"avatar_filename"`
	Role           string
	Token          string `json:"token"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
