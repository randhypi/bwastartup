package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return User{}, err
	}

	user := User{
		Name:           input.Name,
		Occupation:     input.Occupation,
		Email:          input.Email,
		PasswordHash:   string(passwordHash),
		AvatarFileName: "",
		Role:           "user",
		Token:          "",
	}

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
