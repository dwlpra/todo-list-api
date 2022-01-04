package activity

type Activities struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	Title     string      `json:"title"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type InsertActivity struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type RequestBody struct {
	Email string `json:"email"`
	Title string `json:"title"`
}
