package permission

import (
	"context"
	"errors"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

func (r *repo) GetPermissionForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error) {
	rf := PermissionPageSpec.GetResultFilter(queries)
	data, total, err := r.queries.FetchPermissionForInternal(ctx, rf)
	if err != nil {
		return nil, err
	}
	var totalPage int64
	if rf.Limit > 0 {
		totalPage = (total / rf.Limit)
	}

	return &pagespecifier.IntoolsResponse{
		Table: pagespecifier.IntoolsTable{
			Limit:     rf.Limit,
			Page:      rf.Page,
			TotalPage: totalPage + 1,
			Rows:      data,
			Total:     total,
		},
	}, nil
}

func (r *repo) AddPermission(ctx context.Context, req *models.Permission) error {
	_, err := r.queries.InsertPermission(ctx, nil, req)
	return err
}

func (r *repo) UpdatePermission(ctx context.Context, req *models.Permission) error {
	return r.queries.UpdatePermission(ctx, nil, req)
}

func (r *repo) AddRolePermission(ctx context.Context, req *models.InternalRolePermission) error {
	_, err := r.queries.InsertRolePermission(ctx, nil, req)
	return err
}

func (r *repo) AddBulkRolePermission(ctx context.Context, dbTx database.DBTX, req []models.InternalRolePermission, roleID int64) error {
	var err error
	if dbTx == nil {
		return errors.New("database transaction should not be nil if using bulk")
	}
	for _, Permission := range req {
		Permission.RoleID = roleID
		_, err = r.queries.InsertRolePermission(ctx, dbTx, &Permission)
		if err != nil {
			return err
		}
	}

	return err
}
func (r *repo) GetPermissionByRoleID(ctx context.Context, RoleID int64) ([]*models.Permission, error) {
	return r.queries.FetchPermissionByRoleID(ctx, RoleID)
}
func (r *repo) ChangeRolePermissionStatus(ctx context.Context, status constants.RolePermissionStatus, RoleID, PermissionID int64) error {
	return r.queries.ChangeStatusRolePermission(ctx, nil, status, RoleID, PermissionID)
}
