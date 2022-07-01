package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
)

type Storage struct {
	conn *pgxpool.Pool
}

var _ app.Storage = (*Storage)(nil)

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{conn}
}
