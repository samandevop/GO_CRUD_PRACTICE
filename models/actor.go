package models

type ActorPrimarKey struct {
	Id string `json:"actor_id"`
}

type CreateActor struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
type Actor struct {
	Id         string `json:"actor_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateActor struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type GetListActorRequest struct {
	Limit  int32
	Offset int32
}

type GetListActorResponse struct {
	Count  int32    `json:"count"`
	Actors []*Actor `json:"actors"`
}
