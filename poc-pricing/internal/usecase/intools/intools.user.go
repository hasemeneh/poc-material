package intools

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

func (m *module) GetUser(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error) {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeUserView)
	if err != nil {
		return nil, err
	}

	return m.user.GetUserForInternal(ctx, queries)
}

func (m *module) AddUserAccess(ctx context.Context, req *models.UserAccess, token string) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeUserAccessAdd)
	if err != nil {
		return err
	}
	return m.role.AddUserAccess(ctx, req)
}

func (m *module) ChangeStatusUserAccess(ctx context.Context, req *models.UserAccess, token string) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeUserAccessUpdate)
	if err != nil {
		return err
	}
	return m.role.ChangeStatusUserAccess(ctx, req)
}

func (m *module) GetUserAccessByIDInternal(ctx context.Context, userID int64, token string) ([]*models.UserAccess, error) {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeUserAccessView)
	if err != nil {
		return nil, err
	}
	return m.role.FetchUserAccessByIDInternal(ctx, userID)
}
