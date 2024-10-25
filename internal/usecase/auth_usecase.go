package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Register(payload dto.AuthRequestDto) (entity.User, error)
}

type authUseCase struct {
	useCase    UserUsecase
	jwtService service.JwtService
	log        *logger.Logger
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	a.log.Info("Starting to authenticate user in the use case layer", nil)

	user, err := a.useCase.FindUserByUsernamePassword(payload.Username, payload.Password)
	if err != nil {
		a.log.Error("Failed to authenticate user: ", err)
		return dto.AuthResponseDto{}, err
	}

	a.log.Info("User has been authenticated successfully", nil)
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		a.log.Error("Failed to create token: ", err)
		return dto.AuthResponseDto{}, err
	}

	response := dto.AuthResponseDto{
		Token: token.Token,
	}

	a.log.Info("User ID %s has been authenticated successfully", user.Id_user)
	return response, nil
}

func (a *authUseCase) Register(payload dto.AuthRequestDto) (entity.User, error) {
	a.log.Info("Starting to register a new user in the use case layer", nil)
	return a.useCase.RegisterUser(entity.User{Username: payload.Username, Password: payload.Password})
}

func NewAuthUseCase(uc UserUsecase, jwtService service.JwtService, log *logger.Logger) AuthUseCase {
	return &authUseCase{useCase: uc, jwtService: jwtService, log: log}
}
