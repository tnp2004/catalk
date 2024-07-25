package users

import (
	"catalk/config"
	"catalk/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type user struct {
	dbConfig *config.Database
}
type UserService interface {
	InsertUser(req *NewUserModel) error
	FindUserByEmail(email string) (*UserEntity, error)
}

func NewUser(config *config.Database) UserService {
	return &user{dbConfig: config}
}

func (u *user) InsertUser(req *NewUserModel) error {
	db := database.New(u.dbConfig)
	conn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(context.Background(), nil)
	if err != nil {
		log.Printf("error begin transaction. Error: %s", err.Error())
		return fmt.Errorf("insert user failed")
	}
	defer tx.Commit()
	query := `INSERT INTO "users" (email, username, picture_url, provider_id)
	VALUES ($1,$2,$3,$4) RETURNING "id";`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := tx.QueryRowContext(ctx, query,
		req.Email,
		req.Username,
		req.PictureURL,
		req.ProviderID,
	).Err(); err != nil {
		tx.Rollback()
		log.Printf("error insert user failed. Error: %s", err.Error())
		return fmt.Errorf("insert user failed")
	}

	return nil
}

func (u *user) FindUserByEmail(email string) (*UserEntity, error) {
	db := database.New(u.dbConfig)
	conn, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `SELECT * FROM users WHERE email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userData := new(UserEntity)
	if err := conn.QueryRowContext(ctx, query,
		email,
	).Scan(&userData.ID, &userData.Email, &userData.Username, &userData.PictureURL,
		&userData.ProviderID, &userData.CreatedAt, &userData.UpdatedAt); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("user email %s not found", email)
		}
		log.Printf("error insert user failed. Error: %s", err.Error())
		return nil, fmt.Errorf("insert user failed")
	}

	return userData, nil
}
