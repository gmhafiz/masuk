package impl

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/nbutton23/zxcvbn-go"

	"masuk/internal/grpc/service"
)

func (r RepositoryServiceGrpcImpl) CreateUser(ctx context.Context, request *service.CreateUserRequest) (*service.CreateUserResponse, error) {
	if request.User.Password != request.User.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	m := zxcvbn.PasswordStrength(request.User.Password, nil)
	if m.Score <= 3 {
		return nil, errors.New("weak password")
	}

	hash, err := argon2id.CreateHash(request.User.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err)
	}

	cmdtag, err := r.DbConn.DB.Exec("INSERT INTO users(username, email, password) VALUES ($1, $2, $3)", request.User.Username, request.User.Email, hash)
	if err != nil {
		return nil, err
	}
	log.Println(cmdtag)

	token, err := r.createToken(request.User.Email)
	if err != nil {
		return nil, err
	}

	res := &service.CreateUserResponse{Token: token}

	return res, nil
}

func (r RepositoryServiceGrpcImpl) GetUser(ctx context.Context, request *service.GetUserRequest) (*service.GetUserResponse, error) {
	panic("implement me")
}

func (r RepositoryServiceGrpcImpl) UpdateUser(ctx context.Context, request *service.GetUserRequest) (*service.GetUserResponse, error) {
	panic("implement me")
}

func (r RepositoryServiceGrpcImpl) DeleteUser(ctx context.Context, request *service.DeleteUserRequest) (*service.DeleteUserResponse, error) {
	panic("implement me")
}

func (r RepositoryServiceGrpcImpl) Login(ctx context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
	type UserDB struct {
		Id        int64          `db:"id"`
		Username  string         `db:"username"`
		Email     string         `db:"email"`
		Password  string         `db:"password"`
		FirstName sql.NullString `db:"first_name"`
		LastName  sql.NullString `db:"last_name"`
		Mobile    sql.NullString `db:"mobile"`
		Phone     sql.NullString `db:"phone"`
		Active    bool           `db:"active"`
		CreatedAt *time.Time     `db:"created_at"`
		UpdatedAt *time.Time     `db:"updated_at"`
		DeletedAt *time.Time     `db:"deleted_at"`
	}
	var userDB UserDB

	rows := r.DbConn.DB.QueryRow(`SELECT * FROM users where email = $1 LIMIT 1`, request.Email)
	err := rows.Scan(&userDB.Id, &userDB.Username, &userDB.Email, &userDB.Password,
		&userDB.FirstName, &userDB.LastName, &userDB.Mobile, &userDB.Phone,
		&userDB.Active, &userDB.CreatedAt, &userDB.UpdatedAt, &userDB.DeletedAt)

	_, err = argon2id.ComparePasswordAndHash(request.Password, userDB.Password)
	if err != nil {
		return &service.LoginResponse{Token: ""}, err
	}

	token, err := r.createToken(request.Email)
	if err != nil {
		return nil, err
	}

	return &service.LoginResponse{Token: token}, nil
}

func (r RepositoryServiceGrpcImpl) Logout(ctx context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
	panic("implement me")
}

func (r RepositoryServiceGrpcImpl) createToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
	}

	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*60*8))

	_, tokenString, err := r.Jwt.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
