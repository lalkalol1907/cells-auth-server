package DB

import (
	"cells-auth-server/src/Config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var DB *pgxpool.Pool

func InitDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, Config.Cfg.DB.Url)

	if err != nil {
		return err
	}

	DB = db

	return nil
}

func CloseDatabase() {
	DB.Close()
}
