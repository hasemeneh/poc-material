package authentication

import (
	"context"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (u *module) Register(ctx context.Context, req *models.RegisterReq) error {
	err := u.ValidateData(ctx, req)
	if err != nil {
		return fmt.Errorf("[authentication.Register] failed on validate data :%w", err)
	}
	err = u.UserRepo.RegisterUser(ctx, req.AuthID, req.Password)
	if err != nil {
		return fmt.Errorf("[authentication.Register] failed to register user : %w", err)
	}
	return nil
}
func (u *module) ValidateData(ctx context.Context, req *models.RegisterReq) error {
	// add more validation here
	if req.AuthID == "" {
		return custerr.NewInvalidInput("username should not be empty")
	}
	if req.Password == "" {
		return custerr.NewInvalidInput("password should not be empty")
	}
	if len(req.Password) < constants.MinPasswordLength {
		return custerr.NewInvalidInput(fmt.Sprintf("password length must be more than or equal to %d", constants.MinPasswordLength))
	}
	if req.Password != req.ConfirmPassword {
		return custerr.NewInvalidInput("confirm password and password mismatch")
	}

	return nil
}
