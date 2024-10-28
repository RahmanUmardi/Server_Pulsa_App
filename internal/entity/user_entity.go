package entity

type (
	User struct {
		Id_user  string `json:"id_user"`
		Username string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UserReqUpdate struct {
		Username string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UserResponse struct {
		Id_user  string `json:"id_user"`
		Username string `json:"name"`
		Password string `json:"password,omitempty"`
		Role     string `json:"role"`
	}
	UserErrorResponse struct {
		Error string `json:"error" example:"Invalid product"`
	}
)
