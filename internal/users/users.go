package users

import (
	"catalk/config"
	"catalk/internal/database"
	"context"
	"fmt"
	"log"
	"time"
)

func InsertUser(userEntity *UserEntity) error {
	config := config.GetConfig().Database
	dbClient := database.New(config)

	ctx := context.Background()

	tx, err := dbClient.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("error begin transaction. Error: %s", err.Error())
		return fmt.Errorf("insert user failed")
	}
	defer tx.Commit()
	query := `INSERT INTO "users" (email, username, picture_url, provider_id) VALUES ($1,$2,$3,$4) RETURNING "id";`
	queryCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := tx.QueryRowContext(queryCtx, query,
		userEntity.Email,
		userEntity.Username,
		userEntity.PictureURL,
		userEntity.ProviderID,
	).Scan(&userEntity.ID); err != nil {
		tx.Rollback()
		log.Printf("error insert user failed. Error: %s", err.Error())
		return fmt.Errorf("insert user failed")
	}

	return nil
}
