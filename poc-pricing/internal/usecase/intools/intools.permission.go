package intools

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

func (m *module) GetPermission(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error) {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodePermissionsView)
	if err != nil {
		return nil, err
	}

	return m.permission.GetPermissionForInternal(ctx, queries)
}

func (m *module) AddPermission(ctx context.Context, req *models.Customer, token string) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodePermissionsAdd)
	if err != nil {
		return err
	}

	return m.permission.AddPermission(ctx, req)
}

func (m *module) UpdatePermission(ctx context.Context, req *models.Customer, token string) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodePermissionsUpdate)
	if err != nil {
		return err
	}

	return m.permission.UpdatePermission(ctx, req)
}
