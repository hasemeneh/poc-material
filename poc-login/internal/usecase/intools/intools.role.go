package intools

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

func (m *module) GetRole(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error) {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRoleView)
	if err != nil {
		return nil, err
	}

	return m.role.GetRoleForInternal(ctx, queries)
}

func (m *module) UpdateRole(ctx context.Context, token string, req *models.InternalRole) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRoleUpdate)
	if err != nil {
		return err
	}

	return m.role.UpdateRole(ctx, req)
}

func (m *module) AddRole(ctx context.Context, req *models.AddRoleReq, token string) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRoleAdd)
	if err != nil {
		return err
	}
	dbTx, err := m.permission.StartTransaction(ctx)
	if err != nil {
		return err
	}

	defer dbTx.Rollback()

	roleID, err := m.role.AddRole(ctx, dbTx, &req.Role)
	if err != nil {
		return err
	}
	if len(req.PermissionIDs) > 0 {
		err = m.permission.AddBulkRolePermission(ctx, dbTx, req.PermissionIDs, roleID)
		if err != nil {
			return err
		}
	}
	return dbTx.Commit()
}

func (m *module) AddPermissionToRole(ctx context.Context, token string, req *models.InternalRolePermission) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRolePermissionAdd)
	if err != nil {
		return err
	}

	return m.permission.AddRolePermission(ctx, req)
}

func (m *module) UpdateRolePermission(ctx context.Context, token string, req *models.InternalRolePermission) error {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRolePermissionUpdate)
	if err != nil {
		return err
	}

	return m.permission.ChangeRolePermissionStatus(ctx, req.Status, req.RoleID, req.PermissionID)
}

func (m *module) GetPermissionByRole(ctx context.Context, roleID int64, token string) ([]*models.Permission, error) {
	err := m.IsValidUser(ctx, token, constants.IntoolsPermissionCodeRolePermissionView)
	if err != nil {
		return nil, err
	}

	return m.permission.GetPermissionByRoleID(ctx, roleID)
}
