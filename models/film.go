package models

type FilmPrimarKey struct {
	Id string `json:"film_id"`
}

type CreateFilm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	Duration    int32  `json:"duration"`
}
type Film struct {
	Id          string `json:"film_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	Duration    int32  `json:"duration"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateFilm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	Duration    int32  `json:"duration"`
}

type GetListFilmRequest struct {
	Limit  int32
	Offset int32
}

type GetListFilmResponse struct {
	Count int32   `json:"count"`
	Films []*Film `json:"films"`
}
