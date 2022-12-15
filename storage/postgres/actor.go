package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
	"crud/pkg/helper"
)

type actorRepo struct {
	db *pgxpool.Pool
}

func NewActorRepo(db *pgxpool.Pool) *actorRepo {
	return &actorRepo{
		db: db,
	}
}

func (f *actorRepo) Create(ctx context.Context, actor *models.CreateActor) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO actor(
			actor_id,
			first_name,
			last_name,
			updated_at
		) VALUES ( $1, $2, $3, now() )
	`

	_, err := f.db.Exec(ctx, query,
		id,
		actor.First_name,
		actor.Last_name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *actorRepo) GetByPKey(ctx context.Context, pkey *models.ActorPrimarKey) (*models.Actor, error) {

	var (
		id         sql.NullString
		first_name sql.NullString
		last_name  sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		SELECT
			actor_id,
			first_name,
			last_name,
			created_at,
			updated_at
		FROM
			actor
		WHERE actor_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&first_name,
			&last_name,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Actor{
		Id:         id.String,
		First_name: first_name.String,
		Last_name:  last_name.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (f *actorRepo) GetList(ctx context.Context, req *models.GetListActorRequest) (*models.GetListActorResponse, error) {

	var (
		resp   = models.GetListActorResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			actor_id,
			first_name,
			last_name,
			created_at,
			updated_at
		FROM
			actor
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id         sql.NullString
			first_name sql.NullString
			last_name  sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&last_name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Actors = append(resp.Actors, &models.Actor{
			Id:         id.String,
			First_name: first_name.String,
			Last_name:  last_name.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})

	}

	return &resp, err
}

func (f *actorRepo) Update(ctx context.Context, id string, req *models.UpdateActor) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			actor
		SET
			first_name = :first_name,
			last_name = :last_name,
			updated_at = now()
		WHERE actor_id = :actor_id
	`

	params = map[string]interface{}{
		"actor_id":   id,
		"first_name": req.First_name,
		"last_name":  req.Last_name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *actorRepo) Delete(ctx context.Context, req *models.ActorPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM actor WHERE actor_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
