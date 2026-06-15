package repository

import (
	"context"

	db "go-user-api/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{Queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob pgtype.Date) (db.User, error) {
	return r.Queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetUser(ctx context.Context, id int32) (db.User, error) {
	return r.Queries.GetUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.Queries.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name string, dob pgtype.Date) (db.User, error) {
	return r.Queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.Queries.DeleteUser(ctx, id)
}

func (r *UserRepository) ListUsersPaginated(ctx context.Context, arg db.ListUsersPaginatedParams) ([]db.User, error) {
	return r.Queries.ListUsersPaginated(ctx, arg)
}
