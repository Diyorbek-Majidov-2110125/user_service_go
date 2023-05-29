package postgres

import (
	"context"
	"practice1/user_service_go/genproto/user_service"

	// "project2/user_service_go/vendor/github.com/jackc/pgx/v4"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.Pkey, err error) {
	query := `INSERT INTO users (id,first_name, last_name, phone_number) VALUES ($1, $2, $3, $4)`

	id := uuid.New().String()

	_, err = u.db.Exec(ctx, query, id, req.FirstName, req.LastName, req.PhoneNumber)
	if err != nil {
		return resp, err
	}

	resp = &user_service.Pkey{
		Id: id,
	}

	return resp, nil
}

func (u userRepo) GetById(ctx context.Context, req *user_service.Pkey) (resp *user_service.User, err error) {
	query := `SELECT id, first_name, last_name, phone_number FROM users WHERE id = $1`

	resp = &user_service.User{}
	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.PhoneNumber,
	)

	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (u userRepo) GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error) {
	query := `Select id, first_name, last_name, phone_number FROM users`

	rows, err := u.db.Query(ctx, query)

	if err != nil {
		return resp, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user_service.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
		)

		if err != nil {
			return resp, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (u userRepo) Update(ctx context.Context, req *user_service.UpdateUserRequest) (rowsAffected int64, err error) {
	query := `UPDATE users SET first_name = $1, last_name = $2, phone_number = $3 WHERE id = $4`

	result, err := u.db.Exec(ctx, query, req.FirstName, req.LastName, req.PhoneNumber, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (u userRepo) Delete(ctx context.Context, req *user_service.Pkey) (rowsAffected int64, err error) {
	query := `delete from users where id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
