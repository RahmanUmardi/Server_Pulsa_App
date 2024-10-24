package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/entity/dto"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/shared/service"
)

var logAuth = logger.GetLogger()

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Register(payload dto.AuthRequestDto) (entity.User, error)
}

type authUseCase struct {
	useCase    UserUsecase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	logAuth.Info("Starting to authenticate user in the use case layer")

	user, err := a.useCase.FindUserByUsernamePassword(payload.Username, payload.Password)
	if err != nil {
		logAuth.Error("Failed to authenticate user: ", err)
		return dto.AuthResponseDto{}, err
	}

	logAuth.Info("User has been authenticated successfully")
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		logAuth.Error("Failed to create token: ", err)
		return dto.AuthResponseDto{}, err
	}

	response := dto.AuthResponseDto{
		Token: token.Token,
	}

	logAuth.Infof("User ID %s has been authenticated successfully", user.Id_user)
	return response, nil
}

func (a *authUseCase) Register(payload dto.AuthRequestDto) (entity.User, error) {
	logAuth.Info("Starting to register a new user in the use case layer")
	return a.useCase.RegisterUser(entity.User{Username: payload.Username, Password: payload.Password})
}

func NewAuthUseCase(uc UserUsecase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{useCase: uc, jwtService: jwtService}
}
