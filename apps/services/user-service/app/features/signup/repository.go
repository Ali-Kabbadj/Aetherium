package signup

import (
	"context"

	db "aetherium.com/user-service/app/externals/database/sqlc"
)

type DBTX interface {
	db.DBTX
}

type SignUpRepository struct {
	q *db.Queries
}

func NewSignUpRepository(dbtx DBTX) *SignUpRepository {
	return &SignUpRepository{q: db.New(dbtx)}
}

func (r *SignUpRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *SignUpRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}