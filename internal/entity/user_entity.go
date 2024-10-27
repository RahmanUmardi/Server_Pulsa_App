package entity

type (
	User struct {
		Id_user  string `json:"id_user"`
		Username string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UserErrorResponse struct {
		Error string `json:"error" example:"Invalid product"`
	}
)
