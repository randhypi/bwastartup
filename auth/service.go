package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() Service {
	return &jwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3t_k3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil

}
