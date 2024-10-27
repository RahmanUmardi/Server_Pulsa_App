package dto

type AuthRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}

type (
	AuthRequest struct {
		Username string `json:"username" binding:"required" example:"john_doe"`
		Password string `json:"password" binding:"required" example:"secret123"`
	}

	AuthResponse struct {
		Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
	}

	AuthRegisterRes struct {
		Password string `json:"password" example:"Hashed Password"`
	}

	ErrorResponse struct {
		Error string `json:"error" example:"Invalid credentials"`
	}
)
