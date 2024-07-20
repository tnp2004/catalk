package users

import (
	"catalk/config"
	"catalk/internal/database"
	"context"
	"fmt"
	"log"
	"time"
)

type user struct {
	dbConfig *config.Database
}
type UserService interface {
	InsertUser(req *NewUserModel) error
}

func NewUser(config *config.Database) UserService {
	return &user{dbConfig: config}
}

func (u *user) InsertUser(req *NewUserModel) error {
	config := config.GetConfig().Database
	db := database.New(config)
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
