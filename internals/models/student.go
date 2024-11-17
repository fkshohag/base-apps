package models

type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Roll       string `json:"roll"`
	Email      string `json:"email"`
	Semester   int    `json:"semester"`
	BatchYear  int    `json:"batch_year"`
}

type ListParams struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	SortBy   string `json:"sort_by,omitempty"`
	SortDesc bool   `json:"sort_desc,omitempty"`
	Search   string `json:"search,omitempty"`
}

type StudentQueryParams struct {
	ListParams
	Grade  *int   `form:"grade,omitempty"`
	Class  string `form:"class,omitempty"`
	Status string `form:"status,omitempty"`
	MinAge *int   `form:"min_age,omitempty"`
	MaxAge *int   `form:"max_age,omitempty"`
}
