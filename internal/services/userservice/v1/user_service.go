package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dmzsz/duozhuayu/internal/constants"
	Repositories "github.com/dmzsz/duozhuayu/internal/datasources/repositories/postgres/v1"
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
	"github.com/dmzsz/duozhuayu/pkg/helpers"
	"github.com/dmzsz/duozhuayu/pkg/jwt"
	"github.com/dmzsz/duozhuayu/pkg/mail"
	"github.com/yhagio/go_api_boilerplate/services/emailservice"
)

type userUsecase struct {
	repo   Repositories.UserRepository
	mailer mail.MailImpl
}

// type user struct {
// 	username string
// }

func NewUserUsecase(repo Repositories.UserRepository, mailer emailservice.EmailService) V1Domains.UserUsecase {
	return &userUsecase{
		repo:   repo,
		mailer: mailer,
	}
}
func (userUC *userUsecase) Delete(ctx context.Context, inDom *V1Domains.UserDomain) (statusCode int, err error) {
	user, err := userUC.repo.GetByWithRoleByField(ctx, "Id", inDom.Id)
	if err != nil {
		return http.StatusNotFound, errors.New("user not found")
	}
	userUC.repo.Delete(ctx, &user)
	return http.StatusOK, nil
}
func (userUC *userUsecase) Store(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error) {
	inDom.Password, err = helpers.GenerateHash(inDom.Password)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	inDom.CreatedAt = time.Now().In(constants.GMT8)
	fmt.Println(time.Now().In(constants.GMT8))
	err = userUC.repo.Store(ctx, inDom)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	outDom, err = userUC.repo.GetByWithRoleByField(ctx, "email", inDom.Email)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	return outDom, http.StatusCreated, nil
}

func (userUC *userUsecase) Login(ctx context.Context, inDom *V1Domains.UserDomain) (outDom V1Domains.UserDomain, statusCode int, err error) {
	userDomain, err := userUC.repo.GetByWithRoleByField(ctx, "username", inDom.Username)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusUnauthorized, errors.New("invalid username or password") // for security purpose better use generic error message
	}

	if !userDomain.IsActive {
		return V1Domains.UserDomain{}, http.StatusForbidden, errors.New("account is not activated")
	}

	if !helpers.ValidateHash(inDom.Password, userDomain.Password) {
		return V1Domains.UserDomain{}, http.StatusUnauthorized, errors.New("invalid username or password")
	}

	userDomain.AccessToken, err = jwt.GenerateToken(userDomain.Id, userDomain.Username, userDomain.Email, *userDomain.Roles, jwt.AccessToken)

	if err != nil {
		return V1Domains.UserDomain{}, http.StatusInternalServerError, err
	}

	return userDomain, http.StatusOK, nil
}

func (userUC *userUsecase) SendOTP(ctx context.Context, email string) (otpCode string, statusCode int, err error) {
	domain, err := userUC.repo.GetByWithRoleByField(ctx, "email", email)
	if err != nil {
		return "", http.StatusNotFound, errors.New("email not found")
	}

	if domain.IsActive {
		return "", http.StatusBadRequest, errors.New("account already activated")
	}

	code, err := helpers.GenerateOTPCode(6)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if err = userUC.mailer.SendOTP(code, email); err != nil {
		return "", http.StatusInternalServerError, err
	}

	return code, http.StatusOK, nil
}

func (userUC *userUsecase) VerifOTP(ctx context.Context, email string, userOTP string, otpRedis string) (statusCode int, err error) {
	domain, err := userUC.repo.GetByWithRoleByField(ctx, "email", email)
	if err != nil {
		return http.StatusNotFound, errors.New("email not found")
	}

	if domain.IsActive {
		return http.StatusBadRequest, errors.New("account already activated")
	}

	if otpRedis != userOTP {
		return http.StatusBadRequest, errors.New("invalid otp code")
	}

	return http.StatusOK, nil
}

func (userUC *userUsecase) ActivateUser(ctx context.Context, email string) (statusCode int, err error) {
	user, err := userUC.repo.GetByWithRoleByField(ctx, "email", email)
	if err != nil {
		return http.StatusNotFound, errors.New("email not found")
	}

	if err = userUC.repo.ChangeActiveUser(ctx, &V1Domains.UserDomain{Id: user.Id, IsActive: true}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (userUC *userUsecase) GetByEmail(ctx context.Context, inDom *V1Domains.UserDomain, decryptEmail bool) (outDom V1Domains.UserDomain, statusCode int, err error) {
	user, err := userUC.repo.GetByEmail(ctx, inDom, decryptEmail)
	if err != nil {
		return V1Domains.UserDomain{}, http.StatusNotFound, errors.New("email not found")
	}

	return user, http.StatusOK, nil
}
