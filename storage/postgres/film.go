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

type filmRepo struct {
	db *pgxpool.Pool
}

func NewFilmRepo(db *pgxpool.Pool) *filmRepo {
	return &filmRepo{
		db: db,
	}
}

func (f *filmRepo) Create(ctx context.Context, film *models.CreateFilm) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO film(
			film_id,
			title,
			description,
			release_year,
			duration,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, now() )
	`

	_, err := f.db.Exec(ctx, query,
		id,
		film.Title,
		film.Description,
		film.ReleaseYear,
		film.Duration,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *filmRepo) GetByPKey(ctx context.Context, pkey *models.FilmPrimarKey) (*models.Film, error) {

	var (
		id          sql.NullString
		title       sql.NullString
		description sql.NullString
		releaseYear sql.NullString
		duration    sql.NullInt32
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			film_id,
			title,
			description,
			TO_CHAR(release_year, 'YYYY-MM-DD'),
			duration,
			created_at,
			updated_at
		FROM
			film
		WHERE film_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&title,
			&description,
			&releaseYear,
			&duration,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Film{
		Id:          id.String,
		Title:       title.String,
		Description: description.String,
		ReleaseYear: releaseYear.String,
		Duration:    duration.Int32,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *filmRepo) GetList(ctx context.Context, req *models.GetListFilmRequest) (*models.GetListFilmResponse, error) {

	var (
		resp   = models.GetListFilmResponse{}
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
			film_id,
			title,
			description,
			TO_CHAR(release_year, 'YYYY-MM-DD'),
			duration,
			created_at,
			updated_at
		FROM
			film
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			title       sql.NullString
			description sql.NullString
			releaseYear sql.NullString
			duration    sql.NullInt32
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&description,
			&releaseYear,
			&duration,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Films = append(resp.Films, &models.Film{
			Id:          id.String,
			Title:       title.String,
			Description: description.String,
			ReleaseYear: releaseYear.String,
			Duration:    duration.Int32,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})

	}

	return &resp, err
}

func (f *filmRepo) Update(ctx context.Context, id string, req *models.UpdateFilm) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			film
		SET
			title = :title,
			description = :description,
			release_year = :release_year,
			duration = :duration,
			updated_at = now()
		WHERE film_id = :film_id
	`

	params = map[string]interface{}{
		"actor_id":     id,
		"title":        req.Title,
		"description":  req.Description,
		"release_year": req.ReleaseYear,
		"duration":     req.Duration,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *filmRepo) Delete(ctx context.Context, req *models.FilmPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM film WHERE film_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
