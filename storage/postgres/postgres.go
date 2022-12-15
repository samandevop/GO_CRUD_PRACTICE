package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"crud/config"
	"crud/storage"
)

type Store struct {
	db       *pgxpool.Pool
	film     *filmRepo
	actor    *actorRepo
	category *categoryRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pool,
		film:     NewFilmRepo(pool),
		actor:    NewActorRepo(pool),
		category: NewCategoryRepo(pool),
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Film() storage.FilmRepoI {

	if s.film == nil {
		s.film = NewFilmRepo(s.db)
	}

	return s.film
}

func (s *Store) Actor() storage.ActorRepoI {

	if s.actor == nil {
		s.actor = NewActorRepo(s.db)
	}

	return s.actor
}

func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}
