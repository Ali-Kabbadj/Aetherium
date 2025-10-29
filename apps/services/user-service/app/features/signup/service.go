package signup

import (
	"context"
	"fmt"

	db "aetherium.com/user-service/app/externals/database/sqlc"
	"github.com/jackc/pgx/v5" // Required for pgx.ErrNoRows
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
}

type SignUpService struct {
	repo Repository
}

func NewSignUpService(repo Repository) *SignUpService {
	return &SignUpService{repo: repo}
}

func (s *SignUpService) CreateUser(ctx context.Context, email, password string) (db.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, fmt.Errorf("could not hash password: %w", err)
	}

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "creator",
	})
	if err != nil {
		return db.User{}, fmt.Errorf("repository could not create user: %w", err)
	}

	return user, nil
}

func (s *SignUpService) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	_, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("repository could not check email: %w", err)
	}

	return false, nil
}