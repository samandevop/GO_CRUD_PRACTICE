package models

type CategoryPrimarKey struct {
	Id string `json:"category_id"`
}

type CreateCategory struct {
	Name string `json:"name"`
}
type Category struct {
	Id        string `json:"category_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCategory struct {
	Name string `json:"name"`
}

type GetListCategoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoryResponse struct {
	Count     int32       `json:"count"`
	Categorys []*Category `json:"categorys"`
}
