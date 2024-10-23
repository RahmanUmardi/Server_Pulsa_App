package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Register(payload dto.AuthRequestDto) (entity.User, error)
}

type authUseCase struct {
	useCase    UserUsecase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.useCase.FindUserByUsernamePassword(payload.Username, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	response := dto.AuthResponseDto{
		Token: token.Token,
	}
	return response, nil
}

func (a *authUseCase) Register(payload dto.AuthRequestDto) (entity.User, error) {
	return a.useCase.RegisterUser(entity.User{Username: payload.Username, Password: payload.Password})
}

func NewAuthUseCase(uc UserUsecase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{useCase: uc, jwtService: jwtService}
}
